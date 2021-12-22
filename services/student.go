package services

import (
	"database/sql"
	"github.com/mertingen/go-samples/models"
)

type Student struct {
	db *sql.DB
}

func NewStudent(db *sql.DB) Student {
	return Student{db: db}
}

func (s *Student) Insert(data models.Student) (int64, error) {
	// perform a db.Query insert
	stmt, err := s.db.Prepare("INSERT INTO students(fullname,email,age) VALUES (?, ?, ?)")

	// if there is an error inserting, handle it
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(data.Fullname, data.Email, data.Age)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
