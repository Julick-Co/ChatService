package chat

import (
	pb "ChatService/proto/pkg/api/chat"
)

type ChatServiceImpl struct {
	pb.UnimplementedChatServiceServer
}
