package service

import (
	"mygram-railway/core"
	"mygram-railway/repository"
)

type UserService struct {
	userRepo *repository.UserRepo
}

func NewUserService(userRepo *repository.UserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) FindOneUser(id int) (*core.FindOneUserResponse, error) {
	user, err := s.userRepo.FindUserByID(uint(id))
	if err != nil {
		return nil, err
	}

	resp := &core.FindOneUserResponse{
		UserName:       user.Name,
		Email:          user.Email,
		Age:            user.Age,
		ProfilImageURL: user.ProfilImageURL,
	}

	return resp, nil
}
