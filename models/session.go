package models

import (
	"fmt"
	"net/http"

	"github.com/ftsog/auth/customerrors"
)

func (dbExec *Model) CreateSession(w http.ResponseWriter, r *http.Request, user string) error {
	session, err := dbExec.RediStore.Get(r, "session_id")
	if err != nil {
		return err
	}

	if session.Values["user"] != nil {
		return customerrors.SessionExistsError
	}

	session.Options.Path = "/"
	session.Values["user"] = user

	err = session.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}

func (dbExec *Model) DeleteSession(w http.ResponseWriter, r *http.Request) error {
	session, err := dbExec.RediStore.Get(r, "session_id")
	if err != nil {
		return err
	}

	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}

func (dbExec Model) CheckSession(w http.ResponseWriter, r *http.Request) (*bool, error) {
	right := true
	wrong := false
	session, err := dbExec.RediStore.Get(r, "session_id")
	if err != nil {
		return nil, err
	}

	_, ok := session.Values["user"]
	if !ok {
		return &wrong, customerrors.InvalidSessionError
	}

	return &right, nil

}

func (dbExec *Model) GetUserFromSession(w http.ResponseWriter, r *http.Request) (*string, error) {
	session, err := dbExec.RediStore.Get(r, "session_id")
	if err != nil {
		return nil, err
	}

	user, ok := session.Values["user"]
	if !ok {
		return nil, nil
	}

	usr := fmt.Sprintf("%v", user)

	return &usr, nil
}
