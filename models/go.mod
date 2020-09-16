module github.com/BNext-IQT/covid-test-db/models

replace github.com/BNext-IQT/covid-test-db/models/company => ./company/
replace github.com/BNext-IQT/covid-test-db/models/poc => ./poc/
replace github.com/BNext-IQT/covid-test-db/models/diagnostic  => ./diagnostic/
replace github.com/BNext-IQT/covid-test-db/models/diagnostic_type  => ./diagnostic-type/
replace github.com/BNext-IQT/covid-test-db/models/diagnostic_performance  => ./diagnostic-performance/
replace github.com/BNext-IQT/covid-test-db/models/test_target_type => ./diagnostic-target-type/
replace github.com/BNext-IQT/covid-test-db/models/regulatory_approval_type => ./regulatory-approval-type/
replace github.com/BNext-IQT/covid-test-db/models/sample_type  => ./sample-type/
replace github.com/BNext-IQT/covid-test-db/models/pcr_platform  => ./pcr-platform/

go 1.14
