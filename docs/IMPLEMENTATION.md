# Implementation Details

Technical deep-dive into the KULTUR platform architecture, decisions, and lessons learned.

---

## Table of Contents

- [Architecture Overview](#architecture-overview)
- [Tech Stack](#tech-stack)
- [Frontend](#frontend)
- [Backend](#backend)
- [Database](#database)
- [Email System](#email-system)
- [Deployment](#deployment)
- [Issues & Solutions](#issues--solutions)
- [Security](#security)
- [Performance](#performance)

---

## Architecture Overview

```
┌────────────────────────────────────────────────────────────────────────────┐
│                              PRODUCTION                                     │
│                                                                            │
│  ┌──────────────────────┐         ┌──────────────────────┐                │
│  │                      │         │                      │                │
│  │   kultur-tt.app      │  HTTPS  │  kultur-api.fly.dev  │                │
│  │   ──────────────     │────────▶│  ─────────────────   │                │
│  │                      │         │                      │                │
│  │   SvelteKit 2        │         │   Go 1.24 + Echo     │                │
│  │   Svelte 5           │         │                      │                │
│  │   TailwindCSS 4      │         │                      │                │
│  │   shadcn-svelte      │         │                      │                │
│  │                      │         │                      │                │
│  └──────────────────────┘         └──────────┬───────────┘                │
│           │                                  │                            │
│           │ Vercel                           │ Fly.io                     │
│           │                                  │                            │
│           │                       ┌──────────┴───────────┐                │
│           │                       │                      │                │
│           │                       │   PostgreSQL         │                │
│           │                       │   ────────────       │                │
│           │                       │   Neon (Serverless)  │                │
│           │                       │                      │                │
│           │                       └──────────────────────┘                │
│           │                                                               │
│           │                       ┌──────────────────────┐                │
│           │                       │                      │                │
│           └──────────────────────▶│   Resend             │                │
│             (logo in emails)      │   ────────────       │                │
│                                   │   noreply@           │                │
│                                   │   kultur-tt.app      │                │
│                                   │                      │                │
│                                   └──────────────────────┘                │
│                                                                            │
└────────────────────────────────────────────────────────────────────────────┘
```

### Data Flow

1. **User visits** `kultur-tt.app` → Vercel serves SvelteKit app
2. **App fetches data** from `kultur-api.fly.dev/api/*`
3. **Backend queries** Neon PostgreSQL database
4. **Email actions** trigger Resend API calls
5. **Emails include** logo from `kultur-tt.app/logo.png`

---

## Tech Stack

### Frontend

| Technology | Version | Purpose |
|:-----------|:--------|:--------|
| SvelteKit | 2.50 | Full-stack framework, SSR, routing |
| Svelte | 5.48 | Reactive UI with runes (`$state`, `$derived`) |
| TailwindCSS | 4.1 | Utility-first CSS |
| shadcn-svelte | 1.1 | UI component library |
| bits-ui | 2.15 | Headless UI primitives |
| Biome | 1.9 | Linting and formatting |
| TypeScript | 5.9 | Type safety |
| Vite | 7.3 | Build tool and dev server |

### Backend

| Technology | Version | Purpose |
|:-----------|:--------|:--------|
| Go | 1.24 | Primary language |
| Echo | 4.15 | HTTP framework |
| pgx | 5.8 | PostgreSQL driver |
| sqlc | - | Type-safe SQL code generation |
| Resend SDK | 2.28 | Email delivery |
| godotenv | 1.5 | Environment variable loading |

### Infrastructure

| Service | Provider | Purpose |
|:--------|:---------|:--------|
| Frontend Hosting | Vercel | CDN, automatic deploys |
| Backend Hosting | Fly.io | Container hosting, global edge |
| Database | Neon | Serverless PostgreSQL |
| Email | Resend | Transactional emails |
| Domain | Name.com | kultur-tt.app |
| DNS | Vercel + Resend | Frontend + email routing |

---

## Frontend

### Project Structure

```
frontend/
├── src/
│   ├── lib/
│   │   ├── api.ts              # API client
│   │   ├── config.ts           # Environment config
│   │   ├── utils.ts            # Utility functions
│   │   ├── components/
│   │   │   ├── ui/             # shadcn-svelte components
│   │   │   ├── calendar/       # Calendar components
│   │   │   └── newsletter/     # Newsletter signup
│   │   ├── data/
│   │   │   ├── festivals.ts    # Festival data + API calls
│   │   │   └── memories.ts     # Memory data
│   │   └── types/
│   │       └── festival.ts     # TypeScript types
│   └── routes/
│       ├── +page.svelte        # Home page
│       ├── +layout.svelte      # Root layout
│       └── festivals/
│           ├── +page.svelte    # Calendar view
│           └── [slug]/
│               ├── +page.svelte    # Festival detail
│               └── +page.ts        # Data loading
├── static/
│   └── logo.png                # Optimized logo (48KB)
├── biome.json                  # Linting config
└── package.json
```

### Key Design Decisions

**Svelte 5 Runes**
We use the new Svelte 5 runes API (`$state`, `$derived`, `$props`) instead of the legacy reactive syntax. This provides better TypeScript inference and more explicit reactivity.

```svelte
<script lang="ts">
    let email = $state('');
    let isSubmitting = $state(false);
    const isValid = $derived(email.includes('@'));
</script>
```

**Mock Data Fallback**
The frontend can operate in two modes controlled by `PUBLIC_DATA_SOURCE`:
- `api` - Fetches from backend API
- `mock` - Uses local mock data (for development without backend)

```typescript
// src/lib/config.ts
export const config = {
    apiUrl: env.PUBLIC_API_URL || 'http://localhost:8080',
    dataSource: (env.PUBLIC_DATA_SOURCE || 'mock') as 'mock' | 'api',
    get useApi() {
        return this.dataSource === 'api';
    }
};
```

**Component Library**
We use shadcn-svelte, which provides unstyled, accessible components that we customize with Tailwind. This gives us full control over styling while maintaining accessibility.

---

## Backend

### Project Structure

```
backend/
├── cmd/
│   └── server/
│       └── main.go             # Entry point, route setup
├── internal/
│   ├── config/
│   │   └── config.go           # Environment loading
│   ├── db/
│   │   ├── db.go               # Connection pool
│   │   ├── models.go           # sqlc generated types
│   │   └── queries.sql.go      # sqlc generated queries
│   ├── email/
│   │   ├── service.go          # Email service
│   │   └── templates.go        # HTML email templates
│   ├── handler/
│   │   ├── handler.go          # Handler struct
│   │   ├── festivals.go        # Festival endpoints
│   │   ├── memories.go         # Memory endpoints
│   │   ├── subscriptions.go    # Subscription endpoints
│   │   └── email_testing.go    # Test email endpoints
│   ├── middleware/
│   │   ├── auth.go             # API key authentication
│   │   └── ratelimit.go        # Rate limiting
│   └── service/
│       ├── festival.go         # Festival business logic
│       ├── memory.go           # Memory business logic
│       └── subscription.go     # Subscription logic
├── sql/
│   ├── migrations/             # Database migrations
│   └── queries/                # sqlc query definitions
├── Dockerfile
├── fly.toml
└── go.mod
```

### Key Design Decisions

**Echo Framework**
Chose Echo over Chi/Gin for its middleware ecosystem and cleaner error handling. Echo's context-based handlers made it easy to implement rate limiting and authentication.

**sqlc for Database**
Instead of an ORM, we use sqlc which generates type-safe Go code from SQL queries. This gives us:
- Full SQL control
- Compile-time type checking
- Zero runtime reflection

```sql
-- queries/festivals.sql
-- name: GetFestivalBySlug :one
SELECT * FROM festivals WHERE slug = $1 AND is_published = true;
```

Generates:
```go
func (q *Queries) GetFestivalBySlug(ctx context.Context, slug string) (Festival, error)
```

**Service Layer Pattern**
Business logic is separated into services, keeping handlers thin:

```
Handler (HTTP) → Service (Business Logic) → Queries (Database)
```

---

## Database

### Schema

```sql
-- Festivals
CREATE TABLE festivals (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    slug VARCHAR(100) UNIQUE NOT NULL,
    name VARCHAR(200) NOT NULL,
    date_type VARCHAR(20) NOT NULL,        -- fixed, lunar, movable
    usual_month VARCHAR(50),
    region VARCHAR(50) NOT NULL,           -- north, south, east, west, central, nationwide
    heritage_type VARCHAR(50) NOT NULL,    -- african, indian, indigenous, mixed, christian
    festival_type VARCHAR(50) NOT NULL,    -- religious, cultural, national
    summary TEXT NOT NULL,
    story TEXT,
    what_to_expect TEXT,
    how_to_participate TEXT,
    practical_info TEXT,
    cover_image_url TEXT,
    gallery_images JSONB DEFAULT '[]',
    video_embeds JSONB DEFAULT '[]',
    is_published BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- Festival Dates (supports multiple years)
CREATE TABLE festival_dates (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    festival_id UUID REFERENCES festivals(id) ON DELETE CASCADE,
    year INTEGER NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE,
    UNIQUE(festival_id, year)
);

-- Memories
CREATE TABLE memories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    festival_id UUID REFERENCES festivals(id) ON DELETE CASCADE,
    author_name VARCHAR(100),
    author_email VARCHAR(255),
    content TEXT NOT NULL,
    year_of_memory VARCHAR(20),
    status VARCHAR(20) DEFAULT 'pending',  -- pending, approved, rejected
    submitted_at TIMESTAMPTZ DEFAULT NOW()
);

-- Subscriptions
CREATE TABLE subscriptions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    confirmed BOOLEAN DEFAULT FALSE,
    confirmation_token VARCHAR(100),
    unsubscribe_token VARCHAR(100) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);
```

### Why Neon?

- **Serverless**: Scales to zero, no idle costs
- **Branching**: Can create database branches for testing
- **PostgreSQL**: Full Postgres compatibility
- **Connection pooling**: Built-in, handles Fly.io's connection patterns

---

## Email System

### Architecture

```
┌─────────────────┐      ┌─────────────────┐      ┌─────────────────┐
│                 │      │                 │      │                 │
│  User Action    │─────▶│  Backend        │─────▶│  Resend API     │
│  (subscribe)    │      │  (Go service)   │      │                 │
│                 │      │                 │      │                 │
└─────────────────┘      └─────────────────┘      └────────┬────────┘
                                                           │
                                                           ▼
                                                  ┌─────────────────┐
                                                  │                 │
                                                  │  User Inbox     │
                                                  │                 │
                                                  └─────────────────┘
```

### Email Templates

Three email types are implemented:

1. **Welcome Email** - Sent on subscription
2. **Festival Reminder** - "Festival X is in 7 days"
3. **Weekly Digest** - List of upcoming festivals

All emails feature:
- KULTUR logo header (loaded from `kultur-tt.app/logo.png`)
- Responsive HTML design
- Unsubscribe link in footer
- Trinidad & Tobago themed colors

### DNS Configuration

For custom domain emails (`noreply@kultur-tt.app`):

| Type | Name | Value |
|:-----|:-----|:------|
| MX | @ | `feedback-smtp.us-east-1.amazonses.com` |
| TXT | @ | SPF record |
| CNAME | resend._domainkey | DKIM key |

---

## Deployment

### Frontend (Vercel)

**Configuration:**
- Framework: SvelteKit
- Build Command: `npm run build`
- Output Directory: `.svelte-kit`
- Node Version: 20.x

**Environment Variables:**
```
PUBLIC_API_URL=https://kultur-api.fly.dev
PUBLIC_DATA_SOURCE=api
```

**Custom Domain:**
- `kultur-tt.app` → Vercel
- A record: `76.76.21.21`
- CNAME: `cname.vercel-dns.com`

### Backend (Fly.io)

**fly.toml:**
```toml
app = "kultur-api"
primary_region = "mia"  # Miami (closest to Trinidad)

[build]
  dockerfile = "Dockerfile"

[env]
  PORT = "8080"

[http_service]
  internal_port = 8080
  force_https = true

[[vm]]
  size = "shared-cpu-1x"
  memory = "256mb"
```

**Dockerfile:**
```dockerfile
FROM golang:1.24-alpine AS builder
WORKDIR /app
RUN apk add --no-cache ca-certificates
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o server ./cmd/server

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
```

**Secrets:**
```bash
fly secrets set DATABASE_URL="postgres://..."
fly secrets set RESEND_API_KEY="re_..."
fly secrets set ADMIN_API_KEY="..."
fly secrets set FROM_EMAIL="noreply@kultur-tt.app"
fly secrets set BASE_URL="https://kultur-api.fly.dev"
fly secrets set ALLOWED_ORIGINS="https://kultur-tt.app"
```

---

## Issues & Solutions

### 1. Logo Background Removal

**Problem:** Logo PNG had white background that looked bad on dark sections.

**Attempted Solutions:**
- Simple white pixel removal → Removed ALL white, including text
- Threshold-based removal → Still too aggressive

**Final Solution:** Python Pillow flood fill from corners:
```python
from PIL import Image
img = Image.open("logo.png").convert("RGBA")
# Flood fill from each corner to remove only connected background
```

**Result:** 988KB → 48KB, transparent background, white text preserved.

### 2. Environment Variable Newlines

**Problem:** Vercel env vars had `\n` appended, causing API calls to fail silently.

**Symptom:** `PUBLIC_API_URL="https://kultur-api.fly.dev\n"` → fetch failed

**Solution:** Re-added env vars using `echo -n` to avoid newlines:
```bash
echo -n "https://kultur-api.fly.dev" | vercel env add PUBLIC_API_URL production
```

### 3. Resend Test Domain Limitation

**Problem:** Emails not being delivered from frontend subscription.

**Investigation:**
- Backend logs showed 201 success
- Resend dashboard showed emails sent
- But nothing in inbox

**Cause:** Resend test domains (`onboarding@resend.dev`) can only send to verified account email.

**Solution:** Set up custom domain `kultur-tt.app` with proper DNS (DKIM, SPF, MX records).

### 4. Festival Card Height Inconsistency

**Problem:** Cards in grid had varying heights due to different title/description lengths.

**Solution:** Added flexbox utilities:
```svelte
<a class="block group h-full w-full">
    <Card.Root class="h-full">
        <Card.Content class="flex-grow flex flex-col">
            <p class="min-h-[2.5rem] line-clamp-2">{summary}</p>
            <div class="mt-auto">{date/location}</div>
        </Card.Content>
    </Card.Root>
</a>
```

### 5. Biome Configuration for Svelte

**Problem:** Initial Biome config used schema 2.x syntax that didn't work with 1.9.x.

**Errors:**
- `includes` should be `include`
- `assist` should be `organizeImports`

**Solution:** Used correct 1.9.x schema:
```json
{
    "$schema": "https://biomejs.dev/schemas/1.9.4/schema.json",
    "files": {
        "include": ["**/*.ts", "**/*.svelte", "**/*.json"]
    },
    "formatter": {
        "indentStyle": "space",
        "indentWidth": 4
    }
}
```

---

## Security

### Rate Limiting

Implemented per-IP rate limiting for public write endpoints:

| Endpoint | Limit |
|:---------|:------|
| `POST /api/memories` | 5/hour |
| `POST /api/subscribe` | 10/hour |

Implementation uses in-memory token bucket (resets on deploy).

### Admin Authentication

Admin routes require `X-API-Key` header:

```go
func APIKeyAuth(validKey string) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            key := c.Request().Header.Get("X-API-Key")
            if key != validKey {
                return echo.NewHTTPError(401, "invalid API key")
            }
            return next(c)
        }
    }
}
```

### CORS

Strict origin whitelist:
```go
AllowOrigins: []string{"https://kultur-tt.app", "http://localhost:5173"}
```

### Input Validation

- Email validation on subscribe
- Content length limits on memories
- SQL injection prevented by parameterized queries (sqlc)

---

## Performance

### Frontend Optimizations

1. **Logo Optimization**: 988KB → 48KB PNG
2. **Preloading**: Critical assets preloaded in `<head>`
3. **Image lazy loading**: Non-critical images use `loading="lazy"`
4. **Code splitting**: SvelteKit automatic route-based splitting

### Backend Optimizations

1. **Connection pooling**: pgx pool with configurable limits
2. **Minimal Docker image**: Multi-stage build, `scratch` base (5.1MB total)
3. **No ORM overhead**: sqlc generates direct SQL calls

### Caching Strategy

Currently minimal caching (hackathon scope). Future improvements:
- CDN caching for static assets (Vercel handles this)
- API response caching with Redis
- Stale-while-revalidate for festival data
