package dao

import (
	"opatutorial/models"
)

const (
	qUserSelectAll = "SELECT * FROM users"
)

// IUser ...
type IUser interface {
	GetAllUser() ([]*models.User, error)
}

type userDAO struct {
}

func (u *userDAO) GetAllUser() ([]*models.User, error) {
	users := []*models.User{}
	db := ConfigurationManager.GetDB()
	if err := db.Select(&users, qUserSelectAll); err != nil {
		return nil, err
	}
	return users, nil
}

// UserDAO ...
var UserDAO IUser = &userDAO{}
