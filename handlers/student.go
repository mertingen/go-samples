package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mertingen/go-samples/models"
	"github.com/mertingen/go-samples/services"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Student struct {
	studentService services.Student
}

func NewStudent(studentService services.Student) Student {
	return Student{studentService: studentService}
}

func (s *Student) FetchOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	//it converts string to int64
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Fatalln(err)
	}
	student, err := s.studentService.FetchOneById(intId)
	if err != nil {
		log.Fatalln(err)
	}

	//update content type
	w.Header().Set("Content-Type", "application/json")

	//specify HTTP status code
	w.WriteHeader(http.StatusOK)

	if (models.Student{}) == student {
		resp := make(map[string]string)
		resp["error"] = "Student is not found!"

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
		return
	}

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

func (s *Student) Update(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	student := models.Student{}
	err = json.Unmarshal(body, &student)
	if err != nil {
		log.Fatalln(err)
	}

	params := mux.Vars(r)
	id := params["id"]

	//it converts string to int64
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Fatalln(err)
	}

	isExist, err := s.studentService.FetchOneById(intId)
	if err != nil {
		log.Fatalln(err)
	}

	if (models.Student{}) == isExist {
		resp := make(map[string]string)
		resp["error"] = "Student is not found!"

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
		return
	}

	_, err = s.studentService.Update(student)
	if err != nil {
		log.Fatalln(err)
	}

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

func (s *Student) Insert(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	student := models.Student{}
	err = json.Unmarshal(body, &student)
	if err != nil {
		log.Fatalln(err)
	}

	isExist, err := s.studentService.FetchOneByEmail(student.Email)
	if err != nil {
		log.Fatalln(err)
	}

	//update content type
	w.Header().Set("Content-Type", "application/json")

	//specify HTTP status code
	w.WriteHeader(http.StatusOK)

	if (models.Student{}) != isExist {
		resp := make(map[string]string)
		resp["error"] = "This e-mail is already taken!"

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
		return
	}

	id, err := s.studentService.Insert(student)
	if err != nil {
		log.Fatalln(err)
	}
	student.Id = id

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
