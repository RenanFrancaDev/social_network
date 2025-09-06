package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/utils"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	if err := json.Unmarshal(req, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	registeredUser, err := repository.GetUserByEmail(user.Email)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = utils.CheckPassword(registeredUser.Password, user.Password); err != nil {
		log.Printf("hash: %s sem HAsh: %s", registeredUser.Password, user.Password)
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	token, _ := utils.CreateToken(registeredUser.ID)
	w.Write([]byte(token))

	// token, err := repository.Signin(user)
	// if err != nil {
	// 	responses.Error(w, http.StatusInternalServerError, err)
	// }

	// responses.JSON(w, http.StatusOK, err)

}
