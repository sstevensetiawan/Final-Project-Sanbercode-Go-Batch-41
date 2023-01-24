package controllers

import (
	"Final-Project-Sanbercode-Go-Batch-41/database"
	"Final-Project-Sanbercode-Go-Batch-41/repository"
	"Final-Project-Sanbercode-Go-Batch-41/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllProduct(c *gin.Context) {
	var (
		result gin.H
	)

	users, err := repository.GetAllProduct(database.DBConnection)

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

func InsertProduct(c *gin.Context) {
	var product structs.Product

	err := c.ShouldBindJSON(&product)
	if err != nil {
		panic(err)
	}

	err = repository.InsertProduct(database.DBConnection, product)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Product",
	})
}

func UpdateProduct(c *gin.Context) {
	var product structs.Product
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	err = c.ShouldBindJSON(&product)
	if err != nil {
		panic(err)
	}
	product.ID = uint64(id)

	err = repository.UpdateProduct(database.DBConnection, product)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Product",
	})
}

func DeleteProduct(c *gin.Context) {
	var product structs.Product
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		panic(err)
	}

	product.ID = uint64(id)
	product.Status = uint8(0)

	err = repository.DeleteProduct(database.DBConnection, product)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Product",
	})
}
