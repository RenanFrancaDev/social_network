package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Funcionou"))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Funcionou"))
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
