package dao

//queries string
const (
	insertUser        = `INSERT INTO users (firstname, lastname, username, password, role_id) VALUES ($1, $2, $3, $4, $5)`
	selectAllusers    = `SELECT firstname, lastname, username, role_name, created_on, is_active FROM users INNER JOIN roles USING(role_id)`
	updateUser        = `UPDATE users SET firstname=$1, lastname=$2, modified_by=(SELECT user_id FROM users WHERE username=$3) WHERE username=$3`
	adminUpdateUser   = `UPDATE users SET firstname=$1, lastname=$2, username=$3, role_id=(SELECT role_id FROM roles WHERE role_name=$4), modified_by=(SELECT user_id FROM users WHERE username=$6), is_active=$5 WHERE username=$7`
	updatePassword    = `UPDATE users SET password=$1 WHERE username=$2`
	desactivateUser   = `UPDATE users set is_active=false WHERE username=$1`
	getUser           = `SELECT firstname, lastname, username, role_name, is_active, created_on, password FROM users INNER JOIN roles USING(role_id) WHERE username=$1`
	selectRoleID      = `SELECT role_id FROM roles where role_name=$1`
	getALLPermissions = `SELECT permission_name FROM roles INNER JOIN role_permission USING(role_id) INNER JOIN permissions USING(permission_id) where role_name = $1`
)
