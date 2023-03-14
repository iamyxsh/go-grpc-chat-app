FROM golang:1.19-alpine

WORKDIR /go/src/github.com/iamyxsh/go-grpc-chat-app/backend/services/auth

RUN go install github.com/cosmtrek/air@latest

COPY go.mod ./
COPY go.sum ./

CMD ["air"]