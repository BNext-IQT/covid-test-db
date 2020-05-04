package main 

import (
	"io"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"

	"gitlab.iqt.org/rashley/covid-test-db/models/poc"
    // "gitlab.iqt.org/rashley/covid-test-db/models/test_type"
    "gitlab.iqt.org/rashley/covid-test-db/models/test_target_type"
    // "gitlab.iqt.org/rashley/covid-test-db/models/regulatory_approval_type"
)

func getDB () *gorm.DB {
	const addr = "postgresql://covid_bug@localhost:26257/covid_tests?sslmode=disable"
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

func getTargetTypes(name string)([]test_target_type.TestTargetType, error){
	db := getDB()
	defer db.Close()
	var validTypes []test_target_type.TestTargetType
	allTypes, err := test_target_type.FetchList(db)

	for _, t := range allTypes{
		if(strings.Contains(strings.ToLower(Name), strings.ToLower(t.name))){
			validTypes = append(validTypes, t)
		}
	}

	return validTypes, err
}

func main() {
	logFile, err := os.OpenFile("scraper.log", os.O_CREATE | os.O_APPEND | os.O_RDWR, 0666)
	if err != nil {
	    panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	log.Println("Logging Started")
	
	// Instantiate default collector
	c := colly.NewCollector(
		// Allow requests only to specified domains
		colly.AllowedDomains("www.finddx.org"),
	)

	// Extract product details
	c.OnHTML("li.diagnostic", func(e *colly.HTMLElement) {
		log.Println("diagnostic found")
		gq := e.DOM
		company_name := e.ChildText("a.labo")
		website_url := e.ChildAttr("a.labo", "href")
		test := gq.Contents().Not("span").Not("a").Text()
		test = strings.TrimSpace(test)
		contact, _ := gq.Find("span.contact").Children().Find("a[href]").Attr("href")
		contact = strings.TrimSpace(strings.Replace(contact, "mailto:", "", -1))
		idx := strings.IndexByte(contact, '@')
		contact_name := ""
		if idx > -1 {
			contact_name = contact[:idx]
		}
		
		log.Println(company_name + " : " + website_url)
		p, err := getOrCreatePoc(contact_name, contact)
		if err != nil {
	        log.Fatal(err)
	    }
		log.Println(p)
	})

	log.Println("Scraping Started")

	c.Visit("https://www.finddx.org/covid-19/pipeline/?avance=Commercialized&type=all&status=all&section=immunoassays#diag_tab")

	log.Println("Scraping Complete")

}