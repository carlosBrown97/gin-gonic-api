package controllers

import (
	"gin-gonic-api/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userPostRequest struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

func UsersGet(newUser *models.User) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := GetAllUsers()
		c.JSON(http.StatusOK, results)
	}
}

func UsersPost(newUser *models.User) gin.HandlerFunc {
	return func(c *gin.Context) {

		reqBody := userPostRequest{}
		c.Bind(&reqBody)

		user := models.User{
			Name:     reqBody.Name,
			LastName: reqBody.LastName,
		}

		AddUser(user)

		c.Status(http.StatusNoContent)
	}
}

func NewUser() *models.User {
	return &models.User{}
}

func AddUser(user models.User) {
	var r []models.User
	r = append(r, user) // (all users, new user)
}

func GetAllUsers() []models.User {
	var r []models.User
	return r
}
