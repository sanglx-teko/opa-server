package dao

import (
	"fmt"
	"opatutorial/models"
)

var (
	qRoleSelectAllWithPermissions = fmt.Sprintf(`
	select %s.name as action_name, %s.name as resource_name, %s.name as role_name 
	from %s
	inner join %s
	on %s.id = %s.role_id
	inner join %s
	on %s.id = %s.permission_id
	inner join %s
	on %s.id = %s.action_id
	inner join %s
	on %s.id = %s.resource_id
	order by %s.name`, models.TableOPAActions, models.TableOPAResource, models.TableOPARole, models.TableOPARolePermissions,
		models.TableOPARole, models.TableOPARole, models.TableOPARolePermissions, models.TableOPAPermission, models.TableOPAPermission,
		models.TableOPARolePermissions, models.TableOPAActions, models.TableOPAActions, models.TableOPAPermission, models.TableOPAResource,
		models.TableOPAResource, models.TableOPAPermission, models.TableOPARole)
	qRoleSelectAll                      = "select * from " + models.TableOPARole
	qRoleGetRolePermissionWithServiceID = fmt.Sprintf(`
	SELECT
		%s.name AS action_name,
		%s.name AS resource_name,
		%s.name AS role_name
	FROM
		%s
	INNER JOIN %s ON %s.id = %s.role_id
	INNER JOIN %s ON %s.id = %s.permission_id
	INNER JOIN %s ON %s.id = %s.action_id
	INNER JOIN %s ON %s.id = %s.resource_id
	WHERE
		%s.service_id = ?
	ORDER BY
		%s.name`, models.TableOPAActions, models.TableOPAResource, models.TableOPARole, models.TableOPARolePermissions, models.TableOPARole, models.TableOPARole,
		models.TableOPARolePermissions, models.TableOPAPermission, models.TableOPAPermission, models.TableOPARolePermissions, models.TableOPAActions, models.TableOPAActions, models.TableOPAPermission,
		models.TableOPAResource, models.TableOPAResource, models.TableOPAPermission, models.TableOPAResource, models.TableOPARole)
)

// IRole ...
type IRole interface {
	GetAllRoles() ([]*models.Role, error)
	GetAllRolesWithPermission() ([]*models.RolePermission, error)
	GetAllRolePermissionWithServiceID(serviceID int64) ([]*models.RolePermission, error)
}

type roleDAO struct {
}

func (u *roleDAO) GetAllRoles() ([]*models.Role, error) {
	db := ConfigurationManager.GetDB()
	roles := []*models.Role{}
	if err := db.Select(&roles, qRoleSelectAll); err != nil {
		return nil, err
	}
	return roles, nil
}

func (u *roleDAO) GetAllRolesWithPermission() ([]*models.RolePermission, error) {
	db := ConfigurationManager.GetDB()
	rolePermissions := []*models.RolePermission{}
	if err := db.Select(&rolePermissions, qRoleSelectAllWithPermissions); err != nil {
		return nil, err
	}
	return rolePermissions, nil
}

func (u *roleDAO) GetAllRolePermissionWithServiceID(serviceID int64) ([]*models.RolePermission, error) {
	db := ConfigurationManager.GetDB()
	rolePermissions := []*models.RolePermission{}
	// fmt.Println("qRoleGetRolePermissionWithServiceID", qRoleGetRolePermissionWithServiceID)
	if err := db.Select(&rolePermissions, qRoleGetRolePermissionWithServiceID, serviceID); err != nil {
		return nil, err
	}
	return rolePermissions, nil
}

// RoleDAO ...
var RoleDAO IRole = &roleDAO{}
