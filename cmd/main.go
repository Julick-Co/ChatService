package main

import (
	"chatservice/internal/chat"
	pb "chatservice/proto/api/v1/chat"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	address := ":9093"

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("listen error: %v", err)
	}

	server := grpc.NewServer()

	pb.RegisterChatServiceServer(server, &chat.ChatServiceImpl{})

	log.Printf("gRPC server running on %s\n", address)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("serve error: %v", err)
	}
}
