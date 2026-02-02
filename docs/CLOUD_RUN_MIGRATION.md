# Cloud Run Migration Guide

Documentation of the migration from Fly.io to Google Cloud Run, including setup steps, configuration, and issues encountered.

---

## Table of Contents

- [Why We Migrated](#why-we-migrated)
- [Cloud Run Free Tier](#cloud-run-free-tier)
- [Prerequisites](#prerequisites)
- [GCP Setup](#gcp-setup)
- [GitHub Actions Configuration](#github-actions-configuration)
- [Environment Variables](#environment-variables)
- [Deployment](#deployment)
- [Issues Encountered](#issues-encountered)
- [Useful Commands](#useful-commands)

---

## Why We Migrated

Fly.io suspended our API and began requiring payment. To keep the application completely free, we migrated to Google Cloud Run which offers a generous free tier suitable for hobby projects.

**Previous Setup:**
- Backend hosted on Fly.io at `kultur-api.fly.dev`
- Configured via `fly.toml`
- Deployed using `flyctl deploy`

**New Setup:**
- Backend hosted on Cloud Run at `kultur-api-971304624476.us-central1.run.app`
- Deployed using `gcloud run deploy` or GitHub Actions
- Secrets managed via Google Secret Manager

---

## Cloud Run Free Tier

Cloud Run's free tier includes (per month):

| Resource | Free Allowance |
|:---------|:---------------|
| Requests | 2 million |
| CPU | 180,000 vCPU-seconds |
| Memory | 360,000 GB-seconds |
| Networking | 1 GB egress to North America |

Key benefits:
- **Scale to zero**: No charges when idle
- **No credit card required** to stay within free tier
- **Container-based**: Minimal migration effort from Fly.io

---

## Prerequisites

### Install Google Cloud SDK

**macOS:**
```bash
brew install google-cloud-sdk
```

**Linux:**
```bash
curl https://sdk.cloud.google.com | bash
exec -l $SHELL
```

### Authenticate

```bash
gcloud auth login
```

### Set Project

```bash
gcloud config set project kulturtt-api
```

---

## GCP Setup

### 1. Enable Required APIs

```bash
gcloud services enable \
  run.googleapis.com \
  artifactregistry.googleapis.com \
  secretmanager.googleapis.com \
  iamcredentials.googleapis.com \
  cloudresourcemanager.googleapis.com \
  iam.googleapis.com \
  cloudbuild.googleapis.com
```

### 2. Create Artifact Registry Repository

```bash
gcloud artifacts repositories create kultur \
  --repository-format=docker \
  --location=us-central1 \
  --description="Kultur API container images"
```

### 3. Create Service Account for GitHub Actions

```bash
# Create service account
gcloud iam service-accounts create github-actions \
  --display-name="GitHub Actions" \
  --description="Service account for GitHub Actions deployments"

# Grant Cloud Run admin role
gcloud projects add-iam-policy-binding kulturtt-api \
  --member="serviceAccount:github-actions@kulturtt-api.iam.gserviceaccount.com" \
  --role="roles/run.admin"

# Grant Artifact Registry writer role
gcloud projects add-iam-policy-binding kulturtt-api \
  --member="serviceAccount:github-actions@kulturtt-api.iam.gserviceaccount.com" \
  --role="roles/artifactregistry.writer"

# Grant service account user role
gcloud projects add-iam-policy-binding kulturtt-api \
  --member="serviceAccount:github-actions@kulturtt-api.iam.gserviceaccount.com" \
  --role="roles/iam.serviceAccountUser"

# Grant secret accessor role
gcloud projects add-iam-policy-binding kulturtt-api \
  --member="serviceAccount:github-actions@kulturtt-api.iam.gserviceaccount.com" \
  --role="roles/secretmanager.secretAccessor"
```

### 4. Set Up Workload Identity Federation

This allows GitHub Actions to authenticate to GCP without storing service account keys.

```bash
# Create workload identity pool
gcloud iam workload-identity-pools create "github" \
  --location="global" \
  --display-name="GitHub Actions Pool"

# Create OIDC provider
gcloud iam workload-identity-pools providers create-oidc "github" \
  --location="global" \
  --workload-identity-pool="github" \
  --display-name="GitHub" \
  --attribute-mapping="google.subject=assertion.sub,attribute.actor=assertion.actor,attribute.repository=assertion.repository" \
  --attribute-condition="assertion.repository=='aidantrabs/kultur'" \
  --issuer-uri="https://token.actions.githubusercontent.com"

# Allow GitHub repo to impersonate service account
gcloud iam service-accounts add-iam-policy-binding \
  "github-actions@kulturtt-api.iam.gserviceaccount.com" \
  --role="roles/iam.workloadIdentityUser" \
  --member="principalSet://iam.googleapis.com/projects/971304624476/locations/global/workloadIdentityPools/github/attribute.repository/aidantrabs/kultur"
```

### 5. Create Secrets in Secret Manager

```bash
# Database URL (Neon PostgreSQL)
echo -n 'postgresql://user:pass@host/db?sslmode=require' | \
  gcloud secrets create DATABASE_URL --data-file=-

# Resend API Key
echo -n 're_xxxxx' | \
  gcloud secrets create RESEND_API_KEY --data-file=-

# Admin API Key
echo -n 'your-admin-key' | \
  gcloud secrets create ADMIN_API_KEY --data-file=-

# Grant Cloud Run default service account access to secrets
PROJECT_NUMBER=$(gcloud projects describe kulturtt-api --format='value(projectNumber)')

for SECRET in DATABASE_URL RESEND_API_KEY ADMIN_API_KEY; do
  gcloud secrets add-iam-policy-binding $SECRET \
    --member="serviceAccount:${PROJECT_NUMBER}-compute@developer.gserviceaccount.com" \
    --role="roles/secretmanager.secretAccessor"
done
```

---

## GitHub Actions Configuration

### Required Secrets

Add these secrets to your GitHub repository at `Settings > Secrets and variables > Actions`:

| Secret | Value |
|:-------|:------|
| `GCP_PROJECT_ID` | `kulturtt-api` |
| `GCP_SERVICE_ACCOUNT` | `github-actions@kulturtt-api.iam.gserviceaccount.com` |
| `GCP_WORKLOAD_IDENTITY_PROVIDER` | `projects/971304624476/locations/global/workloadIdentityPools/github/providers/github` |

### Set Secrets via CLI

```bash
gh secret set GCP_PROJECT_ID --repo aidantrabs/kultur --body "kulturtt-api"

gh secret set GCP_SERVICE_ACCOUNT --repo aidantrabs/kultur \
  --body "github-actions@kulturtt-api.iam.gserviceaccount.com"

gh secret set GCP_WORKLOAD_IDENTITY_PROVIDER --repo aidantrabs/kultur \
  --body "projects/971304624476/locations/global/workloadIdentityPools/github/providers/github"

# Remove old Fly.io secret
gh secret delete FLY_API_TOKEN --repo aidantrabs/kultur
```

### Workflow File

The workflow at `.github/workflows/deploy-backend.yml` handles automatic deployments:

```yaml
name: Deploy Backend

on:
    push:
        branches: [main]
        paths:
            - "backend/**"
            - ".github/workflows/deploy-backend.yml"
    workflow_dispatch:

jobs:
    deploy:
        name: Deploy to Cloud Run
        runs-on: ubuntu-latest

        permissions:
            contents: read
            id-token: write

        steps:
            - name: Checkout
              uses: actions/checkout@v4

            - name: Authenticate to Google Cloud
              uses: google-github-actions/auth@v2
              with:
                  workload_identity_provider: ${{ secrets.GCP_WORKLOAD_IDENTITY_PROVIDER }}
                  service_account: ${{ secrets.GCP_SERVICE_ACCOUNT }}

            - name: Set up Cloud SDK
              uses: google-github-actions/setup-gcloud@v2

            - name: Configure Docker for Artifact Registry
              run: gcloud auth configure-docker us-central1-docker.pkg.dev --quiet

            - name: Build and Push Container
              run: |
                  docker build -t us-central1-docker.pkg.dev/${{ secrets.GCP_PROJECT_ID }}/kultur/api:${{ github.sha }} ./backend
                  docker push us-central1-docker.pkg.dev/${{ secrets.GCP_PROJECT_ID }}/kultur/api:${{ github.sha }}

            - name: Deploy to Cloud Run
              run: |
                  gcloud run deploy kultur-api \
                    --image us-central1-docker.pkg.dev/${{ secrets.GCP_PROJECT_ID }}/kultur/api:${{ github.sha }} \
                    --region us-central1 \
                    --platform managed \
                    --allow-unauthenticated \
                    --min-instances 0 \
                    --max-instances 2 \
                    --memory 256Mi \
                    --cpu 1 \
                    --port 8080 \
                    --set-secrets="DATABASE_URL=DATABASE_URL:latest,RESEND_API_KEY=RESEND_API_KEY:latest,ADMIN_API_KEY=ADMIN_API_KEY:latest" \
                    --set-env-vars="..." # See env vars section
```

---

## Environment Variables

### Cloud Run Service Configuration

| Variable | Value | Type |
|:---------|:------|:-----|
| `DATABASE_URL` | Neon connection string | Secret |
| `RESEND_API_KEY` | Resend API key | Secret |
| `ADMIN_API_KEY` | Admin authentication key | Secret |
| `FROM_EMAIL` | `noreply@kultur-tt.app` | Env var |
| `BASE_URL` | `https://kultur-api-971304624476.us-central1.run.app` | Env var |
| `ALLOWED_ORIGINS` | `https://kultur-tt.app,https://www.kultur-tt.app,http://localhost:5173` | Env var |

### Update Environment Variables

Due to commas in `ALLOWED_ORIGINS`, use an env-vars-file:

```yaml
# env-vars.yaml
FROM_EMAIL: noreply@kultur-tt.app
BASE_URL: https://kultur-api-971304624476.us-central1.run.app
ALLOWED_ORIGINS: "https://kultur-tt.app,https://www.kultur-tt.app,http://localhost:5173"
```

```bash
gcloud run services update kultur-api --region us-central1 \
  --set-secrets="DATABASE_URL=DATABASE_URL:latest,RESEND_API_KEY=RESEND_API_KEY:latest,ADMIN_API_KEY=ADMIN_API_KEY:latest" \
  --env-vars-file="env-vars.yaml"
```

### Vercel Frontend Configuration

Update the frontend environment variable to point to Cloud Run:

```bash
# Remove old URL
vercel env rm PUBLIC_API_URL production -y

# Add new URL
echo -n "https://kultur-api-971304624476.us-central1.run.app" | \
  vercel env add PUBLIC_API_URL production

# Redeploy frontend
vercel --prod --force
```

---

## Deployment

### Manual Deployment (from source)

```bash
cd backend

gcloud run deploy kultur-api \
  --source . \
  --region us-central1 \
  --platform managed \
  --allow-unauthenticated
```

### View Logs

```bash
# Recent logs
gcloud run services logs read kultur-api --region us-central1 --limit 50

# Stream logs
gcloud run services logs tail kultur-api --region us-central1

# Error logs only
gcloud logging read "resource.type=cloud_run_revision AND resource.labels.service_name=kultur-api AND severity>=ERROR" --limit=20
```

### Check Service Status

```bash
gcloud run services describe kultur-api --region us-central1
```

---

## Issues Encountered

### 1. Billing Required

**Problem:** Cloud Run APIs require billing to be enabled, even for free tier usage.

**Solution:** Enable billing at https://console.cloud.google.com/billing/linkedaccount?project=kulturtt-api

The free tier limits are still honored - you won't be charged if you stay within them.

### 2. Cloud Build "Directory Not Found" Error

**Problem:** `gcloud run deploy --source` failed with `stat /app/cmd/server: directory not found`

**Cause:** The default `.gcloudignore` behavior was excluding the `cmd/` directory because gcloud reads `.gitignore` patterns by default.

**Solution:** Create a minimal `.gcloudignore` file:

```
.env
.env.local
.git
```

This overrides the default behavior and ensures all necessary files are uploaded.

### 3. Docker Build Failed on Apple Silicon

**Problem:** Building with `--platform linux/amd64` on Apple Silicon Macs failed with "illegal instruction" errors due to QEMU emulation issues.

**Solution:** Use `gcloud run deploy --source` which builds on Google Cloud's infrastructure instead of locally. Alternatively, use `gcloud builds submit` to build remotely.

### 4. Environment Variables with Commas

**Problem:** `gcloud run services update --set-env-vars` failed when values contained commas (like `ALLOWED_ORIGINS`).

**Solution:** Use the `--env-vars-file` flag with a YAML file instead:

```yaml
ALLOWED_ORIGINS: "https://kultur-tt.app,https://www.kultur-tt.app,http://localhost:5173"
```

### 5. Duplicate Email Subscription Returns 500

**Problem:** Subscribing with an already-registered email returned a 500 Internal Server Error instead of a user-friendly message.

**Cause:** The database has a `UNIQUE` constraint on the email column, and the error wasn't being handled gracefully.

**Solution:** Updated the subscription service to check for existing emails and return a 200 OK with a generic message (for privacy, we don't reveal if an email is already subscribed):

```go
if errors.Is(err, service.ErrEmailAlreadyExists) {
    return c.JSON(http.StatusOK, map[string]any{
        "message": "if this email is not already subscribed, you will receive a confirmation email",
    })
}
```

### 6. CORS Configuration

**Problem:** Frontend requests failed with NetworkError.

**Solution:** Ensure `ALLOWED_ORIGINS` includes all frontend domains:
- `https://kultur-tt.app` (production)
- `https://www.kultur-tt.app` (www subdomain)
- `http://localhost:5173` (local development)

### 7. PORT Environment Variable Reserved

**Problem:** Setting `PORT=8080` in `--set-env-vars` failed because Cloud Run reserves this variable.

**Solution:** Remove `PORT` from env vars - Cloud Run automatically sets it and the app reads it correctly.

---

## Useful Commands

### Service Management

```bash
# List services
gcloud run services list

# Describe service
gcloud run services describe kultur-api --region us-central1

# Delete service (careful!)
gcloud run services delete kultur-api --region us-central1
```

### Secret Management

```bash
# List secrets
gcloud secrets list

# View secret versions
gcloud secrets versions list DATABASE_URL

# Update a secret
echo -n 'new-value' | gcloud secrets versions add DATABASE_URL --data-file=-
```

### Artifact Registry

```bash
# List images
gcloud artifacts docker images list us-central1-docker.pkg.dev/kulturtt-api/kultur

# Delete old images
gcloud artifacts docker images delete us-central1-docker.pkg.dev/kulturtt-api/kultur/api:TAG
```

### Debugging

```bash
# Test health endpoint
curl https://kultur-api-971304624476.us-central1.run.app/health

# Test with CORS headers
curl -X POST "https://kultur-api-971304624476.us-central1.run.app/api/subscribe" \
  -H "Origin: https://kultur-tt.app" \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com"}'

# Check CORS preflight
curl -I -X OPTIONS "https://kultur-api-971304624476.us-central1.run.app/api/subscribe" \
  -H "Origin: https://kultur-tt.app" \
  -H "Access-Control-Request-Method: POST"
```

---

## Files Changed in Migration

| File | Change |
|:-----|:-------|
| `backend/fly.toml` | Deleted |
| `backend/.gcloudignore` | Created |
| `.github/workflows/deploy-backend.yml` | Updated for Cloud Run |
| `frontend/.env` | Updated API URL |
| `frontend/.env.example` | Updated API URL |
| `docs/SETUP.md` | Updated CLI instructions |
| `docs/IMPLEMENTATION.md` | Updated architecture docs |
| `docs/PLANNING.md` | Updated deploy command |
| `backend/README.md` | Updated deployment section |
| `README.md` | Updated tech stack |
| `docs/assets/tech-icons/flyio.svg` | Deleted |
| `docs/assets/tech-icons/cloudrun.svg` | Created |

---

## Resources

- [Cloud Run Documentation](https://cloud.google.com/run/docs)
- [Cloud Run Pricing](https://cloud.google.com/run/pricing)
- [Workload Identity Federation](https://cloud.google.com/iam/docs/workload-identity-federation)
- [GitHub Actions for Google Cloud](https://github.com/google-github-actions)
