package objects

type FuetterungsZeiten struct {
	ID      int
	Zeit    Zeit
	Gebaude Gebaude
}

func (f *FuetterungsZeiten) GetFutterZeitFrom(id string) (string, []interface{}) {
	query := `
		SELECT 
		    fuetterungszeit.id,
		    fuetterungszeit.zeit_id,
		    zeit.uhr_zeit,
		    fuetterungszeit.gebaude_id,
		    gebaude.name,
		    gebaude.revier_id
		FROM
		    fuetterungszeit
		JOIN zeit ON fuetterungszeit.zeit_id = zeit.id
		JOIN gebaude ON fuetterungszeit.gebaude_id = gebaude.id
		where fuetterungszeit.id = ?
		`

	args := []interface{}{id}
	return query, args
}

func (f *FuetterungsZeiten) AllFromFuetterungsZeiten() (string, []interface{}) {
	query := `
		SELECT 
			fuetterungszeit.id,
		    fuetterungszeit.zeit_id,
		    zeit.uhr_zeit,
		    fuetterungszeit.gebaude_id,
		    gebaude.name,
		    gebaude.revier_id
		FROM
		    fuetterungszeit
		JOIN zeit ON fuetterungszeit.zeit_id = zeit.id
		JOIN gebaude ON fuetterungszeit.gebaude_id = gebaude.id
		`
	args := []interface{}{}
	return query, args
}

func (f *FuetterungsZeiten) CountFuetterungsZeiten() (string, []interface{}) {
	query := `SELECT COUNT(*) FROM fuetterungszeit`
	args := []interface{}{}
	return query, args
}

func (f *FuetterungsZeiten) InsertFuetterungsZeiten(zeit_id int, gebaude_id int) (string, []interface{}) {
	query := `INSERT INTO fuetterungszeit (zeit_id, gebaude_id) VALUES (?, ?)`
	args := []interface{}{zeit_id, gebaude_id}
	return query, args
}

func (f *FuetterungsZeiten) DeleteFuetterungsZeiten(gebaude_id int) (string, []interface{}) {
	query := `DELETE FROM fuetterungszeit WHERE gebaude_id = ?;`
	args := []interface{}{gebaude_id}
	return query, args
}
