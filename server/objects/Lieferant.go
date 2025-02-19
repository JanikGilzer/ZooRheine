package objects

type Lieferant struct {
	ID      int
	Name    string
	Ort     Ort
	Adresse string
}

func (l *Lieferant) GetLieferantFrom(id int) (string, interface{}) {
	query := "SELECT * from lieferant where id= ?"
	args := []interface{}{id}
	return query, args
}
