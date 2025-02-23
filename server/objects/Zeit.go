package objects

type Zeit struct {
	ID      int
	Uhrzeit string
}

func (z *Zeit) GetZeitFrom(id string) (string, []interface{}) {
	query := "Select * from zeit where id = ?"
	args := []interface{}{id}
	return query, args
}

func (z *Zeit) GetAllZeiten() (string, []interface{}) {
	query := "Select * from zeit"
	args := []interface{}{}
	return query, args
}

func (z *Zeit) InsertZeit() (string, []interface{}) {
	query := "Insert into zeit (uhr_zeit) values (?)"
	args := []interface{}{z.Uhrzeit}
	return query, args
}

func (z *Zeit) GetZeitFromUhrzeit(uhrzeit string) (string, []interface{}) {
	query := "Select * from zeit where uhr_zeit = ?"
	args := []interface{}{uhrzeit}
	return query, args
}

func (z *Zeit) CountZeit() (string, []interface{}) {
	query := `SELECT COUNT(*) FROM zeit`
	args := []interface{}{}
	return query, args
}
