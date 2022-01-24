package config

import (
	"fmt"
	"github.com/mertingen/go-samples/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var Database *gorm.DB
var ConnInfo string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DB"))

func Connect() error {
	var err error

	Database, err = gorm.Open(mysql.Open(ConnInfo), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		log.Fatalln(err)
	}

	err = Database.AutoMigrate(&entities.Student{})
	if err != nil {
		log.Fatalln(err)
	}

	return nil
}
