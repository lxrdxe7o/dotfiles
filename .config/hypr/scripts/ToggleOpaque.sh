#!/bin/bash
# Toggle full opacity (no transparency/blur) on active window

# Get active window info
WINDOW_JSON=$(hyprctl activewindow -j)
ADDR=$(echo "$WINDOW_JSON" | jq -r '.address')

if [[ -z "$ADDR" || "$ADDR" == "null" ]]; then
    notify-send "No active window"
    exit 1
fi

# Use a state file to track toggle state per window
STATE_DIR="/tmp/hypr-opaque-state"
mkdir -p "$STATE_DIR"
STATE_FILE="$STATE_DIR/$ADDR"

if [[ -f "$STATE_FILE" ]]; then
    # Currently opaque, restore transparency
    hyprctl dispatch setprop address:$ADDR opaque 0
    rm "$STATE_FILE"
    notify-send -t 1500 " Window" "Transparency restored"
else
    # Make fully opaque
    hyprctl dispatch setprop address:$ADDR opaque 1
    touch "$STATE_FILE"
    notify-send -t 1500 " Window" "Full opacity enabled"
fi
