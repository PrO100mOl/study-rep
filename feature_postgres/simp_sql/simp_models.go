package simpsql

import "time"

type Task struct {
	ID           int
	Title        string
	Description  string
	Completed    bool
	Created_at   time.Time
	Completed_at *time.Time
}
