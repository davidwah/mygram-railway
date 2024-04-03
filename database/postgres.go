package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mygram-railway/core"
	"os"
)

type Postgres struct {
	DB  *gorm.DB
	Err error
}

func NewPostgres() *Postgres {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("PGHOST"), os.Getenv("PGUSER"), os.Getenv("PGPASS"),
		os.Getenv("PGDBNAME"), os.Getenv("PGPORT"))

	db, err := gorm.Open(postgres.Open(dsn), nil)
	db.Debug().AutoMigrate(&core.Product{}, &core.User{})

	return &Postgres{
		DB:  db,
		Err: err,
	}
}
