package test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sanglx-teko/opa-server/middleware/bundler"
	manager "github.com/sanglx-teko/opa-server/middleware/configurationmanager"
	"github.com/sanglx-teko/opa-server/models/dao"
	"github.com/stretchr/testify/assert"
)

// Create DB Connection
func Initialize() {
	if err := manager.Instance.ConnectDB(os.Getenv("SQL_DIALECT"), os.Getenv("SQL_DSN")); err != nil {
		fmt.Println(err)
		panic(err)
	}

	if err := os.MkdirAll("static", 0700); err != nil {
		panic(err)
	}
	bundler.InitCFManager(manager.Instance)
	dao.InitCFManager(manager.Instance)
}

// Create DB Connection before testing
func TestMain(m *testing.M) {
	Initialize()
	setDB(dao.ConfigurationManager.GetDB())
	code := m.Run()
	os.Exit(code)
}

// Test godotenv.Load()
func TestLoadEnv(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("mysql", os.Getenv("SQL_DIALECT"))
	assert.Equal("root:123@tcp(localhost:3306)/opadb?parseTime=true&charset=utf8", os.Getenv("SQL_DSN"))
	assert.Equal(":3000", os.Getenv("PORT"))
}

// Test configurationManager.ConnectDB()
func TestConnectDB(t *testing.T) {
	assert := assert.New(t)
	err := manager.Instance.ConnectDB(os.Getenv("SQL_DIALECT"), os.Getenv("SQL_DSN"))
	assert.Nil(err, "Must connect to database")
}

// Test RoleDAO.GetAllRoles()
func TestGetAllRoles(t *testing.T) {
	assert := assert.New(t)
	InitRoleTable()

	roles, err := dao.RoleDAO.GetAllRoles()
	assert.Nil(err, "Must be able to get all roles")
	assert.Equal(4, len(roles), "Example data must have 4 roles")

	xRoles := []string{"developer", "administrator", "user", "super admin"}
	for i, role := range roles {
		assert.Equal(xRoles[i], role.Name)
	}
	DropAllDB()
}

// Test RoleDAO.GetAllRolesWithPermission()
func TestGetAllRolesWithPermission(t *testing.T) {
	assert := assert.New(t)
	InitRoleTable()
	InitActionTable()
	InitResourceTable()
	InitPermissionTable()
	InitRolePermissionTable()

	rolePermissions, err := dao.RoleDAO.GetAllRolesWithPermission()
	assert.Nil(err, "Must be able to get all role_permissions")
	assert.Equal(2, len(rolePermissions), "Example data must have 2 role_permissions")

	xRolePermission := make([][]string, 2)
	xRolePermission[0] = []string{"developer", "write", "permissions"}
	xRolePermission[1] = []string{"user", "read", "users"}
	for i, rolePermission := range rolePermissions {
		assert.Equal(xRolePermission[i][0], rolePermission.RoleName)
		assert.Equal(xRolePermission[i][1], rolePermission.Action)
		assert.Equal(xRolePermission[i][2], rolePermission.Resource)
	}
	DropAllDB()
}

// Test RoleDAO.GetAllRolePermissionWithServiceID()
func TestGetAllRolePermissionWithServiceID(t *testing.T) {
	assert := assert.New(t)
	InitRoleTable()
	InitActionTable()
	InitResourceTable()
	InitPermissionTable()
	InitRolePermissionTable()
	// rolePermissions, err := dao.RoleDAO.GetAllRolePermissionWithServiceID(1)
	var serviceID int64 = 2
	rolePermissions, err := dao.RoleDAO.GetAllRolePermissionWithServiceID(serviceID)
	assert.Nil(err, fmt.Sprintf("Must be able to get all role_permissions in service id %v", serviceID))
	assert.Equal(0, len(rolePermissions), "must have 0 role_permissions")
	DropAllDB()
}

// Test ServiceDAO.GetAllService()
func TestGetAllServices(t *testing.T) {
	assert := assert.New(t)
	InitServiceTable()

	services, err := dao.ServiceDAO.GetAllService()
	assert.Nil(err, "Must be able to get all services")
	assert.Equal(2, len(services), "Must have 2 services")

	xServices := make([][]string, 2)
	xServices[0] = []string{"Identity API Platform", "1"}
	xServices[1] = []string{"Payment API", "1"}
	for i, service := range services {
		assert.Equal(xServices[i][0], service.ServiceInfo)
		assert.Equal(xServices[i][1], fmt.Sprintf("%v", service.ServiceGroupID))
	}
	DropAllDB()
}

