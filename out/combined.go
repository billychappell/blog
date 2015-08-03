package main

import (
	"net"
	"strings"
	"time"
)

func httpsHost() string {
	if *prod {
		return "" // TODO: Fill in prod address
	}
	if v := *httpAddr; strings.HasPrefix(v, ":") {
		return "localhost" + v
	} else {
		return v
	}
}

func httpHost() string {
	if *prod {
		return "" // TODO: Fill in prod address
	}
	if v := *httpAddr; strings.HasPrefix(v, ":") {
		return "localhost" + v
	} else {
		return v
	}
}

type tcpKeepAlive struct {
	*net.TCPListener
}

func (ln tcpKeepAlive) Accept() (c net.Conn, err error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return
	}
	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(3 * time.Minute)
	return tc, nil
}

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func gzipHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// if header doesn't indicate gzip is accepted, return as is.
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			fn(w, r)
			return
		}
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Set("Content-Type", "text/html")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		fn(gzipResponseWriter{Writer: gz, ResponseWriter: w}, r)
	}
}

var t = template.Must(template.ParseGlob("tmpl/*"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if err := t.ExecuteTemplate(w, "index", &p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
