package models

type user struct {
	username string
	email    string
}

func (dbExec *Model) UsernameExists(username string) error {
	usr := user{}

	sql := "SELECT username FROM users WHERE username=$1"

	stmt, err := dbExec.PgDBConn.Prepare(sql)
	if err != nil {
		return err
	}
	err = stmt.QueryRow(username).Scan(&usr.username)
	if err != nil {
		return err
	}

	err = stmt.Close()
	if err != nil {
		return err
	}

	return nil
}

func (dbExec *Model) EmailExists(email string) error {
	usr := user{}

	sql := "SELECT email FROM users WHERE email=$1"

	stmt, err := dbExec.PgDBConn.Prepare(sql)
	if err != nil {
		return err
	}
	err = stmt.QueryRow(email).Scan(&usr.email)
	if err != nil {
		return err
	}

	return nil
}
