#!/usr/bin/env bash
# tmux-help.sh — readable, colourised key-binding reference
# Pipe through "less -R" or "bat --color=always -p" for paging.
set -euo pipefail

# ── Gruvbox truecolor ───────────────────────────────────
# Uses \033[38;2;R;G;Bm (foreground) so these match your
# terminal's gruvbox theme exactly, not some 256-col approximation.
readonly F_ORANGE="\033[38;2;214;93;14m"    # #d65d0e
readonly F_GREEN="\033[38;2;152;151;26m"    # #98971a
readonly F_BLUE="\033[38;2;69;133;136m"     # #458588
readonly F_PURPLE="\033[38;2;177;98;134m"   # #b16286
readonly F_YELLOW="\033[38;2;215;153;33m"   # #d79921
readonly F_AQUA="\033[38;2;104;157;106m"    # #689d6a
readonly F_RED="\033[38;2;204;36;29m"       # #cc241d
readonly F_GREY="\033[38;2;168;153;132m"    # #a89984 (fg4)
readonly F_FG_DIM="\033[38;2;189;174;147m"  # #bdae93 (fg3, for dim text)
readonly C_RESET="\033[0m"
readonly F_BOLD="\033[1m"
readonly F_DIM="\033[2m"

# ── Strip bind-key prefix from a list-keys line ──────────
# Handles: bind-key -T prefix K cmd   and   bind-key -r -T prefix K cmd
clean_line() {
  printf "%s" "$1" | sed -E 's/^bind-key(\s+-[a-zA-Z])*\s+-T\s+\S+\s+//'
}
key_of()   { clean_line "$1" | sed -E 's/^(\S+)\s+.*/\1/'; }
cmd_of()   { clean_line "$1" | sed -E 's/^\S+\s+//'; }

# Human-readable descriptions ──────────────────────────────

desc_prefix() {
  local k="$1" cmd="$2"
  case "$k" in
    "Space") echo "Cycle through pane layouts" ;;
    "!")     echo "Break pane into its own window" ;;
    '"')     echo "Split pane horizontally (tmux default)" ;;
    "#")     echo "List paste buffers" ;;
    '$')     echo "Rename current session" ;;
    "&")     echo "Kill current window (with prompt)" ;;
    "'")     echo "Select window by index" ;;
    "(")     echo "Previous session" ;;
    ")")     echo "Next session" ;;
    ",")     echo "Rename current window" ;;
    "-")     echo "Split pane vertically (custom)" ;;
    ".")     echo "Move window to another session" ;;
    "/")     echo "Look up which key runs a command" ;;
    [0-9])   echo "Select window ${k}" ;;
    ":")     echo "Open tmux command prompt" ;;
    ";")     echo "Jump to last active pane" ;;
    "<")     echo "Window operations menu" ;;
    ">")     echo "Pane operations menu" ;;
    "=")     echo "Choose a paste buffer" ;;
    "?")     echo "This help screen" ;;
    "C")     echo "Customize mode (interactive key edit)" ;;
    "D")     echo "Choose a client to detach" ;;
    "E")     echo "Select a preset pane layout" ;;
    "I")     echo "Install TPM plugins" ;;
    "L")     echo "Switch to last session" ;;
    "M")     echo "Toggle pane mark" ;;
    "R")     echo "Rename session" ;;
    "U")     echo "Update TPM plugins" ;;
    "[")     echo "Enter copy mode" ;;
    "]")     echo "Paste from buffer" ;;
    "c")     echo "Create a new window" ;;
    "d")     echo "Detach from session" ;;
    "f")     echo "Find / search windows" ;;
    "h")     echo "Select pane  ←  (vim)" ;;
    "i")     echo "Show pane info" ;;
    "j")     echo "Select pane  ↓  (vim)" ;;
    "k")     echo "Select pane  ↑  (vim)" ;;
    "l")     echo "Select pane  →  (vim)" ;;
    "m")     echo "Toggle pane marking" ;;
    "n")     echo "Next window" ;;
    "o")     echo "Cycle through panes" ;;
    "p")     echo "Previous window" ;;
    "q")     echo "Show pane numbers" ;;
    "r")     echo "Reload tmux config" ;;
    "s")     echo "Choose session interactively" ;;
    "t")     echo "Display clock in pane" ;;
    "w")     echo "Choose window interactively" ;;
    "x")     echo "Kill current pane (with prompt)" ;;
    "z")     echo "Toggle pane zoom" ;;
    "{")     echo "Swap pane with previous" ;;
    "}")     echo "Swap pane with next" ;;
    "~")     echo "Show tmux message log" ;;
    "C-a")   echo "Attach to session by name" ;;
    "C-d")   echo "Detach client" ;;
    "C-k")   echo "Kill session (with prompt)" ;;
    "C-l")   echo "Choose session interactively" ;;
    "C-n")   echo "Create a new session" ;;
    "C-o")   echo "Rotate panes forward" ;;
    "C-r")   echo "Run default command in pane" ;;
    "C-s")   echo "Set mark in copy mode" ;;
    "C-z")   echo "Suspend tmux client" ;;
    "DC")    echo "Refresh client" ;;
    "Up"|"Down"|"Left"|"Right") echo "Select pane with arrow key" ;;
    "PPage") echo "Enter copy mode (page up first)" ;;
    "#"|"\\#")     echo "List paste buffers" ;;
    "$"|"\\$")     echo "Rename current session" ;;
    "'"|"\\'")     echo "Select window by index" ;;
    ";"|"\\;")     echo "Jump to last active pane" ;;
    "{"|"\\{")     echo "Swap pane with previous" ;;
    "}"|"\\}")     echo "Swap pane with next" ;;
    "~"|"\\~")     echo "Show tmux message log" ;;
    "|")     echo "Split pane horizontally (custom)" ;;
    "M-Up")    echo "Resize pane  ↑  5 lines" ;;
    "M-Down")  echo "Resize pane  ↓  5 lines" ;;
    "M-Left")  echo "Resize pane  ←  5 cols" ;;
    "M-Right") echo "Resize pane  →  5 cols" ;;
    "M-1"|"M-2"|"M-3"|"M-4"|"M-5") echo "Switch to ${k#M-}-pane even layout" ;;
    "M-6")   echo "Switch to main-horizontal layout" ;;
    "M-7")   echo "Switch to main-vertical layout" ;;
    "M-n")   echo "Next window (all sessions)" ;;
    "M-o")   echo "Reverse cycle through panes" ;;
    "M-p")   echo "Previous window (all sessions)" ;;
    "M-u")   echo "Respawn pane (no prompt)" ;;
    *)       echo "${cmd}" ;;
  esac
}

