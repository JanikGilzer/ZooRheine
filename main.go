package main

import (
	"ZooDaBa/server"
	"ZooDaBa/server/core"
	"ZooDaBa/server/objects"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"html/template"
	"log"
	"log/slog"
	"mime"
	"net/http"
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
	http.HandleFunc("/server/template/read/gebaude-banner", serveGebaudeBanner)
	http.HandleFunc("/server/template/read/tierart-banner", serveTierartBanner)
	http.HandleFunc("/server/template/read/pfleger", core.RequireAuth(core.RequireRole("Verwaltung", servePflegerTemplate)))

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
		fmt.Println(err)
	}
	slog.Info("Server successfully started on port 8090")
}

// #endregion

// #region serve

func serveLogin(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err == nil {
		// Validate token
		claims := &core.Claims{}
		token, err := jwt.ParseWithClaims(cookie.Value, claims, func(token *jwt.Token) (interface{}, error) {
			return core.JwtSecret, nil
		})

		if err == nil && token.Valid {
			// Redirect to appropriate page based on role
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
	http.ServeFile(w, req, "./html/animals.html")
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
	id := req.URL.Query().Get("id")
	tier := server.GetTier(db2, id)

	tmpl, err := template.ParseFiles("./html/templates/update/tier.html")
	if err != nil {
		log.Println(err)
	}
	err = tmpl.Execute(w, tier)
	if err != nil {
		log.Println(err)
	}
}

func serveUpdatePfleger(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	pfleger := server.GetPfleger(db2, id)
	tmpl, _ := template.ParseFiles("./html/templates/update/pfleger.html")
	tmpl.Execute(w, pfleger)
}

func serveUpdateGebaude(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	gebaude := server.GetGebaeude(db2, id)
	tmpl, _ := template.ParseFiles("./html/templates/update/gebaude.html")
	tmpl.Execute(w, gebaude)
}

func serveRevier(w http.ResponseWriter, req *http.Request) {
	revierID := req.URL.Query().Get("id")

	rev := server.GetRevier(db2, revierID)

	tmpl, err := template.ParseFiles("./html/templates/read/revier.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = tmpl.Execute(w, rev)
	if err != nil {
		fmt.Println(err)
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

	tmpl, err := template.ParseFiles("./html/templates/read/revier_banner.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = tmpl.Execute(w, rev)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func serveGebaudeBanner(w http.ResponseWriter, req *http.Request) {
	gebaudeID := req.URL.Query().Get("id")

	geb := server.GetGebaeude(db2, gebaudeID)

	tmpl, err := template.ParseFiles("./html/templates/read/gebaude_banner.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = tmpl.Execute(w, geb)
	if err != nil {
		fmt.Println(err)
		return
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
	fmt.Println(tierId)
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
	fmt.Println(tierId)
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
		fmt.Println(benoetigtesFutter)
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
		fmt.Println(futter)
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
		fmt.Println(pfleger)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(pfleger)
	} else {
		pfleger := server.GetPfleger(db2, pflegerId)
		fmt.Println(pfleger)
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
	fmt.Println(tier)
	tFutter := []objects.Futter{}
	for _, futter := range tier.Futter {
		f := server.GetFutterFromName(db2, futter)
		tFutter = append(tFutter, f)
	}
	fmt.Println(tFutter)
	server.CreateTier(db2, tier.Tier, tFutter)
}

func createPfleger(w http.ResponseWriter, req *http.Request) {
	type newPfleger struct {
		Pfleger objects.Pfleger
	}
	np := newPfleger{}
	json.NewDecoder(req.Body).Decode(&np)
	fmt.Println(np)
	server.CreatePfleger(db2, np.Pfleger)
}

func createZeit(w http.ResponseWriter, req *http.Request) {
	z := objects.Zeit{}
	json.NewDecoder(req.Body).Decode(&z)
	fmt.Println(z)
	server.CreateZeit(db2, z)
}

func createTierArt(w http.ResponseWriter, req *http.Request) {
	t := objects.TierArt{}
	json.NewDecoder(req.Body).Decode(&t)
	fmt.Println(t)
	server.CreateTiertArt(db2, t)
}

func createRevier(w http.ResponseWriter, req *http.Request) {
	r := objects.Revier{}
	json.NewDecoder(req.Body).Decode(&r)
	fmt.Println(r)
	server.CreateRevier(db2, r)
}

func createOrt(w http.ResponseWriter, req *http.Request) {
	o := objects.Ort{}
	json.NewDecoder(req.Body).Decode(&o)
	fmt.Println(o)
	server.CreateOrt(db2, o)
}

func createLieferant(w http.ResponseWriter, req *http.Request) {
	l := objects.Lieferant{}
	json.NewDecoder(req.Body).Decode(&l)
	fmt.Println(l)
	server.CreateLieferant(db2, l)
}

func createGebaude(w http.ResponseWriter, req *http.Request) {
	type newGebaude struct {
		Gebaude objects.Gebaude
		Zeit    []string
	}
	g := newGebaude{}
	json.NewDecoder(req.Body).Decode(&g)
	fmt.Println(g)
	fmt.Println(g.Zeit)
	fmt.Println(g.Gebaude)
	gZeit := []objects.Zeit{}
	for _, zeit := range g.Zeit {
		fmt.Println(zeit)
		z := server.GetZeitFromUhrzeit(db2, zeit)
		gZeit = append(gZeit, z)
	}
	fmt.Println(gZeit)
	server.CreateGebaude(db2, g.Gebaude, gZeit)
}

func createFutter(w http.ResponseWriter, req *http.Request) {
	f := objects.Futter{}
	json.NewDecoder(req.Body).Decode(&f)
	fmt.Println(f)
	server.CreateFutter(db2, f)
}

// #endregion

// #region update
func updateTier(w http.ResponseWriter, req *http.Request) {
	var tier objects.Tier
	fmt.Println(tier)
	json.NewDecoder(req.Body).Decode(&tier)
	fmt.Println(tier)
	server.UpdateTier(db2, tier)
}

func updatePfleger(w http.ResponseWriter, req *http.Request) {
	var pfleger objects.Pfleger
	fmt.Println(pfleger)
	json.NewDecoder(req.Body).Decode(&pfleger)
	fmt.Println(pfleger)
	server.UpdatePfleger(db2, pfleger)
}

func updateGebaude(w http.ResponseWriter, req *http.Request) {
	var g objects.Gebaude
	json.NewDecoder(req.Body).Decode(&g)
	server.UpdateGebaude(db2, g)
}

// #endregion
