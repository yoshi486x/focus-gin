package main

import (
	"net/http"
	"news/app/views"
	"news/app/models"

	"github.com/gin-gonic/gin"
)

// Set env before execution. `export PORT=9090`
func main() {
	r := gin.Default()
	
	models.ConnectDB()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World! I'm Gin!")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/tickets", views.FindTickets)
	v1 := r.Group("api/v1")
	{
		v1.GET("/tickets", GetTickets)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080"
}

