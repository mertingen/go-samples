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

type Lecture struct {
	lectureService services.Lecture
}

func NewLecture(lectureService services.Lecture) Lecture {
	return Lecture{lectureService: lectureService}
}

func (l *Lecture) FetchAll(w http.ResponseWriter, r *http.Request) {
	lectures, err := l.lectureService.FetchAll()
	if err != nil {
		log.Fatalln(err)
	}

	//update content type
	w.Header().Set("Content-Type", "application/json")

	//specify HTTP status code
	w.WriteHeader(http.StatusOK)

	resp := make(map[string][]models.Lecture)
	resp["data"] = lectures

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

func (l *Lecture) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	strId := params["id"]

	//it converts string to int64
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		log.Fatalln(err)
	}
	lecture, err := l.lectureService.FetchOneById(id)
	if err != nil {
		log.Fatalln(err)
	}

	//update content type
	w.Header().Set("Content-Type", "application/json")

	//specify HTTP status code
	w.WriteHeader(http.StatusOK)

	if (models.Lecture{}) == lecture {
		resp := make(map[string]string)
		resp["error"] = "Lecture is not found!"

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

	err = l.lectureService.Delete(id)
	if err != nil {
		log.Fatalln(err)
	}

	resp := make(map[string]bool)
	resp["status"] = true

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

func (l *Lecture) FetchOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	strId := params["id"]

	//it converts string to int64
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		log.Fatalln(err)
	}
	lecture, err := l.lectureService.FetchOneById(id)
	if err != nil {
		log.Fatalln(err)
	}

	//update content type
	w.Header().Set("Content-Type", "application/json")

	//specify HTTP status code
	w.WriteHeader(http.StatusOK)

	if (models.Lecture{}) == lecture {
		resp := make(map[string]string)
		resp["error"] = "Lecture is not found!"

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

	resp := make(map[string]models.Lecture)
	resp["data"] = lecture

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

func (l *Lecture) Update(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	lecture := models.Lecture{}
	err = json.Unmarshal(body, &lecture)
	if err != nil {
		log.Fatalln(err)
	}

	params := mux.Vars(r)
	strId := params["id"]

	//it converts string to int64
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		log.Fatalln(err)
	}

	if id != lecture.Id {
		resp := make(map[string]string)
		resp["error"] = "Ids are not matched!"

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

	isExist, err := l.lectureService.FetchOneById(id)
	if err != nil {
		log.Fatalln(err)
	}

	if (models.Lecture{}) == isExist {
		resp := make(map[string]string)
		resp["error"] = "Lecture is not found!"

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

	_, err = l.lectureService.Update(lecture)
	if err != nil {
		log.Fatalln(err)
	}

	//update content type
	w.Header().Set("Content-Type", "application/json")

	//specify HTTP status code
	w.WriteHeader(http.StatusOK)

	resp := make(map[string]models.Lecture)
	resp["data"] = lecture

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

func (l *Lecture) Insert(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	lecture := models.Lecture{}
	err = json.Unmarshal(body, &lecture)
	if err != nil {
		log.Fatalln(err)
	}

	isExist, err := l.lectureService.FetchOneByName(lecture.Name)
	if err != nil {
		log.Fatalln(err)
	}

	//update content type
	w.Header().Set("Content-Type", "application/json")

	//specify HTTP status code
	w.WriteHeader(http.StatusOK)

	if (models.Lecture{}) != isExist {
		resp := make(map[string]string)
		resp["error"] = "This name is already taken!"

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

	id, err := l.lectureService.Insert(lecture)
	if err != nil {
		log.Fatalln(err)
	}
	lecture.Id = id

	resp := make(map[string]models.Lecture)
	resp["data"] = lecture

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
