package main

import (
	"html/template"
	"net/http"
)

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
