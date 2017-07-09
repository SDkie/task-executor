package main

import (
	"net/http"

	"github.com/SDkie/task-executor/config"
	"github.com/SDkie/task-executor/controller"
	"github.com/SDkie/task-executor/runner"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	r := initRoutes()
	runner.Init()
	http.ListenAndServe(":"+config.Get().Port, r)
}

func initRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/task", controller.CreateTask).Methods("POST")
	r.HandleFunc("/runner/status", controller.RunnerStatus).Methods("GET")
	logrus.Info("Route : Initialized")
	return r
}
