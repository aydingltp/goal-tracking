package models

import "gorm.io/gorm"

type Classroom struct {
	gorm.Model
	Name              string             `json:"name"`
	Goals             []Goal             `json:"goals"`
	SchoolID          int64              `json:"school_id"`
	School            School             `json:"school"`
	ClassroomStudents []ClassroomStudent `json:"classroom_students"`
}

type ClassroomStudent struct {
	gorm.Model
	UserID      int64     `json:"user_id"` //ogrencinin id'si
	User        User      `json:"user"`
	ClassroomID int64     `json:"classroom_id"`
	Classroom   Classroom `json:"classroom"`
}

type ClassroomTeacher struct {
	gorm.Model
	UserID      int64     `json:"user_id"` //ogretmenin id'si
	User        User      `json:"user"`
	ClassroomID int64     `json:"classroom_id"`
	Classroom   Classroom `json:"classroom"`
	SchoolID    int64     `json:"school_id"`
	School      School    `json:"school"`
}

type ClassroomDirectory struct {
	gorm.Model
	UserID int64 `json:"user_id"` //müdürün id'si
	User   User  `json:"user"`
}