// Test ServiceDAO.GetAllServiceWithServiceGroupNameAndURL()
func TestGetAllServiceWithServiceGroupNameAndURL(t *testing.T) {
	assert := assert.New(t)
	InitServiceTable()
	InitServiceGroupTable()

	services, err := dao.ServiceDAO.GetAllServiceWithServiceGroupNameAndURL()
	assert.Nil(err, "Must be able to get all services with serviceGroup name and URI")
	assert.Equal(3, len(services), "Must have 2 serviceModels")

	xServices := make([][]string, 2)
	xServices[0] = []string{"1", "1", "core_services", "IAM", "http://localhost:8181/v1/data/rbac/authz/allow"}
	xServices[1] = []string{"1", "2", "core_services", "PMAPI", "http://localhost:8181/v1/data/rbac/authz/allow"}
	for i, service := range services[:1] {
		assert.Equal(xServices[i][0], fmt.Sprintf("%v", service.ID))
		assert.Equal(xServices[i][1], fmt.Sprintf("%v", service.ServiceID.Int64))
		assert.Equal(xServices[i][2], service.ServiceGroupName)
		assert.Equal(xServices[i][3], service.Name.String)
		assert.Equal(xServices[i][4], service.URI.String)
	}
	DropAllDB()
}

// Test ServiceGroupDAO.GetAllServiceGroup()
func TestGetAllServiceGroup(t *testing.T) {
	assert := assert.New(t)
	InitServiceGroupTable()

	serviceGroups, err := dao.ServiceGroupDAO.GetAllServiceGroup()
	assert.Nil(err, "Must be able to get all service groups")
	assert.Equal(2, len(serviceGroups), "Must have 2 serviceGroups")

	xSerivceGroups := make([][]string, 2)
	xSerivceGroups[0] = []string{"1", "core_services", "http://localhost:8181/v1/data/rbac/authz/allow"}
	xSerivceGroups[1] = []string{"2", "online_sales", "http://localhost:8181/v1/data/rbac/authz/allow"}
	for i, serviceGroup := range serviceGroups {
		assert.Equal(xSerivceGroups[i][0], fmt.Sprintf("%v", serviceGroup.ID))
		assert.Equal(xSerivceGroups[i][1], serviceGroup.Name)
		assert.Equal(xSerivceGroups[i][2], serviceGroup.URI)
	}
	DropAllDB()
}

// Test UserDAO.GetAllUser()
func TestGetAllUser(t *testing.T) {
	assert := assert.New(t)
	InitUserTable()

	users, err := dao.UserDAO.GetAllUser()
	assert.Nil(err, "Must be able to get all users")
	assert.Equal(2, len(users), "Must have 2 users")

	xUsers := make([][]string, 2)
	xUsers[0] = []string{"U01", "Ly Xuan Sang", "sang.lxuan@gmail.com", "+84347942877", "HaNoi", "1991-04-02"}
	xUsers[1] = []string{"U02", "Le Hai Nam", "lehainam.dev@gmail.com", "", "Quang Ninh", "1995-04-02"}
	for i, user := range users {
		assert.Equal(xUsers[i][0], user.ID)
		assert.Equal(xUsers[i][1], user.Name)
		assert.Equal(xUsers[i][2], user.Email)
		assert.Equal(xUsers[i][3], user.PhoneNumber)
		assert.Equal(xUsers[i][4], user.Address)
		assert.Equal(xUsers[i][5], strings.Split(user.Birthday.String(), " ")[0])
	}
}

// Test UserDAO.GetALlUserWithRoles()
func TestGetAllUserWithRoles(t *testing.T) {
	assert := assert.New(t)
	InitUserTable()
	InitRoleTable()
	InitUserRoleTable()

	users, err := dao.UserDAO.GetAllUserWithRoles()
	assert.Nil(err, "Must eb able to get all users with roles")
	assert.Equal(2, len(users), "Must have 2 users with roles")

	xUsers := make([][]string, 2)
	xUsers[0] = []string{"1", "U01", "Ly Xuan Sang", "developer"}
	xUsers[1] = []string{"2", "U02", "Le Hai Nam", "user"}
	for i, user := range users {
		assert.Equal(xUsers[i][0], user.ID)
		assert.Equal(xUsers[i][1], user.UserID)
		assert.Equal(xUsers[i][2], user.UserName)
		assert.Equal(xUsers[i][3], user.RoleName)
	}
}

// func TestCreateBundleFile(t *testing.T) {
// 	_, err := dao.ServiceDAO.GetAllServiceWithServiceGroupNameAndURL()
// 	assert.Nil(t, err)
// }
