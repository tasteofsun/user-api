package handlers

import (
	"net/http"
	"user-api/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var users = make(map[string]models.User)

func GetUsers(c *gin.Context) {
	userList := make([]models.User, 0, len(users))
	for _, user := range users {
		userList = append(userList, user)
	}

	c.JSON(http.StatusOK, gin.H{
		"users": userList,
		"total": len(userList),
	})
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	user, exists := users[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var request models.CreateUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON format",
		})
		return
	}

	if request.Name == "" || request.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Name and email are required",
		})
		return
	}

	user := models.User{
		ID:    uuid.New().String(),
		Name:  request.Name,
		Email: request.Email,
	}

	users[user.ID] = user
	c.JSON(http.StatusCreated, user)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	existingUser, exists := users[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	var request models.CreateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON format",
		})
		return
	}

	existingUser.Name = request.Name
	existingUser.Email = request.Email
	users[id] = existingUser

	c.JSON(http.StatusOK, existingUser)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	_, exists := users[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	delete(users, id)
	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}
