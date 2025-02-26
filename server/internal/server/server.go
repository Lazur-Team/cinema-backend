package server

import (
	"fmt"
	"net/http"
	"server/internal/config"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})
	return http.ListenAndServe(":"+config.Cfg.ServerPort, nil)
}
