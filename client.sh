#!/usr/bin/env bash

set -e

STARTUP_COUNTDOWN=5

SERVER_ADDRESS=${SERVER_ADDRESS:-server:8050}
CONN_COUNT=${CONN_COUNT:-60000}
CONN_PERIOD_MS=${CONN_PERIOD_MS:-10}

for i in $(seq ${STARTUP_COUNTDOWN}); do
	echo "Starting in $(($STARTUP_COUNTDOWN-$i+1)) sec..."
	sleep 1
done

if [[ -n "${ATTACK_KILLSWITCH}" ]]; then
	echo "Not attacking, unset \$ATTACK_KILLSWITCH to re-enable"
	exit 0
fi

echo "Launching 'attack'"

/usr/bin/debug-client "${SERVER_ADDRESS}" "${CONN_COUNT}" "${CONN_PERIOD_MS}"

# Pause forever
cat /dev/zero
