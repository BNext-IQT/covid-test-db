#!/bin/bash

EXTREME=0

function show_help()
{
    echo "covid test db start (uses sudo)

    Usage: start [option]
    Options:
      -ep, extreme_prejudice     delete database files. ALL DATA WILL BE LOST!
      -h,  help       			 print this help"
}

function check_args()
{
    while [ $# -gt 0 ]; do
        case $1 in
            -ep|extreme_prejudice)
                EXTREME=1
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

if [[ "$EXTREME" -gt 0 ]]; then
	sudo rm -rf cockroach-data/
fi

docker-compose -f docker-compose.yaml down 
