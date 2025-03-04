package main

import (
	"ZooDaBa/server"
	"ZooDaBa/server/core"
	"ZooDaBa/server/objects"
	"encoding/json"
	"fmt"
	"html/template"
	"mime"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

var db2 core.DB_Handler

// #region Server
func main() {
	core.Logger_init()
	db2.Init()

	mime.AddExtensionType(".js", "application/javascript")
	mime.AddExtensionType(".css", "text/css")

	// Websites
	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			core.LoginHandler(w, r, db2)
		} else {
			serveLogin(w, r)
		}
	})
	http.HandleFunc("/logout", core.LogoutHandler)

	http.HandleFunc("/revier", serveRevier)
	http.HandleFunc("/about", serveAbout)
	http.HandleFunc("/animals", serveAnimals)
	http.HandleFunc("/contact", serveContact)
	http.HandleFunc("/admin-panel", core.RequireAuth(serveAdminPanel))
	http.HandleFunc("/futterplan", serveFutterplan)

	// internal
	http.HandleFunc("/server/template/create/tier", core.RequireAuth(core.RequireRole("Verwaltung", serveCreateTier)))
	http.HandleFunc("/server/template/update/tier", core.RequireAuth(core.RequireRole("Verwaltung", serveUpdateTier)))
	http.HandleFunc("/server/template/create/pfleger", core.RequireAuth(core.RequireRole("Verwaltung", serveCreatePfleger)))

	http.HandleFunc("/server/template/update/pfleger", core.RequireAuth(core.RequireRole("Verwaltung", serveUpdatePfleger)))
	http.HandleFunc("/server/template/create/zeit", core.RequireAuth(serveCreateZeit))

	http.HandleFunc("/server/template/create/tierart", core.RequireAuth(core.RequireRole("Verwaltung", serveCreateTierArt)))
	http.HandleFunc("/server/template/create/revier", core.RequireAuth(core.RequireRole("Verwaltung", serveCreateRevier)))
	http.HandleFunc("/server/template/create/ort", core.RequireAuth(core.RequireRole("Verwaltung", serveCreateOrt)))
	http.HandleFunc("/server/template/create/lieferant", core.RequireAuth(core.RequireRole("Verwaltung", serveCreateLieferant)))
	http.HandleFunc("/server/template/create/gebaude", core.RequireAuth(core.RequireRole("Verwaltung", serveCreateGebaude)))
	http.HandleFunc("/server/template/update/gebaude", core.RequireAuth(core.RequireRole("Verwaltung", serveUpdateGebaude)))
	http.HandleFunc("/server/template/create/futter", core.RequireAuth(core.RequireRole("Verwaltung", serveCreateFutter)))

	// Templates
	http.HandleFunc("/server/template/read/revier", serveRevierTemplate)
	http.HandleFunc("/server/template/read/gebaude-icon", serveGebaudeIcon)
	http.HandleFunc("/server/template/read/gebaude-banner", serveGebaudeBanner)
	http.HandleFunc("/server/template/read/pfleger-banner", core.RequireAuth(core.RequireRole("Verwaltung", servePflegerBanner)))

	http.HandleFunc("/server/template/read/futterplan-banner", serveFutterplanBanner)
	http.HandleFunc("/server/template/update/futterplan", core.RequireAuth(core.RequireRole("Verwaltung", serveUpdateFutterplan)))

	http.HandleFunc("/server/template/read/tierart-banner", serveTierartBanner)
	http.HandleFunc("/server/template/read/pfleger", core.RequireAuth(core.RequireRole("Verwaltung", servePflegerTemplate)))
	http.HandleFunc("/server/template/read/tier-banner", core.RequireAuth(serveTierBanner))
	// Header/Footer
	http.HandleFunc("/server/template/header", serveHeader)
	http.HandleFunc("/server/template/footer", serveFooter)

	// JSON
	http.HandleFunc("/server/json/revier", jsonRevier)
	http.HandleFunc("/server/json/gebaude", jsonGebaude)
	http.HandleFunc("/server/json/tier", jsonTier)
	http.HandleFunc("/server/json/tierart", jsonTierArt)
	http.HandleFunc("/server/json/zeit", jsonZeit)
	http.HandleFunc("/server/json/fuetterungszeiten", jsonfuetterungszeiten)
	http.HandleFunc("/server/json/benoetigtesfutter", jsonBenoetigtesFutter)
	http.HandleFunc("/server/json/futter", jsonFutter)
	http.HandleFunc("/server/json/pfleger", core.RequireAuth(jsonPfleger))
	http.HandleFunc("/server/json/ort", jsonOrt)
	http.HandleFunc("/server/json/lieferant", core.RequireAuth(core.RequireRole("Verwaltung", jsonLieferant)))

	// how many of these are in the database?
	http.HandleFunc("/server/count/reviere", countReviere)
	http.HandleFunc("/server/count/gebaude", countGebeude)
	http.HandleFunc("/server/count/tiere", countTiere)

	// create new Object in database
	http.HandleFunc("/server/create/tier", core.RequireAuth(core.RequireRole("Verwaltung", createTier)))
	http.HandleFunc("/server/create/pfleger", core.RequireAuth(core.RequireRole("Verwaltung", createPfleger)))
	http.HandleFunc("/server/create/zeit", core.RequireAuth(createZeit))
	http.HandleFunc("/server/create/tierart", core.RequireAuth(core.RequireRole("Verwaltung", createTierArt)))
	http.HandleFunc("/server/create/revier", core.RequireAuth(core.RequireRole("Verwaltung", createRevier)))
	http.HandleFunc("/server/create/ort", core.RequireAuth(core.RequireRole("Verwaltung", createOrt)))
	http.HandleFunc("/server/create/lieferant", core.RequireAuth(core.RequireRole("Verwaltung", createLieferant)))
	http.HandleFunc("/server/create/gebaude", core.RequireAuth(core.RequireRole("Verwaltung", createGebaude)))
	http.HandleFunc("/server/create/futter", core.RequireAuth(core.RequireRole("Verwaltung", createFutter)))

	// update
	http.HandleFunc("/server/update/tier", core.RequireAuth(core.RequireRole("Verwaltung", updateTier)))
	http.HandleFunc("/server/update/pfleger", core.RequireAuth(core.RequireRole("Verwaltung", updatePfleger)))
	http.HandleFunc("/server/update/gebaude", core.RequireAuth(core.RequireRole("Verwaltung", updateGebaude)))
	http.HandleFunc("/server/update/futterplan", core.RequireAuth(core.RequireRole("Verwaltung", updateFutterplan)))

	http.HandleFunc("/server/send/contact", sendContact)

	http.Handle("/bilder/", http.StripPrefix("/bilder/", http.FileServer(http.Dir("./html/bilder"))))

	// Serve static HTML files
	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("./html"))))
	http.Handle("/html/templates/create", http.StripPrefix("/html/templates/create", http.FileServer(http.Dir("./html/templates/create"))))

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))

	// Serve static JavaScript files
	http.Handle("/script/", http.StripPrefix("/script/", http.FileServer(http.Dir("./script"))))

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		core.Logger.Error("Server konnte nicht Starten: ", "err", err)
	}
	core.Logger.Info("Server successfully started on port 8090")
}

