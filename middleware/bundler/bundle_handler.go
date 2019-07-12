package bundler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"opatutorial/models/dao"
	"opatutorial/utils/tarball"
	"os"
	"strings"
)

const (
	userRoleDirectoryFormat       = "static/%s/%s/user/role"
	rolePermissionDirectoryFormat = "static/%s/%s/role/permission"
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

	for serviceGroup, service := range mServices {
		if err = os.MkdirAll(fmt.Sprintf("static/%s/rbac/authz", serviceGroup), 0700); err != nil {
			fmt.Println("could not create service group directory")
			return err
		}

		// Copy file .rego and use gzip tarball file
		regoFilePath := fmt.Sprintf("static/%s/rbac/authz/rbac.rego", serviceGroup)
		nBytes, err := tarball.CopyFile("rbac.rego", regoFilePath)
		if err != nil {
			fmt.Println("could not create rbac.rego file")
			return err
		}
		fmt.Println(nBytes, " bytes copied!")
		for _, serviceInfo := range service {
			if !serviceInfo.Name.Valid || !serviceInfo.ServiceID.Valid || !serviceInfo.URI.Valid {
				continue
			}
			serviceName := strings.ToLower(serviceInfo.Name.String)
			userRoleDirectory := fmt.Sprintf(userRoleDirectoryFormat, serviceGroup, serviceName)
			rolePermissionDirectory := fmt.Sprintf(rolePermissionDirectoryFormat, serviceGroup, serviceName)

			// create user role directory ...
			if err = os.MkdirAll(userRoleDirectory, 0700); err != nil {
				fmt.Println("could not create user/role directory")
				return err
			}

			// create user/role/data.json file
			if err = createUserRoleDataFile(userRoleDirectory, serviceInfo.ServiceID.Int64); err != nil {
				fmt.Println("could not create user/role/data.json file")
				return err
			}

			// // create role permission directory ...
			if err = os.MkdirAll(rolePermissionDirectory, 0700); err != nil {
				fmt.Println("Could not create role/permission directory")
				return err
			}

			// create role/permission/data.json file
			if err = createRolePermissionDataFile(rolePermissionDirectory, serviceInfo.ServiceID.Int64); err != nil {
				fmt.Println("Could not create role/permission/data.json file")
				return err
			}
		}
		// Compress gzip file
		if err = tarball.CompressTarball("static/"+serviceGroup+".tar.gz", "static/"+serviceGroup); err != nil {
			fmt.Println("Could not create bundle file")
			return err
		}
	}

	return nil
}
