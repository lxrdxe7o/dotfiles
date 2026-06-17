#!/bin/bash

# Gruvbox colors for the help screen
YELLOW='\033[1;33m'
BLUE='\033[1;34m'
GREEN='\033[1;32m'
PURPLE='\033[1;35m'
CYAN='\033[1;36m'
RED='\033[1;31m'
NC='\033[0m' # No Color

echo -e "${PURPLE}======================================================"
echo -e "       󰋖 TMUX CHEAT SHEET - KEYBINDINGS        "
echo -e "======================================================${NC}"
echo ""
echo -e "${YELLOW}PREFIX KEY:${NC}  ${BLUE}Ctrl + a${NC}"
echo ""
echo -e "${GREEN}--- PANES ---${NC}"
echo -e "${CYAN}|${NC}             Split pane horizontally"
echo -e "${CYAN}-${NC}             Split pane vertically"
echo -e "${CYAN}h, j, k, l${NC}    Select pane (Vim style)"
echo -e "${CYAN}Alt + Arrows${NC}  Resize pane"
echo -e "${CYAN}x${NC}             Kill current pane"
echo -e "${CYAN}z${NC}             Zoom toggle (Full screen pane)"
echo ""
echo -e "${GREEN}--- MOVE PANE ---${NC}"
echo -e "${CYAN}!${NC}             Break pane to a new window"
echo -e "${CYAN}M${NC}             Move pane to existing window (prompts for #)"
echo -e "${CYAN}S${NC}             Move pane to another session (prompts for session:window)"
echo ""
echo -e "${GREEN}--- WINDOWS ---${NC}"
echo -e "${CYAN}c${NC}             Create new window"
echo -e "${CYAN},${NC}             Rename current window"
echo -e "${CYAN}&${NC}             Kill current window"
echo -e "${CYAN}n / p${NC}         Next / Previous window"
echo -e "${CYAN}0-9${NC}           Switch to window by number"
echo ""
echo -e "${GREEN}--- SESSIONS ---${NC}"
echo -e "${CYAN}C-n${NC}           Create new session"
echo -e "${CYAN}C-l${NC}           List/Choose session"
echo -e "${CYAN}R${NC}             Rename current session"
echo -e "${CYAN}C-k${NC}           Kill current session"
echo -e "${CYAN}d${NC}             Detach from session"
echo ""
echo -e "${GREEN}--- MISC ---${NC}"
echo -e "${CYAN}r${NC}             Reload tmux.conf"
echo -e "${CYAN}?${NC}             Show this help screen"
echo -e "${CYAN}[${NC}             Enter Copy Mode (vi keys)"
echo -e "  - ${YELLOW}Space${NC}     Start selection"
echo -e "  - ${YELLOW}Enter${NC}     Copy selection"
echo -e "  - ${YELLOW}q${NC}         Exit copy mode"
echo ""
echo -e "${PURPLE}======================================================"
echo -e "    Press ${RED}q${NC} to close this help screen          "
echo -e "======================================================${NC}"

# Keep the window open until the user presses 'q'
read -n 1 -s key
if [[ $key == "q" ]]; then
    exit 0
fi
