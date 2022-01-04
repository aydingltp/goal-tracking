package models

import "gorm.io/gorm"

type Classroom struct {
	gorm.Model
	Name   string `json:"name"`
	UserID int64  `json:"user_id"`
	User   User   `json:"user"`
	Goals  []Goal
}
