package views

import (
	"net/http"
	"news/app/models"

	"github.com/gin-gonic/gin"
)

// FindTickets is ... GET /tickets
func FindTickets(c *gin.Context) {
	var tickets []models.Ticket
	models.DB.Find(&tickets)

	c.JSON(http.StatusOK, gin.H{"data": tickets})
}

func GetTickets(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	var tickets []Tickets
	// SELECT * FROM users
	db.Find(&tickets)

	// Display JSON result
	c.JSON(200, tickets)

	// curl -i http://localhost:8080/api/v1/users
}
