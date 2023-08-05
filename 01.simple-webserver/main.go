/*
Very simple package to understand:
1. configuration basic HTTP server with Graceful shutdown
2. simple middleware to setup the logging & HTTP header
3. simple JSON encoding
*/
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// declare the struct for user
type User struct {
	ID    int    `json:"id"`
	NAME  string `json:"name"`
	EMAIL string `json:"email"`
}

// greet funtion when root '/' is called
func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Gopher!!\nThe time is : %v", time.Now().UTC())
}

// create one user using the struct and display as JSON when '/user' is called
func getUser(w http.ResponseWriter, r *http.Request) {

	// create a new user goher with the struct User
	gopher := User{
		ID:    0,
		NAME:  "Gopher",
		EMAIL: "gopher@go.dev",
	}

	// encode the struct as JSON and write it to the http.ResponseWriter
	json.NewEncoder(w).Encode(gopher)
}

// middleware funtion to setup logging & HTTP header for JSON output
func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v is called from the client", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		f(w, r)
	}
}

func main() {
	log.Println("initializing the http server.")
	// get the listen address from the OS environment variable if not set the default
	addr := os.Getenv("HTTP_ADDR")
	if addr == "" {
		log.Println("HTTP_ADDR variable is not set in OS. Setting the value to default value")
		addr = ":8080"
	}
	log.Printf("Staring the server on - \"%v\"", addr)

	// initialize http mux, handlers & server
	mux := http.NewServeMux()

	mux.HandleFunc("/", logging(greet))
	mux.HandleFunc("/user", logging(getUser))

	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	log.Println("starting up the http server.")

	// create new channel that will be used to notify is the HTTP server is shutdown
	done := make(chan struct{})

	//  new go routine that takes the HTTP server & the channel as input for gracefulshutdown function
	go gracefulShutdown(srv, done)

	// starting the HTTP sever
	err := srv.ListenAndServe()
	if err != nil {
		log.Printf("Error starting the HTTP server - %v", err)
	}

	// channel waits until its return from the shutdown routine
	<-done

	log.Println("server is shutdown")
}

func gracefulShutdown(srv *http.Server, done chan struct{}) {
	// new channel that listens to the SIGTERM & OS Interrupt signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	<-quit

	// crating new context that sets a connection timeout value of 25 seconds to shutdown the HTTP server
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*25)
	defer cancel()

	log.Println("Caught signal... shutting down the server!")
	srv.Shutdown(ctx)

	//  closing the channel in the main funtion to notify that the server us shutdown.
	close(done)

}
