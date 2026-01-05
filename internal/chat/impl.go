package chat

import (
	pb "ChatService/proto/pkg/api/chat"
	"context"
	"log"
)

type ChatServiceImpl struct {
	pb.UnimplementedChatServiceServer
}

func (s *ChatServiceImpl) CreateDirectChat(ctx context.Context,
	req *pb.CreateDirectChatRequest) (*pb.CreateDirectChatResponse, error) {

	log.Println("Достучался до метода CreateDirectChat")

	return &pb.CreateDirectChatResponse{
		ChatId: "generated_id",
	}, nil
}

func (s *ChatServiceImpl) GetChat(
	ctx context.Context,
	req *pb.GetChatRequest,
) (*pb.GetChatResponse, error) {

	log.Println("Достучался до метода GetChat")

	return &pb.GetChatResponse{
		Chat: &pb.Chat{
			ChatId: req.GetChatId(),
			Type:   pb.Chat_DIRECT,
		},
	}, nil
}

func (s *ChatServiceImpl) ListUserChats(
	ctx context.Context,
	req *pb.ListUserChatsRequest,
) (*pb.ListUserChatsResponse, error) {

	log.Println("Достучался до метода ListUserChats")

	return &pb.ListUserChatsResponse{
		Chats: []*pb.Chat{
			{
				ChatId: "chat_1",
				Type:   pb.Chat_DIRECT,
			},
			{
				ChatId: "chat_2",
				Type:   pb.Chat_GROUP,
			},
		},
	}, nil
}
