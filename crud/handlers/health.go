package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	//update content type
	w.Header().Set("Content-Type", "application/json")

	//specify HTTP status code
	w.WriteHeader(http.StatusOK)

	//convert struct to JSON
	jsonResponse, err := json.Marshal(map[string]bool{"status": true})
	if err != nil {
		return
	}

	//update response
	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Fatalln(err)
	}

}
