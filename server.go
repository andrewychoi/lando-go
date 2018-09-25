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
  user     = "lando	"
  password = "lando"
  dbname   = "lando"
)


func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
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


    log.Fatal(http.ListenAndServe(":8080", nil))
}