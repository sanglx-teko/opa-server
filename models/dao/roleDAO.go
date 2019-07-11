package dao

import (
	"fmt"
	"opatutorial/models"
)

const (
	qRoleSelectAllWithPermissions = `
	select actions.name as action_name, resources.name as resource_name, roles.name as role_name 
	from role_permissions
	inner join roles
	on roles.id = role_permissions.role_id
	inner join permissions
	on permissions.id = role_permissions.permission_id
	inner join actions
	on actions.id = permissions.action_id
	inner join resources
	on resources.id = permissions.resource_id
	order by roles.name`
	qRoleSelectAll                      = "select * from roles"
	qRoleGetRolePermissionWithServiceID = `
	SELECT
		actions.name AS action_name,
		resources.name AS resource_name,
		roles.name AS role_name
	FROM
		role_permissions
	INNER JOIN roles ON roles.id = role_permissions.role_id
	INNER JOIN permissions ON permissions.id = role_permissions.permission_id
	INNER JOIN actions ON actions.id = permissions.action_id
	INNER JOIN resources ON resources.id = permissions.resource_id
	WHERE
		resources.service_id = ?
	ORDER BY
		roles.name`
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
		fmt.Println(err)
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
	fmt.Println("service id:", serviceID)
	db := ConfigurationManager.GetDB()
	rolePermissions := []*models.RolePermission{}
	if err := db.Select(&rolePermissions, qRoleGetRolePermissionWithServiceID, serviceID); err != nil {
		return nil, err
	}
	return rolePermissions, nil
}

// RoleDAO ...
var RoleDAO IRole = &roleDAO{}
