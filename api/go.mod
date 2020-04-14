module gitlab.iqt.org/rashley/covid-test-db/api

go 1.14

replace gitlab.iqt.org/rashley/covid-test-db/api/models => ./models/

require (
	github.com/google/uuid v1.1.1
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.4
	github.com/jinzhu/gorm v1.9.12
	github.com/lib/pq v1.3.0
)
