package proxy

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	target string
	router *chi.Mux
}

func NewServer(t string) *Server {
	return &Server{
		target: t,
		router: chi.NewRouter(),
	}

}

func (s *Server) initRoute() error {

	s.router.Get("/", s.handler)
	s.router.Post("/api/v1", s.postHandler)

	return nil

}

func (s *Server) Listen() error {

	s.initRoute()

	addr := "localhost:8080"
	port := strings.Split(addr, ":")[1]
	fmt.Printf("Server Listening on %s\n ", port)
	return http.ListenAndServe(addr, s.router)

}

func (s *Server) postHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Could not read body ", http.StatusBadRequest)
	}

	data := string(body)

	fmt.Println(data)
}

func (s *Server) handler(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error handling File", http.StatusInternalServerError)
	}
	proxyAddr := string(body)
	fmt.Print(proxyAddr)
	w.Write([]byte("Testing "))

}
