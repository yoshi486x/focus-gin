package models

import(
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Tickets struct {
	ID			int 	`gorm:"AUTO_INCREMENT" json:"id" form: "id`
	Title 		string 	`json:"title" form: "title"`
	Description	string 	`json:"description" form: "description"`
}

// var DB *gorm.DB

// func ConnectDB() {
// 	var err error

// 	db, err := gorm.Open("sqlite3", "focus_dev")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	db.AutoMigrate(&Ticket{})

// 	DB = db
// }

func InitDb() *gorm.DB {
	// Openning file
	db, err := gorm.Open("sqlite3", "./data.db")
	// Display SQL queries
	db.LogMode(true)

	// Error
	if err != nil {
		panic(err)
	}
	// Creating the table
	if !db.HasTable(&Tickets{}) {
		db.CreateTable(&Tickets{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Tickets{})
	}

	return db
}