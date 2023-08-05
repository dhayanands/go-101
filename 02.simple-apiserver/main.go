/*
Very simple package to understand echo framework to implement simple restapi service.
Package allows to create a user registry. can:
- create user entries
- get a specific entry
- list all the user entries
- update a specific user entry
- delete a specific user entry
*/
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

// user struct
type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// slice of struct for string all the users
var users []User

// funtion definitions
func greet(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Gopher!! 😀🥳")
}

// return the slice of structs
func listUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

// if userid is not already available, save the user data to the slice of struct
func saveUser(c echo.Context) error {
	id := c.FormValue("id")
	name := c.FormValue("name")
	email := c.FormValue("email")

	for _, user := range users {
		if user.Id == id {
			return c.String(http.StatusOK, "user id in use")
		}
	}
	u := User{
		Id:    id,
		Name:  name,
		Email: email,
	}

	users = append(users, u)

	return c.String(http.StatusOK, "user saved successfully")
}

// search for the userid and return the specific struct from slice
func getUser(c echo.Context) error {
	id := c.Param("id")

	for _, user := range users {
		if user.Id == id {
			return c.JSON(http.StatusOK, user)
		}
	}
	return c.String(http.StatusNotFound, "user not found")
}

// search for the userid, delete the old entry and create new entry
func updateUser(c echo.Context) error {
	id := c.Param("id")
	name := c.FormValue("name")
	email := c.FormValue("email")

	for i, user := range users {
		if user.Id == id {
			users = append(users[:i], users[i+1:]...)
			u := User{
				Id:    id,
				Name:  name,
				Email: email,
			}

			users = append(users, u)
			return c.String(http.StatusOK, "user details updated successfully")
		}
	}
	return c.String(http.StatusOK, "user id not found")
}

// search for the userid and remove the struct from slice using the append function
func deleteUser(c echo.Context) error {
	id := c.Param("id")
	for i, user := range users {
		if user.Id == id {
			users = append(users[:i], users[i+1:]...)
			log.Println(users)
		}
	}
	return c.JSON(http.StatusOK, users)
}

func main() {
	// define new echo instance
	e := echo.New()

	addr := os.Getenv("HTTP_ADDR")
	if addr == "" {
		log.Println("HTTP_ADDR variable is not set in OS. Setting the value to default value")
		addr = ":8080"
	}
	log.Printf("Staring the server on - \"%v\"", addr)

	// define routes
	e.GET("/", greet)
	e.POST("/user", saveUser)
	e.GET("/listusers", listUsers)
	e.GET("/user/:id", getUser)
	e.PUT("/user/:id", updateUser)
	e.DELETE("/user/:id", deleteUser)

	done := make(chan struct{})
	go gracefulShutdown(e, done)

	e.Logger.Fatal(e.Start(addr))

	<-done

}

func gracefulShutdown(e *echo.Echo, done chan struct{}) {

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	<-quit
	log.Println("caught signal. shutting down the server.")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*25)
	defer cancel()

	e.Shutdown(ctx)
	log.Println("server is shutdown")

	close(done)
}
