# Introduction

Very simple package to understand:
1. configuration basic HTTP server with Graceful shutdown
2. simple middleware to setup the logging & HTTP header
3. simple JSON encoding

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
  - http://<HTTP_ADDR>/user - gives user info in JSON format