module github.com/iamyxsh/go-grpc-chat-app/backend/services/user

go 1.20

replace github.com/iamyxsh/go-grpc-chat-app/backend/packages/db => ../../packages/db

replace github.com/iamyxsh/go-grpc-chat-app/backend/packages/models => ../../packages/models

replace github.com/iamyxsh/go-grpc-chat-app/backend/packages/utils => ../../packages/utils

replace github.com/iamyxsh/go-grpc-chat-app/backend/packages/kafka => ../../packages/kafka

replace github.com/iamyxsh/go-grpc-chat-app/backend/packages/logger => ../../packages/logger

require (
	github.com/iamyxsh/go-grpc-chat-app/backend/packages/db v0.0.0-00010101000000-000000000000
	github.com/iamyxsh/go-grpc-chat-app/backend/packages/kafka v0.0.0-00010101000000-000000000000
	github.com/iamyxsh/go-grpc-chat-app/backend/packages/models v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.8.2
	golang.org/x/net v0.7.0
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.28.1
	gorm.io/driver/sqlite v1.4.4
	gorm.io/gorm v1.24.6
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.3.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/mattn/go-sqlite3 v1.14.15 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/segmentio/kafka-go v0.4.39 // indirect
	golang.org/x/crypto v0.6.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/driver/postgres v1.4.8 // indirect
)
