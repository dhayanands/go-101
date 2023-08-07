/*
Very simple package to understand how to read configration for the application from environment variables.
Read the database connection string from the environment variable, make connectiion to a postgres database, run a query and print the result.
*/

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/dhayanands/go-101/03.reading-config-from-env/config"
	"github.com/jackc/pgx/v4"
)

const (
	MsgConfigSetupError string = "config.dependencies.error"
	// connStrExample1 = "postgres://username:password@localhost:5432/database_name?sslmode=disable&pool_max_conns=10"
	// connStrExample2 = "user=%s password=%s database=%s host=%s sslmode=disable port=%s"
	connStr string = "postgres://%s:%s@%s:%s/%s"
)

func main() {
	log.Println("reading configurations.")

	cfg, key, err := config.New(config.WithDBConfig())
	if err != nil {
		log.Fatalf("%s : %s %v \n", MsgConfigSetupError, key, err)
	}

	conn, err := prepareConnString(cfg)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("reading configurations completed.")

	log.Println("connecting to database")
	db, err := pgx.Connect(context.Background(), conn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())

	err = db.Ping(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to ping to database: %v\n", err)
		os.Exit(1)
	}
	log.Println("Successfully connected to db.")

	var quote string
	err = db.QueryRow(context.Background(), "select quote from quote where author ilike 'robert%'").Scan(&quote)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("QUOTE: %s\n", quote)

	// fmt.Println(cfg, "\n", conn)
}

func prepareConnString(cfg *config.Config) (string, error) {
	// "postgres://postgres:mypassword@localhost:5432/postgres"
	return fmt.Sprintf(connStr, cfg.DBUser(), cfg.DBPass(), cfg.DBHost(), cfg.DBPort(), cfg.DBSchema()), nil
}
