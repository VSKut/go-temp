package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/gommon/log"
)

func FormPostHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	name := r.FormValue("name")
	address := r.FormValue("address")
	if name == "" || address == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	cookie := &http.Cookie{
		Name:  "token",
		Value: fmt.Sprintf("%s:%s", name, address),
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/", 302)
}
