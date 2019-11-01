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
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"strings"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

// Represents a user account
type Account struct {
	gorm.Model
	Email    string  `json:"email"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Token    string  `gorm:"-" json:"token"`
	Notes    []Notes `gorm:"foreignkey:AccountID" json:"-"`
}

// Validate incoming user details
func (acc *Account) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(acc.Email, "@") {
		return u.Message(false, "Missing/Malformed email"), false
	}

	if len(acc.Username) <= 4 || len(acc.Username) > 60 {
		return u.Message(false, "Username length not in bounds"), false
	}

	if len(acc.Password) < 8 || len(acc.Password) > 60 {
		return u.Message(false, "Password length not in bounds"), false
	}

	acc.Username = strings.ToLower(acc.Username)

	temp := &Account{}

	// check for duplicate email and username
	err := GetDB().Where("email = ?", acc.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error, try again"), false
	}
	if temp.Email != "" {
		return u.Message(false, "Email address already in use!"), false
	}

	err = GetDB().Where("username = ?", acc.Username).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error, try again"), false
	}
	if temp.Username != "" {
		return u.Message(false, "Username already in use!"), false
	}

	return u.Message(true, "Validated successfully"), true
}

func (acc *Account) Create() map[string]interface{} {
	if res, ok := acc.Validate(); !ok {
		return res
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(acc.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln(err)
	}
	acc.Password = string(hashedPass)

	GetDB().Create(acc)

	if acc.ID <= 0 {
		return u.Message(false, "Failed to create account, connection error")
	}

	// create a jwt token for the newly registered account
	tk := &Token{UserId: acc.ID}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tk)
	tokenString, err := token.SignedString([]byte(os.Getenv("token_password")))
	if err != nil {
		log.Fatalln(err)
	}
	acc.Token = tokenString

	acc.Password = "" // delete password

	response := u.Message(true, "Account has been created")
	response["account"] = acc
	return response
}

func Login(email, pass string) map[string]interface{} {
	acc := &Account{}
	err := GetDB().Where("email = ?", email).First(acc).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error, try again")
	}

	err = bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(pass))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { // password doesn't match
		return u.Message(false, "Invalid login credentials")
	}

	// login successful
	acc.Password = ""

	// create jwt token
	tk := &Token{UserId: acc.ID}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tk)
	tokenString, err := token.SignedString([]byte(os.Getenv("token_password")))
	if err != nil {
		log.Fatalln(err)
	}
	acc.Token = tokenString
	res := u.Message(true, "Login successful")
	res["account"] = acc
	return res
}

func LoginUsername(username, pass string) map[string]interface{} {
	acc := &Account{}
	err := GetDB().Where("username = ?", strings.ToLower(username)).First(acc).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Username not found")
		}
		return u.Message(false, "Connection error, try again")
	}

	err = bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(pass))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { // password doesn't match
		return u.Message(false, "Invalid login credentials")
	}

	// login successful
	acc.Password = ""

	// create jwt token
	tk := &Token{UserId: acc.ID}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tk)
	tokenString, err := token.SignedString([]byte(os.Getenv("token_password")))
	if err != nil {
		log.Fatalln(err)
	}
	acc.Token = tokenString
	res := u.Message(true, "Login successful")
	res["account"] = acc
	return res
}

func DeleteAccount(ctx context.Context) map[string]interface{} {
	acc := GetUser(ctx.Value("user").(uint))
	if acc == nil {
		return u.Message(false, "Account not found")
	}

	err := GetDB().Delete(acc).Error
	if err != nil {
		return u.Message(false, err.Error())
	}

	return u.Message(true, "Account deleted")
}

func AccInfo(ctx context.Context) map[string]interface{} {
	acc := GetUser(ctx.Value("user").(uint))
	if acc == nil {
		return u.Message(false, "Account not found")
	}

	res := u.Message(true, "Account found")
	res["account"] = acc
	return res
}

func ModifyUsername(ctx context.Context, username string) map[string]interface{} {
	acc := GetUser(ctx.Value("user").(uint))
	if acc == nil {
		return u.Message(false, "Account not found")
	}

	temp := &Account{}

	if len(username) <= 4 || len(username) > 60 {
		return u.Message(false, "New username out of bounds")
	}

	err := GetDB().Where("username = ?", username).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error, try again")
	}
	if temp.Username != "" {
		if temp.Username == username {
			return u.Message(true, "Username updated")
		}
		return u.Message(false, "Username already in use!")
	}

	err = GetDB().Model(acc).Update("username", username).Error
	if err != nil {
		return u.Message(false, err.Error())
	}

	return u.Message(true, "Username updated")
}

func ModifyPassword(ctx context.Context, password string) map[string]interface{} {
	acc := GetUser(ctx.Value("user").(uint))
	if acc == nil {
		return u.Message(false, "Account not found")
	}

	if len(password) < 8 || len(password) > 60 {
		return u.Message(false, "New password out of bounds")
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln(err)
	}

	err = GetDB().Model(acc).Update("password", hashedPass).Error
	if err != nil {
		return u.Message(false, err.Error())
	}

	return u.Message(true, "Password updated")
}

func GetUser(u uint) *Account {
	acc := &Account{}
	GetDB().Where("id = ?", u).First(acc)
	if acc.Email == "" {
		return nil
	}

	acc.Password = ""
	return acc
}