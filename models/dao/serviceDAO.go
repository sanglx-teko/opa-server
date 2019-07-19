package dao

import (
	"fmt"
	"github.com/sanglx-teko/opa-server/models"
)

var (
	qServiceSelectAll                           = "SELECT * FROM " + models.TableOPAService
	qServiceSelectAllWithServiceGroupNameAndURL = fmt.Sprintf(`
	SELECT
	%s.id,
	%s.id as service_id,
	%s.name AS service_group_name,
	%s.uri,
	%s.name,
	%s.created_at,
	%s.updated_at
FROM
	%s
	LEFT JOIN %s ON %s.id = %s.service_group_id
ORDER BY
	%s.id`, models.TableOPAServiceGroup, models.TableOPAService, models.TableOPAServiceGroup, models.TableOPAServiceGroup, models.TableOPAService, models.TableOPAService, models.TableOPAService, models.TableOPAServiceGroup, models.TableOPAService,
		models.TableOPAServiceGroup, models.TableOPAService, models.TableOPAServiceGroup)
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
