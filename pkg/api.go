package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/onyewuenyi/rest_api/pkg/util/configuration"
)

// gorilla/mux is a package that enhances net/http and let you
// create routes with named parameters, GET/POST handlers and domain

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// func (a *App) Initialize(cfg *configuration.Configuration) error {
// 	db, err := a.initStore(cfg.Database)
// 	if err != nil {
// 		return err
// 	}
// 	defer db.Close()
// }

// func (a *App) initStore(db *configuration.Database) (*sql.DB, error) {
// 	// TODO what is sslmode=disable in regards to the db
// 	pgConnString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
// 		os.Getenv("PGHOST"),
// 		os.Getenv("PGPORT"),
// 		os.Getenv("PGDATABASE"),
// 		os.Getenv("PGUSER"),
// 		os.Getenv("PGPASSWORD"),
// 	)

// 	var (
// 		DB  *sql.DB
// 		err error
// 	)
// 	openDB := func() error {
// 		a.DB, err := sql.Open("postgres", pgConnString)
// 		return err
// 	}

// 	err = backoff.Retry(openDB, backoff.NewExponentialBackOff())
// 	if err != nil {
// 		return nil, err
// 	}

// 	const tableCreationQuery = `CREATE TABLE IF NOT EXISTS subscriptions
// 	(
// 		id uuid NOT NULL,
// 		PRIMARY KEY (id),
// 		email TEXT NOT NULL UNIQUE,
// 		name TEXT NOT NULL,
// 		subscribed_as timestamptz NOT NULL
// 	)`

// 	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
// 		return nil, err
// 	}

// 	return a.DB, nil
// }

func Start() error {
	// app := App{}
	cfg, err := configuration.LoadConfig("rest_api/pkg/util/configuration")
	if err != nil {
		// fmt.Errorf("fatal error reading config file: %w", err)
		return err
	}
	// app.Initialize(cfg)

	// Create a req router that will receive all http conn. Used to pass to registered req handlers
	mux := mux.NewRouter()

	// Registered req handler receives all incoming HTTP req from clients: http clients, or API req
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	}).Methods("GET")

	mux.HandleFunc("/health_check", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")

	// POST / HTTP/2.0
	// Host: http://localhost:8000/subscriptions
	// Content-Type: application/x-www-form-urlencoded
	// Content-Length: 46
	// Res: 200 OK
	// Body: name=charles%20senpai&email=senpai%40gmail.com
	mux.HandleFunc("/subscriptions", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Printf("email => %s\n", r.FormValue("email"))
		fmt.Printf("name => %s\n", r.FormValue("name"))
		fmt.Println("path", r.URL.Path)
	}).Methods("POST")

	// Start go HTTP server and listen on a port N. Pass conn to handlers
	// nill as an arg  instructs the server to use the default router.
	log.Println("Listening...")
	err = http.ListenAndServe(cfg.Server.Port, mux)
	if err != nil {
		return err
	}
	return nil

}
