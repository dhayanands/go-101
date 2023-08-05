# Introduction

Very simple package to understand echo framework to implement simple restapi service.
Package allows to create a user registry. can:
- create user entries
- get a specific entry
- list all the user entries
- update a specific user entry
- delete a specific user entry

## How to use:

- Setup the HTTP_ADDR env variable (if not default value ":8080" will be used)
```bash
export HTTP_ADDR=":8080"
```

- start the webserver
```bash
go run main.go
```

- use browser / curl / API clients like Postman to access the below URLs:
  - http://<HTTP_ADDR>/ - gives a greeting message
  - http://<HTTP_ADDR>/listusers - gives list of all users info in JSON format

  ```bash
  # add user
  # curl -X "POST" -d "id=<value>" -d "name=<value>" -d "email=<value>" http://localhost:8080/user/

  curl -X "POST" -d "id=1" -d "name=dhayanand" -d "email=dhayanand@dhayanand.com" http://localhost:8080/user
  curl -X "POST" -d "id=2" -d "name=support" -d "email=support@dhayanand.com" http://localhost:8080/user

  # list all users
  curl -X "GET" http://localhost:8080/listusers

  # get specific user
  # curl -X "GET" http://localhost:8080/user/{id}
  
  curl -X "GET" http://localhost:8080/user/2

  # update user
  # curl -X "PUT" -d "name=<new_value>" -d "email=<new_value>" http://localhost:8080/user/{id}

  curl -X "PUT" -d "name=user" -d "email=user@example.com" http://localhost:8080/user/2
  curl -X "GET" http://localhost:8080/listusers

  # delete user
  # curl -X "DELETE" http://localhost:8080/user/{id}

  curl -X "DELETE" http://localhost:8080/user/2
  curl -X "GET" http://localhost:8080/listusers


  ```
