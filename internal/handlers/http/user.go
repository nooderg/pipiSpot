package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	command "github.com/nooderg/pipiSpot/internal/application/command/user"
	query "github.com/nooderg/pipiSpot/internal/application/query/user"
	"github.com/nooderg/pipiSpot/pkg/responses"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var userCommand command.RegisterCommand
	err := json.NewDecoder(r.Body).Decode(&userCommand)
	if err != nil {
		responses.WriteError(w, err)
		return
	}

	handler := command.RegisterCommandHandler.New(command.RegisterCommandHandler{})
	user, err := handler.Handle(userCommand)
	if err != nil {
		responses.WriteError(w, err)
		return
	}

	responses.WriteData(w, user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		responses.WriteError(w, errors.New("cannot get id"))
		return
	}

	userQuery := query.GetUserQuery{ID: uint(userID)}

	handler := query.GetUserQueryHandler.New(query.GetUserQueryHandler{})
	user, err := handler.Handle(userQuery)
	if err != nil {
		responses.WriteError(w, err)
		return
	}

	responses.WriteData(w, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var userCommand command.RegisterCommand
	err := json.NewDecoder(r.Body).Decode(&userCommand)
	if err != nil {
		responses.WriteError(w, err)
		return
	}

	handler := command.RegisterCommandHandler.New(command.RegisterCommandHandler{})
	user, err := handler.Handle(userCommand)
	if err != nil {
		responses.WriteError(w, err)
		return
	}

	responses.WriteData(w, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var userCommand command.DeleteCommand
	err := json.NewDecoder(r.Body).Decode(&userCommand)
	if err != nil {
		responses.WriteError(w, err)
		return
	}

	handler := command.DeleteCommandHandler.New(command.DeleteCommandHandler{})
	err = handler.Handle(userCommand)
	if err != nil {
		responses.WriteError(w, err)
		return
	}

	responses.WriteData(w, err == nil)
}
