# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Working style — read this first

This repo is a learning exercise. **The user drives the implementation.** Do not
write, scaffold, or "get a head start on" the next section or chapter of the
book on your own, and do not fill in empty placeholder directories just because
they are empty. The point is for the user to type the new code themselves and
study it.

Only touch implementation code when explicitly asked, and keep the change to
exactly what was asked. When you spot something worth doing next, say so in a
sentence and stop — do not do it. Answering questions, explaining concepts,
reviewing code the user wrote, and tooling/config work (Makefile, CI, docs) are
fine when requested.

## Project

Greenlight — a JSON API for retrieving and managing movie records, built by
following Alex Edwards' book *Let's Go Further*. Module path is
`github.com/go-spass/lets-go-greenlight` (Go 1.27.1). The repo is currently a
scaffold: only `cmd/api/main.go` exists; `internal/`, `migrations/`, and
`remote/` are empty placeholders that fill in as the book progresses.

## Commands

All work goes through the `Makefile`; targets are chained so each one runs the
checks before it.

```sh
make            # == make all == make api
make api        # vet -> lint -> fmt, then build to ./bin/api
make vet        # lint -> fmt, then go vet ./...
make lint       # fmt, then golangci-lint run
make fmt        # go fmt ./...
make test       # go test -v -cover ./...
make bench      # go test -bench . ./...
make clean      # rm -r ./bin
```

`BUILD_DIR` overrides the output directory (default `./bin`).

Because `make api` re-runs `fmt`/`lint`/`vet`, a `golangci-lint` failure blocks
the build. For a quick compile-only check use `go build ./...` directly.

Single test / single package:

```sh
go test -v -run TestName ./internal/data/
go test -v -cover ./cmd/api/
```

## Architecture

The directory layout is the book's standard structure and the import direction
is a hard rule:

- `cmd/api/` — application-specific code: the HTTP server, handlers, routing,
  middleware, request/response helpers, and authentication. This is the only
  place that knows about HTTP.
- `internal/` — reusable, non-application-specific packages: database access,
  validation, mailer, and so on. **`cmd/api` imports `internal`; `internal`
  never imports `cmd`.**
- `migrations/` — SQL migration files (up/down pairs), applied with `migrate`.
- `remote/` — production server configuration and setup scripts.
- `bin/` — compiled binaries; git-ignored.

## Conventions

- Conventional commit prefixes (`feat:`, `fix:`, `test:`, `refactor:`,
  `docs:`, `chore:`).
- `.env` is git-ignored; keep DSNs and SMTP credentials out of the repo.
