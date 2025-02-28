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
		SELECT 
		    tier.id,
		    tier.name,
		    tier.geburtstag,
		    tier.gebaude_id,
			gebaude.name,
			gebaude.revier_id,
			revier.name,
			revier.beschreibung,
			tier.tierart_ID,
			tierart.name
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
		SELECT 
			tier.id,
		    tier.name,
		    tier.geburtstag,
		    tier.gebaude_id,
			gebaude.name,
			gebaude.revier_id,
			revier.name,
			revier.beschreibung,
			tier.tierart_ID,
			tierart.name
		FROM 
			tier
		JOIN tierart ON tier.tierart_id = tierart.id
		JOIN gebaude ON tier.gebaude_id = gebaude.id
		JOIN revier ON gebaude.revier_id = revier.id
		ORDER BY tier.id
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
	query := "INSERT INTO tier (name, geburtstag, gebaude_id, tierart_ID) VALUES (?,?,?,?)"
	args := []interface{}{name, geburtstag, gebaude, tierart_id}
	return query, args
}

func (t *Tier) UpdateTier(id int, name string, geburtstag string, gebaude int) (string, []interface{}) {
	query := "UPDATE tier SET name = ?, geburtstag = ?, gebaude_id = ? WHERE id = ?"
	args := []interface{}{name, geburtstag, gebaude, id}
	return query, args
}
