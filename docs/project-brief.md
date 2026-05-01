# Project Brief: Aura-SQM

## Purpose

Aura-SQM is an adaptive Smart Queue Management engine for OpenWrt. It controls network traffic on a small router so latency stays low during gaming, voice, and normal household use.

The first target is a JCG Q20 router with a MediaTek MT7621 processor, 128 MB RAM, 16 MB flash, and a MyRepublic 100 Mbps fiber connection. The main user experience target is stable latency for high tick-rate competitive games such as Valorant.

## Confirmed Facts

- The daemon must run on OpenWrt.
- The target CPU family is little-endian MIPS, built with `GOARCH=mipsle`.
- The first shaper backend is Linux CAKE.
- The control loop must adjust upload and download bandwidth limits dynamically.
- The control logic must use latency error, smoothing, anti-windup, and safe fallback behavior.
- The repository is currently in documentation-first mode. No application code has been created yet.

## Product Objectives

1. Keep queueing delay low when the line is saturated.
2. Avoid throughput oscillation from overly aggressive control changes.
3. Give the gaming laptop a clear priority path without starving other clients.
4. Keep the binary small enough for 16 MB flash devices.
5. Expose enough observability to debug latency, rate changes, drops, and probe health.

## Non-Goals For The First Build

- Aura-SQM will not replace the router firewall.
- Aura-SQM will not provide a web dashboard in the first build.
- Aura-SQM will not depend on a database server.
- Aura-SQM will not require Docker on the router.
- Aura-SQM will not flush live queues during normal rate updates.
- Aura-SQM will not perform real CAKE updates until hardware validation confirms the interface and qdisc behavior.

## Primary Features

### Dynamic Bandwidth Shaper

The daemon adjusts CAKE bandwidth limits from the measured latency state. The shaper must respect a configured floor and ceiling so the network remains usable during bad probes or traffic spikes.

### Type C PID Governor

The governor computes rate adjustments from latency error. It must include anti-windup clamping and derivative damping to avoid sawtooth throughput behavior.

### Multi-Protocol Probing

The probe layer sends ICMP and UDP probes to several reflectors, including public resolvers and the ISP gateway when available. Results must be filtered before they affect the governor.

### Priority Traffic Mode

Farid-Mode marks or routes traffic from a configured gaming IP or MAC address into a high-priority CAKE tin. This must be explicit, auditable, and reversible.

### Panic Fallback

If all reflectors become unreachable, Aura-SQM must stop tightening bandwidth limits and move to a safe fallback state. The fallback must avoid trapping the user behind a broken shaper.

### Observability

Aura-SQM must expose a Prometheus-compatible `/metrics` endpoint and a terminal user interface for low-overhead monitoring over Secure Shell (SSH).

### Offline Simulator

Aura-SQM includes a simulator path that runs the governor against scripted latency and an in-memory shaper. This allows development to continue before the router arrives.

## Constraints

- Device memory is limited. The daemon must avoid heavy dependencies, large buffers, and high-cardinality metrics.
- Flash storage is limited. Builds must use linker stripping and optional UPX compression.
- OpenWrt service management must use `procd`.
- CAKE support must be present through `kmod-sched-cake`.
- The control loop target is 20 Hz to 50 Hz, but the implementation must validate CPU cost on the MT7621 device before enabling the highest rate.

## Evidence Sources

| Source | What It Confirms | Fetched At |
| --- | --- | --- |
| https://go.dev/wiki/GoMips | Go supports MIPS and MIPS little-endian cross-compilation patterns. | 2026-05-01 |
| https://go.dev/src/cmd/go/internal/help/helpdoc.go | `GOMIPS=hardfloat` and `softfloat` map to MIPS feature build tags. | 2026-05-01 |
| https://openwrt.org/packages/pkgdata/kmod-sched-cake | OpenWrt provides `kmod-sched-cake` as the CAKE kernel module package. | 2026-05-01 |
| https://openwrt.org/docs/techref/procd | OpenWrt `procd` manages daemon processes started from init scripts. | 2026-05-01 |
| https://man7.org/linux/man-pages/man8/tc-cake.8.html | CAKE supports bandwidth shaping, Diffserv presets, and `triple-isolate`. | 2026-05-01 |
| https://pkg.go.dev/github.com/vishvananda/netlink | `vishvananda/netlink` is a maintained Go module for Linux netlink access. | 2026-05-01 |
| https://prometheus.io/docs/instrumenting/exposition_formats/ | Prometheus metrics are exposed over HTTP using a text exposition format. | 2026-05-01 |

## Assumptions To Validate

- The JCG Q20 firmware image includes or can install `kmod-sched-cake`.
- The MT7621 floating point behavior works with `GOMIPS=hardfloat` on the target OpenWrt build.
- The ISP gateway responds consistently enough to be used as one reflector.
- The gaming device can be identified reliably by static IP, DHCP lease, or MAC address.
- The current OpenWrt image has enough flash space for the compressed binary and service files.

## Next Validation Action

Install the Go toolchain locally, run the offline simulator and tests, then confirm the router output for `uname -a`, `opkg list-installed | grep cake`, WAN interface names, CPU model, and available flash space when the hardware arrives.
