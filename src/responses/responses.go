package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

// return response in JSON
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(dados); err != nil {
		log.Fatal(err)
	}
}

// return Error in Json
func Error(w http.ResponseWriter, statusCode int, err error) {

	response := ErrorResponse{
		Error: err.Error(),
	}

	JSON(w, statusCode, response)

}
