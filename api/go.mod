module github.com/BNext-IQT/covid-test-db/api

go 1.14

replace github.com/BNext-IQT/covid-test-db/models/poc => ../models/poc/

replace github.com/BNext-IQT/covid-test-db/models/company => ../models/company/

replace github.com/BNext-IQT/covid-test-db/models/diagnostic => ../models/diagnostic/

replace github.com/BNext-IQT/covid-test-db/models/diagnostic_type => ../models/diagnostic-type/

replace github.com/BNext-IQT/covid-test-db/models/diagnostic_target_type => ../models/diagnostic-target-type/

replace github.com/BNext-IQT/covid-test-db/models/regulatory_approval_type => ../models/regulatory-approval-type/

replace github.com/BNext-IQT/covid-test-db/models/sample_type => ../models/sample-type/

replace github.com/BNext-IQT/covid-test-db/models/pcr_platform => ../models/pcr-platform/

require (
	github.com/BNext-IQT/covid-test-db/models/company v0.0.1 // indirect
	github.com/BNext-IQT/covid-test-db/models/diagnostic v0.0.1
	github.com/BNext-IQT/covid-test-db/models/diagnostic_target_type v0.0.1 // indirect
	github.com/BNext-IQT/covid-test-db/models/diagnostic_type v0.0.1
	github.com/BNext-IQT/covid-test-db/models/pcr_platform v0.0.1
	github.com/BNext-IQT/covid-test-db/models/poc v0.0.1
	github.com/BNext-IQT/covid-test-db/models/regulatory_approval_type v0.0.1
	github.com/BNext-IQT/covid-test-db/models/sample_type v0.0.1
	github.com/google/uuid v1.1.1
	github.com/gorilla/handlers v1.5.0
	github.com/gorilla/mux v1.8.0
	github.com/jinzhu/gorm v1.9.12
	github.com/lib/pq v1.8.0
)
