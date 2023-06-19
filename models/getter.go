package models

import (
	"github.com/gofrs/uuid"
)

type User struct {
	UserID   string
	Username string
	Email    string
	Password string
	Salt     string
	Role     string
	Verified bool
}

func (dbExec *Model) GetUserByEmail(email string) (*User, error) {
	usr := User{}

	sql := `SELECT user_id, username, email, password, salt, role, verified FROM users WHERE email=$1`

	stmt, err := dbExec.PgDBConn.Prepare(sql)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(email).Scan(&usr.UserID, &usr.Username, &usr.Email, &usr.Password, &usr.Salt, &usr.Role, &usr.Verified)
	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return &usr, nil
}

func (dbExec *Model) GetUserByUsername(username string) (*User, error) {
	usr := User{}

	sql := `SELECT user_id, username, email, password, salt, role, verified FROM users WHERE username=$1`

	stmt, err := dbExec.PgDBConn.Prepare(sql)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(username).Scan(&usr.UserID, &usr.Username, &usr.Email, &usr.Password, &usr.Salt, &usr.Role, &usr.Verified)
	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return &usr, nil
}

func (dbExec *Model) GetUserIDByUsername(username string) (*uuid.UUID, error) {
	var userID *uuid.UUID

	sql := `SELECT user_id FROM users WHERE username=$1`

	stmt, err := dbExec.PgDBConn.Prepare(sql)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(username).Scan(&userID)
	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return userID, nil
}

func (dbExec *Model) GetUserEmailByUsername(username string) (*string, error) {
	var email string

	sql := `SELECT email FROM users WHERE username=$1`

	stmt, err := dbExec.PgDBConn.Prepare(sql)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(username).Scan(&email)
	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return &email, nil
}

func (dbExec *Model) GetUserEmailByUserID(usrID string) (*string, error) {
	var userID string

	sql := `SELECT email FROM users WHERE user_id=$1`

	stmt, err := dbExec.PgDBConn.Prepare(sql)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(usrID).Scan(&userID)
	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return &userID, nil
}

func (dbExec *Model) GetUserVerificationStatusByUsername(username string) (*bool, error) {
	var verified bool

	sql := `SELECT verified FROM users WHERE username=$1`

	stmt, err := dbExec.PgDBConn.Prepare(sql)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(username).Scan(&verified)
	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return &verified, nil
}
