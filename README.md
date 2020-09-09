# covid test db

Database application to catalog tests for COVID-19 and the companies that make them.


# Deploy Instructions - Manual
  - clone repo
  - scp repo to home dir `scp -r ./covid-test-db/ user@ip_address:~/`
  - ssh  to machine
  - `cd covid-test-db`
  - `sudo ./stop.sh -ep` - stops any existing instances and clears the old DB
  - `sudo ./start.sh -c`
