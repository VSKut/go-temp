package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vskut/go-temp/internal/handlers"
)

type server struct {
	router *mux.Router
}

func newServer() *server {
	s := &server{
		router: mux.NewRouter(),
	}

	s.router.HandleFunc("/", handlers.FormHandler()).Methods("GET")
	s.router.HandleFunc("/", handlers.FormPostHandler).Methods("POST")
	s.router.NotFoundHandler = http.HandlerFunc(handlers.NotFoundHandler())

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func Start(c *Config) error {
	srv := newServer()

	log.Println("Launched on", c.GetAddr())

	return http.ListenAndServe(c.GetAddr(), srv)
}
