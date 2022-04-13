package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mgjules/experiments/grpchello/hello"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":4000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not dial to server: %v", err)
	}
	c := hello.NewHelloWorldClient(conn)

	res, err := getHello(context.Background(), c)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res.Text)
}

func getHello(ctx context.Context, client hello.HelloWorldClient) (*hello.Hello, error) {
	c, err := client.GetHello(ctx, &hello.Void{})
	if err != nil {
		return nil, fmt.Errorf("could not rpc hello: %v", err)
	}

	return c, nil
}
