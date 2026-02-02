# Email Testing Demo

This guide walks through testing all email templates in the KULTUR platform.

## Prerequisites

- Access to the admin API key
- `curl` installed (or use Postman/Insomnia)

## API Endpoint

All test email endpoints are protected and require the admin API key header:

```
X-API-Key: <ADMIN_API_KEY>
```

Base URL: `https://kultur-api.fly.dev`

## Email Templates

### 1. Welcome Email

Sent when a user subscribes to the newsletter. Includes a welcome message and links to explore festivals.

```bash
curl -X POST "https://kultur-api.fly.dev/api/admin/test-email/welcome" \
  -H "Content-Type: application/json" \
  -H "X-API-Key: <ADMIN_API_KEY>" \
  -d '{"email": "your-email@example.com"}'
```

**Response:**
```json
{
  "message": "welcome email sent",
  "to": "your-email@example.com"
}
```

### 2. Festival Reminder Email

Sent to remind users about an upcoming festival. Includes festival name, days until the event, and a link to learn more.

```bash
curl -X POST "https://kultur-api.fly.dev/api/admin/test-email/reminder" \
  -H "Content-Type: application/json" \
  -H "X-API-Key: <ADMIN_API_KEY>" \
  -d '{"email": "your-email@example.com"}'
```

**Test Data:**
- Festival: Trinidad Carnival
- Days until: 7 days

**Response:**
```json
{
  "message": "festival reminder email sent",
  "to": "your-email@example.com"
}
```

### 3. Weekly Digest Email

Sent weekly with a list of upcoming festivals. Includes festival names, dates, heritage type, and regions.

```bash
curl -X POST "https://kultur-api.fly.dev/api/admin/test-email/digest" \
  -H "Content-Type: application/json" \
  -H "X-API-Key: <ADMIN_API_KEY>" \
  -d '{"email": "your-email@example.com"}'
```

**Test Data:**
| Festival | Date | Heritage | Region |
|----------|------|----------|--------|
| Trinidad Carnival | February 16-17, 2026 | Mixed Heritage | Nationwide |
| Hosay | February 20, 2026 | Indian Heritage | St. James |
| Phagwa | March 14, 2026 | Indian Heritage | Central Trinidad |

**Response:**
```json
{
  "message": "weekly digest email sent",
  "to": "your-email@example.com"
}
```

## Quick Test Script

Run all three tests at once (replace `<ADMIN_API_KEY>` and `<EMAIL>`):

```bash
API_KEY="<ADMIN_API_KEY>"
EMAIL="your-email@example.com"

echo "Sending Welcome Email..."
curl -s -X POST "https://kultur-api.fly.dev/api/admin/test-email/welcome" \
  -H "Content-Type: application/json" \
  -H "X-API-Key: $API_KEY" \
  -d "{\"email\": \"$EMAIL\"}"

echo -e "\n\nSending Festival Reminder..."
curl -s -X POST "https://kultur-api.fly.dev/api/admin/test-email/reminder" \
  -H "Content-Type: application/json" \
  -H "X-API-Key: $API_KEY" \
  -d "{\"email\": \"$EMAIL\"}"

echo -e "\n\nSending Weekly Digest..."
curl -s -X POST "https://kultur-api.fly.dev/api/admin/test-email/digest" \
  -H "Content-Type: application/json" \
  -H "X-API-Key: $API_KEY" \
  -d "{\"email\": \"$EMAIL\"}"

echo -e "\n\nDone! Check your inbox."
```

## Email Features

All emails include:
- KULTUR logo header
- Responsive design (mobile-friendly)
- Unsubscribe link in footer
- Links to kultur-tt.app

## Troubleshooting

**"invalid API key" error**
- Verify you're using the correct admin API key
- Ensure the header is `X-API-Key` (case-sensitive)

**Email not received**
- Check spam/junk folder
- Verify the email address is correct
- Check Resend dashboard for delivery status

## Production Notes

- Emails are sent via Resend from `noreply@kultur-tt.app`
- Custom domain DNS (DKIM, SPF, DMARC) is configured for deliverability
- Rate limits apply to prevent abuse (pls don't spam)
