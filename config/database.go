package config

import (
	"bwa-startup/helpers"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	dbUser := helpers.GetKeyValue("DB_USER", "root")
	dbPass := helpers.GetKeyValue("DB_PASS", "password")
	dbHost := helpers.GetKeyValue("DB_HOST", "localhost")
	dbPort := helpers.GetKeyValue("DB_PORT", "3306")
	dbSchema := helpers.GetKeyValue("DB_SCHEMA", "schemaName")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbSchema)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connection Database Success!")
	return db
}