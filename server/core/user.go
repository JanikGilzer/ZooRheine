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
		mitarbeiter.id,
		mitarbeiter.name,
		mitarbeiter.gruppen_id,
		gruppe.name
	FROM
		mitarbeiter
	JOIN gruppe ON mitarbeiter.gruppen_id = gruppe.id
	WHERE
		mitarbeiter.name = ? 
		AND mitarbeiter.password = ?;
	`
	args := []interface{}{name, password}
	return query, args
}
