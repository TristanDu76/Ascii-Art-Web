package utils

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func Serve404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	tmpl, err := template.ParseFiles(filepath.Join("templates", "404.html"))
	if err != nil {
		http.Error(w, "404 - Page Not Found", http.StatusNotFound)
		return
	}
	tmpl.Execute(w, nil)
}

func ServeError(w http.ResponseWriter, r *http.Request, statusCode int, tmplFile string) {
	w.WriteHeader(statusCode)
	tmpl, err := template.ParseFiles(filepath.Join("templates", tmplFile))
	if err != nil {
		http.Error(w, http.StatusText(statusCode), statusCode)
		return
	}
	tmpl.Execute(w, nil)
}
