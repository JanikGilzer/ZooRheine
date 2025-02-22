package server

import (
	"ZooDaBa/server/core"
	"ZooDaBa/server/objects"
	"fmt"
	"log"
	"strconv"
)

// #region Tier

func GetTier(db core.DB_Handler, id string) objects.Tier {
	var t objects.Tier

	query, args := t.GetTierFrom(id)
	fmt.Println(query)
	var tRow = db.QueryRow(query, args...)

	err := tRow.Scan(
		&t.ID,                          // tier_id
		&t.Name,                        // tier_name
		&t.Geburtsdatum,                // tier_geburtstag
		&t.Gebaude.ID,                  // tier_gebaude_id
		&t.Gebaude.Name,                // gebaude_name
		&t.Gebaude.Revier.ID,           // gebaude_revier_id
		&t.Gebaude.Revier.Name,         // revier_name
		&t.Gebaude.Revier.Beschreibung, // revier_beschreibung
		&t.Tierart.ID,
		&t.Tierart.Name,
	)
	if err != nil {
		core.Logger.Error("Error getting tier from DB", "tier id:", &t.ID, "err", err)
	}
	if err := tRow.Err(); err != nil {
		core.Logger.Error("Error encountered during test iterating", "tier id:", id, "err", err)
	}
	return t
}

func GetAllTiere(db core.DB_Handler) []objects.Tier {
	var tiereArr []objects.Tier
	var tier objects.Tier

	query, args := tier.GetAllTiere()
	var tRows = db.Query(query, args...)

	for tRows.Next() {
		var t objects.Tier
		err := tRows.Scan(
			&t.ID,                          // tier_id
			&t.Name,                        // tier_name
			&t.Geburtsdatum,                // tier_geburtstag
			&t.Gebaude.ID,                  // tier_gebaude_id
			&t.Gebaude.Name,                // gebaude_name
			&t.Gebaude.Revier.ID,           // gebaude_revier_id
			&t.Gebaude.Revier.Name,         // revier_name
			&t.Gebaude.Revier.Beschreibung, // revier_beschreibung
			&t.Tierart.ID,
			&t.Tierart.Name,
		)
		if err != nil {
			core.Logger.Error("Error getting all tiere from DB", "tier id:", &t.ID, "err", err)
		}
		tiereArr = append(tiereArr, t)
	}

	if err := tRows.Err(); err != nil {
		core.Logger.Error("Error encountered during test iterating", "Object", "Tier", "err", err)
	}
	return tiereArr
}

func CreateTier(db core.DB_Handler, tier objects.Tier, futter []objects.Futter) {
	query, args := tier.InsertTier(tier.Name, tier.Geburtsdatum, tier.Tierart.ID, tier.Gebaude.ID)
	db.Exec(query, args...)
	count := CountTiere(db)
	println(count)
	t := GetTier(db, strconv.Itoa(count))
	fmt.Println(t)
	var bFutter objects.BenoetigtesFutter
	for _, f := range futter {
		fmt.Println(f.ID)
		fmt.Println(t.ID)
		query, args := bFutter.InsertBenoetigtesFutter(t.ID, f.ID)
		fmt.Println(query)
		fmt.Println(args)
		db.Exec(query, args...)
	}
}

func UpdateTier(db core.DB_Handler, tier objects.Tier) {
	query, args := tier.UpdateTier(tier.ID, tier.Name, tier.Geburtsdatum, tier.Gebaude.ID)
	db.Exec(query, args...)
}

func CountTiere(db core.DB_Handler) int {
	var tier objects.Tier
	query, args := tier.CountTiere()
	var tRow = db.QueryRow(query, args...)
	count := 0
	err := tRow.Scan(&count)
	if err != nil {
		core.Logger.Error("Error getting tier count", "err", err)
	}
	return count
}

// #endregion

// #region Revier

