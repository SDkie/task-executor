package model

import (
	"github.com/SDkie/task-executor/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

var db *gorm.DB

func init() {
	var err error
	cfg := config.Get()
	connString := cfg.DbUser + ":" + cfg.DbPassword + "@tcp(" + cfg.DbHost + ":" + cfg.DbPort + ")/" + cfg.DbName
	db, err = gorm.Open("mysql", connString+"?parseTime=true")
	if err != nil {
		logrus.Panic(err)
	}

	db.SingularTable(true)
	db = db.LogMode(true)

	if err := db.DB().Ping(); err != nil {
		panic(err)
	}

	db.AutoMigrate(Task{})

	logrus.Info("Mysql : Initialized")
}
