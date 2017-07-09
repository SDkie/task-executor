package config

import (
	"os"

	"github.com/asaskevich/govalidator"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Port       string `valid:"required"`
	DbHost     string `valid:"required"`
	DbName     string `valid:"required"`
	DbUser     string `valid:"required"`
	DbPassword string
	DbPort     string `valid:"required"`
}

var config Config

func init() {
	config.Port = os.Getenv("PORT")
	config.DbHost = os.Getenv("DB_HOST")
	config.DbName = os.Getenv("DB_NAME")
	config.DbUser = os.Getenv("DB_USER")
	config.DbPassword = os.Getenv("DB_PASSWORD")
	config.DbPort = os.Getenv("DB_PORT")

	_, err := govalidator.ValidateStruct(&config)
	if err != nil {
		logrus.Panicf("Environment Variables not set:, %s", err)
	}

	logrus.Info("Config : Initialized")
}

func Get() *Config {
	return &config
}
