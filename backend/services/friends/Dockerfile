FROM golang:1.19-alpine

WORKDIR /go/src/github.com/iamyxsh/go-grpc-chat-app/backend/services/friends

RUN go install github.com/cosmtrek/air@latest

COPY go.mod ./
COPY go.sum ./

CMD ["air"]