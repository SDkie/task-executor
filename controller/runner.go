package controller

import (
	"html/template"
	"net/http"

	"github.com/SDkie/task-executor/model"
	"github.com/SDkie/task-executor/utils"
	"github.com/sirupsen/logrus"
)

func RunnerStatus(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/Status.html")
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	tasks, err := model.GetAllTasks()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = t.Execute(w, *tasks)
	if err != nil {
		logrus.Error(err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}
