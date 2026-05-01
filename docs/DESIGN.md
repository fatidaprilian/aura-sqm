# TUI Design Contract

## Scope

Aura-SQM does not plan a browser UI in the first build. The user-facing interface is an SSH-friendly terminal user interface plus Prometheus metrics.

The TUI must show real-time network control state without adding meaningful CPU, memory, or rendering overhead on the router.

## Design Intent

The interface should feel like an instrument panel for live network control, not a decorative dashboard. It must help an operator answer four questions quickly:

1. Is latency healthy?
2. What rate is the shaper applying now?
3. Are probes healthy?
4. Is fallback or priority mode active?

## Information Layout

The first TUI view should include:

- current upload and download shaper rates
- latency fast and slow EWMA values
- jitter estimate
- packet drop and mark counters when available
- probe health by reflector group
- fallback state
- Farid-Mode state
- current WAN interface
- loop interval and overrun count

## Interaction Model

The first build should be read-only.

Planned keys:

- `q`: quit
- `r`: refresh now
- `1`: overview view
- `2`: probe view
- `3`: shaper view
- `4`: controller view

No TUI action should mutate traffic control state in the first build.

## Visual Rules

- Use high-contrast text that works in common SSH terminals.
- Do not rely on color alone. Pair status color with labels such as `OK`, `WARN`, and `FAIL`.
- Keep the screen stable. Values may update, but layout should not jump.
- Prefer compact tables and fixed-width numeric fields.
- Avoid animation. The update cadence should be readable over SSH.

## Product-Specific Anchor

Anchor reference: a latency oscilloscope used for live signal tuning.

This means the UI should emphasize signal quality, stable traces, thresholds, and control response. It should not look like a sales dashboard.

## Next Validation Action

After the daemon state model exists, create a terminal mock using recorded sample state before wiring it to live control data.
