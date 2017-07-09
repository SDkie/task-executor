package model

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

//go:generate easytags task.go json
//go:generate easytags task.go sql

var retryTimeInterval = 10 * time.Second

type Task struct {
	ID uint64 `gorm:"primary_key" sql:"id" json:"id"`

	Url        string      `sql:"url" json:"url"`
	Method     string      `sql:"method" json:"method"`
	RetryUntil time.Time   `sql:"retry_until" json:"retry_until"`
	Data       interface{} `sql:"-" json:"data"`
	DataByte   []byte      `sql:"data_byte" json:"data_byte"`

	Status     int    `sql:"status" json:"-"`
	StatusMsg  string `sql:"status_msg" json:"-"`
	StatusCode int    `sql:"status_code" json:"-"`

	LastRun time.Time `sql:"last_run" json:"-"`
	NextRun time.Time `sql:"next_run" json:"-"`

	TotalRetry uint `sql:"total_retry" json:"-"`

	// JSON String
	CreatedAt time.Time  `sql:"created_at" json:"-"`
	UpdatedAt time.Time  `sql:"updated_at" json:"-"`
	DeletedAt *time.Time `sql:"deleted_at" json:"-"`
}

func (t *Task) Create() error {
	var err error
	t.Status = TASK_SCHEDULED
	t.DataByte, err = json.Marshal(t.Data)
	if err != nil {
		logrus.Error(err)
		return err
	}

	t.RetryUntil = time.Now().Add(1 * time.Hour * 24)
	err = db.Create(t).Error
	if err != nil {
		logrus.Error(err)
	}
	return err
}

func (t *Task) Update() error {
	return db.Save(t).Error
}

const (
	TASK_SCHEDULED = iota
	TASK_EXECUTED
	TASK_ABORTED
)

func (task Task) Run() {
	task.LastRun = time.Now()
	task.TotalRetry++

	var statusMsg string
	var statusCode int
	client := &http.Client{}
	task.LastRun = time.Now()

	req, err := http.NewRequest(task.Method, task.Url, bytes.NewBuffer(task.DataByte))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		logrus.Error(err)
	}

	resp, err := client.Do(req)
	if err == nil && resp.StatusCode == http.StatusOK {
		task.Status = TASK_EXECUTED
		task.StatusCode = http.StatusOK
		task.StatusMsg = "Successfully Executed"
		task.Update()
		return
	}

	if err != nil {
		logrus.Error(err)
		statusMsg = err.Error()
	} else {
		statusCode = resp.StatusCode
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Errorf("Error: error while reading Response Body, %s", err)
		} else {
			statusMsg = string(body)
		}
	}

	task.Status = TASK_SCHEDULED
	task.StatusCode = statusCode
	task.StatusMsg = statusMsg
	task.NextRun = time.Now().Add(retryTimeInterval)
	task.Update()
}
