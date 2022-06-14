package dao

//allows to change a password
//accept the new password and the username as argument
func EditPassword(newPassword, username string) error {
	db := GetDB()

	//set username to lowewrcase
	result, err := db.Exec(updatePassword, newPassword, username)
	if err != nil {
		return err
	}

	if affected, err := result.RowsAffected(); affected == 0 {
		return err
	}

	return nil
}
