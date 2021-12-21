package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	//specify status code
	w.WriteHeader(http.StatusOK)

	//update response writer
	_, err := fmt.Fprintf(w, "API is up and running...")
	if err != nil {
		log.Fatalln(err)
	}

}
