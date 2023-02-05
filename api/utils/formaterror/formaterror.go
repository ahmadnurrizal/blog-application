package formaterror

import (
	"errors"
	"strings"
)

// format error messages returned from the database
func FormatError(err string) error {

	if strings.Contains(err, "nickname") {
		return errors.New("Nickname Already Taken")
	}

	if strings.Contains(err, "email") {
		return errors.New("Email Already Taken")
	}

	if strings.Contains(err, "title") {
		return errors.New("Title Already Taken")
	}
	if strings.Contains(err, "hashedPassword") {
		return errors.New("Incorrect Password")
	}

	// If the err does not contain any of these keywords, it returns a new error with the message "Incorrect Details"
	return errors.New("Incorrect Details")
}
