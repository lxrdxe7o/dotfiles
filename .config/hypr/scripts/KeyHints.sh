#!/bin/bash
# /* ---- 💫 https://github.com/JaKooLit 💫 ---- */  ##

# GDK BACKEND. Change to either wayland or x11 if having issues
BACKEND=wayland

# Check if rofi or yad is running and kill them if they are
if pidof rofi >/dev/null; then
  pkill rofi
fi

if pidof yad >/dev/null; then
  pkill yad
  exit 0
fi

# Launch yad with calculated width and height
GDK_BACKEND=$BACKEND yad \
  --center \
  --title="KooL Quick Cheat Sheet" \
  --no-buttons \
  --list \
  --column=Key: \
  --column=Description: \
  --column=Command: \
  --timeout-indicator=bottom \
  "ESC" "close this app" "" \
  "" "SUPER KEY (Windows Key)" "" \
  " SHIFT K" "Searchable Keybinds" "(Search all Keybinds via rofi)" \
  " SHIFT E" "KooL Hyprland Settings Menu" "" \
  "" "" "" \
  "───────────────" "APPS & LAUNCHERS" "───────────────" \
  " Return" "Terminal" "(kitty)" \
  " ALT Return" "Terminal + tmux" "(kitty --hold tmux attach)" \
  " SHIFT Return" "Floating Terminal" "(centered)" \
  " D" "Application Launcher" "(rofi)" \
  " A" "Desktop Overview" "(quickshell expose)" \
  " E" "File Manager" "(thunar)" \
  " B" "Default Browser" "(xdg-open)" \
  " S" "Google Search" "(rofi)" \
  "" "" "" \
  "───────────────" "WINDOW MANAGEMENT" "───────────────" \
  " Q" "Close Active Window" "(not kill)" \
  " SHIFT Q" "Kill Active Process" "(force kill)" \
  " SHIFT F" "Fullscreen" "(toggle)" \
  " CTRL F" "Fake Fullscreen" "(toggle)" \
  " SPACE" "Toggle Float" "(single window)" \
  " ALT SPACE" "Toggle All Float" "(all windows)" \
  " G" "Toggle Group" "(window group)" \
  " CTRL Tab" "Change Group Active" "" \
  "ALT Tab" "Cycle Windows" "(bring to top)" \
  "" "" "" \
  "───────────────" "LAYOUTS" "───────────────" \
  " ALT L" "Toggle Layout" "(Master/Dwindle)" \
  " I" "Add Master" "(master layout)" \
  " CTRL D" "Remove Master" "(master layout)" \
  " J / K" "Cycle Next/Prev" "(master layout)" \
  " CTRL Return" "Swap with Master" "(master layout)" \
  " SHIFT I" "Toggle Split" "(dwindle layout)" \
  " P" "Pseudo" "(dwindle layout)" \
  " M" "Split Ratio 0.3" "" \
  "" "" "" \
  "───────────────" "MOVE & RESIZE" "───────────────" \
  " Arrows" "Move Focus" "(left/right/up/down)" \
  " SHIFT Arrows" "Resize Window" "(±50 pixels)" \
  " CTRL Arrows" "Move Window" "(to direction)" \
  " ALT Arrows" "Swap Window" "(with neighbor)" \
  " + Left Click" "Move Window" "(drag)" \
  " + Right Click" "Resize Window" "(drag)" \
  "" "" "" \
  "───────────────" "WORKSPACES" "───────────────" \
  " 1-0" "Switch Workspace" "(1-10)" \
  " SHIFT 1-0" "Move Window to Workspace" "(follow)" \
  " CTRL 1-0" "Move Window Silently" "(stay)" \
  " Tab" "Next Workspace" "" \
  " SHIFT Tab" "Previous Workspace" "" \
  " Mouse Scroll" "Scroll Workspaces" "" \
  " , / ." "Prev/Next Workspace" "(comma/period)" \
  " SHIFT [ / ]" "Move to Prev/Next WS" "(follow)" \
  " CTRL [ / ]" "Move Silently" "(stay)" \
  " U" "Toggle Special WS" "" \
  " SHIFT U" "Move to Special WS" "" \
  "" "" "" \
  "───────────────" "SCREENSHOTS" "───────────────" \
  " Print" "Screenshot" "(full screen)" \
  " SHIFT Print" "Screenshot Region" "(select area)" \
  " SHIFT S" "Screenshot + Edit" "(swappy)" \
  " CTRL Print" "Screenshot 5s Delay" "" \
  " CTRL SHIFT Print" "Screenshot 10s Delay" "" \
  "ALT Print" "Screenshot Active Window" "" \
  "" "" "" \
  "───────────────" "THEMING & APPEARANCE" "───────────────" \
  " W" "Wallpaper Select" "(choose wallpaper)" \
  " SHIFT W" "Wallpaper Effects" "(imagemagick + swww)" \
  "CTRL ALT W" "Random Wallpaper" "(via swww)" \
  " CTRL ALT B" "Hide/Unhide Waybar" "(toggle)" \
  " CTRL B" "Waybar Styles" "(style menu)" \
  " ALT B" "Waybar Layout" "(layout menu)" \
  " ALT R" "Refresh All" "(waybar, swaync, rofi)" \
  " SHIFT A" "Animations Menu" "(choose animations)" \
  " CTRL R" "Rofi Theme Selector" "" \
  " CTRL SHIFT R" "Rofi Theme Selector v2" "(modified)" \
  " ALT O" "Toggle Blur" "(normal/less)" \
  " CTRL O" "Toggle Opaque" "(active window)" \
  " SHIFT O" "Zsh Theme Change" "(oh-my-zsh)" \
  "" "" "" \
  "───────────────" "UTILITIES" "───────────────" \
  " N" "Night Mode Toggle" "(hyprsunset)" \
  " ALT E" "Emoji Menu" "(rofi)" \
  " ALT V" "Clipboard Manager" "(cliphist)" \
  " ALT C" "Calculator" "(qalculate)" \
  " SHIFT M" "Online Music" "(RofiBeats)" \
  " SHIFT N" "Notification Panel" "(swaync)" \
  " SHIFT P" "Power Profile" "(cycle profiles)" \
  " SHIFT G" "Game Mode" "(toggle animations)" \
  " ALT Mouse Scroll" "Desktop Zoom" "(magnifier)" \
  "ALT SHIFT" "Switch Keyboard Layout" "" \
  "" "" "" \
  "───────────────" "SYSTEM" "───────────────" \
  " CTRL L" "Lock Screen" "(hyprlock)" \
  " CTRL P" "Power Menu" "(wlogout)" \
  "CTRL ALT Delete" "Exit Hyprland" "(immediate)" \
  " H" "This Help Menu" "" \
  "" "" ""
