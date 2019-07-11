package dao

import (
	"fmt"
	"opatutorial/models"
)

const (
	qServiceSelectAll                           = "SELECT * FROM services"
	qServiceSelectAllWithServiceGroupNameAndURL = `
	SELECT
	service_groups.id,
	services.id as service_id,
	service_groups.name AS service_group_name,
	service_groups.uri,
	services.name,
	services.created_at,
	services.updated_at
FROM
	service_groups
	LEFT JOIN services ON service_groups.id = services.service_group_id
ORDER BY
	service_groups.id`
)

// IService ...
type IService interface {
	GetAllService() ([]*models.Service, error)
	GetAllServiceWithServiceGroupNameAndURL() ([]*models.ServiceModel, error)
}

type serviceDAO struct {
}

func (u *serviceDAO) GetAllServiceWithServiceGroupNameAndURL() ([]*models.ServiceModel, error) {
	db := ConfigurationManager.GetDB()
	services := []*models.ServiceModel{}
	if err := db.Select(&services, qServiceSelectAllWithServiceGroupNameAndURL); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return services, nil
}

func (u *serviceDAO) GetAllService() ([]*models.Service, error) {
	db := ConfigurationManager.GetDB()
	services := []*models.Service{}
	if err := db.Select(&services, qServiceSelectAll); err != nil {
		return nil, err
	}
	return services, nil
}

// ServiceDAO ...
var ServiceDAO IService = &serviceDAO{}
