package main 

import (
	"io"
	"log"
	"os"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"

	"gitlab.iqt.org/rashley/covid-test-db/models/poc"
	"gitlab.iqt.org/rashley/covid-test-db/models/company"
	"gitlab.iqt.org/rashley/covid-test-db/models/diagnostic"
    "gitlab.iqt.org/rashley/covid-test-db/models/diagnostic_type"
    "gitlab.iqt.org/rashley/covid-test-db/models/diagnostic_target_type"
    "gitlab.iqt.org/rashley/covid-test-db/models/regulatory_approval_type"
)

func molecularColumnMapping () map[string]int{
	return map[string]int{
		"company_name":  0,
		"test_name":   1,
		"regulatory_status": 2,
		"gene_target": 3,
		"sample_type": 4,
		"verfified_lod": 5,
		"avg_ct": 6,
		"sensitivity": 7,
		"specificity": 8,
		"source": 9,
		"turn_around_time": 10,
		"prep integrated": 11,	
		"tests_per_run": 12,
		"tests_per_kit": 13,
		"product_no": 14,
		"product_name": 15,
		"pcr_platform": 16,	
		"ct_cut_off": 17,
		"company_poc": 18,
		"company_poc_phone": 19,
		"company_street": 20,
		"company_city": 21,
		"ompany_state": 22,
		"company_country": 23,
		"company_postal_code": 24,
		"cost_per_kit": 25,
		"in_stock": 26,
		"lead_time": 27,
		"company_stage": 28,
		"company_valuation": 29,
		"company_entered_market": 30,
		"test_counts": 31,
		"production_rate": 32,
		"expanding_capacity": 33,
	}
} 

func getDB () *gorm.DB {
	const addr = "postgresql://covid_bug@localhost:26257/covid_diagnostics?sslmode=disable"
    db, err := gorm.Open("postgres", addr)
    if err != nil {
        log.Fatal(err)
    }
    //defer db.Close()

    // Set to `true` and GORM will print out all DB queries.
    db.LogMode(true)

    return db
}

func getOrCreatePoc(name string, email string)(*poc.Poc, error){
	db := getDB()
	defer db.Close()
	var result *poc.Poc = nil
	existing, err := poc.FetchByNameAndEmail(db, name, email)
	if(existing != nil && !gorm.IsRecordNotFoundError(err)){
		result = existing
	} else {
		result, err = poc.Create(db, name, email, "")
	}

	return result, err
}

func getOrCreateCompany(name string, streetAddress string, city string, state string, 
        country string, postalCode string, stage string, valuation string, poc poc.Poc)(*company.Company, error){
	db := getDB()
	defer db.Close()
	var result *company.Company = nil
	existing, err := company.FetchByName(db, name)
	if(existing != nil && !gorm.IsRecordNotFoundError(err)){
		result = existing
	} else {
		result, err = company.Create(db, name, streetAddress, city, state, country, postalCode, stage, valuation, poc)
	}

	return result, err
}

func createDiagnostic( name string, description string, company company.Company,
			 diagnosticType diagnostic_type.DiagnosticType, poc poc.Poc, 
			 approvals []regulatory_approval_type.RegulatoryApprovalType, 
             targets []diagnostic_target_type.DiagnosticTargetType)(*diagnostic.Diagnostic, error){
	db := getDB()
	defer db.Close()

	var result *diagnostic.Diagnostic = nil
	result, err := diagnostic.Create(db, name, description, company, diagnosticType, poc, approvals, targets)	

	return result, err
}

func getTargetTypes(name string)([]diagnostic_target_type.DiagnosticTargetType, error){
	db := getDB()
	defer db.Close()
	var validTypes []diagnostic_target_type.DiagnosticTargetType
	allTypes, err := diagnostic_target_type.FetchList(db)

	for _, t := range allTypes{
		if(strings.Contains(strings.ToLower(name), strings.ToLower(t.Name))){
			validTypes = append(validTypes, t)
		}
	}

	return validTypes, err
}

func getApprovals(name string)([]regulatory_approval_type.RegulatoryApprovalType, error){
	db := getDB()
	defer db.Close()
	var validApprovals []regulatory_approval_type.RegulatoryApprovalType
	allApprovals, err := regulatory_approval_type.FetchList(db)

	for _, a := range allApprovals{
		if(strings.Contains(strings.ToLower(name), strings.ToLower(a.Name))){
			validApprovals = append(validApprovals, a)
		}
	}

	return validApprovals, err
}

func getDiagnosticType(name string)(*diagnostic_type.DiagnosticType, error){
	db := getDB()
	defer db.Close()

	result, err := diagnostic_type.FetchByName(db, name)

	return result, err
}

func getDiagnosticFromRow(row []string)(*diagnostic.Diagnostic, error){
	var mapping = molecularColumnMapping()
	//PoC Data
	var contact_name string = ""
	var contact_email string = ""
	contact_text := row[mapping["company_poc"]]
	idx := strings.Index(contact_text, "mailto:")
	if( idx > -1){
		contact_email = contact_text[:idx + len("mailto:")]
	}
	contact_name = contact_text
	poc, err := getOrCreatePoc(contact_name, contact_email)

	//company data
	company, err := getOrCreateCompany(
        row[mapping["company_name"]],
        row[mapping["company_street"]],
        row[mapping["company_city"]],
        row[mapping["company_state"]],
        row[mapping["company_country"]],
        row[mapping["company_postal_code"]],
    	row[mapping["company_stage"]],
    	row[mapping["company_valuation"]],
    	*poc,
    )
    
    targetTypes, err := getTargetTypes("IgM")
    diagnosticType, err := getDiagnosticType("molecular assays")
    approvals, err := getApprovals(row[mapping["regulatory_status"]])

    if(err != nil){
    	return nil, err
    }
    
	dx, dxErr := createDiagnostic(
		row[mapping["test_name"]],
		row[mapping["test_name"]],
		*company,
		*diagnosticType,
		*poc, 
		approvals,
        targetTypes,
    )

    return dx, dxErr
}

func main() {
	logFile, err := os.OpenFile("scraper.log", os.O_CREATE | os.O_APPEND | os.O_RDWR, 0666)
	if err != nil {
	    panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	log.Println("Logging Started")
	log.Println("Excel file processing started")
	
	// open excel file
	f, err := excelize.OpenFile("Database_Molecular_and_Sero.xlsx")
    if err != nil {
        log.Println(err.Error())
        return
    }
    rows := f.GetRows("Molecular test fields")
    for _, row := range rows {
    	getDiagnosticFromRow(row)
    }

	log.Println("Excel file processing complete")

}