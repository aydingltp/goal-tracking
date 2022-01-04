package models

import (
	"gorm.io/gorm"
	"time"
)

type Goal struct {
	gorm.Model
	Name           string           `json:"name"`
	GoalDoneStatus []GoalDoneStatus `json:"goal_done_status"`
	UserID         int64            `json:"user_id"`
	User           User             `json:"user"`
	ClassroomID    int64            `json:"classroom_id"`
	Classroom      Classroom        `json:"classroom"`
}

type GoalDoneStatus struct {
	gorm.Model
	Date   time.Time `json:"dates"`
	IsDone bool      `json:"is_done"`
	GoalId int64     `json:"goal_id"`
	Goal   Goal      `json:"goal"`
}