desc_copy_vi() {
  local k="$1" cmd="$2"
  case "$k" in
    "v")        echo "Begin selection (visual mode)" ;;
    "V")        echo "Select entire line" ;;
    "y")        echo "Copy selection & exit copy mode" ;;
    "Y")        echo "Copy entire line & exit" ;;
    "Enter")    echo "Copy selection & exit" ;;
    "Space")    echo "Start selection at cursor" ;;
    "Escape")   echo "Cancel selection / exit" ;;
    "C-v")      echo "Toggle rectangular (block) selection" ;;
    "C-u")      echo "Scroll  ½ page up" ;;
    "C-d")      echo "Scroll  ½ page down" ;;
    "C-b")      echo "Scroll  1 page up" ;;
    "C-f")      echo "Scroll  1 page down" ;;
    "C-y")      echo "Scroll down one line" ;;
    "C-e")      echo "Scroll up one line" ;;
    "h")        echo "Cursor  ←" ;;
    "j")        echo "Cursor  ↓" ;;
    "k")        echo "Cursor  ↑" ;;
    "l")        echo "Cursor  →" ;;
    "w")        echo "Jump  →  start of word" ;;
    "W")        echo "Jump  →  start of WORD (punct skip)" ;;
    "b")        echo "Jump  ←  start of word" ;;
    "B")        echo "Jump  ←  start of WORD (punct skip)" ;;
    "e")        echo "Jump  →  end of word" ;;
    "E")        echo "Jump  →  end of WORD (punct skip)" ;;
    "0")        echo "Jump to start of line" ;;
    "^")        echo "Jump to first non-whitespace" ;;
    '$')        echo "Jump to end of line" ;;
    "/")        echo "Search forward" ;;
    "?")        echo "Search backward" ;;
    "n")        echo "Repeat search forward" ;;
    "N")        echo "Repeat search backward" ;;
    "g")        echo "Go to top of pane history" ;;
    "G")        echo "Go to bottom of pane history" ;;
    "%")        echo "Jump to matching bracket" ;;
    "f")        echo "Jump  →  next char f" ;;
    "F")        echo "Jump  ←  char F" ;;
    "t")        echo "Jump to just before char t" ;;
    "T")        echo "Jump to just after char T (←)" ;;
    "o")        echo "Go to other end of selection" ;;
    "O")        echo "Swap selection anchor" ;;
    "q")        echo "Exit copy mode" ;;
    "C-c")      echo "Cancel & exit" ;;
    "C-g")      echo "Show cursor position" ;;
    "C-r")      echo "Toggle regex in search prompt" ;;
    "M")        echo "Jump to middle of visible pane" ;;
    "H")        echo "Jump to top of visible pane" ;;
    "L")        echo "Jump to bottom of visible pane" ;;
    "{")        echo "Jump  ←  blank-line paragraph" ;;
    "}")        echo "Jump  →  blank-line paragraph" ;;
    *)          echo "${cmd}" ;;
  esac
}

