package models

import "gorm.io/gorm"

type Hello struct {
	gorm.Model
	Message string
}
