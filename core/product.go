package core

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name     string
	Quantity int64
}
