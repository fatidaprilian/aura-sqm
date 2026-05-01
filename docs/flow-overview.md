# Flow Overview

## Startup Flow

1. Load the configuration file from the OpenWrt config path.
2. Validate bandwidth limits, latency target, reflector list, and gaming device identity.
3. Detect WAN and ingress shaping interfaces.
4. Confirm CAKE availability.
5. Install or reconcile qdisc state.
6. Start probe workers.
7. Start the control loop.
8. Start `/metrics` and the terminal status publisher.

If startup cannot verify the shaper backend, Aura-SQM must exit with a clear error. It must not leave a half-configured traffic control state.

## Control Loop Flow

The control loop runs at a configured interval between 20 ms and 50 ms.

1. Read the latest probe samples.
2. Drop invalid samples and extreme outliers.
3. Update fast and slow EWMA windows.
4. Compute latency error against the target baseline.
5. Run the Type C PID controller.
6. Clamp the result to the bandwidth floor and ceiling.
7. Apply a maximum per-tick rate delta.
8. Push the new rate through netlink when the change is meaningful.
9. Publish metrics and status snapshots.

The loop must tolerate missed probe samples. Missing one tick is not a panic condition.

## Probe Flow

Probe workers send ICMP and UDP probes to multiple reflectors. Each sample includes:

- reflector ID
- protocol
- send timestamp
- receive timestamp
- round-trip time
- timeout state
- error state

The filter layer must compare reflectors and protocols before changing the governor input. One bad reflector must not force a rate reduction by itself.

## Panic Fallback Flow

Panic fallback starts when all reflectors fail for a configured window.

1. Mark probe health as failed.
2. Freeze or relax the shaper according to configuration.
3. Stop further rate decreases.
4. Emit a warning metric and log entry.
5. Continue probing at a safe interval.
6. Recover only after enough reflectors become healthy again.

The fallback must be boring and predictable. It should keep the network reachable for repair.

## Priority Traffic Flow

Farid-Mode is a controlled priority policy for one gaming device.

1. Match the configured device by IP address, MAC address, or both.
2. Mark packets for the intended CAKE tin.
3. Keep the rule visible in logs and metrics.
4. Remove or disable the rule when the config changes.

The policy must not silently priority-mark the whole LAN.

## Observability Flow

The daemon emits:

- structured logs for state transitions and errors
- Prometheus metrics for scrape-based monitoring
- TUI snapshots for SSH monitoring

Metrics and TUI data must be read-only in the first build. Control actions should stay in the daemon config and service lifecycle.

## Deployment Flow

1. Cross-compile from WSL2 or Linux.
2. Optionally compress the binary with UPX.
3. Copy the binary to `/usr/bin/aura-sqm`.
4. Install the `procd` init script.
5. Enable and start the service.
6. Confirm metrics and qdisc state.

Planned commands:

```bash
GOOS=linux GOARCH=mipsle GOMIPS=hardfloat CGO_ENABLED=0 \
  go build -trimpath -ldflags="-s -w" -o aura-sqm ./cmd/aurad

upx --brute aura-sqm
scp aura-sqm root@192.168.10.1:/usr/bin/
```

## Next Validation Action

Capture real router interface names and current qdisc output before the shaper module is coded.
