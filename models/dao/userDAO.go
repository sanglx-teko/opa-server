package dao

import (
	"fmt"
	"github.com/sanglx-teko/opa-server/models"
)

var (
	qUserSelectAll          = "SELECT * FROM opa_users" + models.TableOPAUser
	qUserSelectAllWithRoles = fmt.Sprintf(`SELECT %s.id, %s.id as user_id, %s.name as user_name, %s.name as role_name, %s.created_at, %s.updated_at 
		FROM %s
	INNER JOIN %s
	ON %s.id = %s.user_id
	INNER JOIN %s
	ON %s.id = %s.role_id
	ORDER BY %s.id`, models.TableOPARUserRole, models.TableOPAUser, models.TableOPAUser, models.TableOPARole,
		models.TableOPARUserRole, models.TableOPARUserRole, models.TableOPARUserRole, models.TableOPAUser, models.TableOPAUser, models.TableOPARUserRole,
		models.TableOPARole, models.TableOPARole, models.TableOPARUserRole, models.TableOPAUser)
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