desc_search() {
  local k="$1" cmd="$2"
  case "$k" in
    "C-r")   echo "Toggle regex in search" ;;
    "C-p")   echo "Previous entry in prompt history" ;;
    "C-n")   echo "Next entry in prompt history" ;;
    "Up")    echo "Previous entry in prompt history" ;;
    "Down")  echo "Next entry in prompt history" ;;
    "Tab")   echo "Complete search from history" ;;
    "C-f")   echo "Forward to next match" ;;
    "C-b")   echo "Back to previous match" ;;
    "Enter") echo "Confirm search" ;;
    "Escape") echo "Cancel search" ;;
    *)       echo "${cmd}" ;;
  esac
}

desc_root() {
  local k="$1" cmd="$2"
  case "$k" in
    "M-Up")       echo "Resize pane  ↑" ;;
    "M-Down")     echo "Resize pane  ↓" ;;
    "M-Left")     echo "Resize pane  ←" ;;
    "M-Right")    echo "Resize pane  →" ;;
    "M-m")        echo "Toggle mouse mode" ;;
    "M-1"|"M-2"|"M-3"|"M-4"|"M-5"|"M-6"|"M-7"|"M-8"|"M-9")
                  echo "Switch to window ${k#M-}" ;;
    "M-0")        echo "Switch to window 10" ;;
    "M-PPage")    echo "Scroll up in pane" ;;
    "C-Left"|"C-Right"|"C-Up"|"C-Down")
                  echo "Resize pane (alternate)" ;;
    *)            echo "${cmd}" ;;
  esac
}

# ── Render utils ─────────────────────────────────────────

print_header() {
  printf "${F_ORANGE}${F_BOLD}\n"
  printf "  ╭──────────────────────────────────────────────────────────────────╮\n"
  printf "  │                        󰘞  TMUX KEYS                            │\n"
  printf "  ╰──────────────────────────────────────────────────────────────────╯\n"
  printf "${C_RESET}"
}