// #endregion

// #region serve

func serveLogin(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err == nil {
		claims := &core.Claims{}
		token, err := jwt.ParseWithClaims(cookie.Value, claims, func(token *jwt.Token) (interface{}, error) {
			return core.JwtSecret, nil
		})

		if err == nil && token.Valid {
			if claims.Role == "Admin" {
				http.Redirect(w, r, "/", http.StatusFound)
			} else {
				http.Redirect(w, r, "/", http.StatusFound)
			}
			return
		}
	}

	// Serve login page
	tmpl := template.Must(template.ParseFiles("./html/login.html"))
	tmpl.Execute(w, nil)
}

func serveIndex(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./html/index.html")
}

func serveAbout(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./html/about.html")
}

func serveContact(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./html/contact.html")
}

func serveAnimals(w http.ResponseWriter, req *http.Request) {
	tierart := server.GetAllTierArt(db2)

	tmpl, _ := template.ParseFiles("./html/animals.html")
	err := tmpl.Execute(w, tierart)
	if err != nil {
		core.Logger.Error("template execute animals.html", "err", err)
	}
}

func serveTierartBanner(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if id == "" {
		return
	}
	tierart := server.GetTier(db2, id)
	tmpl, _ := template.ParseFiles("./html/templates/read/tierart_banner.html")
	err := tmpl.Execute(w, tierart)
	if err != nil {
		core.Logger.Error("template execute tierart_banner.html", "err", err)
	}
}

