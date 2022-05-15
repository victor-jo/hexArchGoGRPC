package api

import (
	"hex/internal/application/domain"
	"hex/internal/ports"
)

type UserApplication struct {
	db   ports.UserDbPorts
	user User
}

func NewUserApplication(db ports.UserDbPorts, user User) *UserApplication {
	return &UserApplication{
		db:   db,
		user: user,
	}
}

func (app UserApplication) Join(username string, password string) (uint64, error) {
	id, err := app.db.CreateUser(username, password)
	if err != nil {
		return 0, err
	}
	return id, err
}

func (app UserApplication) Find(id uint64) (domain.UserEntity, error) {
	return app.db.FindUser(id)
}
