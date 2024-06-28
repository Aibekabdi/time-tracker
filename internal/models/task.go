package models

import "time"

type Task struct {
	UserID      uint
	Title       string
	StartTime   time.Time
	EndTime     time.Time
	Description string
}
