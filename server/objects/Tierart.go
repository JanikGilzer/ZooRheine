package objects

type TierArt struct {
	ID   int
	Name string
}

func (t *TierArt) GetTierArtFrom(id string) (string, []interface{}) {
	query := "SELECT * from tierart where id=?"
	args := []interface{}{id}
	return query, args
}

func (t *TierArt) GetAllTierArt() (string, []interface{}) {
	query := "SELECT * FROM tierart"
	args := []interface{}{}
	return query, args
}

func (t *TierArt) InsertTierArt() (string, []interface{}) {
	query := "Insert INTO tierart (name) values (?)"
	args := []interface{}{t.Name}
	return query, args
}
