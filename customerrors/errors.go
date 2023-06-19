package customerrors

import (
	"errors"
)

var (

	// general error message for bad request
	BadRequest = errors.New("Bad json request")

	// email error message
	InvalidEmail = errors.New("Invalid email")

	// username error message
	InvalidUsername  = errors.New("Invalid Username, Username should be greater than 6 and less than 20")
	AllFieldRequired = errors.New("All field Required")

	// password error messages
	InvalidPassword = errors.New("password should be minimum 8 in length and Password should contain at least a single uppercase letter, lowercase letter, single digit and a special character")

	// server error message
	InternalServerError = errors.New("Sorry, something went wrong on our end, try again later")

	// user error messages
	EmailExist = errors.New("Email exists, we already have an accounted linked to this email")
	UserExist  = errors.New("Username exists, we already have an accounted linked to this username")

	// session
	SessionExistsError  = errors.New("Session already exists")
	InvalidSessionError = errors.New("Invalid session")
)
