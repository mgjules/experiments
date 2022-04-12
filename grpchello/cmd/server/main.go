package main

import (
	"context"
	"log"
	"net"

	"github.com/mgjules/toys/grpchello/hello"
	"google.golang.org/grpc"
)

func main() {
	l, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	helloSrv := helloServer{}
	hello.RegisterHelloWorldServer(srv, helloSrv)
	log.Fatalln(srv.Serve(l))
}

type helloServer struct{}

func (s helloServer) GetHello(ctx context.Context, void *hello.Void) (*hello.Hello, error) {
	out := &hello.Hello{Text: "Hello World w/ gRPC ;)"}
	return out, nil
}
