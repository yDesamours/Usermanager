package dao

//this function return the id of a role
func GetRoleID(role_name string) string {
	db := GetDB()
	var role_ID string
	err := db.QueryRow(selectRoleID, role_name).Scan(&role_ID)
	if err != nil {
		return "role doesn't exists"
	}

	return role_ID
}

//this function return all the permissions attached to a role
func GetRolePermissions(roleName string) ([]string, error) {
	db := GetDB()
	var permissions []string
	rows, err := db.Query(getALLPermissions, roleName)

	if err != nil {
		return permissions, err
	}
	defer rows.Close()

	for rows.Next() {
		var permission string

		rows.Scan(&permission)
		permissions = append(permissions, permission)
	}

	return permissions, nil
}
