package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	pkg "project/pkg/models"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type config struct {
	port string
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
	flag.StringVar(&cfg.port, "port", ":5432", "API server port")
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

	app := &application{
		config: cfg,
		models: pkg.NewModels(db),
	}

	app.run()
}
func (app *application) run() {
	r := mux.NewRouter()

	v1 := r.PathPrefix("/api/v1").Subrouter()

	// Menu Singleton
	// Create a new menu
	v1.HandleFunc("/clubs", app.createClubHandler).Methods("POST")
	// Get a specific menu
	v1.HandleFunc("/clubs/{clubId:[0-9]+}", app.getClubHandler).Methods("GET")
	// Update a specific menu
	v1.HandleFunc("/clubs/{clubId:[0-9]+}", app.updateClubHandler).Methods("PUT")
	// Delete a specific menu
	v1.HandleFunc("/clubs/{clubId:[0-9]+}", app.deleteClubHandler).Methods("DELETE")

	log.Printf("Starting server on %s\n", app.config.port)
	err := http.ListenAndServe(app.config.port, r)
	log.Fatal(err)
}
func openDB(cfg config) (*sql.DB, error) {
	// Use sql.Open() to create an empty connection pool, using the DSN from the config // struct.
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
