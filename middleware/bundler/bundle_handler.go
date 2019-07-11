package bundler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"opatutorial/models/dao"
	"os"
	"strings"
)

const (
	userRoleDirectoryFormat       = "static/%s/user/role"
	rolePermissionDirectoryFormat = "static/%s/role/permission"
)

func createUserRoleDataFile(path string, serviceID int64) error {
	userRoles, err := dao.UserDAO.GetAllUserWithRoles()
	if err != nil {
		return err
	}
	m := make(map[string][]string)
	for _, userRole := range userRoles {
		if _, ok := m[userRole.UserName]; !ok {
			m[userRole.UserName] = []string{userRole.RoleName}
		} else {
			m[userRole.UserName] = append(m[userRole.UserName], userRole.RoleName)
		}
	}
	if len(m) > 0 {
		filePath := fmt.Sprintf("%s/data.json", path)
		jsonObject, err := json.Marshal(m)
		if err != nil {
			fmt.Println(err)
			return err
		}

		f, err := os.Create(filePath)
		defer f.Close()
		if err != nil {
			fmt.Println(err)
			return err
		}
		_, err = f.Write(jsonObject)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

func createRolePermissionDataFile(path string, serviceID int64) error {
	rolePermissions, err := dao.RoleDAO.GetAllRolePermissionWithServiceID(serviceID)
	if err != nil {
		return err
	}
	type tempStruct struct {
		Action string `json:"action"`
		Object string `json:"object"`
	}

	m := make(map[string][]*tempStruct)
	for _, rolePermission := range rolePermissions {
		t := &tempStruct{
			Action: rolePermission.Action,
			Object: rolePermission.Resource,
		}
		if _, ok := m[rolePermission.RoleName]; !ok {
			m[rolePermission.RoleName] = []*tempStruct{t}
		} else {
			m[rolePermission.RoleName] = append(m[rolePermission.RoleName], t)
		}
	}
	if len(m) > 0 {
		filePath := fmt.Sprintf("%s/data.json", path)
		jsonObject, err := json.Marshal(m)
		if err != nil {
			return err
		}

		f, err := os.Create(filePath)
		defer f.Close()
		if err != nil {
			return err
		}
		_, err = f.Write(jsonObject)
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateBundleFile ...
func CreateBundleFile() (err error) {
	// Get all service and service_group_name
	services, err := dao.ServiceDAO.GetAllServiceWithServiceGroupNameAndURL()
	if err != nil {
		return err
	}
	type tempServiceInfo struct {
		URI       sql.NullString `json:"uri"`
		Name      sql.NullString `json:"name"`
		ServiceID sql.NullInt64  `json:"service_id"`
	}
	mServices := make(map[string][]*tempServiceInfo)
	for _, service := range services {
		serviceInfo := &tempServiceInfo{
			URI:       service.URI,
			Name:      service.Name,
			ServiceID: service.ServiceID,
		}
		if _, ok := mServices[service.ServiceGroupName]; !ok {
			mServices[service.ServiceGroupName] = []*tempServiceInfo{serviceInfo}
		} else {
			mServices[service.ServiceGroupName] = append(mServices[service.ServiceGroupName], serviceInfo)
		}
	}

	for k, service := range mServices {

		serviceName := strings.ToLower(service.Name.String)
		userRoleDirectory := fmt.Sprintf(userRoleDirectoryFormat, serviceName)
		rolePermissionDirectory := fmt.Sprintf(rolePermissionDirectoryFormat, serviceName)
		// create user role directory
		if err = os.MkdirAll(userRoleDirectory, 0700); err != nil {
			fmt.Println("could not create user/role directory")
			return err
		}
		// create role permission directory
		if err = os.MkdirAll(rolePermissionDirectory, 0700); err != nil {
			fmt.Println("could not create role/permission directory")
			return err
		}

		createUserRoleDataFile(userRoleDirectory, service.ServiceID.Int64)
		createRolePermissionDataFile(rolePermissionDirectory, service.ServiceID.Int64)
	}

	return nil
}
