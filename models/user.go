package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name"`
	Surname  string `gorm:"not null" json:"surname"`
	Email    string `gorm:"not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Role     Role   `json:"role" validate:"required"` // 0 : user 1: teacher 2: chair
	Goals    []Goal `json:"goals"`
}

type Role int

const (
	Student Role = 1
	Teacher Role = 10
	Chair   Role = 20
)
