package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	// Import GORM-related packages.
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"

    "github.com/BNext-IQT/covid-test-db/models/poc"
    "github.com/BNext-IQT/covid-test-db/models/diagnostic"
    "github.com/BNext-IQT/covid-test-db/models/diagnostic_type"
    "github.com/BNext-IQT/covid-test-db/models/sample_type"
    "github.com/BNext-IQT/covid-test-db/models/regulatory_approval_type"
    "github.com/BNext-IQT/covid-test-db/models/pcr_platform"
)

func getDB () *gorm.DB {
	const addr = "postgresql://covid_bug@roach:26257/covid_diagnostics?sslmode=disable"
    db, err := gorm.Open("postgres", addr)
    if err != nil {
        log.Fatal(err)
    }
    //defer db.Close()

    // Set to `true` and GORM will print out all DB queries.
    db.LogMode(false)

    return db
}

func sendJsonResponse(w http.ResponseWriter, payload interface{}) {
    response, _ := json.Marshal(payload)

    w.Header().Set("Content-Type", "application/json")
    w.Write(response)
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	log.Println("homelink called")
	fmt.Fprintf(w, "Welcome home!")
}

//PoC endpoints
func createPoc(w http.ResponseWriter, r *http.Request) {
	db := getDB()
	defer db.Close()
	var p poc.Poc

	err := json.NewDecoder(r.Body).Decode(&p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	created, err := poc.Create(db, p.Name, p.Email, p.Phone)
	if err != nil {
        log.Print(err)
    }

   	sendJsonResponse(w, created)
}

func updatePoc(w http.ResponseWriter, r *http.Request) {
	pocID, err := uuid.Parse(mux.Vars(r)["id"])
	db := getDB()
	defer db.Close()
	var p poc.Poc

	err = json.NewDecoder(r.Body).Decode(&p)

    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    p.Id = pocID

	created, err := poc.Update(db, &p)
	if err != nil {
        log.Print(err)
    }
	
    sendJsonResponse(w, created)
}

func getPocList(w http.ResponseWriter, r *http.Request) {
	db := getDB()
	defer db.Close()

	results, err := poc.FetchList(db)
	if err != nil {
        log.Print(err)
    }
	
    sendJsonResponse(w, results)
}

func getPoc(w http.ResponseWriter, r *http.Request) {
	pocID, err := uuid.Parse(mux.Vars(r)["id"])
	db := getDB()
	defer db.Close()

	result, err := poc.FetchById(db, pocID)
	if err != nil {
        log.Print(err)
    }
	
    sendJsonResponse(w, result)
}

//diagnostic endpoints
func getDiagnosticList(w http.ResponseWriter, r *http.Request) {
	db := getDB()
	defer db.Close()

	results, err := diagnostic.FetchList(db)
	if err != nil {
        log.Print(err)
    }
	
    sendJsonResponse(w, results)
}

func getDiagnostic(w http.ResponseWriter, r *http.Request) {
	dxID, err := uuid.Parse(mux.Vars(r)["id"])
	db := getDB()
	defer db.Close()

	result, err := diagnostic.FetchById(db, dxID)
	if err != nil {
        log.Print(err)
    }
	
    sendJsonResponse(w, result)
}

//diagnosticType endpoints
func getDiagnosticTypeList(w http.ResponseWriter, r *http.Request) {
	db := getDB()
	defer db.Close()

	results, err := diagnostic_type.FetchList(db)
	if err != nil {
        log.Print(err)
    }
	
    sendJsonResponse(w, results)
}

//sampleType endpoints
func getSampleTypeList(w http.ResponseWriter, r *http.Request) {
	db := getDB()
	defer db.Close()

	results, err := sample_type.FetchList(db)
	if err != nil {
        log.Print(err)
    }
	
    sendJsonResponse(w, results)
}

//Regulatory Approvals
func getRegulatoryApprovalList(w http.ResponseWriter, r *http.Request) {
	db := getDB()
	defer db.Close()

	results, err := regulatory_approval_type.FetchList(db)
	if err != nil {
        log.Print(err)
    }
	
    sendJsonResponse(w, results)
}

//PCR Platforms
func getPcrPlatformList(w http.ResponseWriter, r *http.Request) {
	db := getDB()
	defer db.Close()

	results, err := pcr_platform.FetchList(db)
	if err != nil {
        log.Print(err)
    }
	
    sendJsonResponse(w, results)
}

func main() {
	logFile, err := os.OpenFile("/var/log/cvdb/api.log", os.O_CREATE | os.O_APPEND | os.O_RDWR, 0666)
	if err != nil {
	    panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	log.Println("Logging Started")

	router := mux.NewRouter().StrictSlash(false)

	// headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	// originsOk := handlers.AllowedOrigins([]string{"*"})
	// methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	router.Methods("OPTIONS").HandlerFunc(
	    func(w http.ResponseWriter, r *http.Request){
	    headers := w.Header()
	    headers.Add("Access-Control-Allow-Origin", "*")
	    headers.Add("Vary", "Origin")
	    headers.Add("Vary", "Access-Control-Request-Method")
	    headers.Add("Vary", "Access-Control-Request-Headers")
	    headers.Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token")
	    headers.Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	})
	router.HandleFunc("/api/", homeLink)
	router.HandleFunc("/api/pocs", getPocList).Methods("GET")
	router.HandleFunc("/api/pocs/{id}", getPoc).Methods("GET")
	router.HandleFunc("/api/diagnostics", getDiagnosticList).Methods("GET")
	router.HandleFunc("/api/diagnostics/{id}", getDiagnostic).Methods("GET")
	router.HandleFunc("/api/diagnostictypes", getDiagnosticTypeList).Methods("GET")
	router.HandleFunc("/api/sampletypes", getSampleTypeList).Methods("GET")
	router.HandleFunc("/api/regulatoryapprovals", getRegulatoryApprovalList).Methods("GET")
	router.HandleFunc("/api/pcrplatforms", getPcrPlatformList).Methods("GET")
	log.Fatal(http.ListenAndServe(":5000", handlers.CORS()(router)))
}