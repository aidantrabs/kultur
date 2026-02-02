# KULTUR Backend

Go API server for the KULTUR platform.

## Tech Stack

- **Go 1.24** - Primary language
- **Echo** - HTTP framework
- **pgx** - PostgreSQL driver
- **sqlc** - Type-safe SQL code generation
- **Resend** - Email delivery

## Setup

### Prerequisites

- Go 1.24+
- PostgreSQL (or Docker)

### Install Dependencies

```bash
go mod download
```

### Environment Variables

Create a `.env` file (or copy from `.env.example`):

```env
PORT=8080
DATABASE_URL=postgres://postgres:postgres@localhost:5432/hackathon?sslmode=disable
RESEND_API_KEY=re_...
ALLOWED_ORIGINS=http://localhost:5173
ADMIN_API_KEY=your-secret-admin-key
BASE_URL=http://localhost:8080
FROM_EMAIL=noreply@kultur-tt.app
```

| Variable | Description |
|:---------|:------------|
| `PORT` | Server port |
| `DATABASE_URL` | PostgreSQL connection string |
| `RESEND_API_KEY` | Resend API key for emails |
| `ALLOWED_ORIGINS` | CORS allowed origins (comma-separated) |
| `ADMIN_API_KEY` | API key for admin endpoints |
| `BASE_URL` | Base URL for email links |
| `FROM_EMAIL` | Sender email address |

## Development

### Start PostgreSQL (Docker)

```bash
docker-compose up -d
```

### Run Server

```bash
go run ./cmd/server
```

Server starts at http://localhost:8080

### Run Migrations

```bash
goose -dir sql/migrations postgres "$DATABASE_URL" up
```

## Project Structure

```
backend/
├── cmd/
│   └── server/
│       └── main.go             # Entry point
├── internal/
│   ├── config/                 # Environment loading
│   ├── db/                     # Database connection + sqlc
│   ├── email/
│   │   ├── service.go          # Email service
│   │   └── templates.go        # HTML templates
│   ├── handler/
│   │   ├── handler.go          # Handler struct
│   │   ├── festivals.go        # Festival endpoints
│   │   ├── memories.go         # Memory endpoints
│   │   ├── subscriptions.go    # Subscription endpoints
│   │   └── email_testing.go    # Test email endpoints
│   ├── middleware/
│   │   ├── auth.go             # API key auth
│   │   └── ratelimit.go        # Rate limiting
│   └── service/                # Business logic
├── sql/
│   ├── migrations/             # Database migrations
│   └── queries/                # sqlc query definitions
├── Dockerfile
└── go.mod
```

## API Endpoints

### Public

| Method | Endpoint | Description |
|:-------|:---------|:------------|
| GET | `/health` | Health check |
| GET | `/api/festivals` | List all festivals |
| GET | `/api/festivals/upcoming` | Upcoming festivals |
| GET | `/api/festivals/calendar` | Festivals by year |
| GET | `/api/festivals/:slug` | Get festival |
| GET | `/api/festivals/:slug/memories` | Get memories |
| POST | `/api/memories` | Submit memory (5/hr limit) |
| POST | `/api/subscribe` | Subscribe (10/hr limit) |
| GET | `/api/subscribe/confirm/:token` | Confirm subscription |
| GET | `/api/unsubscribe/:token` | Unsubscribe |

### Admin (requires `X-API-Key` header)

| Method | Endpoint | Description |
|:-------|:---------|:------------|
| GET | `/api/admin/memories` | List all memories |
| PATCH | `/api/admin/memories/:id` | Update memory status |
| DELETE | `/api/admin/memories/:id` | Delete memory |
| GET | `/api/admin/subscriptions` | List subscriptions |
| DELETE | `/api/admin/subscriptions/:id` | Delete subscription |
| POST | `/api/admin/festivals` | Create festival |
| PUT | `/api/admin/festivals/:id` | Update festival |
| DELETE | `/api/admin/festivals/:id` | Delete festival |
| POST | `/api/admin/test-email/welcome` | Test welcome email |
| POST | `/api/admin/test-email/reminder` | Test reminder email |
| POST | `/api/admin/test-email/digest` | Test digest email |

## Build

```bash
go build -o bin/server ./cmd/server
```

## Deployment

Deployed to Cloud Run at [kultur-api-971304624476.us-central1.run.app](https://kultur-api-971304624476.us-central1.run.app).

```bash
gcloud run deploy kultur-api \
  --source . \
  --region us-central1 \
  --allow-unauthenticated
```

### Environment Variables

Set via Cloud Run console or CLI:

```bash
gcloud run services update kultur-api --region us-central1 \
  --set-secrets="DATABASE_URL=DATABASE_URL:latest" \
  --set-secrets="RESEND_API_KEY=RESEND_API_KEY:latest" \
  --set-secrets="ADMIN_API_KEY=ADMIN_API_KEY:latest" \
  --set-env-vars="FROM_EMAIL=noreply@kultur-tt.app" \
  --set-env-vars="BASE_URL=https://kultur-api-971304624476.us-central1.run.app" \
  --set-env-vars="ALLOWED_ORIGINS=https://kultur-tt.app"
```
