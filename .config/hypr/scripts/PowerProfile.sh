#!/bin/bash
set -euo pipefail
# Power Profile Switcher with Notifications

iDIR="$HOME/.config/swaync/icons"

# Get current profile and cycle to next
CURRENT=$(powerprofilesctl get)

case "$CURRENT" in
    performance)
        powerprofilesctl set balanced
        ;;
    balanced)
        powerprofilesctl set power-saver
        ;;
    power-saver)
        powerprofilesctl set performance
        ;;
esac

# Get the new active profile
PROFILE=$(powerprofilesctl get)

# Set icon and message based on profile
case "$PROFILE" in
    performance)
        ICON="$iDIR/power.png"
        MSG="Maximum performance"
        ;;
    balanced)
        ICON="$iDIR/power.png"
        MSG="Balanced power & performance"
        ;;
    power-saver)
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