func GetRevier(db core.DB_Handler, id string) objects.Revier {
	var revier objects.Revier
	query, args := revier.GetRevierFrom(id)
	var rRow = db.QueryRow(query, args...)
	err := rRow.Scan(
		&revier.ID,
		&revier.Name,
		&revier.Beschreibung,
	)
	if err != nil {
		core.Logger.Error("Error getting revier from DB", "revier id:", id, "err", err)
	}
	if err := rRow.Err(); err != nil {
		core.Logger.Error("Error encountered during test iterating", "revier id:", id, "err", err)
	}
	return revier
}

func GetAllReviere(db core.DB_Handler) []objects.Revier {
	var reviereArr []objects.Revier
	var revier objects.Revier
	query, args := revier.GetAllReviere()
	var rRows = db.Query(query, args...)
	for rRows.Next() {
		var r objects.Revier
		err := rRows.Scan(
			&r.ID,
			&r.Name,
			&r.Beschreibung,
		)
		if err != nil {
			core.Logger.Error("Error getting revier from DB", "revier id:", r, "err", err)
		}
		reviereArr = append(reviereArr, r)
	}
	if err := rRows.Err(); err != nil {
		core.Logger.Error("Error encountered during test iterating", "Object", "Revier", "err", err)
	}
	return reviereArr
}

func CountReviere(db core.DB_Handler) int {
	var revier objects.Revier
	query, args := revier.CountReviere()
	var rRow = db.QueryRow(query, args...)
	var count int
	rRow.Scan(&count)
	return count
}

// #endregion

// #region FutterrungsZeit

func GetFutterZeit(db core.DB_Handler, id string) objects.FuetterungsZeiten {
	var fz objects.FuetterungsZeiten
	query, args := fz.GetFutterZeitFrom(id)
	fmt.Println(query)
	var fRow = db.QueryRow(query, args...)
	fRow.Scan(
		&fz.ID,
		&fz.Zeit.ID,
		&fz.Zeit.Uhrzeit,
		&fz.Gebaude.ID,
		&fz.Gebaude.Name,
		&fz.Gebaude.Revier.ID,
	)
	return fz
}

func GetAllFutterZeiten(db core.DB_Handler) []objects.FuetterungsZeiten {
	var fzArr []objects.FuetterungsZeiten
	var fz objects.FuetterungsZeiten

	query, args := fz.AllFromFuetterungsZeiten()
	var fRows = db.Query(query, args...)
	for fRows.Next() {
		var fz objects.FuetterungsZeiten
		fRows.Scan(
			&fz.ID,
			&fz.Zeit.ID,
			&fz.Zeit.Uhrzeit,
			&fz.Gebaude.ID,
			&fz.Gebaude.Name,
			&fz.Gebaude.Revier.ID,
		)
		fzArr = append(fzArr, fz)
	}
	return fzArr
}

// #endregion

// #region Tier Art

func GetTierArt(db core.DB_Handler, id string) objects.TierArt {
	var t objects.TierArt
	query, args := t.GetTierArtFrom(id)
	var tRow = db.QueryRow(query, args...)
	err := tRow.Scan(
		&t.ID,
		&t.Name,
	)
	if err != nil {
		fmt.Println(err)
	}
	return t
}

func GetAllTierArt(db core.DB_Handler) []objects.TierArt {
	var tiArr []objects.TierArt
	var tier objects.TierArt
	query, args := tier.GetAllTierArt()
	var tRows = db.Query(query, args...)
	for tRows.Next() {
		var t objects.TierArt
		tRows.Scan(
			&t.ID,
			&t.Name,
		)
		tiArr = append(tiArr, t)
	}
	return tiArr
}

// #endregion

// #region BenoetigtesFutter

func GetBenoetigtesFutter(db core.DB_Handler, id string) objects.BenoetigtesFutter {
	var b objects.BenoetigtesFutter

	// Build the query and arguments
	query, args := b.GetBenoetigtesFutterFrom(id)
	fmt.Println(query)

	// Execute the query
	var bRow = db.QueryRow(query, args...)

	// Scan the result into the struct
	err := bRow.Scan(
		&b.ID,
		&b.Tier.ID,
		&b.Futter.ID,
		&b.Futter.Name,
		&b.Futter.Lieferant.ID,
		&b.Tier.Name,
		&b.Tier.Geburtsdatum,
		&b.Tier.Gebaude.ID,
		&b.Tier.Tierart.ID,
	)
	if err != nil {
		log.Fatal("Error scanning row:", err)
	}

	return b
}

