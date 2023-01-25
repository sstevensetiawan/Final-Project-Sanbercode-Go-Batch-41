package controllers

import (
	"Final-Project-Sanbercode-Go-Batch-41/database"
	"Final-Project-Sanbercode-Go-Batch-41/repository"
	"Final-Project-Sanbercode-Go-Batch-41/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCustomer(c *gin.Context) {
	var (
		result gin.H
	)

	users, err := repository.GetAllCustomer(database.DBConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": users,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertCustomer(c *gin.Context) {
	var customer structs.Customer

	err := c.ShouldBindJSON(&customer)
	if err != nil {
		panic(err)
	}

	err = repository.InsertCustomer(database.DBConnection, customer)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Customer",
	})
}

func UpdateCustomer(c *gin.Context) {
	var customer structs.Customer
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	err = c.ShouldBindJSON(&customer)
	if err != nil {
		panic(err)
	}
	customer.ID = uint64(id)

	err = repository.UpdateCustomer(database.DBConnection, customer)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Customer",
	})
}

func DeleteCustomer(c *gin.Context) {
	var customer structs.Customer
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		panic(err)
	}

	customer.ID = uint64(id)

	err = repository.DeleteCustomer(database.DBConnection, customer)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Customer",
	})
}
