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

var ListNotes = func(w http.ResponseWriter, r *http.Request) {
	res := models.ListNotes(r.Context())
	u.Respond(w, res)
}

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
