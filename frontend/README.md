# KULTUR Frontend

SvelteKit frontend for the KULTUR platform.

## Tech Stack

- **SvelteKit 2** - Full-stack framework
- **Svelte 5** - Reactive UI with runes
- **TailwindCSS 4** - Utility-first CSS
- **shadcn-svelte** - UI components
- **Biome** - Linting and formatting

## Setup

### Prerequisites

- Node.js 20+

### Install Dependencies

```bash
npm install
```

### Environment Variables

Create a `.env` file:

```env
PUBLIC_API_URL=http://localhost:8080
PUBLIC_DATA_SOURCE=api
```

| Variable | Description |
|:---------|:------------|
| `PUBLIC_API_URL` | Backend API URL |
| `PUBLIC_DATA_SOURCE` | `api` for real data, `mock` for local mock data |

## Development

```bash
npm run dev
```

Opens at http://localhost:5173

## Scripts

| Command | Description |
|:--------|:------------|
| `npm run dev` | Start dev server |
| `npm run build` | Production build |
| `npm run preview` | Preview production build |
| `npm run check` | TypeScript type checking |
| `npm run format` | Format with Biome |
| `npm run lint` | Lint with Biome |
| `npm run deploy` | Deploy to Vercel |

## Project Structure

```
src/
├── lib/
│   ├── api.ts              # API client
│   ├── config.ts           # Environment config
│   ├── components/
│   │   ├── ui/             # shadcn-svelte components
│   │   ├── calendar/       # Calendar components
│   │   └── newsletter/     # Newsletter signup
│   ├── data/
│   │   ├── festivals.ts    # Festival data
│   │   └── memories.ts     # Memory data
│   └── types/
│       └── festival.ts     # TypeScript types
└── routes/
    ├── +page.svelte        # Home page
    └── festivals/
        ├── +page.svelte    # Calendar view
        └── [slug]/         # Festival detail
```

## Deployment

Deployed to Vercel at [kultur-tt.app](https://kultur-tt.app).

```bash
npm run deploy
```

Or push to `main` branch for automatic deployment.
