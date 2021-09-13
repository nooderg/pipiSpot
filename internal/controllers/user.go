package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/nooderg/pipiSpot/internal/models"
	"github.com/nooderg/pipiSpot/pkg/responses"
)

func PostUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		responses.WriteError(w, err)
		return
	}
	
}
