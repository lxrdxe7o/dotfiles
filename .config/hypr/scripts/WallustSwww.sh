#!/bin/bash
# 💫 Robust swww → Wallust sync script (by Xero) 💫

set -euo pipefail  # Exit on error, unset variables, or failed pipes

# --- Config Paths ---
CACHE_DIR="$HOME/.cache/swww"
ROFI_LINK="$HOME/.config/rofi/.current_wallpaper"
EFFECT_COPY="$HOME/.config/hypr/wallpaper_effects/.wallpaper_current"

# --- Ensure necessary directories exist ---
mkdir -p "$(dirname "$ROFI_LINK")"
mkdir -p "$(dirname "$EFFECT_COPY")"

# --- Check dependencies ---
for cmd in hyprctl wallust cp ln tail; do
    command -v "$cmd" >/dev/null 2>&1 || { 
        echo "❌ Missing required command: $cmd"; 
        exit 1; 
    }
done

# --- Get focused monitor ---
current_monitor=$(hyprctl monitors | awk '/^Monitor/{name=$2} /focused: yes/{print name}')

if [ -z "$current_monitor" ]; then
    echo "❌ Could not detect focused monitor"
    exit 1
fi

cache_file="$CACHE_DIR/$current_monitor"

if [ ! -f "$cache_file" ]; then
    echo "❌ Cache file not found for monitor '$current_monitor': $cache_file"
    exit 1
fi

# --- Extract wallpaper path ---
raw_path=$(tail -n 1 "$cache_file")
# Remove any prefix before the first /
wallpaper_path=$(echo "$raw_path" | grep -o '/.*')

# Extra check if the path is empty or invalid
if [ -z "$wallpaper_path" ] || [ ! -f "$wallpaper_path" ]; then
    echo "❌ Wallpaper path invalid or file does not exist: '$wallpaper_path'"
    exit 1
fi

echo "🎨 Current wallpaper: $wallpaper_path"

# --- Symlink for Rofi ---
ln -sf "$wallpaper_path" "$ROFI_LINK"

# --- Copy for wallpaper effects ---
cp -f "$wallpaper_path" "$EFFECT_COPY"

# --- Run Wallust to generate color scheme ---
echo "🚀 Running Wallust..."
wallust run "$wallpaper_path" -s &

echo "✅ Done!"
