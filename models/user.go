package models

import "database/sql"

type (
	User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	Person struct {
		FirstName string `db:"first_name"`
		LastName  string `db:"last_name"`
		Email     string
	}

	Place struct {
		Country string
		City    sql.NullString
		TelCode int
	}
)
