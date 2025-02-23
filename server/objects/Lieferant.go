package objects

type Lieferant struct {
	ID      int
	Name    string
	Adresse string
	Ort     Ort
}

func (l *Lieferant) GetLieferantFrom(id int) (string, interface{}) {
	query := "SELECT * from lieferant where id= ?"
	args := []interface{}{id}
	return query, args
}

func (l *Lieferant) GetAllLieferant() (string, []interface{}) {
	query := `
		SELECT
    		lieferant.id,
    		lieferant.name,
    		lieferant.adresse,
    		lieferant.ort_id,
    		ort.stadt,
    		ort.plz
		FROM
		    lieferant
		JOIN ort on lieferant.ort_id = ort.id
		`
	args := []interface{}{}
	return query, args
}

func (l *Lieferant) InsertLieferant() (string, []interface{}) {
	query := "INSERT INTO lieferant (name, adresse, ort_id) VALUES (?, ?,?)"
	args := []interface{}{l.Name, l.Adresse, l.Ort.ID}
	return query, args
}
