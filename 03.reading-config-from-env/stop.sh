docker-compose down
unset DB_PG_HOST
unset DB_PG_USERNAME
unset DB_PG_PASSWORD
unset DB_PG_DATABASE
unset DB_PG_PORT
unset ENV

rm -rf ./postgres-data/*