func serveTierBanner(w http.ResponseWriter, req *http.Request) {
	type tierBanner struct {
		Tier   objects.Tier
		Futter []objects.Futter
	}
	id := req.URL.Query().Get("id")
	if id == "" {
		http.ServeFile(w, req, "./html/templates/read/tier.html")
		return
	}

	tb := tierBanner{}
	tb.Tier = server.GetTier(db2, id)

	bFutter := server.GetAllBenoetigtesFutter(db2)

	for _, b := range bFutter {
		if b.Tier.ID == tb.Tier.ID {
			tb.Futter = append(tb.Futter, b.Futter)
		}
	}
	fmt.Println(tb)
	tmpl, _ := template.ParseFiles("./html/templates/read/tier_banner.html")
	err := tmpl.Execute(w, tb)
	if err != nil {
		core.Logger.Error("template execute tier_banner.html", "err", err)
	}
}

func serveAdminPanel(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./html/admin-panel.html")
}

func serveFutterplan(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./html/futterplan.html")
}

func serveHeader(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./html/templates/header.html")
}

func serveFooter(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./html/templates/footer.html")
}

func serveUpdateTier(w http.ResponseWriter, req *http.Request) {
	type updateTier struct {
		Tier    objects.Tier
		Gebaude []objects.Gebaude
		Futter  []objects.Futter
	}
	ut := updateTier{}
	id := req.URL.Query().Get("id")

	ut.Tier = server.GetTier(db2, id)
	ut.Gebaude = server.GetAllGebaude(db2)
	ut.Futter = server.GetAllFutter(db2)

	tmpl, err := template.ParseFiles("./html/templates/update/tier.html")
	if err != nil {
		core.Logger.Error("template tier.html", "err", err)
	}

	err = tmpl.Execute(w, ut)
	if err != nil {
		core.Logger.Error("template execute tier.html", "err", err)
	}
}

func serveUpdatePfleger(w http.ResponseWriter, req *http.Request) {
	type updatePfleger struct {
		Pfleger objects.Pfleger
		Orte    []objects.Ort
		Reviere []objects.Revier
	}
	up := updatePfleger{}
	id := req.URL.Query().Get("id")
	up.Pfleger = server.GetPfleger(db2, id)
	up.Reviere = server.GetAllReviere(db2)
	up.Orte = server.GettAllOrte(db2)
	tmpl, _ := template.ParseFiles("./html/templates/update/pfleger.html")
	tmpl.Execute(w, up)
}

func serveUpdateGebaude(w http.ResponseWriter, req *http.Request) {
	type updateGebaude struct {
		Gebaude objects.Gebaude
		Reviere []objects.Revier
	}
	id := req.URL.Query().Get("id")
	ug := updateGebaude{}
	ug.Gebaude = server.GetGebaeude(db2, id)
	ug.Reviere = server.GetAllReviere(db2)
	tmpl, _ := template.ParseFiles("./html/templates/update/gebaude.html")
	tmpl.Execute(w, ug)
}

func serveRevier(w http.ResponseWriter, req *http.Request) {
	revierID := req.URL.Query().Get("id")

	rev := server.GetRevier(db2, revierID)

	tmpl, err := template.ParseFiles("./html/templates/read/revier.html")
	if err != nil {
		core.Logger.Error("template revier.html", "err", err)
		return
	}

	err = tmpl.Execute(w, rev)
	if err != nil {
		core.Logger.Error("template execute revier.html", "err", err)
		return
	}
}

func servePflegerTemplate(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./html/templates/read/pfleger.html")
}

func serveCreateTier(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./html/templates/create/createTier.html")
}

func serveCreatePfleger(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./html/templates/create/pfleger.html")
}

func serveCreateZeit(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./html/templates/create/zeit.html")
}

func serveCreateTierArt(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./html/templates/create/tierart.html")
}

func serveCreateRevier(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./html/templates/create/revier.html")
}

func serveCreateOrt(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./html/templates/create/ort.html")
}

func serveCreateLieferant(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./html/templates/create/lieferant.html")
}

func serveCreateGebaude(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./html/templates/create/gebaude.html")
}

