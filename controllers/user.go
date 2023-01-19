package controllers

import (
	"Final-Project-Sanbercode-Go-Batch-41/database"
	"Final-Project-Sanbercode-Go-Batch-41/repository"
	"Final-Project-Sanbercode-Go-Batch-41/structs"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context) {
	var (
		result gin.H
	)

	users, err := repository.GetAllUser(database.DBConnection)

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

func InsertUser(c *gin.Context) {
	var user structs.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
	err = repository.InsertUser(database.DBConnection, user)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert User",
	})
}

func UpdateUser(c *gin.Context) {
	var user structs.User
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	err = c.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}
	user.ID = uint64(id)

	err = repository.UpdateUser(database.DBConnection, user)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update User",
	})
}

func DeleteUser(c *gin.Context) {
	var user structs.User
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		panic(err)
	}

	user.ID = uint64(id)
	user.Status = uint8(0)

	err = repository.DeleteUser(database.DBConnection, user)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete User",
	})
}
