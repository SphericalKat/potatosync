package controllers

import (
	"encoding/json"
	"github.com/ATechnoHazard/potatonotes-api/models"
	u "github.com/ATechnoHazard/potatonotes-api/utils"
	"net/http"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	acc := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(acc) // decode request body
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		return
	}

	res := acc.Create() // create account
	u.Respond(w, res)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	acc := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(acc)
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		return
	}

	if acc.Email == "" {
		if acc.Username == "" {
			u.Respond(w, u.Message(false, "Both username and password missing"))
		}

		res := models.LoginUsername(acc.Username, acc.Password)
		u.Respond(w, res)
		return
	}

	res := models.Login(acc.Email, acc.Password)
	u.Respond(w, res)
}

var Delete = func(w http.ResponseWriter, r *http.Request) {
	res := models.Delete(r.Context())
	u.Respond(w, res)
}

var UserInfo = func(w http.ResponseWriter, r *http.Request) {
	res := models.AccInfo(r.Context())
	u.Respond(w, res)
}
