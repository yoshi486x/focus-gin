package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

// Database is ...
type Database struct {
	DBName			string
	Driver			string
	Addr			string
	Passwd			string
	User			string
}

var DBConfig Database

func init(){
	cfg, err := ini.Load("config.ini")
	if err!= nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	DBConfig = Database{
		DBName:			cfg.Section("db").Key("DBName").String(),
		Driver:     	cfg.Section("db").Key("driver").String(),
		Addr:     		cfg.Section("db").Key("Addr").String(),
		Passwd:     	cfg.Section("db").Key("passwd").String(),
		User:     		cfg.Section("db").Key("user").String(),
	}	
}
