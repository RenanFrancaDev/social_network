package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("[controllers] [msg: Error in the request %v]", err)
	}

	var user models.User

	if err = json.Unmarshal(req, &user); err != nil {
		log.Fatalf("[controllers] [msg: Error in unmarshal %v]", err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("[controllers] [msg: Error in database connect %v]", err)
	}

	repository := repositories.NewUsersRepository(db)
	userID, err := repository.Create(user)
	if err != nil {
		log.Fatalf("[controllers] [msg: Error to create user repository %v]", err)
	}

	w.Write([]byte(fmt.Sprintf("Id inserido: %v", userID)))

}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	db, err := database.Connect()
	if err != nil {
		log.Printf("Request received: %s %s", r.Method, r.URL.Path)
		log.Fatalf("[controllers] [msg: Error in database connect %v]", err)
		http.Error(w, "Database connection error", http.StatusInternalServerError)
	}

	var users []models.User

	repository := repositories.NewUsersRepository(db)
	users, err = repository.GetUsersRepository()
	if err != nil {
		log.Fatalf("[controllers] [msg: Error to create user repository %v]", err)
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err = json.NewEncoder(w).Encode(users); err != nil {
		log.Printf("[controllers] Error encoding JSON: %v", err)
		http.Error(w, "Error generating response", http.StatusInternalServerError)
	}

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Funcionou"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Funcionou"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Funcionou"))
}
