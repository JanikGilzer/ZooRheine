package core

type User struct {
	ID     int
	Name   string
	Passwd string
	Role   Role
}

func (u *User) Verify(name string, password string) (string, []interface{}) {
	query := `
	SELECT
		user.id,
		user.name,
		user.password,
		user.gruppen_id
	FROM
		user
	JOIN gruppe ON user.gruppen_id = gruppe.id
	WHERE
		user.name = ? 
		AND user.password = ?;
	`
	args := []interface{}{name, password}
	return query, args
}
