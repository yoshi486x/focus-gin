package models

import (
	// "focus-gin/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Ticket is a sample table for benchmarking
type Ticket struct {
	// gorm.Model
	ID          int    		`json:"id" form:"id" gorm:"primary_key"`
	Title       string 		`json:"title" form:"title"`
	UpdatedAt 	time.Time 	`json:"updated_at" form:"updated_at"`
	DeletedAt 	time.Time	`json:"deleted_at" form:"deleted_at"`
	CreatedAt 	time.Time 	`json:"created_at" form:"created_at"`
}

// InitDb initializes database via gorm
func InitDb() *gorm.DB {
	// Openning file
	// dsn := "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	dsn := "root:yFJw4b74nts8ybLo@tcp(10.85.64.3:3306)/focus_dev?charset=utf8&parseTime=True&loc=Local"
	// dsn := "root:yFJw4b74nts8ybLo@tcp(34.84.108.14:3306)/focus_dev?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Ticket{})

	return db
}
