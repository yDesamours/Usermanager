package dao

//queries string
const (
	insertUser        = `INSERT INTO users (firstname, lastname, username, password, role_id) VALUES ($1, $2, $3, $4, $5)`
	selectAllusers    = `SELECT firstname, lastname, username, role_name, created_on, is_active FROM users`
	updateUser        = `UPDATE users SET firstname=$1, lastname=$2 WHERE username=$3`
	adminUpdateUser   = `UPDATE users SET firstname=$1, lastname=$2, username=$3, role=$4, is_active=$5 WHERE username=$6`
	updatePassword    = `UPDATE users SET password=$1 WHERE username=$2`
	desactivateUser   = `UPDATE users set is_active=false WHERE username=$1`
	getUser           = `SELECT firstname, lastname, username, role_name, is_active, created_on FROM users INNER JOIN role USING(role_id) WHERE username=$1`
	selectRoleID      = `SELECT role_id FROM roles where role_name=$1`
	getALLPermissions = `SELECT permission_name FROM roles INNER JOIN role_permission USING(permission_id) INNER JOIN roles USING(role_id) where role_name = $1`
)
