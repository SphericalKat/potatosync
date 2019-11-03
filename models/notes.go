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

package models

import (
	"context"

	u "github.com/ATechnoHazard/potatonotes-api/utils"
)

// Notes Represents a note in the database
type Notes struct {
	NoteID          uint   `gorm:"primary_key" json:"note_id"`
	AccountID       uint   `gorm:"column:account_id" json:"id"`
	Title           string `json:"title"`
	Content         string `json:"content"`
	IsStarred       bool   `json:"is_starred"`
	Date            int    `json:"date"`
	Color           int    `json:"color"`
	ImageURL        string `json:"image_url"`
	IsList          bool   `json:"is_list"`
	ListParseString string `json:"list_parse_string"`
	Reminders       string `json:"reminders"`
	HideContent     bool   `json:"hide_content"`
	Pin             string `json:"pin"`
	Password        string `json:"password"`
	IsDeleted       bool   `json:"is_deleted"`
	IsArchived      bool   `json:"is_archived"`
}

// SaveNote Create or update a note
func (note *Notes) SaveNote(ctx context.Context) map[string]interface{} {
	acc := GetUser(ctx.Value("user").(uint))
	if acc == nil {
		return u.Message(false, "UserNotFoundError")
	}

	if note.NoteID <= 0 {
		return u.Message(false, "MissingNoteIdError")
	}

	note.AccountID = acc.ID

	err := GetDB().Save(note).Error
	if err != nil {
		return u.Message(false, err.Error())
	}

	return u.Message(true, "NoteCreationSuccess")
}

// ListNotes List all notes belonging to a user
func ListNotes(ctx context.Context) map[string]interface{} {
	acc := GetUser(ctx.Value("user").(uint))
	if acc == nil {
		return u.Message(false, "UserNotFoundError")
	}
	var notes []Notes
	GetDB().Where("account_id = ?", acc.ID).Find(&notes)

	res := make(map[string]interface{})
	res["notes"] = notes
	return res
}

// DeleteNote Delete a single note belonging to a user
func DeleteNote(ctx context.Context, noteID uint) map[string]interface{} {
	acc := GetUser(ctx.Value("user").(uint))
	if acc == nil {
		return u.Message(false, "UserNotFoundError")
	}

	if noteID <= 0 {
		return u.Message(false, "MissingNoteIdError")
	}

	note := &Notes{}
	err := GetDB().Where("note_id = ?", noteID).Where("account_id = ?", acc.ID).Delete(note).Error
	if err != nil {
		return u.Message(false, err.Error())
	}

	return u.Message(true, "NoteDeleteSuccess")
}

// DeleteAllNotes Deletes all notes belonging to a user
func DeleteAllNotes(ctx context.Context) map[string]interface{} {
	acc := GetUser(ctx.Value("user").(uint))
	if acc == nil {
		return u.Message(false, "UserNotFoundError")
	}

	note := &Notes{}
	err := GetDB().Where("account_id = ?").Delete(note).Error
	if err != nil {
		return u.Message(false, err.Error())
	}

	return u.Message(true, "NotesDeleteSuccess")
}
