package dao

import (
	"opatutorial/models"
)

const (
	qUserSelectAll          = "SELECT * FROM users"
	qUserSelectAllWithRoles = `SELECT user_roles.id, users.id as user_id, users.name as user_name, roles.name as role_name, user_roles.created_at, user_roles.updated_at 
		FROM user_roles
	INNER JOIN users
	ON users.id = user_roles.user_id
	INNER JOIN roles
	ON roles.id = user_roles.role_id
	ORDER BY users.id`
)

// IUser ...
type IUser interface {
	GetAllUser() ([]*models.User, error)
	GetAllUserWithRoles() ([]*models.UserRole, error)
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

func (u *userDAO) GetAllUserWithRoles() ([]*models.UserRole, error) {
	userRoles := []*models.UserRole{}
	db := ConfigurationManager.GetDB()
	if err := db.Select(&userRoles, qUserSelectAllWithRoles); err != nil {
		return nil, err
	}
	return userRoles, nil
}

// UserDAO ...
var UserDAO IUser = &userDAO{}
