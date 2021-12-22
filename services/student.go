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

func (s *Student) FetchOneById(id int64) (models.Student, error) {
	student := models.Student{}

	err := s.db.QueryRow("SELECT * FROM students WHERE id=?", id).Scan(
		&student.Id,
		&student.Fullname,
		&student.Email,
		&student.Age)

	//if there is no row, it shouldn't give an error
	//thus "sql.ErrNoRows" is added
	if err != sql.ErrNoRows && err != nil {
		return student, err
	}

	return student, nil
}

func (s *Student) FetchOneByEmail(email string) (models.Student, error) {
	student := models.Student{}

	err := s.db.QueryRow("SELECT * FROM students WHERE email=?", email).Scan(
		&student.Id,
		&student.Fullname,
		&student.Email,
		&student.Age)

	//if there is no row, it shouldn't give an error
	//thus "sql.ErrNoRows" is added
	if err != sql.ErrNoRows && err != nil {
		return student, err
	}

	return student, nil
}

func (s *Student) Update(data models.Student) (models.Student, error) {
	// perform a db.Query insert
	stmt, err := s.db.Prepare("UPDATE students SET fullname=?, email=?, age=? WHERE id=?")

	// if there is an error inserting, handle it
	if err != nil {
		return data, err
	}

	_, err = stmt.Exec(data.Fullname, data.Email, data.Age, data.Id)
	if err != nil {
		return data, err
	}
	return data, nil
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
