/*
 *     Copyright ATechnoHazard 2019  <amolele@gmail.com>.
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

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
