package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
}

func newServer() *server {
	s := &server{
		router: mux.NewRouter(),
	}

	s.router.PathPrefix("/").HandlerFunc(s.anyRequestHandler())

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) anyRequestHandler() http.HandlerFunc {
	type response struct {
		Host       string      `json:"host"`
		UserAgent  string      `json:"user_agent"`
		RequestUri string      `json:"request_uri"`
		Header     http.Header `json:"headers"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewEncoder(w).Encode(response{
			r.Host,
			r.UserAgent(),
			r.RequestURI,
			r.Header,
		}); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Fatal(err)
		}
	}
}

func Start(c *Config) error {
	srv := newServer()
	log.Println("listen on", c.Port)
	return http.ListenAndServe(c.Port, srv)
}
