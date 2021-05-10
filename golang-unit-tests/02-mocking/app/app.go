package app

import (
	"database/database"
	"fmt"
)

type App struct {
	db database.Database
}

func NewApp(db database.Database) *App {
	return &App{db: db}
}

func (a *App) AddUser(username, email string) error {
	if a.db.Exists(email) {
		return fmt.Errorf("email %s já existe", email)
	}
	return a.db.Insert(email, username)
}
