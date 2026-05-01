# Flow Overview

## Startup Flow

1. Load the JSON configuration file from the OpenWrt config path.
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

The offline runtime uses the same control loop shape with a scripted probe source and an in-memory shaper. This lets the governor, EWMA filters, fallback behavior, and metrics be exercised before router hardware arrives.

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

The first metrics server exposes `/metrics` from the in-memory runtime snapshot. It is safe to run from the simulator because it does not touch router interfaces.

## Offline Simulation Flow

1. Load `config/example.json`.
2. Create an in-memory shaper at the configured ceiling rate.
3. Create a scripted probe source with baseline latency, buffer latency, and periodic spikes.
4. Run the control loop for the requested tick count.
5. Print status snapshots every ten ticks.
6. Optionally serve `/metrics` while the simulation runs.

Planned command:

```bash
go run ./cmd/aurasim --config config/example.json --ticks 120 --serve-metrics
```

## Deployment Flow

1. Cross-compile from WSL2 or Linux.
2. Optionally compress the binary with UPX.
3. Copy the binary to `/usr/bin/aura-sqm`.
4. Install the `procd` init script.
5. Enable and start the service.
6. Confirm metrics and qdisc state.

The initial OpenWrt init skeleton lives at `packaging/openwrt/aura-sqm.init`.

Planned commands:

```bash
GOOS=linux GOARCH=mipsle GOMIPS=hardfloat CGO_ENABLED=0 \
  go build -trimpath -ldflags="-s -w" -o aura-sqm ./cmd/aurad

upx --brute aura-sqm
scp aura-sqm root@192.168.10.1:/usr/bin/
```

Config validation command:

```bash
aura-sqm --config /etc/aura-sqm/config.json --validate-config
```

## Next Validation Action

Run the simulator locally after the Go toolchain is installed, then capture real router interface names and current qdisc output when the hardware arrives.
