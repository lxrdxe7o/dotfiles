mkdir ($nu.data-dir | path join "vendor/autoload")
starship init nu | save -f ($nu.data-dir | path join "vendor/autoload/starship.nu")

# Pokemon colorscript with fastfetch
# pokemon-colorscripts --no-title -s -r | fastfetch -c $"($env.HOME)/.config/fastfetch/config-pokemon.jsonc" --logo-type file-raw --logo-height 10 --logo-width 5 --logo -
fastfetch -c ($env.HOME)/.config/fastfetch/config-compact.jsonc

# Aliases 
alias fetch = fastfetch
alias lg = lazygit

source ~/.zoxide.nu
source ~/.local/share/atuin/init.nu

$env.config.show_banner = false
