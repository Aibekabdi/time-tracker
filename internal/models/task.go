package models

import "time"

type Task struct {
	ID          uint
	UserID      uint
	Title       string
	Description string
	Start_at    time.Time
	End_at      time.Time
}
