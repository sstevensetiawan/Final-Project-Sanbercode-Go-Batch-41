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

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uname, pwd, ok := c.Request.BasicAuth()
		if !ok {
			c.Writer.Write([]byte("Username atau Password tidak boleh kosong"))
			c.Abort()
			return
		}

		if uname != "admin" && pwd != "admin" {
			c.Writer.Write([]byte("Username atau Password salah"))
			c.Abort()
			return
		}
	}
}

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
	// customer
	router.GET("/customer", controllers.GetAllCustomer)
	router.POST("/customer", AuthMiddleware(), controllers.InsertCustomer)
	router.PUT("/customer/:id", AuthMiddleware(), controllers.UpdateCustomer)
	router.DELETE("/customer/:id", AuthMiddleware(), controllers.DeleteCustomer)

	// product
	router.GET("/product", controllers.GetAllProduct)
	router.POST("/product", AuthMiddleware(), controllers.InsertProduct)
	router.PUT("/product/:id", AuthMiddleware(), controllers.UpdateProduct)
	router.DELETE("/product/:id", AuthMiddleware(), controllers.DeleteProduct)

	// invoice
	router.GET("/invoice", controllers.GetAllInvoice)
	router.GET("/invoice/:invoice_id", controllers.GetInvoice)
	router.POST("invoice/", AuthMiddleware(), controllers.InsertInvoice)
	router.DELETE("/invoice/:invoice_id", AuthMiddleware(), controllers.DeleteInvoice)

	router.Run("localhost:8080")
}
