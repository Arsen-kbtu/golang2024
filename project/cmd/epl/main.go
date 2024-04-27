package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	pkg "project/pkg/epl/models"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

type application struct {
	config config
	models pkg.Models
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://postgres:password@localhost:5432/project?sslmode=disable", "PostgreSQL DSN")
	flag.Parse()

	// Connect to DB
	db, err := openDB(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect to the database: ", err)
	} else {
		log.Println("Connected to the database")
	}
	app := &application{
		config: cfg,
		models: pkg.NewModels(db),
	}

	if err := app.serve(); err != nil {
		fmt.Sprintf("error starting server: %v\n", err)
	}
	//app.run()
}

func (app *application) run() http.Handler {
	r := mux.NewRouter()

	v1 := r.PathPrefix("/api/v1").Subrouter()

	// Menu Singleton
	// Create a new menu
	v1.HandleFunc("/clubs", app.createClubHandler).Methods("POST")
	// Get a specific menu
	v1.HandleFunc("/clubs", app.getClubsHandler).Methods("GET")
	v1.HandleFunc("/clubs/{clubId:[0-9]+}", app.requireActivatedUser(app.getClubHandler)).Methods("GET")
	// Update a specific menu
	v1.HandleFunc("/clubs/{clubId:[0-9]+}", app.updateClubHandler).Methods("PUT")
	// Delete a specific menu
	v1.HandleFunc("/clubs/{clubId:[0-9]+}", app.deleteClubHandler).Methods("DELETE")

	v1.HandleFunc("/players", app.createPlayerHandler).Methods("POST")
	v1.HandleFunc("/players", app.getPlayersHandler).Methods("GET")
	v1.HandleFunc("/players/{playerId:[0-9]+}", app.requireActivatedUser(app.getPlayerHandler)).Methods("GET")
	v1.HandleFunc("/players/{playerId:[0-9]+}", app.updatePlayerHandler).Methods("PUT")
	v1.HandleFunc("/players/{playerId:[0-9]+}", app.deletePlayerHandler).Methods("DELETE")

	v1.HandleFunc("/users", app.registerUserHandler).Methods("POST")
	v1.HandleFunc("/users/activated", app.activateUserHandler).Methods("PUT")
	v1.HandleFunc("/tokens/authentication", app.createAuthenticationTokenHandler).Methods("POST")

	log.Printf("Starting server on %s\n", app.config.port)
	//err := http.ListenAndServe(app.config.port, r)
	//if err != nil {
	//
	//}
	return app.authenticate(r)

}
func openDB(cfg config) (*sql.DB, error) {
	// Use sql.Open() to create an empty connection pool, using the DSN from the config // struct.
	db, err := sql.Open("postgres", cfg.db.dsn)
	fmt.Println(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}
