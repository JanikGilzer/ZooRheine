package main

import (
	"ZooDaBa/server"
	"ZooDaBa/server/core"
	"ZooDaBa/server/objects"
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"mime"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db2 core.DB_Handler

// #region Server
func main() {

	db2.Init()

	mime.AddExtensionType(".js", "application/javascript")
	mime.AddExtensionType(".css", "text/css")

	// Websites
	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/revier", serveRevier)
	http.HandleFunc("/about", serveAbout)
	http.HandleFunc("/animals", serveAnimals)
	http.HandleFunc("/contact", serveContact)
	http.HandleFunc("/admin-panel", serveAdminPanel)
	http.HandleFunc("/futterplan", serveFutterplan)

	// internal
	http.HandleFunc("/server/template/create/tier", serveCreateTier)
	http.HandleFunc("/server/template/update/tier", serveUpdateTier)
	http.HandleFunc("/server/template/create/pfleger", serveCreatePfleger)
	http.HandleFunc("/server/template/update/pfleger", serveUpdatePfleger)

	// Templates
	http.HandleFunc("/server/template/read/revier", serveRevierTemplate)
	http.HandleFunc("/server/template/read/gebaude-banner", serveGebaudeBanner)
	http.HandleFunc("/server/template/read/tierart-banner", serveTierartBanner)

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
	http.HandleFunc("/server/json/pfleger", jsonPfleger)
	http.HandleFunc("/server/json/ort", jsonOrt)

	// test
	//http.HandleFunc("/server/clicked", jsonConstructedTier)

	// how many of these are in the database?
	http.HandleFunc("/server/count/reviere", countReviere)
	http.HandleFunc("/server/count/gebaude", countGebeude)
	http.HandleFunc("/server/count/tiere", countTiere)

	// create new Object in database
	http.HandleFunc("/server/create/tier", createTier)
	http.HandleFunc("/server/create/pfleger", createPfleger)

	// update
	http.HandleFunc("/server/update/tier", updateTier)
	http.HandleFunc("/server/update/pfleger", updatePfleger)

	http.HandleFunc("/server/send/contact", sendContact)

	http.Handle("/bilder/", http.StripPrefix("/bilder/", http.FileServer(http.Dir("./html/bilder"))))

	// Serve static HTML files
	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("./html"))))
	http.Handle("/html/templates/create", http.StripPrefix("/html/templates/create", http.FileServer(http.Dir("./html/templates/create"))))

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))

	// Serve static JavaScript files
	http.Handle("/script/", http.StripPrefix("/script/", http.FileServer(http.Dir("./script"))))

	fmt.Println("Listening on :8090...")
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Println(err)
	}
}

// #endregion

// #region serve
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
	tmpl.Execute(w, tierart)
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

	tmpl, _ := template.ParseFiles("./html/templates/update/tier.html")
	tmpl.Execute(w, tier)
}

func serveUpdatePfleger(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	pfleger := server.GetPfleger(db2, id)
	tmpl, _ := template.ParseFiles("./html/templates/update/pfleger.html")
	tmpl.Execute(w, pfleger)
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

func serveCreateTier(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./html/templates/create/createTier.html")
}

func serveCreatePfleger(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./html/templates/create/pfleger.html")
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

// #endregion

func clicked(w http.ResponseWriter, req *http.Request) {
	var tier_list []objects.Tier

	fmt.Println("clicked")
	db := Connect_Database()
	defer db.Close()
	rows, _ := db.Query("SELECT * FROM tier")

	for rows.Next() {
		var tier objects.Tier
		rows.Scan(&tier.ID, &tier.Name, &tier.Geburtsdatum, &tier.Gebaude.ID)
		tier_list = append(tier_list, tier)
	}
	fmt.Print(tier_list)
	json.NewEncoder(w).Encode(tier_list)
}

// ------------------------------------------------------------------------ //

func Connect_Database() *sql.DB {
	// Read database configuration from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Set up MySQL configuration
	database_conf := mysql.NewConfig()
	database_conf.User = dbUser
	database_conf.Passwd = dbPassword
	database_conf.Net = "tcp"
	database_conf.Addr = fmt.Sprintf("%s:%s", dbHost, dbPort) // Use DB_HOST and DB_PORT
	database_conf.DBName = dbName

	// Open the database connection
	db, err := sql.Open("mysql", database_conf.FormatDSN())
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Successfully connected to the database!")
	return db
}
