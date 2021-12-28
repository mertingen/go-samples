package services

import (
	"database/sql"
	"github.com/mertingen/go-samples/models"
)

type Lecture struct {
	db *sql.DB
}

func NewLecture(db *sql.DB) Lecture {
	return Lecture{db: db}
}

func (l *Lecture) Delete(id int64) error {
	stmt, err := l.db.Prepare("DELETE FROM lectures WHERE id=?")

	// if there is an error, handle it
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (l *Lecture) FetchAll() ([]models.Lecture, error) {
	lectures := make([]models.Lecture, 0)
	lecture := models.Lecture{}
	rows, err := l.db.Query("SELECT * FROM lectures ORDER BY name")
	if err != sql.ErrNoRows && err != nil {
		return lectures, err
	}
	for rows.Next() {
		err := rows.Scan(&lecture.Id,
			&lecture.Name)
		if err != nil {
			return lectures, err
		}
		lectures = append(lectures, lecture)
	}
	return lectures, nil
}

func (l *Lecture) FetchOneById(id int64) (models.Lecture, error) {
	lecture := models.Lecture{}

	err := l.db.QueryRow("SELECT * FROM lectures WHERE id=?", id).Scan(
		&lecture.Id,
		&lecture.Name)

	//if there is no row, it shouldn't give an error
	//thus "sql.ErrNoRows" is added
	if err != sql.ErrNoRows && err != nil {
		return lecture, err
	}

	return lecture, nil
}

func (l *Lecture) FetchOneByName(name string) (models.Lecture, error) {
	lecture := models.Lecture{}

	err := l.db.QueryRow("SELECT * FROM lectures WHERE name=?", name).Scan(
		&lecture.Id,
		&lecture.Name)

	//if there is no row, it shouldn't give an error
	//thus "sql.ErrNoRows" is added
	if err != sql.ErrNoRows && err != nil {
		return lecture, err
	}

	return lecture, nil
}

func (l *Lecture) Update(data models.Lecture) (models.Lecture, error) {
	// perform a db.Query insert
	stmt, err := l.db.Prepare("UPDATE lectures SET name=? WHERE id=?")

	// if there is an error inserting, handle it
	if err != nil {
		return data, err
	}

	_, err = stmt.Exec(data.Name, data.Id)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (l *Lecture) Insert(data models.Lecture) (int64, error) {
	// perform a db.Query insert
	stmt, err := l.db.Prepare("INSERT INTO lectures(name) VALUES (?)")

	// if there is an error inserting, handle it
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(data.Name)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
