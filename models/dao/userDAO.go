package dao

type IUser interface {
	GetUserFromID() int
}

type userDAO struct {
}

func (u userDAO) GetUserFromID() int {
	return 1
}

var UserDAO IUser = &userDAO{}
