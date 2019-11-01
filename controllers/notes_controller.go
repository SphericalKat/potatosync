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
* @api {post} /api/notes/save
* @apiName Create or update a note
* @apiGroup Notes
*
* @apiPermission Logged-in users
*
* @apiParam {number} note_id The note's ID
* @apiParam {string} title The note's title
* @apiParam {string} content The note's content
* @apiParam {boolean} is_starred Whether the note is starred or not
* @apiParam {number} date Date of note creation
* @apiParam {number} color The color of the note
* @apiParam {string} image_path Image path
* @apiParam {boolean} is_list Whether the note contains a list or not
* @apiParam {string} list_parse_string List parse string
* @apiParam {string} reminders Reminders associated with the note
* @apiParam {boolean} hide_content Whether to hide the note or not
* @apiParam {string} pin Pin
* @apiParam {string} password Password to lock the note with
* @apiParam {boolean} is_deleted Is deleted or not
* @apiParam {boolean} is_archived Is archived or not
* @apiParamExample {json} request-example
*
* 	{
*		"note_id": 5,
*		"title": "Sample note 2",
*		"content": "This is a sample note 2",
*		"is_starred": true,
*		"date": 20191211,
*		"color": 4,
*		"image_path": "abcd",
*		"is_list": false,
*		"list_parse_string": "abcd",
*		"reminders": "abcd",
*		"hide_content": false,
*		"pin": "smth",
*		"password": "password123",
*		"is_deleted": false,
*		"is_archived": false
* 	}
*
* @apiHeader {string} Authorization JWT token associated with user account
* @apiHeaderExample {string} Header-Example:
*	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjV9.FTIhjfCLND1L-hvhgH9_TC_P7MbGQnjnNnFOjJL8Q1k
*
**/
var CreateNote = func(w http.ResponseWriter, r *http.Request) {
	note := &models.Notes{}
	err := json.NewDecoder(r.Body).Decode(note)
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		return
	}
	res := note.SaveNote(r.Context())
	u.Respond(w, res)
}

/**
* @api {get} /api/notes/list
* @apiName List notes associated with a user
* @apiGroup Notes
*
* @apiPermission Logged-in users
*
* @apiHeader {string} Authorization JWT token associated with user account
* @apiHeaderExample {string} Header-Example:
*	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjV9.FTIhjfCLND1L-hvhgH9_TC_P7MbGQnjnNnFOjJL8Q1k
*
**/
var ListNotes = func(w http.ResponseWriter, r *http.Request) {
	res := models.ListNotes(r.Context())
	u.Respond(w, res)
}

/**
* @api {post} /api/notes/delete
* @apiName Delete a saved note
* @apiGroup Notes
*
* @apiPermission Logged-in users
*
* @apiParam {number} note_id The note's ID
* @apiParamExample {json} request-example
*
* 	{
*		"note_id": 5,
* 	}
*
* @apiHeader {string} Authorization JWT token associated with user account
* @apiHeaderExample {string} Header-Example:
*	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjV9.FTIhjfCLND1L-hvhgH9_TC_P7MbGQnjnNnFOjJL8Q1k
*
**/
var DeleteNote = func(w http.ResponseWriter, r *http.Request) {
	note := &models.Notes{}
	err := json.NewDecoder(r.Body).Decode(note)
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		return
	}
	res := models.DeleteNote(r.Context(), note.NoteID)
	u.Respond(w, res)
}
