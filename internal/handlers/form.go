package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func FormHandler() http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("internal/templates/form.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		if err := tmpl.Execute(w, nil); err != nil {
			log.Fatal(err)
		}
	}
}
