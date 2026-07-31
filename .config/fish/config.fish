# Fastfetch
#fastfetch -c {$HOME}/.config/fastfetch/config-compact.jsonc
#pokemon-colorscripts --no-title -s -r | fastfetch -c {$HOME}/.config/fastfetch/config-compact.jsonc --logo-type file-raw --logo-height 6 --logo-width 5 --logo -

zoxide init fish | source
starship init fish | source
#atuin init fish | source
source {$HOME}/.config/fish/atuin.fish

# Aliases being used
alias fetch='fastfetch'
alias lg='lazygit'
alias pacup='sudo pacman -Syu'
alias yayup='yay -Syu'
alias parupg='paru -Syu'
alias ff='pokemon-colorscripts --no-title -s -n gengar | fastfetch -c {$HOME}/.config/fastfetch/config-compact.jsonc --logo-type file-raw --logo-height 6 --logo-width 5 --logo -'
alias reload='source ~/.config/fish/config.fish'
alias ta='tmux attach'
alias chrome-gravity='google-chrome --enable-features=UseOzonePlatform --ozone-platform=wayland --remote-debugging-port=9222 --user-data-dir=/tmp/antigravity-agent --remote-allow-origins="*"'


# Set-up icons for files/directories in terminal using lsd
alias ls='lsd'
alias l='ls -l'
alias la='ls -a'
alias lla='ls -la'
alias lt='ls --tree'

# Set man pager to nvimpager
set -Ux MANPAGER "nvim +Man!"

# Default editor
set -gx EDITOR nvim
set -gx VISUAL nvim

# opencode
fish_add_path /home/lxrdxe7o/.opencode/bin

# OpenClaw Completion
source "/home/lxrdxe7o/.openclaw/completions/openclaw.fish"

# bun
set --export BUN_INSTALL "$HOME/.bun"
set --export PATH $BUN_INSTALL/bin $PATH
set -gx OBSIDIAN_URL "https://127.0.0.1:27124/"
set -gx OBSIDIAN_API_KEY "a81364af1ffb843d1162c6eb3651656570deef0ff10948e2a3d738bcea79c476"


# Added by Antigravity CLI installer
set -gx PATH "/home/lxrdxe7o/.local/bin" $PATH

# Composio CLI
set --export COMPOSIO_INSTALL_DIR "/home/lxrdxe7o/.composio"
set --export PATH $COMPOSIO_INSTALL_DIR $PATH

# Kiro shell integration — source the static script directly.
# (Do NOT use `kiro --locate-shell-integration-path`; Kiro's CLI launches the
#  full IDE on invocation, which is what caused fish to open Kiro on startup.)
if string match -q -- "$TERM_PROGRAM" kiro
    set -l __kiro_integration /opt/kiro/resources/app/out/vs/workbench/contrib/terminal/common/scripts/shellIntegration.fish
    test -f "$__kiro_integration"; and . "$__kiro_integration"
end
