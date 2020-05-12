module gitlab.iqt.org/rashley/covid-test-db/scraper

replace gitlab.iqt.org/rashley/covid-test-db/models/poc => ../models/poc/

replace gitlab.iqt.org/rashley/covid-test-db/models/diagnostic_type => ../models/diagnostic-type/

replace gitlab.iqt.org/rashley/covid-test-db/models/diagnostic_target_type => ../models/diagnostic-target-type/

replace gitlab.iqt.org/rashley/covid-test-db/models/regulatory_approval_type => ../models/regulatory-approval-type/

go 1.14

require (
	github.com/PuerkitoBio/goquery v1.5.1
	github.com/gocolly/colly/v2 v2.0.1
	github.com/google/uuid v1.1.1 // indirect
	github.com/jinzhu/gorm v1.9.12
	gitlab.iqt.org/rashley/covid-test-db/models/poc v0.0.0-00010101000000-000000000000
	gitlab.iqt.org/rashley/covid-test-db/models/regulatory_approval_type v0.0.0-00010101000000-000000000000
	gitlab.iqt.org/rashley/covid-test-db/models/diagnostic_target_type v0.0.0-00010101000000-000000000000
	gitlab.iqt.org/rashley/covid-test-db/models/diagnostic_type v0.0.0-00010101000000-000000000000
)
