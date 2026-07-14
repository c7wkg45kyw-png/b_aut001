# BAUT001 Backend

Authentication module for decoupled ERP services. It issues a single HMAC SHA256 JWT containing merchant_id, audience, client_id, type, and scopes.

## Run

```bash
cp .env.example .env
go run ./cmd/server
```

Swagger UI: `http://localhost:8081/docs/swagger.html`

## Default Seed User

Configured in `.env`:

- username: `admin`
- password: `admin1234`
- merchant: `mch_555666777`

## Login Example

```json
{
  "username": "admin",
  "password": "admin1234"
}
```

## Swagger Failed to Fetch

Open Swagger through the running backend URL, for example `http://localhost:8081/docs/swagger.html`. Do not open docs/swagger.html directly from the filesystem, because browser CORS requires an `http` or `https` URL scheme. The OpenAPI server URL uses `/api/v1` relative to the current host.
# b_aut001
