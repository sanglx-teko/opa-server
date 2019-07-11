package configurationmanager

import (
	"fmt"
	"opatutorial/middleware/configurationmanager/model"

	"github.com/jmoiron/sqlx"
)

type configurationManager struct {
	db *sqlx.DB
}

func (c *configurationManager) GetDB() (result *sqlx.DB) {
	result = c.db
	return
}

// ConnectDB with connectionstring: mysql, root:@(localhost:3306)/test_sqlx?parseTime=true
func (c *configurationManager) ConnectDB(dbDialect, connectionString string) (erro error) {
	_db, err := sqlx.Connect(dbDialect, connectionString)
	if _db == nil {
		erro = fmt.Errorf("Connection to DB fail %v", err)
		return
	}
	c.db = _db
	return nil
}

// Instance of ConfigurationManager ...
var Instance model.IConfigurationManager = &configurationManager{}
