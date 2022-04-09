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

func CreateComment(w http.ResponseWriter, r *http.Request) {
	spotID, err := strconv.Atoi(mux.Vars(r)["spot_id"])
	if err != nil {
		responses.WriteError(w, errors.New("cannot get spot_id"))
		return
	}

	var commentForm forms.Comment
	err = json.NewDecoder(r.Body).Decode(&commentForm)
	if err != nil {
		responses.WriteError(w, err)
		return
	}

	newComment := commentForm.GetComment(uint(spotID))

	dbClient := configs.GetDBClient()
	spotRepo := repository.SpotRepository{}
	_, err = spotRepo.GetSpotByID(dbClient, uint(spotID))
	if err != nil {
		responses.WriteError(w, errors.New("spot does not exist"))
		return
	}

	commentsRepo := repository.CommentRepository{}
	err = commentsRepo.CreateComment(dbClient, &newComment)
	if err != nil {
		responses.WriteError(w, err)
		return
	}

	responses.WriteData(w, newComment)
}

func GetComment(w http.ResponseWriter, r *http.Request) {
	spotID, err := strconv.Atoi(mux.Vars(r)["spot_id"])
	if err != nil {
		responses.WriteError(w, errors.New("cannot get spot_id"))
		return
	}

	commentID, err := strconv.Atoi(mux.Vars(r)["comm_id"])
	if err != nil {
		responses.WriteError(w, errors.New("cannot get comm_id"))
		return
	}

	dbClient := configs.GetDBClient()

	spotRepo := repository.SpotRepository{}
	_, err = spotRepo.GetSpotByID(dbClient, uint(spotID))
	if err != nil {
		responses.WriteError(w, errors.New("spot does not exist"))
		return
	}

	commentsRepo := repository.CommentRepository{}
	Comment, err := commentsRepo.GetCommentByID(dbClient, uint(spotID), uint(commentID))
	if err != nil {
		responses.WriteError(w, err)
		return
	}

	responses.WriteData(w, Comment)
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {
	spotID, err := strconv.Atoi(mux.Vars(r)["spot_id"])
	if err != nil {
		responses.WriteError(w, errors.New("cannot get spot_id"))
		return
	}

	commentID, err := strconv.Atoi(mux.Vars(r)["comm_id"])
	if err != nil {
		responses.WriteError(w, errors.New("cannot get comm_id"))
		return
	}

	var commentForm forms.Comment
	err = json.NewDecoder(r.Body).Decode(&commentForm)
	if err != nil {
		responses.WriteError(w, err)
		return
	}

	newComment := commentForm.GetComment(uint(spotID))

	dbClient := configs.GetDBClient()

	spotRepo := repository.SpotRepository{}
	_, err = spotRepo.GetSpotByID(dbClient, uint(spotID))
	if err != nil {
		responses.WriteError(w, errors.New("spot does not exist"))
		return
	}

	commentsRepo := repository.CommentRepository{}
	err = commentsRepo.UpdateComment(dbClient, uint(spotID), uint(commentID), &newComment)
	if err != nil {
		responses.WriteError(w, err)
		return
	}

	responses.WriteData(w, newComment)
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	spotID, err := strconv.Atoi(mux.Vars(r)["spot_id"])
	if err != nil {
		responses.WriteError(w, errors.New("cannot get spot_id"))
		return
	}

	commentID, err := strconv.Atoi(mux.Vars(r)["comm_id"])
	if err != nil {
		responses.WriteError(w, errors.New("cannot get comm_id"))
		return
	}

	dbClient := configs.GetDBClient()

	spotRepo := repository.SpotRepository{}
	_, err = spotRepo.GetSpotByID(dbClient, uint(spotID))
	if err != nil {
		responses.WriteError(w, errors.New("spot does not exist"))
		return
	}

	commentsRepo := repository.CommentRepository{}
	_, err = commentsRepo.GetCommentByID(dbClient, uint(spotID), uint(commentID))
	if err != nil {
		responses.WriteError(w, errors.New("comment does not exist"))
		return
	}

	err = commentsRepo.DeleteComment(dbClient, uint(spotID), uint(commentID))
	if err != nil {
		responses.WriteError(w, err)
		return
	}
}
