package model

import "github.com/jmoiron/sqlx"

// IConfigurationManager ...
type IConfigurationManager interface {
	GetDB() *sqlx.DB
	ConnectDB(dbDialect string, connectionString string) error
}
