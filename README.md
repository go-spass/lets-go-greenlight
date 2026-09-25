# lets-go-greenlight
Greenlight application from the book "Let's Go Further" by Alex Edwards

## Project Structure

Let’s take a moment to talk through these files and folders and explain the
purpose they’ll serve in our finished project.

```text
.
├── bin
├── cmd
│   └── api
│       ├── healthcheck.go
│       ├── helpers.go
│       ├── helpers_test.go
│       ├── main.go
│       ├── movies.go
│       └── routes.go
├── go.mod
├── go.sum
├── internal
├── LICENSE
├── Makefile
├── migrations
├── README.md
└── remote
```

* `bin` will contain our compiled application binaries, ready for
  deployment
  to a production server.
* `cmd/api` will contain the application-specific code for our
  Greenlight API
  application. This will include the code for running the server,
  reading and writing HTTP requests, and managing authentication.
* `internal` will contain various ancillary packages used by our
  API. It will
  contain the code for interacting with our database, doing data
  validation, sending emails and so on. Basically, any code that
  isn’t application-specific
  and can potentially be reused will live in here. Our Go code under
  cmd/api will import the packages in the internal directory (but
  never the other way around).
* `migrations` will contain the SQL migration files for our
  database.
* `remote` will contain the configuration files and setup scripts
  for our production server.
* `go.mod` will declare our project dependencies, versions and
  module path.
* The `Makefile` will contain recipes for automating common
  administrative tasks — like auditing our Go code, building binaries,
  and executing database migrations.

## APIs
API endpoints and RESTful routing
Over the next few sections of this book, we’re going to gradually build up our API so that the endpoints start to look like this:

| Method | URL Pattern     | Handler            | Action                                 |
| :----- | :-------------- | :----------------- | :------------------------------------- |
| GET    | /v1/healthcheck | healthcheckHandler | Show application information           |
| GET    | /v1/movies      | listMoviesHandler  | Show the details of all movies         |
| POST   | /v1/movies      | createMovieHandler | Create a new movie                     |
| GET    | /v1/movies/:id  | showMovieHandler   | Show the details of a specific movie   |
| PUT    | /v1/movies/:id  | editMovieHandler   | Update the details of a specific movie |
| DELETE | /v1/movies/:id  | deleteMovieHandler | Delete a specific movie                |

> Note: The book recommends using [httprouter](https://github.com/julienschmidt/httprouter)
> which I will do as I follow along, however I am not convinced this is actually a good choice.
> Something to revisit in the future.
