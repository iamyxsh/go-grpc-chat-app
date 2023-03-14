module github.com/iamyxsh/go-grpc-chat-app/backend/services/chat

replace github.com/iamyxsh/go-grpc-chat-app/backend/packages/db => ../../packages/db

replace github.com/iamyxsh/go-grpc-chat-app/backend/packages/models => ../../packages/models

replace github.com/iamyxsh/go-grpc-chat-app/backend/packages/utils => ../../packages/utils

replace github.com/iamyxsh/go-grpc-chat-app/backend/packages/kafka => ../../packages/kafka

replace github.com/iamyxsh/go-grpc-chat-app/backend/packages/logger => ../../packages/logger

go 1.20

require (
	github.com/google/uuid v1.3.0
	github.com/iamyxsh/go-grpc-chat-app/backend/packages/db v0.0.0-00010101000000-000000000000
	github.com/iamyxsh/go-grpc-chat-app/backend/packages/kafka v0.0.0-00010101000000-000000000000
	github.com/iamyxsh/go-grpc-chat-app/backend/packages/models v0.0.0-00010101000000-000000000000
	github.com/scylladb/gocqlx/v2 v2.8.0
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/gocql/gocql v1.3.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	github.com/scylladb/go-reflectx v1.0.1 // indirect
	github.com/segmentio/kafka-go v0.4.39 // indirect
	golang.org/x/crypto v0.7.0 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
)
