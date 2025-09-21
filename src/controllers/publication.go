package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CreatePublicaton(w http.ResponseWriter, r *http.Request) {
	userID, err := utils.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publication models.Publication

	if err = json.Unmarshal(reqBody, &publication); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
	}
	defer db.Close()

	repository := repositories.NewUsersRepositoryPulications(db)
	publicationID, err := repository.Create(userID, publication)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, publicationID)

}
func CreatePublicatons(w http.ResponseWriter, r *http.Request) {

}
func GetPublication(w http.ResponseWriter, r *http.Request) {

}
func UpdatePublication(w http.ResponseWriter, r *http.Request) {

}
func DeletePublication(w http.ResponseWriter, r *http.Request) {

}
