package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type Postgres struct {
	db  *gorm.DB
	err error
}

func NewPostgres() *Postgres {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("PGHOST"), os.Getenv("PGUSER"), os.Getenv("PGPASS"),
		os.Getenv("PGDBNAME"), os.Getenv("PGPORT"))

	db, err := gorm.Open(postgres.Open(dsn), nil)

	return &Postgres{
		db:  db,
		err: err,
	}
}
