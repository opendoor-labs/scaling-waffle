#!/usr/bin/env bash

set -e

STARTUP_COUNTDOWN=5

SERVER_ADDRESS=${SERVER_ADDRESS:-server:80}
CONN_COUNT=${CONN_COUNT:-10000}
CONN_PERIOD_MS=${CONN_PERIOD_MS:-100}

for i in $(seq ${STARTUP_COUNTDOWN}); do
	echo "Starting in $(($STARTUP_COUNTDOWN-$i+1)) sec..."
	sleep 1
done

if [[ -z ${ATTACK_KILLSWITCH} ]]; then
	echo "Not attacking, unset \$ATTACK_KILLSWITCH to re-enable"
fi

echo "Launching 'attack'"

/usr/bin/debug-client "${SERVER_ADDRESS}" "${CONN_COUNT}" "${CONN_PERIOD_MS}"

# Pause forever
cat /dev/zero
