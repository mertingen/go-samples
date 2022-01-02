package services

import (
	"database/sql"
	"github.com/mertingen/go-samples/models"
)

type Student struct {
	db             *sql.DB
	lectureService Lecture
}

func NewStudent(db *sql.DB, lectureService Lecture) Student {
	return Student{db: db, lectureService: lectureService}
}

func (s *Student) Delete(id int64) error {
	stmt, err := s.db.Prepare("DELETE FROM students WHERE id=?")

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

func (s *Student) FetchAll() ([]models.Student, error) {
	students := make([]models.Student, 0)
	rows, err := s.db.Query(`SELECT s.id, s.fullname, s.email, s.age, l.id, l.name  FROM students s
									LEFT JOIN students_lectures sl ON sl.student_id = s.id
									LEFT JOIN lectures l ON l.id = sl.lecture_id
									ORDER BY s.fullname`)
	if err != sql.ErrNoRows && err != nil {
		return students, err
	}

	var sId int64
	for rows.Next() {
		student := models.Student{}
		student.Lecture = make([]models.Lecture, 0)
		var lectureId sql.NullInt64
		var lectureName sql.NullString
		lecture := models.Lecture{}
		err := rows.Scan(
			&student.Id,
			&student.Fullname,
			&student.Email,
			&student.Age,
			&lectureId,
			&lectureName)
		if err != nil {
			return students, err
		}
		if lectureId.Valid {
			lecture.Id = lectureId.Int64
			lecture.Name = lectureName.String
			if sId == student.Id {
				//if it's same student, it appends the lecture previous student that in the slice
				//thus the same students will have their lectures.
				lastIndex := len(students) - 1
				students[lastIndex].Lecture = append(students[lastIndex].Lecture, lecture)
			} else {
				student.Lecture = append(student.Lecture, lecture)
			}
		}
		if sId != student.Id {
			students = append(students, student)
		}
		sId = student.Id

	}
	return students, nil
}

func (s *Student) FetchOneById(id int64) (models.Student, error) {
	student := models.Student{}
	student.Lecture = make([]models.Lecture, 0)

	rows, err := s.db.Query(`SELECT s.id, s.fullname, s.email, s.age, l.id, l.name  FROM students s
									LEFT JOIN students_lectures sl ON sl.student_id = s.id
									LEFT JOIN lectures l ON l.id = sl.lecture_id
									WHERE s.id=?
									ORDER BY s.fullname`, id)
	if err != nil {
		return student, err
	}

	//if there is no row, it shouldn't give an error
	//thus "sql.ErrNoRows" is added
	if err != sql.ErrNoRows && err != nil {
		return student, err
	}

	for rows.Next() {
		var lectureId sql.NullInt64
		var lectureName sql.NullString
		lecture := models.Lecture{}
		err := rows.Scan(
			&student.Id,
			&student.Fullname,
			&student.Email,
			&student.Age,
			&lectureId,
			&lectureName)
		if err != nil {
			return student, err
		}
		if lectureId.Valid {
			lecture.Id = lectureId.Int64
			lecture.Name = lectureName.String
			student.Lecture = append(student.Lecture, lecture)
		}
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

func (s *Student) AttachLectures(student models.Student, lectureIds []int64) error {
	for _, l := range student.Lecture {
		willBeDeleted := true
		for _, nl := range lectureIds {
			if l.Id == nl {
				willBeDeleted = false
			}
		}
		if willBeDeleted == true {
			// perform a db.Query insert
			stmt, err := s.db.Prepare("DELETE FROM students_lectures WHERE student_id=? AND lecture_id=?")

			// if there is an error inserting, handle it
			if err != nil {
				continue
			}

			_, err = stmt.Exec(student.Id, l.Id)
			if err != nil {
				continue
			}
			continue
		}
	}

	//it's going to continue if there is an error
	//there should be a proper way to fix this situation.
	for _, id := range lectureIds {
		isExistLecture, err := s.lectureService.FetchOneById(id)
		if err != nil {
			continue
		}
		if (models.Lecture{}) == isExistLecture {
			continue
		}

		//if this new id cannot match exist one, it'll be added as a new row.
		willBeAdded := true
		for _, l := range student.Lecture {
			if l.Id == id {
				willBeAdded = false
			}
		}

		if willBeAdded == true {
			// perform a db.Query insert
			stmt, err := s.db.Prepare("INSERT INTO students_lectures(student_id,lecture_id) VALUES (?, ?)")

			// if there is an error inserting, handle it
			if err != nil {
				continue
			}

			_, err = stmt.Exec(student.Id, id)
			if err != nil {
				continue
			}
			continue
		}
	}
	return nil
}
