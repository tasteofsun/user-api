package main

import (
	"user-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "OK",
			"service": "user-api",
		})
	})

	api := r.Group("/api/users")
	{
		api.GET("", handlers.GetUsers)
		api.GET("/:id", handlers.GetUserByID)
		api.POST("", handlers.CreateUser)
		api.PUT("/:id", handlers.UpdateUser)
		api.DELETE("/:id", handlers.DeleteUser)
	}

	r.Run(":8080")
}
