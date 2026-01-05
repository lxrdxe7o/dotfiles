#!/bin/bash
# /* ---- 💫 https://github.com/JaKooLit 💫 ---- */  ##
# Script for toggling blur on/off

notif="$HOME/.config/swaync/images"

# Check if blur is currently enabled
ENABLED=$(hyprctl -j getoption decoration:blur:enabled | jq ".int")

if [ "${ENABLED}" == "1" ]; then
    # Blur is ON, turn it OFF
    hyprctl keyword decoration:blur:enabled false
    notify-send -e -u low -i "$notif/note.png" " Blur" "Disabled"
else
    # Blur is OFF, turn it ON
    hyprctl keyword decoration:blur:enabled true
    hyprctl keyword decoration:blur:size 6
    hyprctl keyword decoration:blur:passes 2
    notify-send -e -u low -i "$notif/ja.png" " Blur" "Enabled"
fi
