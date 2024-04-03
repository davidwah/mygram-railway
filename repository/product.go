package repository

import (
	"mygram-railway/core"
	"mygram-railway/database"
)

type ProductRepo struct {
	db *database.Postgres
}

type ProductImpl interface {
	FindProductByID(ID int64) (*core.Product, error)
}

func NewProductRepo(db *database.Postgres) *ProductRepo {
	return &ProductRepo{
		db: db,
	}
}
