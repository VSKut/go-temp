package handlers

import (
	"html/template"
	"net/http"

	"github.com/labstack/gommon/log"
)

func NotFoundHandler() http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("internal/templates/not_found.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		if err := tmpl.Execute(w, nil); err != nil {
			log.Fatal(err)
		}
	}
}
