package models

import "gorm.io/gorm"

type School struct {
	gorm.Model
	ClassroomDirectoryID int64              `json:"classroom_directory_id"`
	ClassroomDirectory   ClassroomDirectory `json:"classroom_directory"`

	ClassroomTeachers []ClassroomTeacher `json:"classroom_teacher"`
}
