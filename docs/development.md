# Development

Install the Go toolchain matching go.mod, or use the official Go container. PostgreSQL connections require certificate-verified TLS; the HTTP listener requires a certificate and key. No plaintext development exception is built in.

## Verify

```sh
go mod download
go vet ./...
go test -race ./...
go build -o bin/bao ./cmd/bao
```

The optional database integration test uses BAO_TEST_DATABASE_URL. Supply a dedicated PostgreSQL database with sslmode=verify-full and a trusted certificate. Without it the integration test explicitly skips; a unit-test pass is not a successful database integration check.

## Run

Set every environment variable listed in .env.example. The example contains illustrative timeout and capacity settings; validate operational settings for each environment. Inject the database URL from a secret manager. The executable does not load dotenv files.

```sh
go run ./cmd/bao
```

GET /health/live returns 204 while serving. GET /health/ready returns 204 when PostgreSQL responds or 503 when unavailable. Neither endpoint returns dependency credentials or errors. The service accepts TLS 1.3 and newer. Configure trusted server certificates before running.

## Layout

- cmd/bao is the composition root and lifecycle owner.
- internal/platform holds configuration and PostgreSQL adapters.
- internal/health/query holds the readiness query and consumer-owned dependency interface.
- internal/health/httpapi holds HTTP transport.
- migrations holds versioned administrative SQL.

Add domain, application command/query and infrastructure packages inside each bounded context as actual stories require them. Domain code must not import transport, database or provider adapters. Do not introduce placeholder business logic to fill directories.

## Production gates

This initialisation does not implement business authentication, tenant data, a cloud deployment, managed keys, storage encryption, encrypted backups, a broker, an outbox or a migration history runner. Implement and verify these in their linked stories before handling client data. See docs/design/foundation.md for boundaries and failure behaviour.
