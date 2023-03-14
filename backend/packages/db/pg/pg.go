package pg

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ReturnDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