func serveCreateFutter(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./html/templates/create/futter.html")
}

func serveRevierTemplate(w http.ResponseWriter, req *http.Request) {
	revierID := req.URL.Query().Get("id")

	rev := server.GetRevier(db2, revierID)

	tmpl, err := template.ParseFiles("./html/templates/read/revier_icon.html")
	if err != nil {
		core.Logger.Error("template revier_icon.html", "err", err)
		return
	}

	err = tmpl.Execute(w, rev)
	if err != nil {
		core.Logger.Error("template execute revier_icon.html", "err", err)
		return
	}
}

func serveGebaudeIcon(w http.ResponseWriter, req *http.Request) {
	type gebaudeIcon struct {
		Gebaude objects.Gebaude
		Tierart []objects.TierArt
		Zeit    []objects.Zeit
	}
	gi := gebaudeIcon{}
	gebaudeID := req.URL.Query().Get("id")

	gi.Gebaude = server.GetGebaeude(db2, gebaudeID)

	allTier := server.GetAllTiere(db2)
	for _, t := range allTier {
		if t.Gebaude.ID == gi.Gebaude.ID {
			found := false
			for _, ta := range gi.Tierart {
				if t.Tierart == ta {
					found = true
					break
				}
			}
			if !found {
				gi.Tierart = append(gi.Tierart, t.Tierart)
			}
		}
	}

	futterZeiten := server.GetAllFutterZeiten(db2)
	for _, z := range futterZeiten {
		if z.Gebaude.ID == gi.Gebaude.ID {
			gi.Zeit = append(gi.Zeit, z.Zeit)
		}
	}

	tmpl, err := template.ParseFiles("./html/templates/read/gebaude_icon.html")
	if err != nil {
		core.Logger.Error("template gebaude_icon.html", "err", err)
		return
	}

	err = tmpl.Execute(w, gi)
	if err != nil {
		core.Logger.Error("template execute gebaude_icon.html", "err", err)
		return
	}
}

func serveGebaudeBanner(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if id == "" {
		http.ServeFile(w, req, "./html/templates/read/gebaude.html")
		return
	}
	gebaude := server.GetGebaeude(db2, id)
	tmpl, _ := template.ParseFiles("./html/templates/read/gebaude_banner.html")
	err := tmpl.Execute(w, gebaude)
	if err != nil {
		core.Logger.Error("template execute gebaude_banner.html", "err", err)
	}
}

func servePflegerBanner(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if id == "" {
		http.ServeFile(w, req, "./html/templates/read/pfleger.html")
		return
	}
	pfleger := server.GetPfleger(db2, id)
	tmpl, _ := template.ParseFiles("./html/templates/read/pfleger_banner.html")
	err := tmpl.Execute(w, pfleger)
	if err != nil {
		core.Logger.Error("template execute pfleger_banner.html", "err", err)
	}
}

func serveFutterplanBanner(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if id == "" {
		http.ServeFile(w, req, "./html/templates/read/futterplan.html")
		return
	}
	type futterplan struct {
		Gebaude objects.Gebaude
		Zeit    []objects.FuetterungsZeiten
		Futter  []objects.BenoetigtesFutter
	}
	fp := futterplan{}
	fp.Gebaude = server.GetGebaeude(db2, id)

	// fuetterungszeiten hinzufügen
	zeit := server.GetAllFutterZeiten(db2)
	for z := range zeit {
		if zeit[z].Gebaude.ID == fp.Gebaude.ID {
			fp.Zeit = append(fp.Zeit, zeit[z])
		}
	}

	// das futter hinzufügen
	tier := server.GetAllTiere(db2)
	var tiereImGebaude = []objects.Tier{}
	for t := range tier {
		if tier[t].Gebaude.ID == fp.Gebaude.ID {
			tiereImGebaude = append(tiereImGebaude, tier[t])
		}
	}

	futter := server.GetAllBenoetigtesFutter(db2)
	for f := range futter {
		for t := range tiereImGebaude {
			if tiereImGebaude[t].ID == futter[f].Tier.ID {
				fp.Futter = append(fp.Futter, futter[f])
			}
		}
	}

	tmpl, _ := template.ParseFiles("./html/templates/read/futterplan_banner.html")
	err := tmpl.Execute(w, fp)
	if err != nil {
		core.Logger.Error("template execute futterplan_banner.html", "err", err)
	}
}

