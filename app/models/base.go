package models

import (
	"github.com/jinzhu/gorm"
)

// Ticket is a sample table for benchmarking
type Ticket struct {
	ID          int    `json:"id" form:"id" gorm:"primary_key"`
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
}

// InitDb initializes database via gorm
func InitDb() *gorm.DB {
	// Openning file
	db, err := gorm.Open("sqlite3", "./focus_dev.db")

	// Display SQL queries
	db.LogMode(true)

	// Error
	if err != nil {
		panic(err)
	}
	// Creating the table
	if !db.HasTable(&Ticket{}) {
		db.CreateTable(&Ticket{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Ticket{})
	}

	return db
}
