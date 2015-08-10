package main

import (
	"html/template"
	"io"
	"net/http"
	"time"

	"github.com/billychappell/downloader/database"
	"github.com/klauspost/compress/gzip"
	"golang.org/x/tools/godoc/static"

	"strings"
)

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func gzipHandler(fn http.HandlerFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// if header doesn't indicate gzip is accepted, return as is.
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			fn(w, r)
			return
		}
		w.Header().Set("Content-Encoding", "gzip")

		switch {
		case strings.Contains(r.URL.Path, ".css"):
			w.Header().Set("Content-Type", "text/css")
		case strings.Contains(r.URL.Path, ".js"):
			w.Header().Set("Content-Type", "application/javascript")
		case strings.Contains(r.URL.Path, ".html"):
			w.Header().Set("Content-Type", "text/html")
		default:
			w.Header().Set("Content-Type", "text/html")
		}
		gz := gzip.NewWriter(w)
		defer gz.Close()
		fn(gzipResponseWriter{Writer: gz, ResponseWriter: w}, r)
	}
}

var t = template.Must(template.ParseGlob("tmpl/*"))

func indexHandler(p []database.Post) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		if err := t.ExecuteTemplate(w, "index", p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path
	b, ok := static.Files[name]
	if !ok {
		http.NotFound(w, r)
		return
	}
	http.ServeContent(w, r, name, time.Time{}, strings.NewReader(b))
}

func registerHandlers(p []database.Post) {
	m := http.NewServeMux()
	http.HandleFunc("/", gzipHandler(func(w http.ResponseWriter, r *http.Request) {
		/* if r.TLS == nil {
			http.Redirect(w, r, "https://chappellmarketing.com", http.StatusFound)
			return
		}
		if r.ProtoMajor == 1 {
			homeOldHTTP(w, r)
			return
		} */
		m.ServeHTTP(w, r)
	}))
	m.HandleFunc("/", indexHandler(p))

	fs := http.FileServer(http.Dir("static/"))
	m.Handle("/static/", http.StripPrefix("/static/", fs))
}
