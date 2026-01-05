#!/bin/bash
set -euo pipefail
# Power Profile Switcher with Notifications

iDIR="$HOME/.config/swaync/icons"

# Switch to next profile
asusctl profile -n

# Get the new active profile
PROFILE=$(asusctl profile -p 2>/dev/null | grep "Active profile" | awk '{print $NF}')

# Set icon and message based on profile
case "$PROFILE" in
    Performance)
        ICON="$iDIR/power.png"
        MSG="Maximum performance"
        ;;
    Balanced)
        ICON="$iDIR/power.png"
        MSG="Balanced power & performance"
        ;;
    Quiet|LowPower)
        ICON="$iDIR/battery-quarter-solid.svg"
        MSG="Power saving mode"
        ;;
    *)
        ICON="$iDIR/power.png"
        MSG="Profile changed"
        ;;
esac

# Send notification
notify-send -h string:x-canonical-private-synchronous:power-profile \
    -u low \
    -i "$ICON" \
    " Power Profile" \
    "$PROFILE - $MSG"
