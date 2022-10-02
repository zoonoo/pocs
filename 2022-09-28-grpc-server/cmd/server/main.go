package main

import (
	"context"
	"log"
	"net"

	"github.com/zoonoo/pocs/2022-09-28-grpc-server/gen/bookshop/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type inventoryServer struct {
	pb.UnimplementedInventoryServer
}

func getSampleBooks() []*pb.Book {
	books := make([]*pb.Book, 0)
	books = append(books, &pb.Book{
		Title:     "Tao Te Ching",
		Author:    "Lao Zi",
		PageCount: 100000,
	})
	return books
}

func (s *inventoryServer) GetBookList(ctx context.Context, in *pb.GetBookListRequest) (*pb.GetBookListResponse, error) {
	return &pb.GetBookListResponse{
		Books: getSampleBooks(),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterInventoryServer(s, &inventoryServer{})
	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
