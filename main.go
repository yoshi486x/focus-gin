package main

import (
	"net/http"
	"focus-gin/app/views"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// Set env before execution. `export PORT=9090`
func main() {
	r := gin.Default()
	
	// models.ConnectDB()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World! I'm Gin!")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	v1 := r.Group("api/v1")
	{
		v1.GET("/tickets", views.GetTickets)
	}

	r.Run(":9090")
}

