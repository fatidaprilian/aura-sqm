#!/bin/sh
set -eu

HOST="${1:-root@192.168.10.1}"
KEY_PATH="${2:-${HOME}/.ssh/id_ed25519.pub}"

case "$HOST" in
	*[!A-Za-z0-9_@%+=:,./-]*)
		echo "AURA_SSH_HOST: unsupported host characters" >&2
		exit 2
		;;
esac

if [ ! -r "$KEY_PATH" ]; then
	echo "AURA_SSH_KEY: public key not found at $KEY_PATH" >&2
	echo "Create one with: ssh-keygen -t ed25519 -f ${KEY_PATH%.pub}" >&2
	exit 2
fi

ssh "$HOST" '
	set -eu
	umask 077
	mkdir -p /etc/dropbear
	touch /etc/dropbear/authorized_keys
	key="$(cat)"
	if ! grep -qxF "$key" /etc/dropbear/authorized_keys; then
		printf "%s\n" "$key" >> /etc/dropbear/authorized_keys
	fi
	echo "AURA_SSH_KEY_OK: /etc/dropbear/authorized_keys updated"
' < "$KEY_PATH"
