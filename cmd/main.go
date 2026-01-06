package main

import (
	"ChatService/internal/chat"
	pb "ChatService/proto/pkg/api/chat"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	address := ":9091"

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
