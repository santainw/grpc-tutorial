package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/prongbang/grpc-kid/helloworld/helloworld"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "World"
)

func main() {

	// Set up a connection to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)

	rs, errs := c.SayHelloAgain(ctx, &pb.HelloRequest{Name: name})
	if errs != nil {
		log.Fatalf("could not greet: %v", errs)
	}
	log.Printf("Greeting: %s", rs.Message)

}
