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
