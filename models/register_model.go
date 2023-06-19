package models

func (dbExec *Model) CreateUser(userID, verificationID, username, email, password, salt, verificationCode string) error {
	sql := `INSERT INTO users(user_id, username, email, password, salt) VALUES($1, $2, $3, $4, $5)`
	sql2 := `INSERT INTO verifications(verification_id, email, code) VALUES($1, $2, $3)`

	tx, err := dbExec.PgDBConn.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare(sql)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(userID, username, email, password, salt)
	if err != nil {
		return err
	}

	// Second Transaction
	stmt, err = tx.Prepare(sql2)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(verificationID, email, verificationCode)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

/*
func (dbExec *Model) CreateUser(userID, username, email, password, salt string) error {
	sql := `INSERT INTO users(user_id, username, email, password, salt) VALUES($1, $2, $3, $4, $5)`

	stmt, err := dbExec.PgDBConn.Prepare(sql)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(userID, username, email, password, salt)
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
