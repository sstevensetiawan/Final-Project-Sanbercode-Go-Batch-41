package controllers

import (
	"Final-Project-Sanbercode-Go-Batch-41/database"
	"Final-Project-Sanbercode-Go-Batch-41/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllInvoice(c *gin.Context) {
	var (
		result gin.H
	)

	invoices, err := repository.GetAllInvoice(database.DBConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": invoices,
		}
	}

	c.JSON(http.StatusOK, result)
}

// func GetInvoice(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	var (
// 		result gin.H
// 	)

// 	users, err := repository.GetDetailInvoice(database.DBConnection)

// 	if err != nil {
// 		result = gin.H{
// 			"result": err,
// 		}
// 	} else {
// 		result = gin.H{
// 			"result": users,
// 		}
// 	}

// 	c.JSON(http.StatusOK, result)
// }
