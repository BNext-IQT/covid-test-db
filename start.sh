#!/bin/bash

RUN_INIT=0

function show_help()
{
    echo "covid test db start (uses sudo)

    Usage: start [option]
    Options:
      -c,  create     run db initialization
      -h,  help       print this help"
}

function check_args()
{
    while [ $# -gt 0 ]; do
        case $1 in
            -c|create)
                RUN_INIT=1
                shift
                ;;
            -h|\?|help)
                show_help
                exit
                ;;
        esac
        shift
    done
}

if [ $# -gt 0 ]; then
    check_args "$@"
fi

docker-compose -f docker-compose.yaml up -d 

if [ "$RUN_INIT" -gt 0 ]; then
	docker exec -it roach \
	sh -c "/cockroach/cockroach sql --insecure < /sql_scripts/initialize.sql"
fi