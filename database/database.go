package database

import (
	"database/sql"
)

var (
	DBConnection *sql.DB
)

func DbConnect(dbParam *sql.DB) {
	DBConnection = dbParam
}
