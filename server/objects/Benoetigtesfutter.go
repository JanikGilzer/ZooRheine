package objects

type BenoetigtesFutter struct {
	ID     int
	Tier   Tier
	Futter Futter
}

func (b *BenoetigtesFutter) GetBenoetigtesFutterFrom(id string) (string, []interface{}) {
	query := `
		SELECT 
		    benoetigtesfutter.id AS benoetigtesfutter_id,
			benoetigtesfutter.tier_id AS tier_id,
			benoetigtesfutter.futter_id AS futter_id,
			futter.name AS futter_name,
			futter.lieferant_id AS futter_lieferant_id,
			tier.name AS tier_name,
			tier.geburtstag AS tier_geburtstag,
			tier.gebaude_id AS tier_gebaude_id,
			tier.tierart_id AS tier_tierart_id
		FROM 
			benoetigtesfutter
		JOIN futter ON benoetigtesfutter.futter_id = futter.id
		JOIN tier ON benoetigtesfutter.tier_id = tier.id
		WHERE benoetigtesfutter.id = ?
		`
	args := []interface{}{id}
	return query, args
}

func (b *BenoetigtesFutter) GetAllBenoetigtesFutter() (string, []interface{}) {
	query := `
		SELECT 
		    benoetigtesfutter.id AS benoetigtesfutter_id,
			benoetigtesfutter.tier_id AS tier_id,
			benoetigtesfutter.futter_id AS futter_id,
			futter.name AS futter_name,
			futter.lieferant_id AS futter_lieferant_id,
			tier.name AS tier_name,
			tier.geburtstag AS tier_geburtstag,
			tier.gebaude_id AS tier_gebaude_id,
			tier.tierart_id AS tier_tierart_id
		FROM 
			benoetigtesfutter
		JOIN futter ON benoetigtesfutter.futter_id = futter.id
		JOIN tier ON benoetigtesfutter.tier_id = tier.id
					`
	args := []interface{}{}
	return query, args
}

func (b *BenoetigtesFutter) InsertBenoetigtesFutter(tier_id int, futter_id int) (string, []interface{}) {
	query := "INSERT INTO benoetigtesfutter (tier_id, futter_id) VALUES (?,?)"
	args := []interface{}{tier_id, futter_id}
	return query, args
}

func (b *BenoetigtesFutter) DeleteBenoetigtesFutterWhereTier(id int) (string, []interface{}) {
	query := "DELETE FROM benoetigtesfutter WHERE id = ?"
	args := []interface{}{id}
	return query, args
}
