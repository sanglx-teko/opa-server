package models

import (
	"database/sql"
	"database/sql/driver"
	"time"
)

const (
	// TableOPAUser ...
	TableOPAUser = "opa_users"
	// TableOPAActions ...
	TableOPAActions = "opa_actions"
	// TableOPARole ...
	TableOPARole = "opa_roles"
	// TableOPAPermission ...
	TableOPAPermission = "opa_permissions"
	// TableOPAService         = "opa_services"
	TableOPAService = "opa_services"
	// TableOPAServiceGroup    = "opa_service_groups"
	TableOPAServiceGroup = "opa_service_groups"
	// TableOPAResource        = "opa_resources"
	TableOPAResource = "opa_resources"
	// TableOPARolePermissions = "opa_role_permissions"
	TableOPARolePermissions = "opa_role_permissions"
	// TableOPARUserRole       = "opa_user_roles"
	TableOPARUserRole = "opa_user_roles"
)

// NullTime ...
type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

// Scan implements the Scanner interface.
func (nt *NullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

// Value implements the driver Valuer interface.
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

type (
	// User ...
	User struct {
		ID          string    `db:"id" json:"id"`
		Name        string    `json:"name"`
		Email       string    `json:"email"`
		PhoneNumber string    `db:"phone_number" json:"phoneNumber"`
		Address     string    `json:"address"`
		Description string    `json:"description"`
		Birthday    time.Time `json:"birthday"`
		CreatedAt   time.Time `db:"created_at" json:"createdAt"`
		UpdatedAt   time.Time `db:"updated_at" json:"updatedAt"`
	}

	// Role ...
	Role struct {
		ID        int       `db:"id"`
		Name      string    `db:"name"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
	}

	// RolePermission ...
	RolePermission struct {
		RoleName string `db:"role_name"`
		Action   string `db:"action_name"`
		Resource string `db:"resource_name"`
	}
	// UserRole ...
	UserRole struct {
		ID        string    `db:"id"`
		UserID    string    `db:"user_id"`
		UserName  string    `db:"user_name"`
		RoleName  string    `db:"role_name"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
	}
	// Action ...
	Action struct {
		ID          int
		Name        string
		Description string
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
	// Service ...
	Service struct {
		ID              int         `db:"id"`
		ServiceInfo     string      `db:"service_info"`
		ServiceGroupID  int         `db:"service_group_id"`
		ServiceMetaData interface{} `db:"service_metadata"`
		CreatedAt       time.Time   `db:"created_at" json:"createdAt"`
		UpdatedAt       time.Time   `db:"updated_at" json:"updatedAt"`
	}

	// ServiceModel ...
	ServiceModel struct {
		ID               int            `db:"id"`
		ServiceID        sql.NullInt64  `db:"service_id"`
		ServiceGroupName string         `db:"service_group_name"`
		Name             sql.NullString `db:"name"`
		URI              sql.NullString `db:"uri"`
		CreatedAt        NullTime       `db:"created_at" json:"createdAt"`
		UpdatedAt        NullTime       `db:"updated_at" json:"updatedAt"`
	}

	// ServiceGroup ...
	ServiceGroup struct {
		ID        int       `db:"id"`
		Name      string    `db:"name"`
		URI       string    `db:"uri"`
		CreatedAt time.Time `db:"created_at" json:"createdAt"`
		UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
	}
	// Resource ...
	Resource struct {
		ID        int
		Name      string
		ServiceID int `db:"service_id"`
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
