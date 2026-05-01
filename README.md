# Aura-SQM

Aura-SQM is an adaptive Smart Queue Management daemon for OpenWrt routers. It is designed for a JCG Q20 class device with a MediaTek MT7621 processor, 128 MB RAM, 16 MB flash, and a MyRepublic 100 Mbps fiber line.

The goal is simple: keep latency low during gaming and heavy household traffic by adjusting CAKE bandwidth limits in real time. The daemon should protect competitive game traffic, reduce bufferbloat, and recover safely when probing or interface state becomes unreliable.

## Current Status

This repository is in documentation-first preparation. Application code is intentionally not scaffolded yet. The first implementation phase should use the docs in `docs/` as the source of truth.

## Target Runtime

- Language: Go 1.21 or newer, validated before implementation.
- Target OS: OpenWrt with Linux kernel 5.4 or newer.
- Target architecture: `linux/mipsle`.
- Build mode: static binary with `CGO_ENABLED=0`.
- Router dependency: `kmod-sched-cake`.
- Service manager: OpenWrt `procd`.

## Planned Build Command

```bash
GOOS=linux GOARCH=mipsle GOMIPS=hardfloat CGO_ENABLED=0 \
  go build -trimpath -ldflags="-s -w" -o aura-sqm ./cmd/aurad
```

## Local Smoke Commands

The first implementation scaffold supports config validation and one status snapshot:

```bash
go run ./cmd/aurad --config config/example.json --validate-config
go run ./cmd/aurad --config config/example.json --once-status
```

Optional size reduction:

```bash
upx --brute aura-sqm
```

## Planned Deployment

```bash
scp aura-sqm root@192.168.10.1:/usr/bin/
```

The OpenWrt package or init script must install a `procd` service so the daemon starts on boot and restarts on failure.

## Documentation Map

- [Project Brief](docs/project-brief.md)
- [Architecture Decision Record](docs/architecture-decision-record.md)
- [Flow Overview](docs/flow-overview.md)
- [API Contract](docs/api-contract.md)
- [Data And Configuration Model](docs/database-schema.md)
- [TUI Design Contract](docs/DESIGN.md)

## Git Setup

This repository uses `main` as the default branch and the remote below:

```bash
git remote add origin https://github.com/fatidaprilian/aura-sqm.git
```

Local commit identity is expected to use:

```bash
git config user.name biqar
git config user.email biqar@users.noreply.github.com
```

GitHub authentication for pushing is separate from commit identity. Source Control should prompt for a GitHub login or token when credentials are not already available.
