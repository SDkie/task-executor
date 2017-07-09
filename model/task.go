package model

import "time"

//go:generate easytags task.go json
//go:generate easytags task.go sql

type Task struct {
	ID uint64 `gorm:"primary_key" sql:"id" json:"id"`

	Url        string    `sql:"url" json:"url"`
	Method     string    `sql:"method" json:"method"`
	RetryUntil time.Time `sql:"retry_until" json:"retry_until"`
	Data       string    `sql:"data" json:"data"`

	Status    int    `sql:"status" json:"-"`
	StatusMsg string `sql:"status_msg" json:"-"`

	// JSON String
	CreatedAt time.Time  `sql:"created_at" json:"-"`
	UpdatedAt time.Time  `sql:"updated_at" json:"-"`
	DeletedAt *time.Time `sql:"deleted_at" json:"-"`
}

func (t *Task) Create() error {
	t.Status = TASK_NOT_EXECUTED
	return db.Create(t).Error
}

const (
	TASK_NOT_EXECUTED = iota
	TASK_EXECUTED
	TASK_ABORTED
)
