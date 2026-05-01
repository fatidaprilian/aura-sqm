# Data And Configuration Model

## Decision

Aura-SQM will not use a database in the first build. Runtime state is kept in memory, exposed through metrics, and rebuilt after service restart.

The durable data model is configuration only.

## Rationale

The target router has 16 MB flash and 128 MB RAM. A database would add operational cost without solving a first-build requirement. The control loop needs current measurements, not long-term storage. Historical analysis belongs in Prometheus and Grafana outside the router.

## Configuration Entities

### Shaper Settings

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| `wan_interface` | string | no | Optional explicit WAN interface. |
| `auto_discover_wan` | boolean | yes | Enables WAN interface detection. |
| `upload_floor_mbps` | number | yes | Minimum upload shaper rate. |
| `upload_ceiling_mbps` | number | yes | Maximum upload shaper rate. |
| `download_floor_mbps` | number | yes | Minimum download shaper rate. |
| `download_ceiling_mbps` | number | yes | Maximum download shaper rate. |
| `cake_diffserv` | string | yes | Initial value should be `diffserv4`. |
| `cake_isolation` | string | yes | Initial value should be `triple-isolate`. |

### Control Settings

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| `loop_interval_ms` | integer | yes | Must stay between 20 and 50 ms until validated otherwise. |
| `target_latency_ms` | number | yes | Desired latency target above idle baseline. |
| `kp` | number | yes | Proportional gain. |
| `ki` | number | yes | Integral gain. |
| `kd` | number | yes | Derivative gain. |
| `integral_min` | number | yes | Anti-windup lower clamp. |
| `integral_max` | number | yes | Anti-windup upper clamp. |
| `max_rate_delta_mbps` | number | yes | Maximum rate change per tick. |

### Probe Settings

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| `reflectors` | array | yes | Public and ISP reflector list. |
| `protocols` | array | yes | Initial values: `icmp`, `udp`. |
| `timeout_ms` | integer | yes | Probe timeout. |
| `fast_ewma_alpha` | number | yes | Expected range: 0.1 to 0.3. |
| `slow_ewma_alpha` | number | yes | Must be lower than fast alpha. |
| `outlier_threshold` | number | yes | Implementation-defined threshold. |

### Priority Settings

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| `enabled` | boolean | yes | Enables Farid-Mode. |
| `device_ip` | string | no | Gaming device IP address. |
| `device_mac` | string | no | Gaming device MAC address. |
| `target_tin` | string | yes | Initial value should map to CAKE voice or video priority. |

### Observability Settings

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| `metrics_listen` | string | yes | Default should bind to LAN or localhost, not WAN. |
| `log_level` | string | yes | `info`, `warn`, `error`, or `debug`. |
| `tui_enabled` | boolean | yes | Enables terminal monitor. |

## Runtime State

Runtime state is not persisted by Aura-SQM.

Important state snapshots:

- probe health
- filtered latency
- current shaper rate
- PID terms
- fallback state
- priority rule state
- interface discovery state

## Privacy Rules

- Do not store or export raw MAC addresses in metrics.
- Do not store probe packet payloads.
- Do not commit router credentials.
- Keep reflector names generic unless the user config explicitly names them.

## Assumptions To Validate

- UCI is the preferred config interface for the target user workflow.
- The first build can avoid writing historical state to flash.
- Grafana or another Prometheus consumer will run outside the router.

## Next Validation Action

Pick the config file format after confirming whether the router admin workflow should be UCI-native or plain-file based.
