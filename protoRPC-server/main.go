package main

import (
	"context"
	"log"
	"net"

	protogRPC "github.com/lleonesouza/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	grpcServer := grpc.NewServer()
	protogRPC.RegisterAddServiceServer(grpcServer, &server{})
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server on port 9000: %v", err)
	}

}

func (s *server) Add(ctx context.Context, request *protogRPC.Request) (*protogRPC.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b

	return &protogRPC.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, request *protogRPC.Request) (*protogRPC.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a * b

	return &protogRPC.Response{Result: result}, nil
}