print_quickref() {
  printf "\n${F_YELLOW}${F_BOLD}  ⚡ quick reference${C_RESET}\n"
  printf "${F_DIM}${F_YELLOW}  ──────────────────────────────────────────────────────────────────${C_RESET}\n"
  printf "${F_YELLOW}${F_DIM}"
  printf "  ╔═══════════════════════╤═══════════════════════════╗\n"
  printf "${C_RESET}"
  printf "${F_YELLOW}  ║${C_RESET}  ${F_BOLD}C-a${C_RESET}       Prefix     ${F_YELLOW}║${C_RESET}  ${F_BOLD}C-a${C_RESET} ?       This help    ${F_YELLOW}║${C_RESET}\n"
  printf "${F_YELLOW}  ║${C_RESET}  ${F_BOLD}C-a${C_RESET} [     Copy mode  ${F_YELLOW}║${C_RESET}  ${F_BOLD}C-a${C_RESET} ]       Paste buffer ${F_YELLOW}║${C_RESET}\n"
  printf "${F_YELLOW}  ║${C_RESET}  ${F_BOLD}C-a${C_RESET} c     New window ${F_YELLOW}║${C_RESET}  ${F_BOLD}C-a${C_RESET} ,       Rename win   ${F_YELLOW}║${C_RESET}\n"
  printf "${F_YELLOW}  ║${C_RESET}  ${F_BOLD}C-a${C_RESET} d     Detach     ${F_YELLOW}║${C_RESET}  ${F_BOLD}C-a${C_RESET} r       Reload       ${F_YELLOW}║${C_RESET}\n"
  printf "${F_YELLOW}  ║${C_RESET}  ${F_BOLD}C-a${C_RESET} x     Kill pane  ${F_YELLOW}║${C_RESET}  ${F_BOLD}C-a${C_RESET} &       Kill window  ${F_YELLOW}║${C_RESET}\n"
  printf "${F_YELLOW}  ║${C_RESET}  ${F_BOLD}C-a${C_RESET} |     Split →    ${F_YELLOW}║${C_RESET}  ${F_BOLD}C-a${C_RESET} -       Split ↓      ${F_YELLOW}║${C_RESET}\n"
  printf "${F_YELLOW}  ║${C_RESET}  ${F_BOLD}C-a${C_RESET} h/j/k  Pane nav  ${F_YELLOW}║${C_RESET}  ${F_BOLD}C-a${C_RESET} M-Arrows  Resize     ${F_YELLOW}║${C_RESET}\n"
  printf "${F_YELLOW}  ║${C_RESET}  ${F_BOLD}C-a${C_RESET} n/p   Win nav    ${F_YELLOW}║${C_RESET}  ${F_BOLD}C-a${C_RESET} 0-9      Win select  ${F_YELLOW}║${C_RESET}\n"
  printf "${F_YELLOW}  ║${C_RESET}  ${F_BOLD}C-a${C_RESET} s     Sessions   ${F_YELLOW}║${C_RESET}  ${F_BOLD}C-a${C_RESET} w        Windows     ${F_YELLOW}║${C_RESET}\n"
  printf "${F_YELLOW}${F_DIM}"
  printf "  ╚═══════════════════════╧═══════════════════════════╝\n"
  printf "${C_RESET}"
}

print_section() {
  local title="$1" color="$2" table="$3" desc_fn="$4" filter_fn="${5:-}"

  local tmp
  tmp="$(mktemp)"

  tmux list-keys -T "$table" 2>/dev/null | while read -r line; do
    local key desc
    key="$(key_of "$line")"
    [ -z "$key" ] && continue

    # Skip mouse / GUI bindings (they're handled transparently by mouse=on)
    case "$key" in
      Mouse*|Wheel*|TripleClick*|DoubleClick*) continue ;;
    esac

    desc="$("$desc_fn" "$key" "$(cmd_of "$line")")"
    [ -z "$desc" ] && continue

    printf "%s\t%s\n" "$key" "$desc"
  done | sort -f -k1 > "$tmp"

  local count
  count="$(wc -l < "$tmp" | tr -d ' ')"

  printf "\n${color}${F_BOLD}  ══ ${title}  ${F_DIM}(${count} bindings)${C_RESET}\n"
  printf "${color}${F_DIM}  ──────────────────────────────────────────────────────────────────${C_RESET}\n"

  while IFS=$'\t' read -r k d; do
    printf "  ${color}%-12s${C_RESET} %s\n" "$k" "$d"
  done < "$tmp"

  rm -f "$tmp"
}

print_footer() {
  printf "\n${F_FG_DIM}${F_DIM}"
  printf "  ╭──────────────────────────────────────────────────────────────────╯\n"
  printf "  │  \033[1mless\033[22m navigation:  j/k scroll  ·  / search  ·  n/N next/prev  ·  q quit\n"
  printf "  │  try:  \033[3mC-a ?\033[23m  for this screen inside tmux\n"
  printf "  ╰\n"
  printf "${C_RESET}"
}

# ── Main ─────────────────────────────────────────────────

print_header
print_quickref

print_section \
  "PREFIX  (C-a then …)" \
  "$F_ORANGE" \
  "prefix" \
  desc_prefix

print_section \
  "COPY MODE — VI  (C-a [ then …)" \
  "$F_GREEN" \
  "copy-mode-vi" \
  desc_copy_vi

print_section \
  "SEARCH MODE  (/, ? inside copy mode)" \
  "$F_PURPLE" \
  "search" \
  desc_search

print_section \
  "ROOT KEYS  (global, no prefix)" \
  "$F_BLUE" \
  "root" \
  desc_root

print_footer
