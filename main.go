package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/piotrromanowski/twirpapp/rpc"
)

type Server struct{}

func (s *Server) Greet(ctx context.Context, p *rpc.Person) (*rpc.Greeting, error) {
	return &rpc.Greeting{Phrase: fmt.Sprintf("Hi %s", p.Name)}, nil
}

func main() {
	greetServer := rpc.NewGreeterServer(&Server{})

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", greetServer)
}
