package main

import (
    "fmt"
    "log"
    "net/http"
    "database/sql"
    _ "github.com/lib/pq"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "lando"
  password = "lando"
  dbname   = "lando"
)


func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func setUpTables(db *sql.DB) {
	fmt.Println("Setting up tables...")
	tableCreationQueries := getCreationQueries()
	for _, creationQuery := range tableCreationQueries {
		fmt.Println(creationQuery)
		_, err := db.Exec(creationQuery)
		if err != nil {
			panic(err)
		}
	}
}

func loadFixtures(db *sql.DB) {
	fmt.Println("Loading fixtures...")
	fixtureQueries := getFixtureQueries()
	for _, fixtureQuery := range fixtureQueries {
		fmt.Println(fixtureQuery)
		_, err := db.Exec(fixtureQuery)
		if err != nil {
			panic(err)
		}
	}
}

func tearDownTables(db *sql.DB) {
	fmt.Println("Tearing down previous tables...")
	tableDeletionQueries := getDeletionQueries()
	for _, deletionQuery := range tableDeletionQueries {
		fmt.Println(deletionQuery)
		_, err := db.Exec(deletionQuery)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
    http.HandleFunc("/", handler)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    	"password=%s dbname=%s sslmode=disable",
    	host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
	  panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
	  panic(err)
	}
	// at this point, have correctly connected to database!

	// let's start setting up some tables:
    tearDownTables(db)
	setUpTables(db)
	loadFixtures(db)

	fmt.Println("Starting up server...")
    log.Fatal(http.ListenAndServe(":8080", nil))

}