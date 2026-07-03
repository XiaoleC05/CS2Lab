# CS2Lab — Counter-Strike 2 Utility Laboratory

> Master every grenade on every map. Step-by-step guides, lineup screenshots, and smart recommendations.

## Why CS2Lab?

In CS2, knowing how to throw a smoke grenade can win rounds. But finding reliable lineups means watching hours of YouTube videos, pausing frame by frame, and memorizing references that change with every update.

**CS2Lab** is your single source for utility knowledge: clear screenshots with crosshair markers, written throw instructions, and combo guides that show you how grenades work together.

## Features

| Module | What You Get |
|--------|-------------|
| **Utility Database** | Stats for smokes, flashes, molotovs, and HE grenades — damage, radius, duration, price |
| **Throw Guides** | Position marker + crosshair reference screenshot + throw type (left-click/right-click/jump-throw) |
| **Combo Guides** | "Smoke CT spawn, then flash over the wall" — explained with map overlays |
| **Map Coverage** | All 7 Active Duty maps: Dust2, Mirage, Inferno, Nuke, Vertigo, Ancient, Anubis |
| **Smart Search** | Describe your situation (e.g. "pushing mid on Dust2 as T") → get recommended utility |
| **Community Contributions** | Submit your own lineups — published after admin review |
| **Favorites & Notes** | Bookmark your most-used lineups and add personal tips |

## Utility Coverage

| Grenade | Stats | Throw Guides |
|---------|-------|--------------|
| Smoke Grenade | ✅ | ✅ |
| Flashbang | ✅ | ✅ |
| Molotov / Incendiary | ✅ | ✅ |
| HE Grenade | ✅ | ✅ |

## Map Coverage

All Active Duty maps, starting with Dust2 and expanding:

Dust2 → Mirage → Inferno → Nuke → Vertigo → Ancient → Anubis

## Content Style

- **No 3D rendering** — uses official map screenshots with annotation overlays
- **Screenshot-based lineups** — crosshair position clearly marked, no guesswork
- **Written + visual** — every lineup has both text instructions and a marked screenshot

## What CS2Lab Is NOT

- ❌ Not a tactics analyzer — no match replay analysis or team strategy generation
- ❌ Not an AI coach — smart search matches your situation to existing guides, it doesn't generate new tactics
- ❌ Not a full weapon guide — focused on throwable utility only

## Tech Stack

| Environment | Backend | Database | Frontend | Special |
|-------------|---------|----------|----------|---------|
| Online (Oxelia51) | Go | PostgreSQL | React | Image storage |
| Desktop (exe) | Go | SQLite | Embedded React | Same, packaged as exe |

## Getting Started

### Online (via Oxelia51)

1. Visit [oxelia51.com](https://oxelia51.com) and sign in
2. Open CS2Lab from the tools menu
3. Browse maps, search for lineups, save your favorites

### Desktop (exe)

1. Download `CS2Lab.exe` from [GitHub Releases](https://github.com/XiaoleC05/CS2Lab/releases)
2. Run the executable — opens a local web interface
3. All data is bundled in the app, no internet required for browsing

## Contributing

Found a great lineup that's not in the database? Submit it through the app. All submissions go through admin review before being published to ensure quality.

## Status

Concept phase. Development not yet started.
