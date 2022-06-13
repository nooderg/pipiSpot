package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nooderg/pipiSpot/internal/configs"
	"github.com/nooderg/pipiSpot/pkg/responses"
)

func CreateSpot(w http.ResponseWriter, r *http.Request) {
	var spotForm forms.Spot
	err := json.NewDecoder(r.Body).Decode(&spotForm)
	if err != nil {
		responses.WriteError(w, err)
		return
	}

	newSpot := spotForm.GetSpot(uint(1))

	dbClient := configs.GetDBClient()
	repo := repository.SpotRepository{}
	err = repo.CreateSpot(dbClient, &newSpot)
	if err != nil {
		responses.WriteError(w, err)
		return
	}

	responses.WriteData(w, newSpot)
}

func GetSpot(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["spot_id"])
	if err != nil {
		responses.WriteError(w, errors.New("cannot get id"))
		return
	}

	dbClient := configs.GetDBClient()
	repo := repository.SpotRepository{}
	Spot, err := repo.GetSpotByID(dbClient, uint(id))
	if err != nil {
		responses.WriteError(w, err)
		return
	}

	responses.WriteData(w, Spot)
}

func UpdateSpot(w http.ResponseWriter, r *http.Request) {
	spotID, err := strconv.Atoi(mux.Vars(r)["spot_id"])
	if err != nil {
		responses.WriteError(w, errors.New("cannot get id"))
		return
	}

	var spotForm forms.Spot
	err = json.NewDecoder(r.Body).Decode(&spotForm)
	if err != nil {
		responses.WriteError(w, err)
		return
	}

	newSpot := spotForm.GetSpot(uint(1))
	newSpot.ID = uint(spotID)

	dbClient := configs.GetDBClient()
	repo := repository.SpotRepository{}
	err = repo.UpdateSpot(dbClient, uint(spotID), &newSpot)
	if err != nil {
		responses.WriteError(w, err)
		return
	}

	responses.WriteData(w, newSpot)
}

func DeleteSpot(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["spot_id"])
	if err != nil {
		responses.WriteError(w, errors.New("cannot get id"))
		return
	}

	dbClient := configs.GetDBClient()
	repo := repository.SpotRepository{}

	_, err = repo.GetSpotByID(dbClient, uint(id))
	if err != nil {
		responses.WriteError(w, errors.New("spot does not exist"))
		return
	}

	err = repo.DeleteSpot(dbClient, uint(id))
	if err != nil {
		responses.WriteError(w, err)
		return
	}
}
