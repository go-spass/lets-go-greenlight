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
│       └── main.go
├── go.mod
├── internal
├── LICENSE
├── Makefile
├── migrations
├── README.md
└── remote
```

* `bin` will contain our compiled application binaries, ready for deployment
  to a production server.
* `cmd/api` will contain the application-specific code for our Greenlight API
  application. This will include the code   for running the server, reading and
  writing HTTP requests, and managing authentication.
* `internal` will contain various ancillary packages used by our API. It will
  contain the code for interacting with our database, doing data validation,
  sending emails and so on. Basically, any code that isn’t application-specific
  and can potentially be reused will live in here. Our Go code under cmd/api
  will import the packages in the internal directory (but never the other way around).
* `migrations` will contain the SQL migration files for our database.
* `remote` will contain the configuration files and setup scripts for our
  production server.
* `go.mod` will declare our project dependencies, versions and module path.
* The `Makefile` will contain recipes for automating common administrative
  tasks — like auditing our Go code, building binaries, and executing
  database migrations.
