# Introduction

Very simple package to understand how to read configration for the application from environment variables.
Read the database connection string from the environment variable, make connectiion to a postgres database, run a query and print the result.

## How to use:

- Set the environment variables and start the postgres database using the bash script
```bash
export DB_PG_HOST=127.0.0.1
export DB_PG_USERNAME=postgres
export DB_PG_PASSWORD=postgres
export DB_PG_DATABASE=postgres
export DB_PG_PORT=5432
export ENV=DEV

sh start.sh
```

- run the applicatio
```bash
go run main.go
```

- Unset the environment variables and stop the postgres database using the bash script
```bash
sh stop.sh
rm -rf ./postgres-data

unset DB_PG_HOST
unset DB_PG_USERNAME
unset DB_PG_PASSWORD
unset DB_PG_DATABASE
unset DB_PG_PORT
unset ENV
```