func GetAllBenoetigtesFutter(db core.DB_Handler) []objects.BenoetigtesFutter {
	var benoetigtesFutter []objects.BenoetigtesFutter
	var b objects.BenoetigtesFutter
	query, args := b.GetAllBenoetigtesFutter()
	var bRows = db.Query(query, args...)
	for bRows.Next() {
		var b objects.BenoetigtesFutter
		bRows.Scan(
			&b.ID,
			&b.Tier.ID,
			&b.Futter.ID,
			&b.Futter.Name,
			&b.Futter.Lieferant.ID,
			&b.Tier.Name,
			&b.Tier.Geburtsdatum,
			&b.Tier.Gebaude.ID,
			&b.Tier.Tierart.ID,
		)
		benoetigtesFutter = append(benoetigtesFutter, b)
	}
	return benoetigtesFutter
}

// #endregion

// #region Zeit
func GetZeit(db core.DB_Handler, id string) objects.Zeit {
	var zeit objects.Zeit
	query, args := zeit.GetZeitFrom(id)
	var zRow = db.QueryRow(query, args...)
	zRow.Scan(
		&zeit.ID,
		&zeit.Uhrzeit,
	)
	return zeit
}

func GetAllZeiten(db core.DB_Handler) []objects.Zeit {
	var zeitenArr []objects.Zeit
	var zeit objects.Zeit
	query, args := zeit.GetAllZeiten()
	var zRows = db.Query(query, args...)
	for zRows.Next() {
		var z objects.Zeit
		zRows.Scan(
			&z.ID,
			&z.Uhrzeit,
		)
		zeitenArr = append(zeitenArr, z)
	}
	return zeitenArr
}

// #endregion

// #region Gebaude

func GetGebaeude(db core.DB_Handler, id string) objects.Gebaude {
	var gebaude objects.Gebaude
	query, args := gebaude.GetGebaeudeFrom(id)
	var gRow = db.QueryRow(query, args...)
	gRow.Scan(
		&gebaude.ID,
		&gebaude.Name,
		&gebaude.Revier.ID,
		&gebaude.Revier.Name,
		&gebaude.Revier.Beschreibung,
	)
	return gebaude
}

func GetAllGebaude(db core.DB_Handler) []objects.Gebaude {
	var gebaudeArr []objects.Gebaude
	var gebaude objects.Gebaude
	query, args := gebaude.GetAllGebaeude()
	var gRows = db.Query(query, args...)
	for gRows.Next() {
		var g objects.Gebaude
		gRows.Scan(
			&g.ID,
			&g.Name,
			&g.Revier.ID,
			&g.Revier.Name,
			&g.Revier.Beschreibung,
		)
		gebaudeArr = append(gebaudeArr, g)
	}
	return gebaudeArr
}

// #endregion

// #region Futter

func GetFutter(db core.DB_Handler, id string) objects.Futter {
	var futter objects.Futter
	query, args := futter.GetFutterFrom(id)
	var fRow = db.QueryRow(query, args...)
	fRow.Scan(
		&futter.ID,
		&futter.Name,
		&futter.Lieferant.ID,
		&futter.Lieferant.Name,
		&futter.Lieferant.Adresse,
		&futter.Lieferant.Ort.ID,
		&futter.Lieferant.Ort.Stadt,
		&futter.Lieferant.Ort.Plz,
	)
	return futter
}

func GetFutterFromName(db core.DB_Handler, name string) objects.Futter {
	var futter objects.Futter
	query, args := futter.GetFutterFromName(name)
	var fRow = db.QueryRow(query, args...)
	fRow.Scan(
		&futter.ID,
		&futter.Name,
		&futter.Lieferant.ID,
		&futter.Lieferant.Name,
		&futter.Lieferant.Adresse,
		&futter.Lieferant.Ort.ID,
		&futter.Lieferant.Ort.Stadt,
		&futter.Lieferant.Ort.Plz,
	)
	return futter
}

