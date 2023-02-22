package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"blog-application/api/auth"
	"blog-application/api/models"
	"blog-application/api/responses"
	"blog-application/api/utils/formaterror"

	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	// read input from client
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// creates a new empty models.User struct and unmarshals the request body into it using the json.Unmarshal function
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// trim any whitespace and convert the email address to lowercase
	user.Prepare()

	// check that the email address and password meet the required validation criteria.
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// generate token
	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignIn(email, password string) (string, error) {

	var err error

	user := models.User{}

	// query the database for a User record that matches the given email parameter. Take method to retrieve a single record that matches the query and assign to user variable
	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}

	// check if the password provided matches the password stored in the user.Password field
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.ID)
}
