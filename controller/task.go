package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/SDkie/task-executor/model"
	"github.com/SDkie/task-executor/utils"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
)

// CreateTask adds the entry of task in task executor
func CreateTask(w http.ResponseWriter, r *http.Request) {
	input := &struct {
		Url        string      `json:"url" valid:"required"`
		Method     string      `json:"method" valid:"required"`
		Data       interface{} `json:"data"`
		MaxRetry   int64       `json:"max_retry"`
		RetryUntil time.Time   `json:"retry_until"`
	}{}

	err := json.NewDecoder(r.Body).Decode(input)
	if err != nil {
		err = fmt.Errorf("Error: JSON Decode, %s", err)
		logrus.Error(err)
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if input.MaxRetry < 0 {
		err := fmt.Errorf("Error: MaxRetry is less then 0")
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if !input.RetryUntil.IsZero() && input.RetryUntil.Before(time.Now()) {
		err := fmt.Errorf("Error: RetryUnitl is less then current time")
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	task := new(model.Task)
	err = copier.Copy(task, input)
	if err != nil {
		err = fmt.Errorf("Error: Error task copier, %s", err)
		logrus.Error(err)
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = task.Create()
	if err != nil {
		err = fmt.Errorf("Error: Create Task %s", err)
		logrus.Error(err)
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteSuccessResponse(w, task)
}
