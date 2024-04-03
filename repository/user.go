package repository

import (
	"mygram-railway/core"
	"mygram-railway/database"
)

type UserRepo struct {
	postgres *database.Postgres
}

func NewUserRepo(db *database.Postgres) *UserRepo {
	return &UserRepo{
		postgres: db,
	}
}

func (r *UserRepo) FindUserByID(id uint) (*core.User, error) {
	user := core.User{}
	err := r.postgres.DB.First(&user, "id=?", id).Error

	return &user, err
}
