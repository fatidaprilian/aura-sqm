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

If SSH password login works but automated commands fail, install the local public key into OpenWrt Dropbear first:

```bash
sh scripts/router-authorize-key.sh root@192.168.10.1
```

Before installing a persistent service, validate the cross-compiled daemon from `/tmp`:

```bash
sh scripts/router-deploy-validate.sh root@192.168.10.1 config/example.json
```

This is a safe preflight. It does not install `/usr/bin/aura-sqm`, does not enable `procd`, and does not modify qdisc state.

## Required Facts

### Device And Kernel

- Confirmed from LuCI snapshot on 2026-05-05: router model is JCG Q20.
- Confirmed from LuCI snapshot on 2026-05-05: kernel version is 5.10.168.
- Confirmed from LuCI snapshot on 2026-05-05: architecture is MediaTek MT7621 ver:1 eco:3.
- Confirmed from SSH on 2026-05-05: `opkg print-architecture` reports `arch mipsel_24kc 10`.
- Confirmed from SSH on 2026-05-05: `uname -a` reports Linux `5.10.168` built on 2023-03-15.
- Confirmed from SSH on 2026-05-05: `/proc/version` reports `mipsel-openwrt-linux-musl-gcc` for OpenWrt revision `r20069-292184b6aa`.
- Use `GOMIPS=softfloat` as the router preflight default until a hardfloat binary is tested successfully on-device.
- Confirmed from `/tmp` preflight on 2026-05-05: the softfloat `linux/mipsle` Aura-SQM binary validates `config/example.json` successfully on the router.
- Confirmed from `/tmp` preflight on 2026-05-05: `--once-status` runs on the router and reports upload/download rates at `100000000` bps with healthy probe state.
- Confirmed from SSH snapshot on 2026-05-05: overlay filesystem has 70.4 MB total and 66.8 MB available.
- Confirmed from SSH snapshot on 2026-05-05: `/tmp` has 122.2 MB total and 121.9 MB available.
- Confirmed from LuCI: configuration backup archive was downloaded before package or qdisc changes.

### CAKE Support

- Confirmed from SSH on 2026-05-05: `/etc/openwrt_release` reports `DISTRIB_RELEASE='22.03-SNAPSHOT'`, `DISTRIB_REVISION='r20069-292184b6aa'`, `DISTRIB_TARGET='ramips/mt7621'`, and `DISTRIB_ARCH='mipsel_24kc'`.
- Confirmed from SSH on 2026-05-05: `/etc/opkg/distfeeds.conf` points only to `https://downloads.openwrt.org/releases/22.03-SNAPSHOT/targets/ramips/mt7621/packages` and `https://downloads.openwrt.org/releases/22.03-SNAPSHOT/packages/mipsel_24kc/base`.
- Confirmed from SSH snapshot on 2026-05-05: `kmod-sched-cake` is not installed yet.
- Confirmed from SSH snapshot on 2026-05-05: `tc` is not installed yet, so `tc qdisc show` cannot run.
- Confirmed from SSH on 2026-05-05: `opkg update` fails for the configured `22.03-SNAPSHOT` feeds, so no package installation has happened yet.
- Do not switch feeds to a different OpenWrt release for kernel modules unless the kernel ABI match is proven.
- Do not edit `/etc/opkg/distfeeds.conf` as an automatic workaround for Aura-SQM validation.
- Do not perform firmware flash, sysupgrade, or attended sysupgrade as part of Aura-SQM validation.
- Confirm CAKE can be added on a disposable test interface before touching the real WAN path.

### Interface Map

- Confirmed from LuCI snapshot on 2026-05-05: IPv4 upstream uses static address `192.168.1.2/24`, gateway `192.168.1.1`, DNS `1.1.1.1` and `8.8.8.8`.
- Confirmed from LuCI snapshot on 2026-05-05: LuCI reports WAN device as Ethernet Adapter `"wan"`.
- Confirmed from SSH snapshot on 2026-05-05: WAN interface is `wan` with `192.168.1.2/24`.
- Confirmed from SSH snapshot on 2026-05-05: LAN bridge is `br-lan` with `192.168.10.1/24`.
- Confirmed from SSH snapshot on 2026-05-05: default route is `default via 192.168.1.1 dev wan`.
- Record ingress shaping strategy needed for download control.

### Probe Reachability

- Confirm the router can reach the ISP gateway.
- Confirm public reflectors respond.
- Compare ICMP and UDP behavior before choosing probe defaults.

### Priority Traffic

- Confirmed from LuCI snapshot on 2026-05-05: active DHCP leases include `DESKTOP-5F0U9I1` at `192.168.10.135`, `POCO-F1` at `192.168.10.149`, `POCO-F4` at `192.168.10.163`, and an unnamed client at `192.168.10.192`.
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
