package core

type User struct {
	ID     int
	Name   string
	Passwd string
	Role   Role
}

func (u *User) Verify(name string) (string, []interface{}) {
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
	`
	args := []interface{}{name}
	return query, args
}

func (u *User) GetHashFrom(name string) (string, []interface{}) {
	query := `
	SELECT
    	mitarbeiter.password
	FROM
	    mitarbeiter
	where 
	    mitarbeiter.name = ?
	`
	args := []interface{}{name}
	return query, args
}
