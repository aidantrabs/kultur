# API Routes

## Production URLs

- **Frontend**: https://kultur-tt.app
- **Backend API**: https://kultur-api.fly.dev

## Route Reference

### Public Routes

| Route | Method | Description |
|:------|:-------|:------------|
| `/health` | GET | Health check endpoint |
| `/api/festivals` | GET | List all festivals |
| `/api/festivals/upcoming` | GET | List festivals in next 30 days |
| `/api/festivals/calendar` | GET | List festivals by year |
| `/api/festivals/:slug` | GET | Get single festival by slug |
| `/api/festivals/:slug/dates` | GET | Get festival dates |
| `/api/festivals/:slug/memories` | GET | Get memories for a festival |
| `/api/memories` | POST | Submit a memory (5/hour rate limit) |
| `/api/subscribe` | POST | Subscribe to newsletter (10/hour rate limit) |
| `/api/subscribe/confirm/:token` | GET | Confirm email subscription |
| `/api/unsubscribe/:token` | GET | Unsubscribe from emails |

### Admin Routes

All admin routes require `X-API-Key` header.

| Route | Method | Description |
|:------|:-------|:------------|
| `/api/admin/memories` | GET | List all memories |
| `/api/admin/memories/:id` | PATCH | Update memory status |
| `/api/admin/memories/:id` | DELETE | Delete a memory |
| `/api/admin/subscriptions` | GET | List all subscriptions |
| `/api/admin/subscriptions/:id` | DELETE | Delete a subscription |
| `/api/admin/festivals` | POST | Create a festival |
| `/api/admin/festivals/:id` | PUT | Update a festival |
| `/api/admin/festivals/:id` | DELETE | Delete a festival |
| `/api/admin/festival-dates` | POST | Create a festival date |
| `/api/admin/festival-dates/:id` | PUT | Update a festival date |
| `/api/admin/festival-dates/:id` | DELETE | Delete a festival date |
| `/api/admin/test-email/welcome` | POST | Send test welcome email |
| `/api/admin/test-email/reminder` | POST | Send test festival reminder |
| `/api/admin/test-email/digest` | POST | Send test weekly digest |

## Frontend Routes

| Route | Description |
|:------|:------------|
| `/` | Home page with hero, featured festivals, newsletter signup |
| `/festivals` | Festival calendar with filters and grid view |
| `/festivals/:slug` | Festival detail page with full guide |

## Rate Limits

| Endpoint | Limit |
|:---------|:------|
| `POST /api/memories` | 5 requests/hour per IP |
| `POST /api/subscribe` | 10 requests/hour per IP |

## Authentication

Admin routes use API key authentication via the `X-API-Key` header:

```bash
curl -X GET "https://kultur-api.fly.dev/api/admin/memories" \
  -H "X-API-Key: YOUR_ADMIN_API_KEY"
```
