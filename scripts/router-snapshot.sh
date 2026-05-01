#!/bin/sh
set -eu

HOST="${1:-root@192.168.10.1}"

ssh "$HOST" '
echo "## uname"
uname -a

echo "## cpu"
cat /proc/cpuinfo

echo "## storage"
df -h

echo "## memory"
free

echo "## cake packages"
opkg list-installed | grep -E "kmod-sched-cake|tc|ip-full|ip-tiny" || true

echo "## interfaces"
ip link show

echo "## addresses"
ip addr show

echo "## routes"
ip route show

echo "## qdisc"
tc qdisc show || true
'
