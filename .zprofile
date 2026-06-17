# --- Auto-start tmux on TTY1 ---
if [[ -z $DISPLAY ]] && [[ $(tty) == /dev/tty1 ]]; then
  # If already inside tmux, do nothing
  if [[ -z $TMUX ]]; then
    # Reattach if a session exists, else create one
    tmux attach-session -t main 2>/dev/null || tmux new-session -s main
  fi
fi

# Added by Toolbox App
# export PATH="$PATH:/home/xero/.local/share/JetBrains/Toolbox/scripts"


# Added by Antigravity CLI installer
export PATH="/home/lxrdxe7o/.local/bin:$PATH"
