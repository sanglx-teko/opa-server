package models

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func ConnectDatabase() error {
	_db, err := sqlx.Connect("mysql", "root:@(localhost:3306)/test_sqlx?parseTime=true")
	if _db == nil {
		return fmt.Errorf("Connection to DB fail %v", err)
	}
	db = _db
	return nil
}

func MigrateDB() (err error) {
	tx := db.MustBegin()
	tx.NamedExec("INSERT INTO person (first_name, last_name, email) VALUES (:first_name, :last_name, :email)", &Person{"Jason", "Moiron", "jmoiron@jmoiron.net"})
	tx.NamedExec("INSERT INTO person (first_name, last_name, email) VALUES (:first_name, :last_name, :email)", &Person{"John", "Doe", "johndoeDNE@gmail.net"})
	tx.NamedExec("INSERT INTO place (country, city, telcode) VALUES (:country, :city, :telcode)", &Place{"United States", sql.NullString{"New York", true}, 1})
	tx.NamedExec("INSERT INTO place (country, telcode) VALUES (:country, :telcode)", &Place{Country: "Hong Kong", TelCode: 852})
	tx.NamedExec("INSERT INTO place (country, telcode) VALUES (:country, :telcode)", &Place{Country: "Singapore", TelCode: 65})
	tx.NamedExec("INSERT INTO person (first_name, last_name, email) VALUES (:first_name, :last_name, :email)", &Person{"Jane", "Citizen", "jane.citzen@example.com"})
	tx.Commit()
	return nil
}
