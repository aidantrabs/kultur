# Hackathon Planning Document

**TrinbagoTech Hackathon 2026 · January 30 – February 2**

---

## Table of Contents

1. [Team](#team)
2. [The Idea](#the-idea)
3. [Why This Matters](#why-this-matters)
4. [Research Summary](#research-summary)
5. [Technical Stack](#technical-stack)
6. [Content Plan](#content-plan)
7. [Schedule](#schedule)
8. [Demo Strategy](#demo-strategy)
9. [Risks](#risks)
10. [Environment Setup](#environment-setup)
11. [Post-Hackathon](#post-hackathon)

---

## Team

| Person | Focus |
|:-------|:------|
| **Aidan** | Backend, Go API, database, deployment |
| **Tevin** | Frontend, SvelteKit, UI, responsive design |
| **Jahzara** | Content, cultural research, demo narrative |

> [!NOTE]
> These are primary focuses, not hard boundaries. If someone finishes early or another needs help, we pick up each other's tasks. The goal is clear ownership so nothing falls through the cracks.

---

## The Idea

**A first-timer's guide to Trinidad & Tobago's cultural festivals.**

T&T has over 20 cultural festivals spanning African, Indian, Indigenous, European, and Chinese heritage. Most people only know Carnival. We're building a platform that helps anyone — locals who've never attended, diaspora returning home, tourists looking for authentic experiences — discover and attend these festivals with confidence.

The key insight: people don't just need to know *when* a festival happens. They need to know *what it is*, *what to expect*, and *how to participate appropriately*. That's the gap we're filling.

### What We're Building

**Festival Calendar**
- Browse all festivals by month, region, or heritage type
- See what's coming up in the next few weeks
- Filter to find festivals that match your interests

**First-Timer's Guides**
- Each festival gets a dedicated page with real depth
- History and cultural significance (not Wikipedia-level — accessible)
- What you'll actually see, hear, and experience
- How to participate respectfully (dress code, etiquette, what's appropriate for outsiders)
- Practical info: where to park, what to bring, best spots to watch

**"This Week in T&T Culture"**
- Homepage section showing what's happening soon
- Email digest for subscribers (weekly)

**Community Memories**
- People can submit their own festival memories
- Adds authenticity and personal stories
- No account required — low friction

**Email Reminders**
- Subscribe to get notified before festivals you care about
- "Hosay is in 2 weeks" type notifications

**PWA**
- Installable on phone
- Works offline (for when you're at a festival with bad signal)

---

## Why This Matters

### The Problem

T&T has 20+ cultural festivals across African, Indian, Indigenous, and other heritage. But finding information about them is painful:

| What Exists | The Problem |
|:------------|:------------|
| WahWeDoing.com | 300+ Carnival parties, zero indigenous festivals |
| VisitTrinidad.tt | Last updated 2020 |
| DiscoverTnT.com | Calendar frozen on 2020 |
| NALIS archives | Excellent content, buried in academic format |
| Facebook groups | Fragmented, unsearchable |

**Result:** Most Trinis only know Carnival. Tourists miss authentic experiences. Indigenous festivals fade from public consciousness.

### The Opportunity

- **374,500** Trinis in diaspora (USA, Canada, UK) disconnected from culture
- **158,890** tourist arrivals Jan-May 2025 (+13.7% YoY)
- Ministry of Tourism pushing "community-centered experiences"
- No platform serves this need

---

## Research Summary

> [!CAUTION]
> **Cultural sensitivity matters.** We're guests in these traditions. When uncertain, hedge language. Jahzara should flag anything questionable.

### First Peoples

- Santa Rosa First Peoples Community in Arima — last organized Indigenous group in T&T
- ~400 registered members preserving traditions for 50+ years
- 8,000 years of continuous Amerindian presence (archaeological evidence)
- Key events: Smoke Ceremony, Water Ritual, Heritage Week (October), Santa Rosa Festival (August)
- Cultural contributions: cassava bread, barbecue, parang music, place names (Arima, Tunapuna, Chaguanas)

### Hosay

> [!IMPORTANT]
> **Hosay is NOT a "festival."** It's a solemn religious commemoration. Use "observance" not "celebration."

- Brought by Indian indentured laborers starting 1845
- Commemorates martyrdom of Hussein and Hassan at Battle of Karbala (680 AD)
- **1884 Massacre:** British authorities fired on procession in San Fernando — 22 killed, 120+ wounded
- Today: Survives in St. James, Cedros. Tadjahs up to 30 feet tall. Tassa drumming.
- Most Trinis have never heard of the massacre — this is our demo hook

### Other Key Festivals

- **Divali:** National holiday. Divali Nagar in Chaguanas = largest in English-speaking Caribbean
- **Phagwa:** Spring colors festival
- **Emancipation Day:** August 1, fixed date
- **Spiritual Baptist Liberation Day:** March 30
- **Ramleela:** Open-air theatre, 35+ sites

---

## Technical Stack

```
hackathon-project/
├── frontend/                # Tevin owns this
│   └── ...                  # SvelteKit + Tailwind
│
├── backend/                 # Aidan owns this
│   ├── cmd/
│   │   └── server/
│   │       └── main.go
│   ├── internal/
│   │   ├── config/
│   │   ├── db/              # sqlc generated code
│   │   ├── handler/
│   │   ├── service/
│   ├── migrations/
│   │   └── 001_init.sql
│   ├── queries/             # sqlc query files
│   │   └── *.sql
│   ├── sqlc.yaml
│   ├── go.mod
│   ├── go.sum
│   ├── Dockerfile
│   └── .env
│
├── docker-compose.yml
├── Makefile
├── .env.example
└── README.md
```

| Layer | Tech | Owner |
|:------|:-----|:------|
| Frontend | SvelteKit + Tailwind | Tevin |
| Frontend Hosting | Cloudflare Workers/Pages | Tevin |
| Storage | Cloudflare R2 | Tevin |
| Backend | Go + Echo | Aidan |
| Backend Hosting | Fly.io | Aidan |
| Database | PostgreSQL (Docker locally, Fly Postgres in prod) | Aidan |
| SQL Queries | sqlc (type-safe generated code) | Aidan |
| Migrations | goose | Aidan |
| Email | Resend | Aidan |

### How They Connect

```
┌─────────────────────────────────────────────────────────────┐
│                        LOCAL DEV                            │
│                                                             │
│   Frontend (localhost:5173)                                 │
│        │                                                    │
│        │  fetch('http://localhost:8080/api/...')            │
│        ▼                                                    │
│   Backend (localhost:8080)  ◄── go run                      │
│        │                                                    │
│        ▼                                                    │
│   PostgreSQL (localhost:5432)  ◄── Docker                   │
│                                                             │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│                       PRODUCTION                            │
│                                                             │
│   Frontend  ◄── Cloudflare Workers/Pages                    │
│        │                                                    │
│        │  fetch('https://[API].fly.dev/api/...')            │
│        ▼                                                    │
│   Backend  ◄── Fly.io                                       │
│        │                                                    │
│        ▼                                                    │
│   PostgreSQL  ◄── Fly Postgres                              │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### Docker (Local Dev)

Only Postgres runs in Docker. Go backend runs directly for faster iteration.

```yaml
# docker-compose.yml

services:
  db:
    image: postgres:16-alpine
    ports:
      - "5432:5432"
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=postgres
      - POSTGRES_DB=hackathon
    volumes:
      - pgdata:/var/lib/postgresql/data

volumes:
  pgdata:
```

### CORS (Aidan)

```go
e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    AllowOrigins:     []string{"http://localhost:5173", "https://[FRONTEND].pages.dev"},
    AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodOptions},
    AllowHeaders:     []string{echo.HeaderContentType},
    AllowCredentials: true,
}))
```

### API Client (Tevin)

```typescript
// frontend/src/lib/api.ts

const API_URL = import.meta.env.PUBLIC_API_URL || 'http://localhost:8080';

export const api = {
  festivals: {
    list: (filters?: { region?: string; heritage?: string }) =>
      fetch(`${API_URL}/api/festivals?${new URLSearchParams(filters)}`).then(r => r.json()),
    
    get: (slug: string) =>
      fetch(`${API_URL}/api/festivals/${slug}`).then(r => r.json()),
  },
  
  subscribe: (email: string) =>
    fetch(`${API_URL}/api/subscribe`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email }),
    }).then(r => r.json()),
};
```

### Database Schema

Single migration file for hackathon:

```sql
-- migrations/001_init.sql

-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE festivals (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    slug VARCHAR(100) UNIQUE NOT NULL,
    name VARCHAR(200) NOT NULL,
    date_type VARCHAR(20) NOT NULL,
    usual_month VARCHAR(50),
    date_2026_start DATE,
    date_2026_end DATE,
    region VARCHAR(50) NOT NULL,
    heritage_type VARCHAR(50) NOT NULL,
    festival_type VARCHAR(50) NOT NULL,
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

CREATE TABLE memories (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    festival_id UUID REFERENCES festivals(id) ON DELETE CASCADE,
    author_name VARCHAR(100),
    author_email VARCHAR(255),
    content TEXT NOT NULL,
    year_of_memory VARCHAR(20),
    status VARCHAR(20) DEFAULT 'pending',
    submitted_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE subscriptions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) UNIQUE NOT NULL,
    digest_weekly BOOLEAN DEFAULT FALSE,
    festival_reminders JSONB DEFAULT '[]',
    confirmed BOOLEAN DEFAULT FALSE,
    confirmation_token VARCHAR(100),
    unsubscribe_token VARCHAR(100) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS subscriptions;
DROP TABLE IF EXISTS memories;
DROP TABLE IF EXISTS festivals;
```

### API Endpoints

```
GET  /api/festivals              # List (with filters)
GET  /api/festivals/upcoming     # Next 30 days
GET  /api/festivals/{slug}       # Single festival
GET  /api/festivals/{slug}/memories

POST /api/memories               # Submit memory
POST /api/subscribe              # Email signup
GET  /api/subscribe/confirm/{token}
GET  /api/unsubscribe/{token}
```

---

## Content Plan

**Jahzara owns this.** The quality of content makes or breaks the demo.

### Festivals to Cover (10 total)

| Festival | Heritage | Priority |
|:---------|:---------|:---------|
| Hosay | Indian/Islamic | High — demo hook |
| Santa Rosa Festival | Indigenous | High — hackathon track |
| First Peoples Heritage Week | Indigenous | High |
| Divali | Indian/Hindu | Medium |
| Phagwa | Indian/Hindu | Medium |
| Emancipation Day | African | Medium |
| Spiritual Baptist Liberation Day | African | Medium |
| Ramleela | Indian/Hindu | Medium |
| Parang | Mixed | Low |
| Carnival | Mixed | Low — already well-known |

### Content Per Festival

| Section | Purpose |
|:--------|:--------|
| Overview | What, when, where |
| The Story | History, origins, meaning |
| What to Expect | Sights, sounds, atmosphere |
| How to Participate | Etiquette, dress, is it open to outsiders? |
| Practical Info | Parking, food, best spots |

**Time estimate:** 2-3 hours per festival = 25-30 hours total

### Sources

| Source | Use |
|:-------|:----|
| NALIS (nalis.gov.tt) | History, First Peoples |
| Caribbean Beat | Festival profiles |
| Wikipedia | Overview, dates |
| Newsday/Guardian | Current details |
| Wikimedia/Flickr | CC-licensed photos |

> [!CAUTION]
> **Copyright.** Only use CC-licensed images. No Google image search grabs.

### Voice

- Warm, not academic
- "Here's what you need to know" not "The history of..."
- Acknowledge it's okay to be a newcomer
- When in doubt: "some community members say..." not definitive claims

---

## Schedule

### Official Timeline

| Event | Date & Time |
|:------|:------------|
| **Launch Session** | Friday, Jan 30, 9:00 AM – 12:00 PM |
| **Mentorship Session** | Saturday, Jan 31 (30 min, scheduled by mentor) |
| **Presentations** | Monday, Feb 2, 9:00 AM – 12:30 PM (15-min slot) |

> [!NOTE]
> All sessions mandatory. Pitch deck template provided at launch.

---

### Pre-Hackathon (Should Be Done)

**Aidan:**
- [ ] Create repo with `frontend/` and `backend/` folders
- [ ] Set up Go API with Chi boilerplate in `backend/`
- [ ] Write Dockerfile and docker-compose.yml
- [ ] Write database migration
- [ ] Set up Neon database (for production)
- [ ] Set up Fly.io project
- [ ] Get Resend API key

**Tevin:**
- [ ] Set up SvelteKit + Tailwind + shadcn-svelte in `frontend/`
- [ ] Set up design tokens (colors from T&T flag)
- [ ] Create basic API client to connect to backend

**Jahzara:**
- [ ] Research and outline all 10 festivals
- [ ] Collect image URLs and video embeds (CC only)
- [ ] Draft 3 complete festival pages
- [ ] Identify demo hooks (Hosay Massacre, First Peoples history)

---

### Day 1 — Friday, Jan 30 (Today)

**Morning: Launch Session (9 AM – 12 PM)**
- Hackathon rules and submission requirements
- IP, pitching, prototyping overview
- Receive pitch deck template
- Select Monday presentation slot

**After launch:**

| Time | Aidan | Tevin | Jahzara |
|:-----|:------|:------|:--------|
| **PM** | Finish API endpoints, test with Docker | Finish frontend setup, connect to API | Research Hosay, Santa Rosa, Divali |
| **Eve** | Seed DB with 3 festivals, test end-to-end | Navigation, home page, festival card | Draft first 3 festival pages |

**End of Day 1:** Both apps running, connected, 3 festivals viewable.

---

### Day 2 — Saturday, Jan 31

**Mentorship session (30 min) — time scheduled by mentor**

| Time | Aidan | Tevin | Jahzara |
|:-----|:------|:------|:--------|
| **AM** | Memory + subscription endpoints | Festival detail page (full layout) | Complete festivals 4-7 |
| **PM** | Resend email integration | Calendar with filters, forms | Complete festivals 8-10 |
| **Eve** | Load all content into DB | Mobile responsiveness | Write sample memories |

**End of Day 2:** All features working, 10 festivals with content.

> [!WARNING]
> **Day 2 is critical.** Feature-complete by end of day or we're in trouble.

---

### Day 3 — Sunday, Feb 1

| Time | Aidan | Tevin | Jahzara |
|:-----|:------|:------|:--------|
| **AM** | PWA setup (manifest, service worker) | UI polish, loading states | Content proofread |
| **PM** | Deploy backend to Fly.io | Deploy frontend to Cloudflare Pages | Visual review for authenticity |
| **Eve** | Test on real devices | Bug fixes | Start pitch deck, draft demo script |

**End of Day 3:** Deployed, tested, pitch deck started.

> [!NOTE]
> **Deploy Sunday, not Monday.** Gives time to catch deployment bugs.

---

### Day 4 — Monday, Feb 2

**Presentations: 9:00 AM – 12:30 PM (15-min slot)**

| Time | Everyone |
|:-----|:---------|
| **Before slot** | Finalize pitch deck, record demo video |
| **Presentation** | 15-min slot (time TBD) |

**Deliverables:**
- Pre-recorded demo (with live site available)
- Pitch deck (using provided template)

---

## Demo Strategy

**Format:** 15-minute presentation slot. Pre-recorded demo with live site available.

### Opening (30 seconds)

> "On October 30, 1884, British authorities opened fire on a peaceful Hosay procession in San Fernando. 22 people died. It's one of the bloodiest events in Trinidad's colonial history.
>
> Most Trinis have never heard of it.
>
> That's the problem we're solving."

### Demo Flow (3-4 minutes)

1. Home page — show "This Week" section
2. Calendar — filter by heritage type
3. Click Hosay — show full First-Timer's Guide
4. Scroll through sections (What to Expect, How to Participate)
5. Show community memories
6. Subscribe flow
7. PWA install on phone (if time)

### Closing (30 seconds)

> "We built this in 4 days. Imagine every festival documented like this. Imagine diaspora kids learning about Hosay before visiting home. Imagine tourists experiencing something real.
>
> KULTUR. Your guide to Trinidad's cultural festivals."

### Remaining Time

Use for Q&A or expanding on technical/business aspects from pitch deck.

---

## Risks

| Risk | Mitigation |
|:-----|:-----------|
| Not enough content | Prioritize 5 great festivals over 10 incomplete ones |
| Inaccurate cultural info | Cross-reference sources, hedge language |
| CORS issues | Configure early Day 1, test with real domains |
| Deployment problems | Deploy early Day 3, have backups ready |
| Live demo fails | Pre-record backup video |

---

## Environment Setup

**Backend (backend/.env):**
```env
PORT=8080
DATABASE_URL=postgres://postgres:postgres@localhost:5432/hackathon?sslmode=disable
RESEND_API_KEY=re_...
ALLOWED_ORIGINS=http://localhost:5173
```

**Frontend (frontend/.env):**
```env
PUBLIC_API_URL=http://localhost:8080
```

### Makefile

```makefile
.PHONY: dev db-up db-down migrate migrate-down build deploy

# Database
db-up:
	docker-compose up -d

db-down:
	docker-compose down

# Backend
dev-api:
	cd backend && go run ./cmd/server

# Frontend
dev-frontend:
	cd frontend && pnpm dev

# Migrations
migrate:
	goose -dir backend/migrations postgres "$(DATABASE_URL)" up

migrate-down:
	goose -dir backend/migrations postgres "$(DATABASE_URL)" down

migrate-status:
	goose -dir backend/migrations postgres "$(DATABASE_URL)" status

# Build
build-api:
	cd backend && go build -o bin/server ./cmd/server

# Deploy
deploy-api:
	cd backend && gcloud run deploy kultur-api --source . --region us-central1 --allow-unauthenticated
```

### Commands

```bash
# Start postgres
make db-up

# Run backend (terminal 1)
make dev-api

# Run frontend (terminal 2)
make dev-frontend

# Migrations
DATABASE_URL="postgres://postgres:postgres@localhost:5432/hackathon?sslmode=disable" make migrate

# Deploy backend
make deploy-api
```

---

## What Was Built (Final State)

This section documents what was actually implemented during the hackathon.

### Production URLs

| Service | URL |
|:--------|:----|
| Frontend | https://kultur-tt.app |
| Backend API | https://kultur-api.fly.dev |
| Email Domain | noreply@kultur-tt.app |

### Tech Stack (Final)

| Layer | Technology | Hosting |
|:------|:-----------|:--------|
| Frontend | SvelteKit 2, Svelte 5, TailwindCSS 4, shadcn-svelte | Vercel |
| Backend | Go 1.24, Echo framework | Fly.io |
| Database | PostgreSQL | Neon |
| Email | Resend | Custom domain |
| Linting | Biome | - |

### Features Implemented

**Frontend:**
- Home page with hero section, featured festivals, newsletter signup
- Festival calendar with month/region/heritage filters
- Festival detail pages with full first-timer guides
- Community memories section
- Newsletter subscription with email validation
- Responsive design (mobile, tablet, desktop)
- Copy link functionality on festival pages

**Backend:**
- RESTful API with Echo framework
- Festival CRUD operations
- Memory submission with moderation
- Email subscription system
- Rate limiting (5/hour memories, 10/hour subscribe)
- Admin API with API key authentication
- Email templates (welcome, festival reminder, weekly digest)

**Email System:**
- Custom domain (noreply@kultur-tt.app)
- Welcome email on subscription
- Festival reminder template
- Weekly digest template
- Unsubscribe functionality

### What Changed from Plan

| Planned | Actual |
|:--------|:-------|
| Cloudflare Pages | Vercel |
| Fly Postgres | Neon PostgreSQL |
| pnpm | npm |
| Chi router | Echo framework |
| PWA with offline | Not implemented |

### API Routes (Final)

See [ROUTES.md](./ROUTES.md) for complete API documentation.

---

## Post-Hackathon (Reference)

If we want to continue:

- Terraform for infrastructure-as-code (Neon, Fly.io, Vercel all have providers)
- GitHub Actions CI/CD
- Partner outreach (Tourism Trinidad, NALIS, Santa Rosa Community)
- Automated weekly digest emails (cron job)
- PWA with offline support

But that's for later. **Focus on the hackathon.**
