module gitlab.iqt.org/rashley/covid-test-db/api

go 1.14

replace gitlab.iqt.org/rashley/covid-test-db/models/poc => ../models/poc/

replace gitlab.iqt.org/rashley/covid-test-db/models/company => ../models/company/

replace gitlab.iqt.org/rashley/covid-test-db/models/diagnostic => ../models/diagnostic/

replace gitlab.iqt.org/rashley/covid-test-db/models/diagnostic_type => ../models/diagnostic-type/

replace gitlab.iqt.org/rashley/covid-test-db/models/diagnostic_target_type => ../models/diagnostic-target-type/

replace gitlab.iqt.org/rashley/covid-test-db/models/regulatory_approval_type => ../models/regulatory-approval-type/

require (
	github.com/google/uuid v1.1.1
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.4
	github.com/jinzhu/gorm v1.9.12
	github.com/lib/pq v1.3.0
	gitlab.iqt.org/rashley/covid-test-db/models/company v0.0.0-00010101000000-000000000000 // indirect
	gitlab.iqt.org/rashley/covid-test-db/models/diagnostic v0.0.0-00010101000000-000000000000
	gitlab.iqt.org/rashley/covid-test-db/models/diagnostic_target_type v0.0.0-00010101000000-000000000000 // indirect
	gitlab.iqt.org/rashley/covid-test-db/models/diagnostic_type v0.0.0-00010101000000-000000000000 // indirect
	gitlab.iqt.org/rashley/covid-test-db/models/poc v0.0.0-00010101000000-000000000000
	gitlab.iqt.org/rashley/covid-test-db/models/regulatory_approval_type v0.0.0-00010101000000-000000000000 // indirect
)