func serveUpdateFutterplan(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if id == "" {
		return
	}
	type futterplan struct {
		Gebaude objects.Gebaude
		Zeit    []objects.Zeit
	}
	fp := futterplan{}
	fp.Gebaude = server.GetGebaeude(db2, id)
	fp.Zeit = server.GetAllZeiten(db2)
	tmpl, _ := template.ParseFiles("./html/templates/update/futterplan.html")
	err := tmpl.Execute(w, fp)
	if err != nil {
		core.Logger.Error("template execute futterplan.html", "err", err)
	}
}

// #endregion

// #region send

func sendContact(w http.ResponseWriter, req *http.Request) {
	var contact objects.Contact
	json.NewDecoder(req.Body).Decode(&contact)
	server.CreateContact(db2, contact)
}

// #endregion

// #region JSON
func jsonRevier(w http.ResponseWriter, req *http.Request) {
	revierID := req.URL.Query().Get("id")
	if revierID == "" {
		reviere := server.GetAllReviere(db2)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reviere)
	} else {
		revier := server.GetRevier(db2, revierID)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(revier)
	}
}

func jsonTier(w http.ResponseWriter, req *http.Request) {
	tierId := req.URL.Query().Get("id")
	if tierId == "" {
		tiere := server.GetAllTiere(db2)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tiere)

	} else {
		tier := server.GetTier(db2, tierId)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tier)
	}
}

func jsonTierArt(w http.ResponseWriter, req *http.Request) {
	tierId := req.URL.Query().Get("id")
	if tierId == "" {
		tiere := server.GetAllTierArt(db2)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tiere)
	} else {
		tier := server.GetTierArt(db2, tierId)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tier)
	}
}

func jsonfuetterungszeiten(w http.ResponseWriter, req *http.Request) {
	fuetterungszeitId := req.URL.Query().Get("id")

	if fuetterungszeitId == "" {
		fuetterungszeiten := server.GetAllFutterZeiten(db2)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(fuetterungszeiten)
	} else {
		fuetterungszeit := server.GetFutterZeit(db2, fuetterungszeitId)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(fuetterungszeit)
	}

}

func jsonBenoetigtesFutter(w http.ResponseWriter, req *http.Request) {
	benoetigtesFutterId := req.URL.Query().Get("id")
	if benoetigtesFutterId == "" {
		benoetigtesFutter := server.GetAllBenoetigtesFutter(db2)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(benoetigtesFutter)
	} else {
		benoetigtesFutter := server.GetBenoetigtesFutter(db2, benoetigtesFutterId)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(benoetigtesFutter)
	}
}

func jsonZeit(w http.ResponseWriter, req *http.Request) {
	zeitId := req.URL.Query().Get("id")

	if zeitId == "" {
		zeiten := server.GetAllZeiten(db2)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(zeiten)
	} else {
		zeit := server.GetZeit(db2, zeitId)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(zeit)
	}

}

func jsonGebaude(w http.ResponseWriter, req *http.Request) {
	gebaudeId := req.URL.Query().Get("id")
	if gebaudeId == "" {
		gebaude := server.GetAllGebaude(db2)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(gebaude)
	} else {
		gebaude := server.GetGebaeude(db2, gebaudeId)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(gebaude)
	}
}

func jsonFutter(w http.ResponseWriter, req *http.Request) {
	futterId := req.URL.Query().Get("id")
	if futterId == "" {
		futter := server.GetAllFutter(db2)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(futter)
	} else {
		futter := server.GetFutter(db2, futterId)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(futter)
	}

}

func jsonPfleger(w http.ResponseWriter, req *http.Request) {
	pflegerId := req.URL.Query().Get("id")
	if pflegerId == "" {
		pfleger := server.GetAllPfleger(db2)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(pfleger)
	} else {
		pfleger := server.GetPfleger(db2, pflegerId)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(pfleger)
	}
}

func jsonOrt(w http.ResponseWriter, req *http.Request) {
	orterId := req.URL.Query().Get("id")
	if orterId == "" {
		ort := server.GettAllOrte(db2)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ort)
	} else {
		ort := server.GetOrt(db2, orterId)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ort)
	}
}

