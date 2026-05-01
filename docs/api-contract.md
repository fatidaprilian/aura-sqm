# API And Public Contract

## Scope

Aura-SQM has no public business API in the first build. It does have public operational contracts:

- Prometheus `/metrics` endpoint
- OpenWrt service lifecycle
- configuration file
- terminal user interface status fields
- planned command-line flags

These contracts must stay stable once users deploy the daemon.

## Prometheus Metrics Endpoint

### `GET /metrics`

Returns Prometheus text exposition.

Required response:

- Status: `200 OK`
- Content-Type: `text/plain; version=0.0.4; charset=utf-8`
- Body: line-oriented Prometheus metrics

The endpoint must not expose secrets, MAC addresses, private IP addresses, or raw reflector payloads.

### Initial Metric Names

| Metric | Type | Labels | Meaning |
| --- | --- | --- | --- |
| `aura_latency_fast_ewma_seconds` | gauge | `direction` | Fast latency estimate used for short-term reaction. |
| `aura_latency_slow_ewma_seconds` | gauge | `direction` | Slow latency estimate used for baseline comparison. |
| `aura_latency_jitter_seconds` | gauge | `direction` | Filtered jitter estimate. |
| `aura_shaper_rate_bits_per_second` | gauge | `direction` | Current active shaper rate. |
| `aura_shaper_floor_bits_per_second` | gauge | `direction` | Configured minimum shaper rate. |
| `aura_shaper_ceiling_bits_per_second` | gauge | `direction` | Configured maximum shaper rate. |
| `aura_probe_success_total` | counter | `reflector`, `protocol` | Successful probe count. |
| `aura_probe_failure_total` | counter | `reflector`, `protocol`, `reason` | Failed probe count. |
| `aura_probe_health` | gauge | none | `1` when at least one reflector is healthy, else `0`. |
| `aura_panic_fallback_active` | gauge | none | `1` when fallback mode is active, else `0`. |
| `aura_priority_rule_active` | gauge | none | `1` when Farid-Mode priority rule is active, else `0`. |
| `aura_control_loop_tick_total` | counter | none | Number of completed control ticks. |
| `aura_control_loop_overrun_total` | counter | none | Number of ticks that exceeded the interval budget. |

Metric labels must stay low-cardinality. Do not label metrics with raw IP addresses, MAC addresses, or per-flow IDs.

## Configuration Contract

The first config format is JSON. This keeps the first daemon scaffold dependency-free and easy to validate with the Go standard library. A later OpenWrt package may add a UCI adapter, but the daemon must keep one normalized internal config shape.

Default path:

```text
/etc/aura-sqm/config.json
```

Required fields:

- WAN interface selection mode
- upload floor and ceiling
- download floor and ceiling
- latency target
- EWMA alpha values
- control loop interval
- PID constants
- reflector list
- panic fallback policy
- gaming device identity for Farid-Mode
- metrics listen address

Invalid configuration must stop startup with a clear error. The daemon must not guess unsafe bandwidth values.

## Service Contract

The OpenWrt service must support:

```bash
/etc/init.d/aura-sqm start
/etc/init.d/aura-sqm stop
/etc/init.d/aura-sqm restart
/etc/init.d/aura-sqm enable
/etc/init.d/aura-sqm disable
```

The service should run the daemon in the foreground under `procd`; `procd` manages backgrounding and respawn behavior.

## Command-Line Contract

Planned flags:

| Flag | Meaning |
| --- | --- |
| `--config PATH` | Load configuration from an explicit path. |
| `--validate-config` | Validate config and exit. |
| `--once-status` | Print one status snapshot and exit. |
| `--version` | Print version and build metadata. |

The first implementation should keep flags minimal. Router admins should not need many flags for normal service use.

## Error Contract

Startup errors must include:

- stable error code
- human-readable message
- safe context
- suggested repair action when possible

Example:

```text
AURA_CONFIG_BANDWIDTH_RANGE: upload floor must be lower than upload ceiling
```

## Evidence Sources

| Source | What It Confirms | Fetched At |
| --- | --- | --- |
| https://prometheus.io/docs/instrumenting/exposition_formats/ | Prometheus text exposition is HTTP-based, UTF-8, and line-oriented. | 2026-05-01 |
| https://openwrt.org/docs/guide-developer/procd-init-scripts | `procd` services use foreground commands and `USE_PROCD=1`. | 2026-05-01 |

## Next Validation Action

Choose UCI or file-based config before implementing `internal/config`.
