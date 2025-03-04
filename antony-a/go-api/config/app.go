package config

import "os"

type App struct {
	Name string
	Port string
	DB   *Database
}

func NewApp() *App {
	appName := os.Getenv("APP_NAME")
	appPort := os.Getenv("APP_PORT")

	return &App{
		Name: appName,
		Port: appPort,
		DB:   NewDatabase(),
	}
}
