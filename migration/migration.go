package main

import (
	"fmt"
	"go-crud/daos/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitialMigration() {
	DB, err = gorm.Open(postgres.Open(os.Getenv("DATABASEURL")), &gorm.Config{})
	if err != nil{
		fmt.Println(err.Error())
		panic("Failed to connect to DB")
	}
	DB.AutoMigrate(&models.User{})
}

func main(){
	InitialMigration()
}