package main

import (
	"Final-Project-Sanbercode-Go-Batch-41/controllers"
	"Final-Project-Sanbercode-Go-Batch-41/database"
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
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

	database.DbConnect(DB)

	defer DB.Close()

	router := gin.Default()
	router.GET("/user", controllers.GetAllUser)
	router.POST("/user", controllers.InsertUser)
	router.PUT("/user/:id", controllers.UpdateUser)
	router.DELETE("/user/:id", controllers.DeleteUser)

	router.Run("localhost:8080")
}
