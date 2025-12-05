#!/bin/sh
set -e

# Try to resolve postgres hostname
POSTGRES_HOST="postgres"
POSTGRES_IP=""

# Try to get IP address using getent (if available)
if command -v getent >/dev/null 2>&1; then
    POSTGRES_IP=$(getent hosts $POSTGRES_HOST | awk '{ print $1 }' | head -n1)
fi

# If getent didn't work, try using ping to resolve
if [ -z "$POSTGRES_IP" ]; then
    POSTGRES_IP=$(ping -c 1 $POSTGRES_HOST 2>/dev/null | grep -oP '\(\K[^)]+' | head -n1 || echo "")
fi

# If we got an IP, use it instead of hostname
if [ -n "$POSTGRES_IP" ]; then
    echo "Resolved postgres to IP: $POSTGRES_IP"
    export DATABASE_URL=$(echo $DATABASE_URL | sed "s/@${POSTGRES_HOST}:/@${POSTGRES_IP}:/")
fi

# Execute the main command
exec "$@"

