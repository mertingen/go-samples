package handlers

import (
	"encoding/json"
	"github.com/mertingen/go-samples/models"
	"github.com/mertingen/go-samples/services"
	"io/ioutil"
	"log"
	"net/http"
)

type Student struct {
	studentService services.Student
}

func NewStudent(studentService services.Student) Student {
	return Student{studentService: studentService}
}

func (s *Student) Insert(w http.ResponseWriter, r *http.Request) {
	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	student := models.Student{}
	err := json.Unmarshal(body, &student)
	if err != nil {
		log.Fatalln(err)
	}

	id, err := s.studentService.Insert(student)
	if err != nil {
		log.Fatalln(err)
	}
	student.Id = id

	//update content type
	w.Header().Set("Content-Type", "application/json")

	//specify HTTP status code
	w.WriteHeader(http.StatusOK)

	resp := make(map[string]models.Student)
	resp["data"] = student

	//convert struct to JSON
	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		return
	}

	//update response
	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Fatalln(err)
	}
}
