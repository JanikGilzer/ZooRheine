package objects

type Pfleger struct {
	ID            int
	Name          string
	Telefonnummer string
	Adresse       string
	Ort           Ort
	Revier        Revier
}

func (p *Pfleger) GetPflegerFrom(id string) (string, []interface{}) {
	query := `
		SELECT 
		    pfleger.id,
		    pfleger.name,
		    pfleger.telefonnummer,
		    pfleger.adresse,
		    pfleger.revier_id,
		    revier.name,
		    revier.beschreibung,
		    pfleger.ort_id,
		    ort.stadt,
		    ort.plz
		FROM 
		    pfleger
		JOIN ort on pfleger.ort_id = ort.id
		JOIN revier on pfleger.revier_id = revier.id 
		WHERE pfleger.id = ?
		`
	args := []interface{}{id}
	return query, args
}

func (p *Pfleger) GetAllPfleger() (string, []interface{}) {
	query := `
		SELECT 
			pfleger.id,
		    pfleger.name,
		    pfleger.telefonnummer,
		    pfleger.adresse,
		    pfleger.revier_id,
		    revier.name,
		    revier.beschreibung,
		    pfleger.ort_id,
		    ort.stadt,
		    ort.plz
		FROM 
		    pfleger 
		JOIN ort on pfleger.ort_id = ort.id 
		JOIN revier on pfleger.revier_id = revier.id
		`
	args := []interface{}{}
	return query, args
}

func (p *Pfleger) InsertPfleger(name string, telefonnummer string, adresse string, ort_id int, revier_id int) (string, []interface{}) {
	query := "INSERT INTO pfleger (name, telefonnummer, adresse, ort_id, revier_id) VALUES (?,?,?,?,?)"
	args := []interface{}{name, telefonnummer, adresse, ort_id, revier_id}
	return query, args
}

func (p *Pfleger) UpdatePfleger(id int, name string, telefonnummer string, adresse string, ort_id int, revier_id int) (string, []interface{}) {
	query := "UPDATE pfleger SET name =  ? , telefonnummer = ? , adresse = ? , ort_id = ? , revier_id = ? WHERE id = ?"
	args := []interface{}{name, telefonnummer, adresse, ort_id, revier_id, id}
	return query, args
}
