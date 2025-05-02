package main

import (
	"log"
	"net"

	pb "github.com/dzaakk/simple-grpc-app/internal/user/generated"
	"github.com/dzaakk/simple-grpc-app/internal/user/handler"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	userHandler := handler.NewUserHandler()
	pb.RegisterUserServiceServer(grpcServer, userHandler)
	log.Println("gRPC server listening on: 50051")
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
