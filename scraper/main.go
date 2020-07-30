package main 

import (
    "io"
    "log"
    "os"
    "strconv"
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
    "gitlab.iqt.org/rashley/covid-test-db/models/sample_type"
    "gitlab.iqt.org/rashley/covid-test-db/models/pcr_platform"
)

func molecularColumnMapping () map[string]int{
    return map[string]int{
        "company_name":  0,
        "test_name":   1,
        "pcr_platform": 2, 
        "sensitivity": 34,
        "specificity": 35,
        "source": 3,
        "test_url": 4,
        "regulatory_status": 5,
        "sample_type": 6,
        "point_of_care": 7,
        "prep_integrated": 8,
        "product_no": 9,
        "company_support": 10,
        "company_poc": 11,
        "company_poc_phone": 12,
        "company_street": 13,
        "company_city": 14,
        "company_state": 15,
        "company_country": 16,
        "company_postal_code": 17,
        "cost_per_kit": 18,
        "in_stock": 19,
        "lead_time": 20,
        "company_stage": 21,
        "company_valuation": 22,
        "company_entered_market": 23,
        "test_counts": 24,
        "production_rate": 25,
        "expanding_capacity": 26,
        "gene_target": 27,
        "verified_lod": 28,
        "avg_ct": 29,
        "turn_around_time": 30,
        "tests_per_run": 31,
        "tests_per_kit": 32,
        "abi_7500": 33,
        "test_type": 36,
    }
} 

//removes all whitespace and converts to lowercase
//to be used for string comparisons
func StompText(text string) string{
    var stomped string = strings.ToLower(strings.Join(strings.Fields(text),""))
    return stomped
}

func Index(vs []string, t string) int {
    for i, v := range vs {
        if v == t {
            return i
        }
    }
    return -1
}

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

func getOrCreatePoc(name string, email string, phone string)(*poc.Poc, error){
    db := getDB()
    defer db.Close()
    var result *poc.Poc = nil
    existing, err := poc.FetchByNameAndEmail(db, name, email)
    if(existing != nil && !gorm.IsRecordNotFoundError(err)){
        result = existing
    } else {
        result, err = poc.Create(db, name, email, phone)
    }

    return result, err
}

func getOrCreateCompany(name string, streetAddress string, city string, state string, 
        country string, postalCode string, stage string, valuation string, poc poc.Poc)(*company.Company, error){
    db := getDB()
    defer db.Close()
    var result *company.Company = nil
    list, err := company.FetchList(db)
    var existing *company.Company = nil
    stompedName := StompText(name);
    for _,c := range list {
        if(stompedName == StompText(c.Name)){
            existing = &c;
            break;
        }
    }

    if(existing != nil && !gorm.IsRecordNotFoundError(err)){
        result = existing
    } else {
        result, err = company.Create(db, name, streetAddress, city, state, country, postalCode, stage, valuation, poc)
    }

    return result, err
}

func getOrCreateTargetType(name string)(*diagnostic_target_type.DiagnosticTargetType, error){
    db := getDB()
    defer db.Close()
    var result *diagnostic_target_type.DiagnosticTargetType = nil
    list, err := diagnostic_target_type.FetchList(db)
    var existing *diagnostic_target_type.DiagnosticTargetType = nil
    stompedName := StompText(name);
    for _,dt := range list {
        if(stompedName == StompText(dt.Name)){
            existing = &dt;
            break;
        }
    }

    if(existing != nil && !gorm.IsRecordNotFoundError(err)){
        result = existing
    } else {
        result, err = diagnostic_target_type.Create(db, name)
    }

    return result, err
}

func getOrCreateSampleType(name string)(*sample_type.SampleType, error){
    db := getDB()
    defer db.Close()
    var result *sample_type.SampleType = nil
    list, err := sample_type.FetchList(db)
    var existing *sample_type.SampleType = nil
    stompedName := StompText(name);
    for _,st := range list {
        if(stompedName == StompText(st.Name)){
            existing = &st;
            break;
        }
    }

    if(existing != nil && !gorm.IsRecordNotFoundError(err)){
        result = existing
    } else {
        result, err = sample_type.Create(db, name)
    }

    return result, err
}

