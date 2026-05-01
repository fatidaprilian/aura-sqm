# Hardware Validation Checklist

## Purpose

This checklist captures the router facts Aura-SQM needs before real CAKE and netlink control are enabled.

Run it when the JCG Q20 arrives and OpenWrt is installed.

## Snapshot Command

From the development machine:

```bash
sh scripts/router-snapshot.sh root@192.168.10.1
```

Save the output outside the repository if it contains private network details.

## Required Facts

### Device And Kernel

- Confirm the router model.
- Confirm the kernel version is 5.4 or newer.
- Confirm the CPU is MT7621 or compatible with the planned MIPS little-endian build.
- Confirm available flash and memory.

### CAKE Support

- Confirm `kmod-sched-cake` is installed.
- Confirm `tc qdisc show` works.
- Confirm CAKE can be added on a disposable test interface before touching the real WAN path.

### Interface Map

- Record WAN interface name.
- Record LAN bridge name.
- Record PPPoE, VLAN, or DHCP details when present.
- Record ingress shaping strategy needed for download control.

### Probe Reachability

- Confirm the router can reach the ISP gateway.
- Confirm public reflectors respond.
- Compare ICMP and UDP behavior before choosing probe defaults.

### Priority Traffic

- Confirm the gaming laptop static IP or DHCP lease.
- Confirm the MAC address if MAC matching will be used.
- Confirm traffic marking affects only the intended device.

## Safety Gate Before Real Shaping

Do not enable real shaper writes until these checks pass:

- The current qdisc state is captured.
- A rollback command is documented.
- The daemon validates config successfully on-router.
- The operator has SSH access from LAN after service restart.
- The bandwidth floor is high enough to preserve remote access.

## Next Action After Validation

Implement the netlink-backed shaper controller and test a no-op or same-rate update before enabling dynamic rate changes.