func jsonLieferant(w http.ResponseWriter, req *http.Request) {
	lieferantId := req.URL.Query().Get("id")
	if lieferantId == "" {
		lieferant := server.GetAllLieferant(db2)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(lieferant)
	}
}

// #endregion

// #region count
func countReviere(w http.ResponseWriter, req *http.Request) {
	count := server.CountReviere(db2)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(count)
}

func countTiere(w http.ResponseWriter, req *http.Request) {
	count := server.CountTiere(db2)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(count)
}

func countGebeude(w http.ResponseWriter, req *http.Request) {
	count := server.CountGebaude(db2)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(count)
}

// #endregion

// #region create
func createTier(w http.ResponseWriter, req *http.Request) {
	type newTier struct {
		Tier   objects.Tier
		Futter []string
	}
	tier := newTier{}
	json.NewDecoder(req.Body).Decode(&tier)
	tFutter := []objects.Futter{}
	for _, futter := range tier.Futter {
		f := server.GetFutterFromName(db2, futter)
		tFutter = append(tFutter, f)
	}
	server.CreateTier(db2, tier.Tier, tFutter)
}

func createPfleger(w http.ResponseWriter, req *http.Request) {
	type newPfleger struct {
		Pfleger objects.Pfleger
	}
	np := newPfleger{}
	json.NewDecoder(req.Body).Decode(&np)
	server.CreatePfleger(db2, np.Pfleger)
}

func createZeit(w http.ResponseWriter, req *http.Request) {
	z := objects.Zeit{}
	json.NewDecoder(req.Body).Decode(&z)
	server.CreateZeit(db2, z)
}

func createTierArt(w http.ResponseWriter, req *http.Request) {
	t := objects.TierArt{}
	json.NewDecoder(req.Body).Decode(&t)
	server.CreateTiertArt(db2, t)
}

func createRevier(w http.ResponseWriter, req *http.Request) {
	r := objects.Revier{}
	json.NewDecoder(req.Body).Decode(&r)
	server.CreateRevier(db2, r)
}

func createOrt(w http.ResponseWriter, req *http.Request) {
	o := objects.Ort{}
	json.NewDecoder(req.Body).Decode(&o)
	server.CreateOrt(db2, o)
}

func createLieferant(w http.ResponseWriter, req *http.Request) {
	l := objects.Lieferant{}
	json.NewDecoder(req.Body).Decode(&l)
	server.CreateLieferant(db2, l)
}

func createGebaude(w http.ResponseWriter, req *http.Request) {
	type newGebaude struct {
		Gebaude objects.Gebaude
		Zeit    []string
	}
	g := newGebaude{}
	json.NewDecoder(req.Body).Decode(&g)
	gZeit := []objects.Zeit{}
	for _, zeit := range g.Zeit {
		z := server.GetZeitFromUhrzeit(db2, zeit)
		gZeit = append(gZeit, z)
	}
	server.CreateGebaude(db2, g.Gebaude, gZeit)
}

func createFutter(w http.ResponseWriter, req *http.Request) {
	f := objects.Futter{}
	json.NewDecoder(req.Body).Decode(&f)
	server.CreateFutter(db2, f)
}

// #endregion

// #region update
func updateTier(w http.ResponseWriter, req *http.Request) {

	t := objects.Tier{}

	json.NewDecoder(req.Body).Decode(&t)
	server.UpdateTier(db2, t)

}

func updatePfleger(w http.ResponseWriter, req *http.Request) {
	var pfleger objects.Pfleger
	json.NewDecoder(req.Body).Decode(&pfleger)
	server.UpdatePfleger(db2, pfleger)
}

func updateGebaude(w http.ResponseWriter, req *http.Request) {
	var g objects.Gebaude
	fmt.Println(g)
	json.NewDecoder(req.Body).Decode(&g)
	server.UpdateGebaude(db2, g)
}

func updateFutterplan(w http.ResponseWriter, req *http.Request) {
	type newFutterplan struct {
		Gebaude objects.Gebaude
		Zeit    []string
	}
	n := newFutterplan{}
	fmt.Println(req.Body)
	json.NewDecoder(req.Body).Decode(&n)
	fmt.Println(n)
	server.UpdateFutterplan(db2, n.Gebaude, n.Zeit)
}

// #endregion
