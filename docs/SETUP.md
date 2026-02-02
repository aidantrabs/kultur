# Local Development Setup

Complete guide to set up the KULTUR project locally from scratch.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Environment Files](#environment-files)
- [Backend Setup](#backend-setup)
- [Frontend Setup](#frontend-setup)
- [Running the Project](#running-the-project)
- [Common Issues](#common-issues)

---

## Prerequisites

### macOS

#### 1. Install Homebrew

Homebrew is the package manager for macOS. Open Terminal and run:

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

Follow the on-screen instructions. After installation, add Homebrew to your PATH:

```bash
echo 'eval "$(/opt/homebrew/bin/brew shellenv)"' >> ~/.zprofile
eval "$(/opt/homebrew/bin/brew shellenv)"
```

Verify installation:

```bash
brew --version
```

#### 2. Install Git

```bash
brew install git
```

Verify:

```bash
git --version
```

#### 3. Install Go

The backend requires Go 1.24 or later.

```bash
brew install go
```

Verify:

```bash
go version
# Should output: go version go1.24.x darwin/arm64
```

#### 4. Install Node.js

The frontend requires Node.js 20 or later.

```bash
brew install node
```

Verify:

```bash
node --version
# Should output: v20.x.x or higher

npm --version
```

#### 5. Install Fly CLI (Optional)

Only needed if you want to deploy or access production logs.

```bash
brew install flyctl
```

Login to Fly (requires account):

```bash
fly auth login
```

### Windows (WSL2)

We recommend using WSL2 (Windows Subsystem for Linux) for development.

#### 1. Install WSL2

Open PowerShell as Administrator:

```powershell
wsl --install
```

Restart your computer, then open Ubuntu from the Start menu.

#### 2. Install Dependencies in WSL2

```bash
# Update packages
sudo apt update && sudo apt upgrade -y

# Install Git
sudo apt install git -y

# Install Go
sudo rm -rf /usr/local/go
wget https://go.dev/dl/go1.24.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.24.linux-amd64.tar.gz
rm go1.24.linux-amd64.tar.gz

# Add Go to PATH
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc

# Verify Go
go version

# Install Node.js via nvm
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.0/install.sh | bash
source ~/.bashrc
nvm install 20
nvm use 20

# Verify Node
node --version
npm --version
```

### Linux (Ubuntu/Debian)

```bash
# Update packages
sudo apt update && sudo apt upgrade -y

# Install Git
sudo apt install git -y

# Install Go
sudo rm -rf /usr/local/go
wget https://go.dev/dl/go1.24.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.24.linux-amd64.tar.gz
rm go1.24.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc

# Install Node.js
curl -fsSL https://deb.nodesource.com/setup_20.x | sudo -E bash -
sudo apt install -y nodejs
```

---

## Environment Files

**Important:** Environment files contain sensitive credentials and are not committed to the repository.

Contact a codeowner to obtain the `.env` files:

- **Backend:** `/backend/.env`
- **Frontend:** `/frontend/.env`

### Backend Environment Variables

Create `/backend/.env` with the following structure:

```env
PORT=8080
DATABASE_URL=<contact codeowner>
RESEND_API_KEY=<contact codeowner>
ALLOWED_ORIGINS=http://localhost:5173
ADMIN_API_KEY=<contact codeowner>
BASE_URL=http://localhost:8080
FROM_EMAIL=noreply@kultur-tt.app
```

### Frontend Environment Variables

Create `/frontend/.env` with the following:

```env
PUBLIC_API_URL=http://localhost:8080
PUBLIC_DATA_SOURCE=api
```

> **Note:** Set `PUBLIC_DATA_SOURCE=mock` to use mock data without running the backend.

---

## Backend Setup

### 1. Clone the Repository

```bash
git clone https://github.com/aidantrabs/trinbago-hackathon.git
cd trinbago-hackathon
```

### 2. Navigate to Backend

```bash
cd backend
```

### 3. Install Go Dependencies

```bash
go mod download
```

### 4. Add Environment File

Copy the `.env` file provided by a codeowner to `/backend/.env`.

Or create from example:

```bash
cp .env.example .env
# Then edit .env with the actual values from a codeowner
```

### 5. Verify Setup

```bash
go build ./...
```

If this completes without errors, the backend is ready.

---

## Frontend Setup

### 1. Navigate to Frontend

From the project root:

```bash
cd frontend
```

### 2. Install Dependencies

```bash
npm install
```

### 3. Add Environment File

Create `/frontend/.env`:

```env
PUBLIC_API_URL=http://localhost:8080
PUBLIC_DATA_SOURCE=api
```

### 4. Verify Setup

```bash
npm run check
```

---

## Running the Project

### Option 1: Full Stack (Backend + Frontend)

Open two terminal windows/tabs.

**Terminal 1 - Backend:**

```bash
cd backend
go run ./cmd/server
```

You should see:

```
server starting on port 8080
```

**Terminal 2 - Frontend:**

```bash
cd frontend
npm run dev
```

You should see:

```
VITE vX.X.X  ready in XXX ms

➜  Local:   http://localhost:5173/
```

Open http://localhost:5173 in your browser.

### Option 2: Frontend Only (Mock Data)

If you don't need the backend, use mock data:

1. Set `PUBLIC_DATA_SOURCE=mock` in `/frontend/.env`
2. Run the frontend:

```bash
cd frontend
npm run dev
```

---

## Project Structure

```
trinbago-hackathon/
├── backend/               # Go API server
│   ├── cmd/server/        # Main entry point
│   ├── internal/          # Internal packages
│   │   ├── config/        # Configuration loading
│   │   ├── db/            # Database queries (sqlc)
│   │   ├── email/         # Email service (Resend)
│   │   ├── handler/       # HTTP handlers
│   │   ├── middleware/    # Rate limiting, auth
│   │   └── service/       # Business logic
│   └── sql/               # SQL migrations & queries
├── frontend/              # SvelteKit app
│   ├── src/
│   │   ├── lib/           # Components, utils, API client
│   │   └── routes/        # Pages
│   └── static/            # Static assets
└── docs/                  # Documentation
```

---

## Useful Commands

### Backend

```bash
# Run server
go run ./cmd/server

# Build binary
go build -o server ./cmd/server

# Run with hot reload (install air first: go install github.com/air-verse/air@latest)
air
```

### Frontend

```bash
# Development server
npm run dev

# Type checking
npm run check

# Build for production
npm run build

# Preview production build
npm run preview
```

### Deployment (Fly.io)

```bash
# Deploy backend
cd backend
fly deploy

# View logs
fly logs

# SSH into container
fly ssh console
```

---

## Common Issues

### "go: command not found"

Go is not in your PATH. Add it:

```bash
# macOS (Homebrew)
echo 'export PATH=$PATH:/opt/homebrew/bin' >> ~/.zshrc
source ~/.zshrc

# Linux
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

### "npm: command not found"

Node.js is not installed or not in PATH. Reinstall Node.js following the prerequisites.

### Backend won't start: "failed to connect to database"

- Verify `DATABASE_URL` in `/backend/.env` is correct
- Ensure you have network access to the database
- Contact a codeowner if using a remote database

### Frontend shows "Failed to fetch" errors

- Ensure the backend is running on port 8080
- Check `PUBLIC_API_URL` in `/frontend/.env` matches backend URL
- Verify CORS: backend's `ALLOWED_ORIGINS` should include `http://localhost:5173`

### "permission denied" when running commands

```bash
# Fix npm permissions
sudo chown -R $(whoami) ~/.npm

# Fix Go module cache
sudo chown -R $(whoami) ~/go
```

### Port already in use

```bash
# Find process using port 8080
lsof -i :8080

# Kill the process
kill -9 <PID>
```

---

## Codeowners

Contact for environment files and access:

- [@aidantrabs](https://github.com/aidantrabs)

---

## Next Steps

Once set up, see [DEMO.md](./DEMO.md) for testing email functionality.
