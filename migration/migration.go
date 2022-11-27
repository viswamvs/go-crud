package migration

import (
	"go-crud/config"
	"go-crud/daos/models"
)

func init() {
	config.LoadEnvVariables()
	config.InitializeDB()
}

func main() {
	config.DB.AutoMigrate(&models.User{})
}
