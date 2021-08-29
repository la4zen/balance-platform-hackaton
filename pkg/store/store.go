package store

import (
	"log"

	"github.com/la4zen/balance-platform-hackaton/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store struct {
	DB *gorm.DB
}

func NewStore() *Store {
	db, err := gorm.Open(postgres.Open(utils.DB_DSN))
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate()
	return &Store{
		DB: db,
	}
}
