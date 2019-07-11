package dao

import (
	"opatutorial/models"
)

const (
	qServiceGroupSelectAll = "SELECT * FROM service_groups"
)

// IServiceGroup ...
type IServiceGroup interface {
	GetAllServiceGroup() ([]*models.ServiceGroup, error)
}

type serviceGroupDAO struct {
}

func (u *serviceGroupDAO) GetAllServiceGroup() ([]*models.ServiceGroup, error) {
	db := ConfigurationManager.GetDB()
	serviceGroups := []*models.ServiceGroup{}
	if err := db.Select(&serviceGroups, qServiceGroupSelectAll); err != nil {
		return nil, err
	}
	return serviceGroups, nil
}

// ServiceGroupDAO ...
var ServiceGroupDAO IServiceGroup = &serviceGroupDAO{}
