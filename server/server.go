package server

import (
	"clickdata/component"
	"fmt"
	"net/http"
	"log"

	"github.com/a-h/templ"
)

type Server struct {
	address string
	port int
	mux *http.ServeMux
}

func NewServer(address string, port int) Server {
	mux := http.NewServeMux()

	mux.Handle("/", templ.Handler(component.Index()))

	return Server{
		address,
		port,
		mux,
	}
}

func (s Server) Run() {
	address := fmt.Sprintf("%s:%d", s.address, s.port)
	log.Printf("Running server on %s", address)
	http.ListenAndServe(address, s.mux)
}
