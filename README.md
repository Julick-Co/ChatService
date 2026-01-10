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
# Генерация кода из proto-файлов
protoc --proto_path=./api/v1/chat --go_out=. --go-grpc_out=. api/v1/chat/messages.proto api/v1/chat/service.proto
