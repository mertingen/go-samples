package services

import (
	"database/sql"
	"log"
)

type Database struct {
	driver   string
	connInfo string
}

func NewDatabase(driver, connInfo string) Database {
	return Database{driver: driver, connInfo: connInfo}
}

func (d *Database) Connect() *sql.DB {
	// Open up our database connection.
	db, err := sql.Open(d.driver, d.connInfo)
	if err != nil {
		log.Fatalln(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}
	log.Println("Database is up and running...")

	return db
}
