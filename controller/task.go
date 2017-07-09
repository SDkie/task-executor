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
		Url        string    `valid:"required"`
		Method     string    `valid:"required"`
		RetryUntil time.Time `valid:"required"`
		Data       interface{}
	}{}

	err := json.NewDecoder(r.Body).Decode(input)
	if err != nil {
		err = fmt.Errorf("Error: JSON Decode, %s", err)
		logrus.Error(err)
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
}
