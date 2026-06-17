#
# ~/.bashrc
#

# If not running interactively, don't do anything
[[ $- != *i* ]] && return

alias ls='ls --color=auto'
alias grep='grep --color=auto'
PS1='[\u@\h \W]\$ '
. "$HOME/.cargo/env"
export PLAYWRIGHT_BROWSERS_PATH=0

. "$HOME/.local/bin/env"
export OBSIDIAN_URL="https://127.0.0.1:27124/"
export OBSIDIAN_API_KEY="a81364af1ffb843d1162c6eb3651656570deef0ff10948e2a3d738bcea79c476"

#THIS MUST BE AT THE END OF THE FILE FOR SDKMAN TO WORK!!!
export SDKMAN_DIR="$HOME/.sdkman"
[[ -s "$HOME/.sdkman/bin/sdkman-init.sh" ]] && source "$HOME/.sdkman/bin/sdkman-init.sh"


# Added by Antigravity CLI installer
export PATH="/home/lxrdxe7o/.local/bin:$PATH"
export OBSIDIAN_VAULT_PATH="/home/lxrdxe7o/Documents/Library Sanctus"
