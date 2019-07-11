package models

import (
	"time"
)

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
		ID              int
		ServiceInfo     string      `db:"service_info"`
		ServiceGroupID  int         `db:"service_group_id"`
		ServiceMetaData interface{} `db:"service_metadata"`
		CreatedAt       time.Time
		UpdatedAt       time.Time
	}
	// ServiceGroup ...
	ServiceGroup struct {
		ID        int
		Name      string
		URI       string `db:"uri"`
		CreatedAt time.Time
		UpdatedAt time.Time
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
