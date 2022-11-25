package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}

func NewServer() *Server {
	s := &Server{
		router: mux.NewRouter(),
	}
	return s
}

func (s *Server) Routes() {
	s.router.HandleFunc("/", HomeHandler)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("./static"))

}

func GetVideosCountry(w http.ResponseWriter, r *http.Request) {
}
