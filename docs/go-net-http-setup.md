# Go `net/http` Setup Guide

This guide is the practical starting point for building the Go-Trade HTTP/API layer with the Go standard library.

It assumes a standard-library-first approach:

- `net/http` for the server and routing
- `http.ServeMux` for route registration
- `encoding/json` for JSON request and response handling
- `cobra` only for CLI/admin commands outside the API layer

If you are using Go 1.22 or newer, you can use the modern `ServeMux` routing patterns such as `GET /accounts/{id}` and read wildcard values with `r.PathValue("id")`.

## Official Reading Path

Start with these official Go resources:

1. Install Go:
   [Download and install Go](https://go.dev/doc/install)
2. Learn the basics of modules and `go run`:
   [Tutorial: Get started with Go](https://go.dev/doc/tutorial/getting-started)
3. Read the package docs you will use most:
   [`net/http` package documentation](https://pkg.go.dev/net/http)
4. Learn the modern router behavior introduced in Go 1.22:
   [Routing Enhancements for Go 1.22](https://go.dev/blog/routing-enhancements)
5. See the official release notes for the routing changes:
   [Go 1.22 release notes](https://go.dev/doc/go1.22)
6. If you want a longer web-app walkthrough:
   [Writing Web Applications](https://go.dev/doc/articles/wiki/)

## Minimal Project Setup

Create a new Go module:

```bash
mkdir go-trade
cd go-trade
go mod init github.com/yourname/go-trade
```

Create a `main.go` file:

```go
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	mux.HandleFunc("GET /accounts/{id}", func(w http.ResponseWriter, r *http.Request) {
		accountID := r.PathValue("id")

		writeJSON(w, http.StatusOK, map[string]string{
			"account_id": accountID,
		})
	})

	mux.HandleFunc("POST /orders", func(w http.ResponseWriter, r *http.Request) {
		type createOrderRequest struct {
			Symbol string  `json:"symbol"`
			Side   string  `json:"side"`
			Size   float64 `json:"size"`
		}

		var req createOrderRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{
				"error": "invalid JSON body",
			})
			return
		}

		writeJSON(w, http.StatusCreated, map[string]any{
			"accepted": true,
			"order":    req,
		})
	})

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("API listening on http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
```

Run it:

```bash
go run .
```

## Quick Test Commands

Health check:

```bash
curl http://localhost:8080/health
```

Path parameter example:

```bash
curl http://localhost:8080/accounts/abc123
```

JSON POST example:

```bash
curl -X POST http://localhost:8080/orders \
  -H "Content-Type: application/json" \
  -d '{"symbol":"BTCUSDT","side":"buy","size":0.1}'
```

## What `net/http` Gives You

With the standard library you already have:

- an HTTP server
- request and response types
- routing with `ServeMux`
- method-aware route patterns in Go 1.22+
- path wildcards in Go 1.22+
- middleware using plain `http.Handler`
- timeouts and server configuration
- testing support via `net/http/httptest`

That is enough for a clean first API for Go-Trade.

## How Routing Looks in Modern Go

Examples of route patterns with `ServeMux` in Go 1.22+:

- `GET /health`
- `POST /orders`
- `GET /accounts/{id}`
- `GET /leaders/{leaderID}/followers/{followerID}`
- `GET /files/{path...}`

Example of reading a wildcard:

```go
accountID := r.PathValue("id")
```

If you are on a version older than Go 1.22, these newer route patterns will not behave the same way. In that case, upgrade Go or use simpler path matching until you do.

## Recommended Next Packages

After `net/http`, the next standard-library packages worth learning are:

- [`encoding/json`](https://pkg.go.dev/encoding/json) for API bodies
- [`context`](https://pkg.go.dev/context) for request-scoped cancellation and deadlines
- [`net/http/httptest`](https://pkg.go.dev/net/http/httptest) for handler tests
- [`log/slog`](https://pkg.go.dev/log/slog) for structured logging
- [`database/sql`](https://pkg.go.dev/database/sql) for persistence boundaries

## Suggested First API Shape for Go-Trade

Once the server is running, a sensible first set of endpoints would be:

- `GET /health`
- `GET /exchanges`
- `POST /accounts`
- `GET /accounts/{id}`
- `POST /strategies`
- `POST /orders`
- `GET /positions`

That is enough to begin wiring exchange adapters, account records, and risk-checked order flow without introducing unnecessary framework complexity.
