package chat

import (
	pb "chatservice/proto/pkg/api/v1/chat"
	"context"
	"fmt"
	"log"
	"time"
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

func (s *ChatServiceImpl) StreamMessages(request *pb.StreamMessagesRequest,
	stream pb.ChatService_StreamMessagesServer) error {

	log.Println("Достучался до метода StreamMessages")

	for i := 0; i < 3; i++ {
		j := 3
		msg := &pb.StreamMessagesResponse{
			Message: &pb.Message{
				Id:            fmt.Sprintf("%d", i),
				ChatId:        fmt.Sprintf("%d", j),
				SenderId:      "system",
				Text:          "hello",
				CreatedUnixMs: time.Now().UnixMilli(),
			},
		}

		j--

		if err := stream.Send(msg); err != nil {
			return err
		}
	}

	return nil
}

func (s *ChatServiceImpl) ListChatMembers(ctx context.Context, req *pb.ListChatMembersRequest,
) (*pb.ListChatMembersResponse, error) {

	log.Println("Достучался до метода ListChatMembers")

	return &pb.ListChatMembersResponse{
		UserIds: []string{},
	}, nil
}

func (s *ChatServiceImpl) SendMessages(ctx context.Context, req *pb.SendMessagesRequest,
) (*pb.SendMessagesResponse, error) {

	log.Println("Достучался до метода SendMessages")

	return &pb.SendMessagesResponse{
		Messages: &pb.Message{
			Id:            "123",
			ChatId:        "321",
			SenderId:      "1",
			Text:          "qwerty",
			CreatedUnixMs: 1,
		},
	}, nil
}

func (s *ChatServiceImpl) ListMessages(ctx context.Context, req *pb.ListMessagesRequest,
) (*pb.ListMessagesResponse, error) {

	log.Println("Достучался до метода ListMessages")

	nextCursor := "111"

	messages := []*pb.Message{
		{
			Id:            "123",
			ChatId:        req.ChatId,
			SenderId:      "1",
			Text:          "qwerty",
			CreatedUnixMs: 1,
		},
	}

	return &pb.ListMessagesResponse{
		Messages:   messages,
		NextCursor: &nextCursor,
	}, nil
}
