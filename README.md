# CS2Lab

CS2 utility training tool covering every Active Duty map with throw lineups, crosshair references, and combo guides.

## Features

- Utility guides for smoke grenades, flashbangs, molotovs, and HE grenades across all 7 competitive maps
- Each lineup includes position, crosshair reference screenshot, and throw type (left-click/right-click/jump-throw)
- Combo guides with map overlays showing coverage areas for multi-utility setups
- Smart search matching situational descriptions to existing content
- Community submissions with admin review before publication
- Bookmark lineups and add personal notes

## Architecture

```text
Browser
  ↓
React Frontend (Oxelia51 unified UI, image annotation)
  ↓
Go API Layer (content CRUD, search, submission review)
  ↓
PostgreSQL / SQLite (maps, lineups, user data)
```

The online version runs on the Oxelia51 platform. Screenshots and annotation data are managed by the Go backend. The desktop version uses SQLite and bundles map data in the binary.

## Requirements

- Online: Oxelia51 platform (Go, PostgreSQL, React)
- Desktop: standalone executable, no runtime dependencies

## Installation

### Desktop

Download `CS2Lab.exe` from [GitHub Releases](https://github.com/XiaoleC05/CS2Lab/releases).

### Online

Integrated into the Oxelia51 platform. See [Oxelia51 deployment guide](https://github.com/XiaoleC05/Oxelia51).

## Usage

### Online

1. Visit [oxelia51.com](https://oxelia51.com), register and sign in
2. Open CS2Lab from the tools menu
3. Select a map and utility type to browse lineups

### Desktop

1. Double-click `CS2Lab.exe` to start
2. All data is built-in, no internet required

## Roadmap

- [ ] Complete utility coverage for Dust2 (first map)
- [ ] Expand to all 7 Active Duty maps
- [ ] Smart search
- [ ] Community submissions and review

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/xxx`)
3. Commit your changes (`git commit -m 'Add xxx'`)
4. Push the branch (`git push origin feature/xxx`)
5. Open a Pull Request

## License

This project is licensed under the MIT License.
