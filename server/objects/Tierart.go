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
	query := "SELECT * from tierart"
	args := []interface{}{}
	return query, args
}
