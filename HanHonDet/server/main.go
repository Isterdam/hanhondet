package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
)

/*
Will use these to connect to the database
*/
const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "hanhondetdb"
)

/*
Declare router outside of main function to be returned later
*/
func newRouter() *mux.Router {
	r := mux.NewRouter()

	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))

	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	r.HandleFunc("/info", searchHandler).Methods("POST")

	return r
}

func main() {
	// Store database info and connect
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}

	InitPage(&dbPage{db: db})
	// Create router and start server
	r := newRouter()
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
