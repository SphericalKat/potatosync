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

/**
* @api {post} /api/users/manage/username
* @apiName Modify account's username
* @apiGroup Authorization
*
* @apiPermission Logged-in users
*
* @apiParam {string} email The user's email
* @apiParam {string} username The new username
* @apiParamExample {json} request-example
*
* {
*	"email": "john.doe@example.com",
*	"username": "JohnDoe",
* }
*
* @apiHeader {string} Authorization JWT token associated with user account
* @apiHeaderExample {string} Header-Example:
*	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjV9.FTIhjfCLND1L-hvhgH9_TC_P7MbGQnjnNnFOjJL8Q1k
*
**/
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

/**
* @api {post} /api/users/manage/password
* @apiName Modify account's password
* @apiGroup Authorization
*
* @apiPermission Logged-in users
*
* @apiParam {string} password The new password
* @apiParamExample {json} request-example
*
* {
*	"password": "password123",
* }
*
* @apiHeader {string} Authorization JWT token associated with user account
* @apiHeaderExample {string} Header-Example:
*	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjV9.FTIhjfCLND1L-hvhgH9_TC_P7MbGQnjnNnFOjJL8Q1k
*
**/
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

/**
* @api {post} /api/users/manage/image
* @apiName Modify account's password
* @apiGroup Authorization
*
* @apiPermission Logged-in users
*
* @apiParam {string} image_url The url to the user's profile image
* @apiParamExample {json} request-example
*
* {
*	"image_url": "https://example.com/1234567",
* }
*
* @apiHeader {string} Authorization JWT token associated with user account
* @apiHeaderExample {string} Header-Example:
*	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjV9.FTIhjfCLND1L-hvhgH9_TC_P7MbGQnjnNnFOjJL8Q1k
*
**/
var SaveImage = func(w http.ResponseWriter, r *http.Request) {
	acc := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(acc)
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		return
	}

	res := models.SaveAccImage(r.Context(), acc.ImageUrl)
	u.Respond(w, res)
}