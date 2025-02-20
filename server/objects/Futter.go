package objects

type Futter struct {
	ID        int
	Name      string
	Lieferant Lieferant
}

func (f *Futter) GetFutterFrom(id string) (string, []interface{}) {
	query := `
		SELECT 
		    futter.id,
		    futter.name,
		    futter.lieferant_id,
		    lieferant.name,
		    lieferant.adresse,
		    lieferant.ort_id,
			ort.stadt,
			ort.plz
		FROM 
		    futter
		JOIN lieferant on futter.lieferant_id = lieferant.id
		Join ort on lieferant.ort_id = ort.id
	  	WHERE futter.id = ?
	  	`
	args := []interface{}{id}
	return query, args
}

func (f *Futter) GetAllFutter() (string, []interface{}) {
	query := `
		SELECT 
		    futter.id,
		    futter.name,
		    futter.lieferant_id,
		    lieferant.name,
		    lieferant.adresse,
		    lieferant.ort_id,
			ort.stadt,
			ort.plz
		FROM 
		    futter
		JOIN lieferant on futter.lieferant_id = lieferant.id
		Join ort on lieferant.ort_id = ort.id
		`
	args := []interface{}{}
	return query, args
}

func (f *Futter) GetFutterFromName(name string) (string, []interface{}) {
	query := `
		SELECT 
		    futter.id,
		    futter.name,
		    futter.lieferant_id,
		    lieferant.name,
		    lieferant.adresse,
		    lieferant.ort_id,
			ort.stadt,
			ort.plz
		FROM 
		    futter
		JOIN lieferant on futter.lieferant_id = lieferant.id
		Join ort on lieferant.ort_id = ort.id
	  	WHERE futter.name = ?
	  	`
	args := []interface{}{name}
	return query, args
}

func (f *Futter) InsertFutter(name string, lieferant_id int) (string, []interface{}) {
	query := "Insert into futter (name, lieferant_id) values (?,?)"
	args := []interface{}{name, lieferant_id}
	return query, args
}
