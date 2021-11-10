package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nooderg/pipiSpot/internal/configs"
	"github.com/nooderg/pipiSpot/internal/forms"
	"github.com/nooderg/pipiSpot/internal/repository"
	"github.com/nooderg/pipiSpot/pkg/responses"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var userForm forms.User
	err := json.NewDecoder(r.Body).Decode(&userForm)
	if err != nil {
		responses.WriteError(w, err)
		return
	}

	newUser := userForm.GetUser()

	dbClient := configs.GetDBClient()
	repo := repository.UserRepository{}
	err = repo.CreateUser(dbClient, &newUser)
	if err != nil {
		responses.WriteError(w, err)
		return
	}

	responses.WriteData(w, newUser)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		responses.WriteError(w, errors.New("cannot get id"))
		return
	}

	dbClient := configs.GetDBClient()
	repo := repository.UserRepository{}
	user, err := repo.GetUserByID(dbClient, uint(userID))
	if err != nil {
		responses.WriteError(w, err)
		return
	}

	responses.WriteData(w, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		responses.WriteError(w, errors.New("cannot get id"))
		return
	}

	var userForm forms.User
	err = json.NewDecoder(r.Body).Decode(&userForm)
	if err != nil {
		responses.WriteError(w, err)
		return
	}

	newUser := userForm.GetUser()
	newUser.ID = uint(userID)

	dbClient := configs.GetDBClient()
	repo := repository.UserRepository{}
	err = repo.UpdateUser(dbClient, uint(userID), &newUser)
	if err != nil {
		responses.WriteError(w, err)
		return
	}

	responses.WriteData(w, newUser)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		responses.WriteError(w, errors.New("cannot get id"))
		return
	}

	dbClient := configs.GetDBClient()
	repo := repository.UserRepository{}

	_, err = repo.GetUserByID(dbClient, uint(userID))
	if err != nil {
		responses.WriteError(w, errors.New("user does not exist"))
		return
	}

	err = repo.DeleteUser(dbClient, uint(userID))
	if err != nil {
		responses.WriteError(w, err)
		return
	}
}
