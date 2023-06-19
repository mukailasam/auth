package models

type verification struct {
	Code    string
	Expired bool
}

func (dbExec *Model) VerifyUser(username string) error {
	sql := `UPDATE users SET verified=true WHERE username=$1`
	sql2 := `UPDATE verifications SET expired=true WHERE email=(SELECT email FROM users WHERE username=$1)`

	tx, err := dbExec.PgDBConn.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare(sql)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(username)
	if err != nil {
		return err
	}

	stmt, err = tx.Prepare(sql2)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(username)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (dbExec *Model) ReadVerifications(username string) (*verification, error) {
	vr := verification{}

	sql := `SELECT code, expired FROM verifications WHERE email=(SELECT email FROM users WHERE username=$1)`

	stmt, err := dbExec.PgDBConn.Prepare(sql)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(username).Scan(&vr.Code, &vr.Expired)
	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return &vr, nil
}

func (dbExec *Model) IsVerified(username string) (*bool, error) {
	var isVerify bool

	sql := `SELECT verified FROM users WHERE username=$1`

	stmt, err := dbExec.PgDBConn.Prepare(sql)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(username).Scan(&isVerify)
	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return &isVerify, nil

}

/*
func (dbExec *Model) SetEmailVerification(verficationID, email, code string) error {
	sql := `INSERT INTO email_verifications(verification_id, email, code) VALUES($1, $2, $3)`

	stmt, err := dbExec.PgDBConn.Prepare(sql)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(verficationID, email, code)
	if err != nil {
		return err
	}

	err = stmt.Close()
	if err != nil {
		return err
	}

	return nil

}
*/
