package objects

type Ort struct {
	ID    int
	Stadt string
	Plz   string
}

func (o *Ort) GetOrtFrom(id string) (string, []interface{}) {
	query := "SELECT * FROM ort where id = ?"
	args := []interface{}{id}
	return query, args
}

func (o *Ort) GetAllOrte() (string, []interface{}) {
	query := "SELECT * FROM ort"
	args := []interface{}{}
	return query, args
}
