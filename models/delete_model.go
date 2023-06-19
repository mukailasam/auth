package models

func (dbExec *Model) DeleteUser(email string) error {
	sql := `DELETE FROM users WHERE email=$1`
	sql2 := `DELETE FROM verifications WHERE email=$1`

	tx, err := dbExec.PgDBConn.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare(sql)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(email)
	if err != nil {
		return err
	}

	stmt, err = tx.Prepare(sql2)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(email)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
