package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

type Database struct {
	Pool *gorm.DB
}

var dbTimeout = 3 * time.Second

//var ctx = func() context.Context {
//	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
//	go func() {
//		time.Sleep(dbTimeout)
//		cancel()
//	}()
//	return ctx
//}

// InitDb assumes DSN is stored in an environment variable
func InitDb() (*Database, error) {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}
	return &Database{
		db,
	}, nil
}
