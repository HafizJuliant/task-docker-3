package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

var tmpl = `<html><body><h1>{{.}}</h1></body></html>`

func main() {
    http.HandleFunc("/", handleRoot)
    http.HandleFunc("/connect", handleConnect)

    log.Fatal(http.ListenAndServe(":8082", nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Someone hit me!") // Log to console
    wiseWord := "Keep pushing forward, no matter the challenge."

    // Render HTML response
    t, _ := template.New("wiseWord").Parse(tmpl)
    t.Execute(w, wiseWord)
}

func handleConnect(w http.ResponseWriter, r *http.Request) {
    // Connect to PostgreSQL
    connStr := "postgres://postgres:password@localhost:5435/postgres?sslmode=disable"

    var err error
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Println("Failed to connect to db")
        fmt.Fprint(w, "Failed to connect to db")
        return
    }
    defer db.Close()

    // Insert timestamp 
	_, err = db.Exec(`INSERT INTO public.hafizjuliant_access_log (timestamp) VALUES ($1)`, time.Now())
	if err != nil {
		log.Println("Failed to insert timestamp:", err) 
		fmt.Fprint(w, "Failed to connect to db")
		return
	}
	
    log.Println("Success connect to db")
    fmt.Fprint(w, "Success connect to db")
}