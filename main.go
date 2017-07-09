package main

import (
	"net/http"

	"github.com/SDkie/task-executor/config"
	"github.com/SDkie/task-executor/controller"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	r := initRoutes()
	http.ListenAndServe(":"+config.Get().Port, r)
}

func initRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/task", controller.CreateTask).Methods("POST")
	logrus.Info("Route : Initialized")
	return r
}
