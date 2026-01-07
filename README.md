# ChatService
Сервис для управления чатами

## Генерация protobuf и gRPC
Запусти команды в терминале:

```bash
# Установка генераторов Go
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

```bash
$ protoc --proto_path=./api/chat/v1 --go_out=. --go-grpc_out=. api/chat/v1/messages.proto api/chat/v1/service.proto
```