func getOrCreatePcrPlatform(name string)(*pcr_platform.PcrPlatform, error){
    db := getDB()
    defer db.Close()
    var result *pcr_platform.PcrPlatform = nil
    list, err := pcr_platform.FetchList(db)
    var existing *pcr_platform.PcrPlatform = nil
    stompedName := StompText(name);
    for _,p := range list {
        if(stompedName == StompText(p.Name)){
            existing = &p;
            break;
        }
    }

    if(existing != nil && !gorm.IsRecordNotFoundError(err)){
        result = existing
    } else {
        result, err = pcr_platform.Create(db, name)
    }

    return result, err
}

func createDiagnostic( name string, description string, testUrl string, company company.Company, 
             diagnosticType diagnostic_type.DiagnosticType, poc poc.Poc, 
             verifiedLod string, avgCt float64, prepIntegrated bool,
             testsPerRun int64, testsPerKit int64, sensitivity float64, specificity float64,
             sourceOfPerfData string, catalogNo string, pointOfCare bool, costPerKit float64,
             inStock bool, leadTime int64,
             approvals []regulatory_approval_type.RegulatoryApprovalType, 
             targets []diagnostic_target_type.DiagnosticTargetType,
             sampleTypes []sample_type.SampleType,
             pcrPlatforms []pcr_platform.PcrPlatform)(*diagnostic.Diagnostic, error){
    db := getDB()
    defer db.Close()

    var result *diagnostic.Diagnostic = nil
    result, err := diagnostic.Create(db, name, description, testUrl, company, diagnosticType, poc,
                    verifiedLod, avgCt, prepIntegrated, testsPerRun, testsPerKit,
                    sensitivity, specificity, sourceOfPerfData, catalogNo, pointOfCare, costPerKit,
                    inStock, leadTime,
                    approvals, targets, sampleTypes, pcrPlatforms)   

    return result, err
}

func getSampleTypes(names []string)([]sample_type.SampleType, []error){
    db := getDB()
    defer db.Close()
    var types []sample_type.SampleType = nil
    var errs []error = nil

    for _, name := range names{
        st, err := getOrCreateSampleType(strings.TrimSpace(name))
        if(err == nil){
            types = append(types, *st)
        } else {
            errs = append(errs, err)
        }
    }

    return types, errs
}

func getPcrPlatforms(names []string)([]pcr_platform.PcrPlatform, []error){
    db := getDB()
    defer db.Close()
    var pcrs []pcr_platform.PcrPlatform = nil
    var errs []error = nil

    for _, name := range names{
        st, err := getOrCreatePcrPlatform(strings.TrimSpace(name))
        if(err == nil){
            pcrs = append(pcrs, *st)
        } else {
            errs = append(errs, err)
        }
    }

    return pcrs, errs
}

