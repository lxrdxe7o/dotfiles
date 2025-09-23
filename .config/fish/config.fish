# Fastfetch
#fastfetch -c {$HOME}/.config/fastfetch/config-compact.jsonc
pokemon-colorscripts --no-title -s -r | fastfetch -c {$HOME}/.config/fastfetch/config-compact.jsonc --logo-type file-raw --logo-height 6 --logo-width 5 --logo -

zoxide init fish | source
starship init fish | source
atuin init fish | source

# Aliases being used
alias fetch='fastfetch'
alias lg='lazygit'
alias pacup='sudo pacman -Syu'
alias yayup='yay -Syu'

# Set-up icons for files/directories in terminal using lsd
alias ls='lsd'
alias l='ls -l'
alias la='ls -a'
alias lla='ls -la'
alias lt='ls --tree'

# Set man pager to nvimpager
set -Ux MANPAGER "nvim +Man!"
