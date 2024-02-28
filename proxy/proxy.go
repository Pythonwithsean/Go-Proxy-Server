package proxy

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	target string
	router *chi.Mux
}

func NewServer(t string) *Server {
	s := Server{
		target: t,
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

	addr := "localhost:8080"
	port := strings.Split(addr, ":")[1]
	fmt.Printf("Server Listening on %s\n ", port)
	return http.ListenAndServe(addr, s.router)

}
func (s *Server) handler(res http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		http.Error(res, "Error handling File", http.StatusInternalServerError)
	}
	proxyAddr := string(body)
	fmt.Print(proxyAddr)
	res.Write([]byte("Testing "))

}
