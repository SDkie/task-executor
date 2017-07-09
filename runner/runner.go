package runner

import (
	"github.com/SDkie/task-executor/model"
	"github.com/bamzi/jobrunner"
	"github.com/sirupsen/logrus"
)

func Init() {
	jobrunner.Start()
	jobrunner.Every(model.TaskInterval, TaskRunner{})
	logrus.Info("Runner : Initialized")
}

type TaskRunner struct{}

func (TaskRunner) Run() {
	tasks, err := model.FindScheduledTasks()
	if err != nil {
		logrus.Error(err)
	}

	for _, t := range *tasks {
		valid, err := t.CheckTaskValidity()
		if err != nil {
			logrus.Error(err)
			continue
		}
		if valid {
			go t.Run()
		}
	}
}
