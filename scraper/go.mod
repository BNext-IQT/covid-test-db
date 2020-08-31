module github.com/BNext-IQT/covid-test-db/scraper

replace github.com/BNext-IQT/covid-test-db/models/poc => ../models/poc/

replace github.com/BNext-IQT/covid-test-db/models/company => ../models/company/

replace github.com/BNext-IQT/covid-test-db/models/diagnostic => ../models/diagnostic/

replace github.com/BNext-IQT/covid-test-db/models/diagnostic_type => ../models/diagnostic-type/

replace github.com/BNext-IQT/covid-test-db/models/diagnostic_target_type => ../models/diagnostic-target-type/

replace github.com/BNext-IQT/covid-test-db/models/regulatory_approval_type => ../models/regulatory-approval-type/

replace github.com/BNext-IQT/covid-test-db/models/sample_type => ../models/sample-type/

replace github.com/BNext-IQT/covid-test-db/models/pcr_platform => ../models/pcr-platform/

go 1.14

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/360EntSecGroup-Skylar/excelize/v2 v2.1.0 // indirect
	github.com/PuerkitoBio/goquery v1.5.1
	github.com/gocolly/colly/v2 v2.1.0
	github.com/google/uuid v1.1.1 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/BNext-IQT/covid-test-db/models/company v0.0.1
	github.com/BNext-IQT/covid-test-db/models/diagnostic v0.0.1
	github.com/BNext-IQT/covid-test-db/models/diagnostic_target_type v0.0.1
	github.com/BNext-IQT/covid-test-db/models/diagnostic_type v0.0.1
	github.com/BNext-IQT/covid-test-db/models/pcr_platform v0.0.1
	github.com/BNext-IQT/covid-test-db/models/poc v0.0.1
	github.com/BNext-IQT/covid-test-db/models/regulatory_approval_type v0.0.1
	github.com/BNext-IQT/covid-test-db/models/sample_type v0.0.1
)
