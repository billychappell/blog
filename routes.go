package main

import (
	"net/http"
)

func registerHandlers() {
	m := http.NewServeMux()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		/* if r.TLS == nil {
			http.Redirect(w, r, "https://chappellmarketing.com", http.StatusFound)
			return
		}
		if r.ProtoMajor == 1 {
			homeOldHTTP(w, r)
			return
		} */
		m.ServeHTTP(w, r)
	})
	m.HandleFunc("/", gzipHandler(indexHandler))
}
