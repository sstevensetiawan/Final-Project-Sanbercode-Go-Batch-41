package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	errLoadEnv = godotenv.Load("config/.env")
	host       = os.Getenv("DB_HOST")
	port       = os.Getenv("DB_PORT")
	user       = os.Getenv("DB_USER")
	password   = os.Getenv("DB_PASSWORD")
	dbname     = os.Getenv("DB_NAME")
)

var (
	DB    *sql.DB
	dbErr error
)

func main() {
	if errLoadEnv != nil {
		fmt.Println("Failed load env")
		panic(errLoadEnv)
	}

	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		user, password, host, port, dbname)
	DB, dbErr = sql.Open("mysql", sqlInfo)
	if dbErr != nil {
		fmt.Println("Fail Connect to Database")
		panic(dbErr.Error())
	}

	defer DB.Close()
}
