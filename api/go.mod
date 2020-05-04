module gitlab.iqt.org/rashley/covid-test-db/api

go 1.14

replace gitlab.iqt.org/rashley/covid-test-db/models/poc => ../models/poc/
replace gitlab.iqt.org/rashley/covid-test-db/models/test_type  => ../models/test-type/
replace gitlab.iqt.org/rashley/covid-test-db/models/test_target_type => ../models/test-target-type/
replace gitlab.iqt.org/rashley/covid-test-db/models/regulatory_approval_type => ../models/regulatory-approval-type/

require (
	github.com/google/uuid v1.1.1
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.4
	github.com/jinzhu/gorm v1.9.12
	github.com/lib/pq v1.3.0
)
