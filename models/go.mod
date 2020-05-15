module gitlab.iqt.org/rashley/covid-test-db/models

replace gitlab.iqt.org/rashley/covid-test-db/models/company => ./company/
replace gitlab.iqt.org/rashley/covid-test-db/models/poc => ./poc/
replace gitlab.iqt.org/rashley/covid-test-db/models/diagnostic  => ./diagnostic/
replace gitlab.iqt.org/rashley/covid-test-db/models/diagnostic_type  => ./diagnostic-type/
replace gitlab.iqt.org/rashley/covid-test-db/models/test_target_type => ./diagnostic-target-type/
replace gitlab.iqt.org/rashley/covid-test-db/models/regulatory_approval_type => ./regulatory-approval-type/
replace gitlab.iqt.org/rashley/covid-test-db/models/sample_type  => ./sample-type/

go 1.14
