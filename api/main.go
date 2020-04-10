package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	// Import GORM-related packages.
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"

    // Necessary in order to check for transaction retry error codes.
    "github.com/lib/pq"

    "gitlab.iqt.org/rashley/covid-test-db/api/models"
)

func getDB () gorm.DB {
	const addr = "postgresql://covid_bug@localhost:26257/covid_tests?sslmode=disable"
    db, err := gorm.Open("postgres", addr)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Set to `true` and GORM will print out all DB queries.
    db.LogMode(false)

    return db
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func createPoc(w http.ResponseWriter, r *http.Request) {
	db := getDB()
	created = model.create(db, "test", "111.111.1111", "test@test.mail")
	json.NewEncoder(w).Encode(created)
}

func getPoc(w http.ResponseWriter, r *http.Request) {
	pocID := mux.Vars(r)["id"]
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/pocs", createPoc).Methods("POST")
	router.HandleFunc("/pocs/{id}", getPoc).Methods("POST")
	log.Fatal(http.ListenAndServe(":5000", router))
}