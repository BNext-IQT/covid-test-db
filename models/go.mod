module gitlab.iqt.org/rashley/covid-test-db/models

replace gitlab.iqt.org/rashley/covid-test-db/models/company => ./company/
replace gitlab.iqt.org/rashley/covid-test-db/models/poc => ./poc/
replace gitlab.iqt.org/rashley/covid-test-db/models/test_type  => ./test_type/
replace gitlab.iqt.org/rashley/covid-test-db/models/test_target_type => ./test_target_type/
replace gitlab.iqt.org/rashley/covid-test-db/models/regulatory_approval_type => ./regulatory_approval_type/

go 1.14
