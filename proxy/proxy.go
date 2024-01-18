package proxy

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	host   string
	port   string
	router *chi.Mux
}

func NewServer(h string, p string) *Server {
	s := Server{
		host:   h,
		port:   p,
		router: chi.NewRouter(),
	}

	return &s

}

func (s *Server) initRoute() error {

	s.router.HandleFunc("/", s.handler)

	return nil

}

func (s *Server) Listen() error {

	s.initRoute()
	addr := s.host + ":" + s.port
	return http.ListenAndServe(addr, s.router)

}
func (s *Server) handler(res http.ResponseWriter, req *http.Request) {

	res.Write([]byte("Testing "))

}