func CountGebaude(db core.DB_Handler) int {
	var gebaude objects.Gebaude
	query, args := gebaude.CountGebaude()
	var gRow = db.QueryRow(query, args...)
	var count int
	gRow.Scan(&count)
	return count
}

func GetAllFutter(db core.DB_Handler) []objects.Futter {
	var futterArr []objects.Futter
	var futter objects.Futter
	query, args := futter.GetAllFutter()
	var fRows = db.Query(query, args...)
	for fRows.Next() {
		var futter objects.Futter
		fRows.Scan(
			&futter.ID,
			&futter.Name,
			&futter.Lieferant.ID,
			&futter.Lieferant.Name,
			&futter.Lieferant.Adresse,
			&futter.Lieferant.Ort.ID,
			&futter.Lieferant.Ort.Stadt,
			&futter.Lieferant.Ort.Plz,
		)
		futterArr = append(futterArr, futter)
	}
	return futterArr
}

// #endregion

// #region Ort

func GetOrt(db core.DB_Handler, id string) objects.Ort {
	var ort objects.Ort
	query, args := ort.GetOrtFrom(id)
	row := db.QueryRow(query, args...)
	row.Scan(
		&ort.ID,
		&ort.Stadt,
		&ort.Plz,
	)
	return ort
}

func GettAllOrte(db core.DB_Handler) []objects.Ort {
	var ortArr []objects.Ort
	var ort objects.Ort
	query, args := ort.GetAllOrte()
	var oRows = db.Query(query, args...)
	for oRows.Next() {
		var ort objects.Ort
		oRows.Scan(
			&ort.ID,
			&ort.Stadt,
			&ort.Plz,
		)
		ortArr = append(ortArr, ort)
	}
	return ortArr
}

// #endregion

// #region Pfleger

func GetPfleger(db core.DB_Handler, id string) objects.Pfleger {
	var pf objects.Pfleger
	query, args := pf.GetPflegerFrom(id)
	row := db.QueryRow(query, args...)
	row.Scan(
		&pf.ID,
		&pf.Name,
		&pf.Telefonnummer,
		&pf.Adresse,
		&pf.Revier.ID,
		&pf.Revier.Name,
		&pf.Revier.Beschreibung,
		&pf.Ort.ID,
		&pf.Ort.Stadt,
		&pf.Ort.Plz,
	)
	return pf
}

func GetAllPfleger(db core.DB_Handler) []objects.Pfleger {
	var pfArr []objects.Pfleger
	var pf objects.Pfleger
	query, args := pf.GetAllPfleger()
	var pfRows = db.Query(query, args...)
	for pfRows.Next() {
		var pf objects.Pfleger
		pfRows.Scan(
			&pf.ID,
			&pf.Name,
			&pf.Telefonnummer,
			&pf.Adresse,
			&pf.Revier.ID,
			&pf.Revier.Name,
			&pf.Revier.Beschreibung,
			&pf.Ort.ID,
			&pf.Ort.Stadt,
			&pf.Ort.Plz,
		)
		pfArr = append(pfArr, pf)
	}
	return pfArr
}

func CreatePfleger(db core.DB_Handler, pfleger objects.Pfleger) {
	query, args := pfleger.InsertPfleger(pfleger.Name, pfleger.Telefonnummer, pfleger.Adresse, pfleger.Ort.ID, pfleger.Revier.ID)
	db.Exec(query, args...)
}

func UpdatePfleger(db core.DB_Handler, pfleger objects.Pfleger) {
	query, args := pfleger.UpdatePfleger(pfleger.ID, pfleger.Name, pfleger.Telefonnummer, pfleger.Adresse, pfleger.Ort.ID, pfleger.Revier.ID)
	db.Exec(query, args...)
}

// #endregion

// #region Contact
// #TODO: contact schreiben
func CreateContact(db core.DB_Handler, contact objects.Contact) {
	fmt.Println(contact)
}

// #endregion
