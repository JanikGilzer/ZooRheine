package objects

type Revier struct {
	ID           int
	Name         string
	Beschreibung string
}

func (r *Revier) GetRevierFrom(id string) (string, []interface{}) {
	query := "Select * from revier where id = ?"
	args := []interface{}{id}
	return query, args
}

func (r *Revier) GetAllReviere() (string, []interface{}) {
	query := "Select * from revier"
	args := []interface{}{}
	return query, args
}

func (r *Revier) CountReviere() (string, []interface{}) {
	query := "Select count(*) from revier"
	args := []interface{}{}
	return query, args
}

func (r *Revier) InsertRevier(name string, beschreibung string) (string, []interface{}) {
	query := "Insert into revier (name, beschreibung) values (?,?)"
	args := []interface{}{name, beschreibung}
	return query, args
}
