# Architecture Decision Record: Aura-SQM Initial Architecture

## Status

Proposed. This ADR defines the first implementation direction. It must be revisited after hardware validation on the JCG Q20 router.

## Decision

Aura-SQM will be a single Go daemon with clear internal modules. It will run directly on OpenWrt as a `procd` service. The first implementation will not use a database, web application, Docker runtime, or multi-service topology.

The daemon will use native Linux netlink calls for qdisc and filter control. Shelling out to `tc` may be allowed only for diagnostics or emergency fallback paths, not for the normal control loop.

## Rationale

The router has limited memory and flash. A single static daemon is easier to deploy, restart, inspect, and remove than a multi-process system. The control loop also needs predictable timing, so avoiding shell process creation in the hot path is important.

Netlink is the right control boundary because CAKE is managed by the Linux traffic control subsystem. The planned Go dependency is `github.com/vishvananda/netlink`, which gives direct netlink access without forcing the daemon to parse command-line output.

## Confirmed Facts

- CAKE supports shaper bandwidth, Diffserv presets, and flow isolation modes.
- OpenWrt provides CAKE through `kmod-sched-cake`.
- OpenWrt `procd` is the target service manager.
- Prometheus text exposition can be served over HTTP for lightweight metrics.
- The project target is an embedded router, not a cloud service.

## Recommended Module Boundaries

### `cmd/aurad`

Daemon entrypoint. It loads configuration, initializes dependencies, starts the control loop, starts observability endpoints, and handles shutdown.

### `internal/config`

Configuration loading and validation. It owns defaults, units, and user-facing config errors.

### `internal/probe`

ICMP and UDP probing, reflector scheduling, timeout handling, and raw sample collection.

### `internal/filter`

EWMA windows, outlier rejection, jitter suppression, and conversion from samples to control input.

### `internal/control`

PID governor, anti-windup, derivative damping, floor/ceiling enforcement, and panic fallback decisions.

### `internal/shaper`

Netlink qdisc setup, CAKE option updates, interface discovery, and priority marking hooks.

### `internal/priority`

Gaming device matching and traffic priority policy. This module must stay explicit because a bad priority rule can affect every LAN client.

### `internal/observe`

Prometheus metrics, status snapshots, and structured logs.

### `internal/tui`

Terminal monitoring. It reads state snapshots and must not control the shaper directly in the first build.

## Data Flow

1. Probes measure current path latency.
2. Filters remove outliers and maintain fast and slow EWMA windows.
3. The governor compares filtered latency against the target baseline.
4. The governor computes the next safe rate within floor and ceiling limits.
5. The shaper applies the rate update in place through netlink.
6. Observability publishes the current state to logs, `/metrics`, and the TUI.

## Safety Decisions

- The daemon must prefer stale-safe behavior over aggressive shaping when probe health is poor.
- Rate changes must be bounded per control tick.
- Queue flushes are not allowed during normal rate updates.
- The gaming priority policy must be disabled by config when no gaming device is configured.
- Any privileged operation must be documented because the daemon runs on a router.

## Rejected Options

### Shell-First `tc` Control Loop

Rejected for the hot path. It adds process startup overhead, string parsing, and harder error handling. It can remain useful for diagnostics.

### Multi-Service Architecture

Rejected for the first build. The target device does not justify separate services for probing, control, metrics, and UI.

### Embedded Database

Rejected for the first build. Static configuration and in-memory telemetry are enough for the required behavior.

### Web Dashboard First

Rejected for the first build. SSH and Prometheus are lower cost and fit router administration better.

## Assumptions To Validate

- Netlink can update CAKE options without replacing the qdisc in a way that causes packet loss.
- The selected Go netlink package exposes all CAKE attributes needed for the first build.
- ICMP probing permissions are available or can be granted safely on OpenWrt.
- UDP probing reflectors can be selected without creating false congestion signals.

## Next Validation Action

Build a small hardware probe after docs approval: detect WAN interface, read current qdisc, and perform a no-op CAKE netlink update on a test interface before implementing the PID loop.
