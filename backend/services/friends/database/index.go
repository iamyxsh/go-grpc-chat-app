package database

import (
	"github.com/iamyxsh/go-grpc-chat-app/backend/packages/models"
	"gorm.io/gorm"
)

const ENUM_STATEMENT = `
DO $$                                               
    BEGIN                                                        
        IF NOT EXISTS(SELECT oid FROM pg_type WHERE typname = 'friendship_status')
        THEN                                                        
            CREATE TYPE friendship_status AS ENUM (
            'sent',
            'accepted',
            'rejected'
            );
    END IF;                                                        
END $$;
`

func ExecStatements(db *gorm.DB) {
	db.Exec(ENUM_STATEMENT)
	db.AutoMigrate(&models.FriendUser{})
	db.AutoMigrate(&models.Friendship{})
}
