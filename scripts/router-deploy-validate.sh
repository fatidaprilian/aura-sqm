#!/bin/sh
set -eu

HOST="${1:-root@192.168.10.1}"
CONFIG_PATH="${2:-config/example.json}"
BUILD_DIR="${BUILD_DIR:-build}"
GOCACHE="${GOCACHE:-/tmp/aura-sqm-go-cache}"
GOMODCACHE="${GOMODCACHE:-/tmp/aura-sqm-go-mod-cache}"
SCP="${SCP:-scp -O}"
BIN_PATH="$BUILD_DIR/aura-sqm-openwrt-mipsle"
REMOTE_BIN="/tmp/aura-sqm"
REMOTE_CONFIG="/tmp/aura-sqm-config.json"
export GOCACHE GOMODCACHE

case "$HOST" in
	*[!A-Za-z0-9_@%+=:,./-]*)
		echo "AURA_SSH_HOST: unsupported host characters" >&2
		exit 2
		;;
esac

if [ ! -r "$CONFIG_PATH" ]; then
	echo "AURA_CONFIG_READ: cannot read $CONFIG_PATH" >&2
	exit 2
fi

mkdir -p "$BUILD_DIR" "$GOCACHE" "$GOMODCACHE"

echo "## build"
GOOS=linux GOARCH=mipsle GOMIPS="${GOMIPS:-softfloat}" CGO_ENABLED=0 \
	go build -trimpath -ldflags="-s -w" -o "$BIN_PATH" ./cmd/aurad

echo "## copy to router /tmp"
$SCP "$BIN_PATH" "$HOST:$REMOTE_BIN"
$SCP "$CONFIG_PATH" "$HOST:$REMOTE_CONFIG"

echo "## validate on router"
ssh "$HOST" "
	set -eu
	chmod 0755 '$REMOTE_BIN'
	'$REMOTE_BIN' --config '$REMOTE_CONFIG' --validate-config
	'$REMOTE_BIN' --config '$REMOTE_CONFIG' --once-status
	echo 'AURA_ROUTER_VALIDATE_OK: binary runs from /tmp only; persistent service not installed'
"
