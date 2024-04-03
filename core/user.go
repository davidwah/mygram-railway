package core

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name           string
	Email          string
	Password       string
	Age            int64
	ProfilImageURL string
	Quantity       int64
}

type FindOneUserResponse struct {
	UserName       string `json:"user_name"`
	Email          string `json:"email"`
	Age            int64  `json:"age"`
	ProfilImageURL string `json:"profil_image_url"`
}
