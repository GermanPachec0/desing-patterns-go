package http_server

import (
	"fmt"
	"net/http"
)

type Server interface {
	Start() error
}

type HttpServer struct {
	Port string
}

func NewHttpServer(port string) *HttpServer {
	return &HttpServer{
		Port: port,
	}
}

func (s *HttpServer) Start() error {
	http.HandleFunc("/hello", hello)
	err := http.ListenAndServe(fmt.Sprintf(":%s", s.Port), nil)
	if err != nil {
		return err
	}
	fmt.Printf("Server Running on port: %s", s.Port)
	return nil
}

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}
