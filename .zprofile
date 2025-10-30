# Auto starts Hyprland if on tty1 and no X session is running
if [[ -z $DISPLAY ]] && [[ ${XDG_VTNR:-0} -eq 1 ]]; then
    exec Hyprland
fi

# --- Auto-start tmux on TTY1 ---
if [[ -z $DISPLAY ]] && [[ $(tty) == /dev/tty1 ]]; then
  # If already inside tmux, do nothing
  if [[ -z $TMUX ]]; then
    # Reattach if a session exists, else create one
    tmux attach-session -t main 2>/dev/null || tmux new-session -s main
  fi
fi

# Added by Toolbox App
export PATH="$PATH:/home/xero/.local/share/JetBrains/Toolbox/scripts"
