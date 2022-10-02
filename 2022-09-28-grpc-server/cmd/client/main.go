package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/zoonoo/pocs/2022-09-28-grpc-server/gen/bookshop/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:8080", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	fmt.Printf("Hello, this is GRPC client.\n")

	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewInventoryClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetBookList(ctx, &pb.GetBookListRequest{})
	if err != nil {
		log.Fatalf("could not retrieve books : %v", err)
	}
	log.Printf("Book : %s", r.GetBooks())
}
