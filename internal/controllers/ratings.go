package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nooderg/pipiSpot/internal/configs"
	"github.com/nooderg/pipiSpot/pkg/responses"
)

func CreateRatings(w http.ResponseWriter, r *http.Request) {
	var ratingsForm forms.Ratings
	err := json.NewDecoder(r.Body).Decode(&ratingsForm)
	if err != nil {
		responses.WriteError(w, err)
		return
	}

	spotID, err := strconv.Atoi(mux.Vars(r)["spot_id"])
	if err != nil {
		responses.WriteError(w, errors.New("cannot get spot_id"))
		return
	}

	newRatings := ratingsForm.GetRatings(uint(1), uint(spotID))

	dbClient := configs.GetDBClient()
	spotRepo := repository.SpotRepository{}
	_, err = spotRepo.GetSpotByID(dbClient, uint(spotID))
	if err != nil {
		responses.WriteError(w, errors.New("spot does not exist"))
		return
	}

	ratingsRepo := repository.RatingsRepository{}
	err = ratingsRepo.CreateRatings(dbClient, &newRatings)
	if err != nil {
		responses.WriteError(w, err)
		return
	}

	responses.WriteData(w, newRatings)
}

func GetRatings(w http.ResponseWriter, r *http.Request) {
	ratingsID, err := strconv.Atoi(mux.Vars(r)["ratings_id"])
	if err != nil {
		responses.WriteError(w, errors.New("cannot get ratings_id"))
		return
	}

	spotID, err := strconv.Atoi(mux.Vars(r)["spot_id"])
	if err != nil {
		responses.WriteError(w, errors.New("cannot get spot_id"))
		return
	}

	dbClient := configs.GetDBClient()
	spotRepo := repository.SpotRepository{}
	_, err = spotRepo.GetSpotByID(dbClient, uint(spotID))
	if err != nil {
		responses.WriteError(w, errors.New("spot does not exist"))
		return
	}

	ratingsRepo := repository.RatingsRepository{}
	Ratings, err := ratingsRepo.GetRatingsByID(dbClient, uint(spotID), uint(ratingsID))
	if err != nil {
		responses.WriteError(w, err)
		return
	}

	responses.WriteData(w, Ratings)
}

func UpdateRatings(w http.ResponseWriter, r *http.Request) {
	ratingsID, err := strconv.Atoi(mux.Vars(r)["ratings_id"])
	if err != nil {
		responses.WriteError(w, errors.New("cannot get ratings_id"))
		return
	}

	spotID, err := strconv.Atoi(mux.Vars(r)["spot_id"])
	if err != nil {
		responses.WriteError(w, errors.New("cannot get spot_id"))
		return
	}

	var ratingsForm forms.Ratings
	err = json.NewDecoder(r.Body).Decode(&ratingsForm)
	if err != nil {
		responses.WriteError(w, err)
		return
	}

	newRatings := ratingsForm.GetRatings(uint(1), uint(spotID))

	dbClient := configs.GetDBClient()
	spotRepo := repository.SpotRepository{}
	_, err = spotRepo.GetSpotByID(dbClient, uint(spotID))
	if err != nil {
		responses.WriteError(w, errors.New("spot does not exist"))
		return
	}

	repo := repository.RatingsRepository{}
	err = repo.UpdateRatings(dbClient, uint(spotID), uint(ratingsID), &newRatings)
	if err != nil {
		responses.WriteError(w, err)
		return
	}

	responses.WriteData(w, newRatings)
}

func DeleteRatings(w http.ResponseWriter, r *http.Request) {
	ratingsID, err := strconv.Atoi(mux.Vars(r)["ratings_id"])
	if err != nil {
		responses.WriteError(w, errors.New("cannot get ratings_id"))
		return
	}

	spotID, err := strconv.Atoi(mux.Vars(r)["spot_id"])
	if err != nil {
		responses.WriteError(w, errors.New("cannot get spot_id"))
		return
	}

	dbClient := configs.GetDBClient()

	spotRepo := repository.SpotRepository{}
	_, err = spotRepo.GetSpotByID(dbClient, uint(spotID))
	if err != nil {
		responses.WriteError(w, errors.New("spot does not exist"))
		return
	}

	repo := repository.RatingsRepository{}
	_, err = repo.GetRatingsByID(dbClient, uint(spotID), uint(ratingsID))
	if err != nil {
		responses.WriteError(w, errors.New("ratings does not exist"))
		return
	}

	err = repo.DeleteRatings(dbClient, uint(spotID), uint(ratingsID))
	if err != nil {
		responses.WriteError(w, err)
		return
	}
}
