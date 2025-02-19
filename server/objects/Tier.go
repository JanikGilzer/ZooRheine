package objects

type Tier struct {
	ID           int
	Name         string
	Tierart      TierArt
	Geburtsdatum string
	Gebaude      Gebaude
}

func (t *Tier) GetTierFrom(id string) (string, []interface{}) {
	query := `
		SELECT *
		FROM 
			tier
		JOIN tierart ON tier.tierart_id = tierart.id
		JOIN gebaude ON tier.gebaude_id = gebaude.id
		JOIN revier ON gebaude.revier_id = revier.id
		WHERE tier.id = ?
	`
	args := []interface{}{id}
	return query, args
}

func (t *Tier) GetAllTiere() (string, []interface{}) {
	query := `
		SELECT *
		FROM 
			tier
		JOIN tierart ON tier.tierart_id = tierart.id
		JOIN gebaude ON tier.gebaude_id = gebaude.id
		JOIN revier ON gebaude.revier_id = revier.id
	`
	args := []interface{}{} // No arguments needed for this query
	return query, args
}

func (t *Tier) CountTiere() (string, []interface{}) {
	query := "SELECT count(*) FROM tier"
	args := []interface{}{} // No arguments needed for this query
	return query, args
}

func (t *Tier) InsertTier(name string, geburtstag string, tierart_id int, gebaude int) (string, []interface{}) {
	query := "INSERT INTO tier (name, geburtstag, tierart_id, gebaude_id) VALUES (?,?,?,?)"
	args := []interface{}{name, geburtstag, tierart_id, gebaude}
	return query, args
}

func (t *Tier) UpdateTier(id int, name string, geburtstag string, tierart_id int, gebaude int) (string, []interface{}) {
	query := "UPDATE tier SET name = ?, geburtstag = ?, tierart_id = ?, gebaude_id = ? WHERE id = ?"
	args := []interface{}{name, geburtstag, tierart_id, gebaude, id}
	return query, args
}
