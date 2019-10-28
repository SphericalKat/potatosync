package controllers

import (
	"encoding/json"
	"github.com/ATechnoHazard/potatonotes-api/models"
	u "github.com/ATechnoHazard/potatonotes-api/utils"
	"net/http"
)

var ModifyUsername = func(w http.ResponseWriter, r *http.Request) {
	acc := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(acc)
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		return
	}

	res := models.ModifyUsername(r.Context(), acc.Username)
	u.Respond(w, res)
}

var ModifyPassword = func(w http.ResponseWriter, r *http.Request) {
	acc := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(acc)
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		return
	}

	res := models.ModifyPassword(r.Context(), acc.Password)
	u.Respond(w, res)
}