func getTargetTypes(names []string)([]diagnostic_target_type.DiagnosticTargetType, []error){
    db := getDB()
    defer db.Close()
    var types []diagnostic_target_type.DiagnosticTargetType
    var errs []error

    for _, name := range names{
        tt, err := getOrCreateTargetType(strings.TrimSpace(name))
        if(err == nil){
            types = append(types, *tt)
        } else {
            errs = append(errs, err)
        }
    }

    return types, errs
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
    var contact_name string = strings.TrimSpace(row[mapping["company_poc"]])
    var contact_email string = strings.TrimSpace(row[mapping["company_support"]])
    var contact_phone string = strings.TrimSpace(row[mapping["company_poc_phone"]])

    poc, err := getOrCreatePoc(contact_name, contact_email, contact_phone)

    //company data
    company, err := getOrCreateCompany(
        strings.TrimSpace(row[mapping["company_name"]]),
        strings.TrimSpace(row[mapping["company_street"]]),
        strings.TrimSpace(row[mapping["company_city"]]),
        strings.TrimSpace(row[mapping["company_state"]]),
        strings.TrimSpace(row[mapping["company_country"]]),
        strings.TrimSpace(row[mapping["company_postal_code"]]),
        strings.TrimSpace(row[mapping["company_stage"]]),
        strings.TrimSpace(row[mapping["company_valuation"]]),
        *poc,
    )
    
    tts := strings.Split(row[mapping["gene_target"]], ",")
    sts := strings.Split(row[mapping["sample_type"]], ",")
    pcrs := strings.Split(row[mapping["pcr_platform"]], ",")
    targetTypes, errs := getTargetTypes(tts)
    if(len(errs) > 0){
        for _, e := range errs{
            log.Println(e)
        }
        return nil, errs[0]
    }
    sampleTypes, errs := getSampleTypes(sts)
    if(len(errs) > 0){
        for _, e := range errs{
            log.Println(e)
        }
        return nil, errs[0]
    }

    pcrPlatforms, errs := getPcrPlatforms(pcrs)
    if(len(errs) > 0){
        for _, e := range errs{
            log.Println(e)
        }
        return nil, errs[0]
    }
    
    diagnosticType, err := getDiagnosticType(strings.TrimSpace(row[mapping["test_type"]]))
    if(err != nil){
        log.Println(err)
        return nil, err
    }

    approvals, err := getApprovals(strings.TrimSpace(row[mapping["regulatory_status"]]))

    if(err != nil){
        log.Println(err)
        return nil, err
    }
    positives := []string{"y", "yes", "t", "true", "1"}
    avgCt, _ := strconv.ParseFloat(row[mapping["avg_ct"]], 64)
    prepIntegrated := Index(positives, strings.ToLower(strings.TrimSpace(row[mapping["prep_integrated"]]))) >= 0
    tpr, _ := strconv.ParseInt(row[mapping["tests_per_run"]], 10, 64)
    tpk, _ := strconv.ParseInt(row[mapping["tests_per_kit"]], 10, 64)
    sensitivity, _ := strconv.ParseFloat(row[mapping["sensitivity"]], 64)
    specificity, _ := strconv.ParseFloat(row[mapping["specificity"]], 64)
    sourceOfPerfData := strings.TrimSpace(row[mapping["source"]])
    catalogNo := strings.TrimSpace(row[mapping["product_no"]])
    pointOfCare := Index(positives, strings.ToLower(strings.TrimSpace(row[mapping["point_of_care"]]))) >= 0
    costPerKit, _ := strconv.ParseFloat(row[mapping["Cost_per_kit"]], 64)
    inStock := Index(positives, strings.ToLower(strings.TrimSpace(row[mapping["in_stock"]]))) >= 0 
    leadTime, _ := strconv.ParseInt(row[mapping["lead_time"]], 10, 64)

    dx, dxErr := createDiagnostic(
        strings.TrimSpace(row[mapping["test_name"]]),
        strings.TrimSpace(row[mapping["test_name"]]),
        strings.TrimSpace(row[mapping["test_url"]]),
        *company,
        *diagnosticType,
        *poc,
        strings.TrimSpace(row[mapping["verified_lod"]]),
        avgCt,
        prepIntegrated,
        tpr,
        tpk,
        sensitivity,
        specificity,
        sourceOfPerfData,
        catalogNo,
        pointOfCare,
        costPerKit,
        inStock,
        leadTime,
        approvals,
        targetTypes,
        sampleTypes,
        pcrPlatforms,
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
    rows := f.GetRows("Combined Molecular Tests")
    for idx, row := range rows {
        log.Printf("Processing row %d of %d \n", idx, len(rows))
        _, dxErr := getDiagnosticFromRow(row)

        if(dxErr != nil){
            log.Println(dxErr)
        }
    }

    log.Println("Excel file processing complete")

}