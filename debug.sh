#!/bin/bash

#stop existing
docker-compose -f docker-compose.yaml down 
sudo rm -rf cockroach-data/

docker-compose -f debug-compose.yaml up --build
docker build -f scraper/Dockerfile -t cdb_scraper .
docker exec -it roach \
	sh -c "/cockroach/cockroach sql --insecure < /sql_scripts/initialize.sql"

docker run --network="database" cdb_scraper