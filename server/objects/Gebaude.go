package objects

type Gebaude struct {
	ID     int
	Name   string
	Revier Revier
}

func (g *Gebaude) GetGebaeudeFrom(id string) (string, []interface{}) {
	query := `
		SELECT 
		    gebaude.id,
		    gebaude.name,
		    gebaude.revier_id,
		    revier.name,
		    revier.beschreibung
		FROM 
		    gebaude
		JOIN revier on gebaude.revier_id = revier.id 
		WHERE gebaude.id = ?
		`
	args := []interface{}{id}
	return query, args
}

func (g *Gebaude) GetAllGebaeude() (string, []interface{}) {
	query := `
		SELECT 
			gebaude.id,
		    gebaude.name,
		    gebaude.revier_id,
		    revier.name,
		    revier.beschreibung
		FROM 
		    gebaude
		JOIN revier on gebaude.revier_id = revier.id
		`
	args := []interface{}{}
	return query, args
}

func (g *Gebaude) InsertGebaeude(name string, revier_id int) (string, []interface{}) {
	query := "Insert into gebaude (name, revier_id) values (?,?)"
	args := []interface{}{name, revier_id}
	return query, args
}

func (g *Gebaude) CountGebaude() (string, []interface{}) {
	query := "SELECT COUNT(*) FROM gebaude"
	args := []interface{}{}
	return query, args
}

func (g *Gebaude) UpdateGebaude(id string) (string, []interface{}) {
	query := "Update gebaude set name = ?, revier_id = ? where id = ?"
	args := []interface{}{g.Name, g.Revier.ID, id}
	return query, args
}
