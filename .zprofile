# Auto starts Hyprland if on tty1 and no X session is running
if [[ -z $DISPLAY ]] && [[ ${XDG_VTNR:-0} -eq 1 ]]; then
    exec Hyprland
fi

# Added by Toolbox App
export PATH="$PATH:/home/xero/.local/share/JetBrains/Toolbox/scripts"

