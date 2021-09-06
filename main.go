package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/cenkalti/backoff/v4"
	"github.com/gorilla/mux"
)

// gorilla/mux is a package that enhances net/http and let you
// create routes with named parameters, GET/POST handlers and domain

type App struct {
	Router *mux.Router
	DB *sql.DB
}

func (a *App) Initialize() {}
func (a *App) Run(addr string) {}

type struct Settings {
    application ApplicationSettings,
    database DatabaseSettings,
}


type struct ApplicationSettings {
    port int,
    host string,
}


type struct DatabaseSettings {
	uname string,
	password string,
	port int,
	host string,
	database_name string,
}

type struct Environment {
	Local string,
	Test string,
	Production string,
	Stagging string,
}



func (d *DatabaseSettings) conn_str() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", d.uname, d.password, d.host, d.port, d.database_name)
}

func (d *DatabaseSettings) default_conn_str() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/postgres", d.uname, d.password, d.host, d.port)
}


func (a *App) initStore() (*sql.DB, error) {
	pgConnString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGDATABASE"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
	)

	var (
		DB  *sql.DB
		err error
	)
	openDB := func() error {
		a.DB, err := sql.Open("postgres", pgConnString)
		return err
	}

	err = backoff.Retry(openDB, backoff.NewExponentialBackOff())
	if err != nil {
		return nil, err
	}

	const tableCreationQuery = `CREATE TABLE IF NOT EXISTS subscriptions
	(
		id uuid NOT NULL,
		PRIMARY KEY (id),
		email TEXT NOT NULL UNIQUE,
		name TEXT NOT NULL,
		subscribed_as timestamptz NOT NULL
	)`
	
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		return nil, err
	}

	return a.DB, nil
}


func main() {
	app: = App{}
	app.Initialize()
	app.Run()


	db, err:= a.initStore()
	if err != nil {
		return err 
	}
	defer db.Close()

	// Create a req router that will receive all http conn. Used to pass to registered req handlers
	mux := mux.NewRouter()

	// Registered req handler receives all incoming HTTP req from clients: http clients, or API req
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)

	}).Methods("GET")

	mux.HandleFunc("/health_check", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "XXX, you've requested: %s\n", r.URL.Path)
	})

	mux.HandleFunc("/subscriptions", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Printf("email => %s\n", r.FormValue("email"))
		fmt.Printf("name => %s\n", r.FormValue("name"))
		fmt.Println("path", r.URL.Path)
	}).Methods("POST")

	// Serve static assets
	fs := http.FileServer(http.Dir("static/"))

	// One thing to note: In order to serve files correctly, we need to strip away a part of the url path.
	// Usually this is the name of the directory our files live in.
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Listening...")

	// Start go HTTP server and listen on a port N. Pass conn to handlers
	// nill as an arg  instructs the server to use the default router.
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}

}

