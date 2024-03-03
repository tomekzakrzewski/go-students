package db

import (
	"os"

	"github.com/tomekzakrzewski/go-students/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DBconnect() (*gorm.DB, error) {
	dbName := os.Getenv("DB_NAME")
	database, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	database.Migrator().DropTable(&models.Student{})
	err = database.AutoMigrate(&models.Student{})
	if err != nil {
		return nil, err
	}
	return database, nil
}
