<div align="center">
  <img src="assets/header.svg" width="100%" />
</div>

# Xero's Dotfiles

Arch Linux + Hyprland development environment dotfiles. Optimized for stability, performance, and a clean developer workflow.

Based on [JaKooLit's Hyprland-Dots](https://github.com/JaKooLit/Hyprland-Dots) with custom modifications and fixes.

## Table of Contents

- [Overview](#overview)
- [Screenshots](#screenshots)
- [Requirements](#requirements)
- [Dependencies](#dependencies)
- [Installer TUI](#installer-tui)
- [Installation](#installation)
- [File Structure](#file-structure)
- [Keybindings](#keybindings)
- [Customization](#customization)
- [Troubleshooting](#troubleshooting)
- [Credits](#credits)

## Overview

This dotfiles repository includes configurations for:

| Component          | Application                          |
| ------------------ | ------------------------------------ |
| **Window Manager** | Hyprland                             |
| **Terminal**       | Kitty, Ghostty                       |
| **Shell**          | Zsh (with Oh-My-Zsh + Powerlevel10k) |
| **Bar**            | Waybar                               |
| **Launcher**       | Rofi (Wayland fork)                  |
| **Notifications**  | SwayNC                               |
| **File Manager**   | Thunar                               |
| **Wallpaper**      | swww + wallust (dynamic theming)     |
| **Lock Screen**    | Hyprlock                             |
| **Idle Daemon**    | Hypridle                             |
| **Multiplexer**    | Tmux                                 |
| **Editor**         | Neovim                               |

## Screenshots

<!-- Add your screenshots here -->

_Screenshots coming soon..._

## Requirements

- **OS**: Arch Linux (or Arch-based distro)
- **Display Server**: Wayland
- **GPU**: Intel/AMD recommended (NVIDIA requires additional configuration)
- **Hyprland**: v0.40+ (tested on latest stable)

## Dependencies

### Quick Install (Essentials Only)

```bash
# Core Hyprland ecosystem
sudo pacman -S hyprland hypridle hyprlock xdg-desktop-portal-hyprland aquamarine hyprcursor hyprgraphics hyprlang hyprutils hyprwayland-scanner hyprpolkitagent hyprsunset hyprtoolkit hyprwire hyprland-guiutils

# Wayland utilities
sudo pacman -S wl-clipboard cliphist grim slurp swappy swww xwayland-satellite

# Audio/Media
sudo pacman -S pipewire-alsa pipewire-pulse wireplumber pamixer playerctl

# Notifications & Bar
sudo pacman -S swaync waybar rofi

# System utilities
sudo pacman -S network-manager-applet blueman brightnessctl power-profiles-daemon

# File manager & Terminal
sudo pacman -S thunar kitty ghostty alacritty

# Theming
sudo pacman -S kvantum qt5ct qt6ct nwg-look nwg-displays wallust

# Fonts
sudo pacman -S ttf-jetbrains-mono-nerd ttf-meslo-nerd otf-font-awesome ttf-firacode-nerd

# Shell & CLI tools
sudo pacman -S zsh fish nushell zoxide atuin lsd lazygit tmux neovim starship ripgrep yazi
```

### AUR Packages (via yay/paru)

```bash
# Install AUR helper first
sudo pacman -S --needed git base-devel
git clone https://aur.archlinux.org/yay-bin.git && cd yay-bin && makepkg -si && cd .. && rm -rf yay-bin

# Essential AUR packages
yay -S swww wallust pokemon-colorscripts-git matugen-git ttf-material-symbols-variable-git aylurs-gtk-shell-git quickshell ghostty zen-browser-bin
```

### Complete AUR Package List

All AUR/foreign packages installed via yay/paru:

```bash
yay -S --needed \
  ani-cli \
  antigravity \
  aylurs-gtk-shell-git \
  eclipse-java-bin \
  fish-done \
  fresh-editor \
  gem \
  ghostmirror \
  ghostty \
  github-desktop-bin \
  gtk-engine-murrine \
  jetbrains-toolbox \
  libastal-apps-git \
  libastal-battery-git \
  libastal-bluetooth-git \
  libastal-cava-git \
  libastal-hyprland-git \
  libastal-mpris-git \
  libastal-network-git \
  libastal-notifd-git \
  libastal-powerprofiles-git \
  libastal-tray-git \
  libastal-wireplumber-git \
  linutil-git \
  matugen-git \
  mpvpaper \
  opencode-bin \
  pokemon-colorscripts-git \
  spotify \
  ttf-material-symbols-variable-git \
  ttf-ms-win11-auto \
  ttf-victor-mono \
  visual-studio-code-bin \
  visual-studio-code-insiders-bin \
  wallust \
  wazuh-agent \
  webapp-manager \
  wlrobs-hg \
  zen-browser-bin
```

<details>
<summary><b>AUR Packages by Category</b></summary>

| Category             | Packages                                                                                                                             |
| -------------------- | ------------------------------------------------------------------------------------------------------------------------------------ |
| **Shells & Widgets** | `aylurs-gtk-shell-git`, `libastal-*-git` (12 packages), `matugen-git`                                                                |
| **Terminals**        | `ghostty`                                                                                                                            |
| **Browsers**         | `zen-browser-bin`                                                                                                                    |
| **IDEs & Editors**   | `visual-studio-code-bin`, `visual-studio-code-insiders-bin`, `opencode-bin`, `eclipse-java-bin`, `jetbrains-toolbox`, `fresh-editor` |
| **Dev Tools**        | `github-desktop-bin`, `gem`                                                                                                          |
| **Media**            | `spotify`, `mpvpaper`, `wlrobs-hg`, `ani-cli`                                                                                        |
| **Theming**          | `wallust`, `gtk-engine-murrine`, `pokemon-colorscripts-git`                                                                          |
| **Fonts**            | `ttf-material-symbols-variable-git`, `ttf-ms-win11-auto`, `ttf-victor-mono`                                                          |
| **Utilities**        | `linutil-git`, `webapp-manager`, `ghostmirror`, `antigravity`, `fish-done`                                                           |
| **Security**         | `wazuh-agent`                                                                                                                        |

</details>

---

## Full Environment Replication

To create an **exact copy** of my development environment, use the package lists below.

### Method 1: One-Command Install (Recommended)

```bash
# Install yay first if not present
sudo pacman -S --needed git base-devel && git clone https://aur.archlinux.org/yay-bin.git && cd yay-bin && makepkg -si && cd .. && rm -rf yay-bin

# Install ALL official repo packages
sudo pacman -S --needed 7zip accountsservice adobe-source-code-pro-fonts adobe-source-han-sans-cn-fonts adobe-source-han-sans-jp-fonts adobe-source-han-sans-kr-fonts alacritty alsa-firmware alsa-plugins alsa-utils aquamarine asusctl atuin awesome-terminal-fonts base base-devel bash-completion bc bind blueman bluez bluez-hid2hci bluez-libs bluez-utils brightnessctl btop btrfs-assistant btrfs-progs calf cava chromium cliphist cmake cmatrix cpupower cryptsetup dart-sass dbeaver ddcutil device-mapper dhclient dialog diffutils discord distrobox dmidecode dmraid dnsmasq docker docker-compose dosfstools duf e2fsprogs easyeffects efibootmgr efitools ethtool evolution-data-server exfatprogs f2fs-tools ffmpegthumbnailer flatpak freecad freerdp gimp git git-lfs github-cli gjs glances glfw glib2-devel gnome-system-monitor go greetd greetd-tuigreet grim grub grub-hook gst-libav gst-plugin-pipewire gst-plugins-bad gst-plugins-ugly gvfs gvfs-mtp haveged hdparm htop hwdetect hwinfo hyprcursor hyprgraphics hypridle hyprland hyprland-guiutils hyprlang hyprlock hyprpolkitagent hyprsunset hyprtoolkit hyprutils hyprwayland-scanner hyprwire impala inetutils inkscape intel-media-driver intel-media-sdk intel-ucode iptables-nft iwd jdk17-openjdk jdk21-openjdk jfsutils kdeconnect ki18n5 kitty knotifications5 kvantum kwidgetsaddons5 kwindowsystem5 lazygit less lib32-libva-intel-driver lib32-mesa lib32-opencl-mesa lib32-vulkan-intel libdvdcss libgsf libopenraw librecad libreoffice-fresh libspng libva-utils libvirt libwnck3 libxcrypt-compat linux-cachyos linux-cachyos-headers linux-firmware localsend logrotate loupe lsb-release lsd lsp-plugins lsscsi lvm2 man-db man-pages mdadm meld mercurial mesa-utils meson micro mkinitcpio modemmanager mousepad mpv mpv-mpris mtools nano nano-syntax-highlighting ncdu neovim netctl network-manager-applet networkmanager networkmanager-openvpn nfs-utils nilfs-utils niri noctalia-shell noisetorch noto-color-emoji-fontconfig noto-fonts noto-fonts-cjk noto-fonts-emoji nss-mdns ntp nushell nvtop nwg-displays nwg-look obs-studio obs-vaapi obs-vkcapture obsidian octopi okular openbsd-netcat opencl-mesa opendesktop-fonts openssh os-prober otf-font-awesome pacman-contrib pamixer paru pavucontrol perl pipewire-alsa pipewire-pulse playerctl plocate plymouth podman podman-compose polkit-kde-agent poppler-glib power-profiles-daemon pv python python-defusedxml python-flask python-flask-cors python-mysql-connector python-packaging python-pip python-pipx python-pyquery python-requests qalculate-gtk qbittorrent qemu-desktop qemu-emulators-full qpdf qt5ct qt6ct quickshell rebuild-detector reflector ripgrep rofi rsync rtkit rust s-nail samba sg3_utils slurp smartmontools snapper sof-firmware spotify-player starship stow sudo swappy swaync swtpm swww sysfsutils telegram-desktop terminus-font texinfo thunar thunar-archive-plugin thunar-volman thunderbird tmux ttf-bitstream-vera ttf-dejavu ttf-droid ttf-fira-code ttf-firacode-nerd ttf-jetbrains-mono ttf-jetbrains-mono-nerd ttf-liberation ttf-meslo-nerd ttf-opensans tumbler typescript udftools ufw umockdev unrar unzip upower usb_modeswitch usbutils uwsm v4l2loopback-dkms vala vesktop vi vim virt-manager vlc vulkan-headers vulkan-intel waybar weechat wget which wine winetricks wireless-regdb wireplumber wlogout wlsunset wpa_supplicant xarchiver xdg-desktop-portal-gnome xdg-desktop-portal-gtk xdg-desktop-portal-hyprland xdg-user-dirs xf86-input-libinput xfsprogs xfwm4 xl2tpd xorg-server xorg-xdpyinfo xorg-xhost xorg-xinit xorg-xinput xorg-xkill xorg-xrandr xwayland-satellite yad yarn yazi yt-dlp zathura zathura-pdf-poppler zed zoxide

# Install ALL AUR packages
paru -S --needed ani-cli antigravity aylurs-gtk-shell-git eclipse-java-bin fish-done fresh-editor gem ghostmirror ghostty github-desktop-bin gtk-engine-murrine jetbrains-toolbox libastal-apps-git libastal-battery-git libastal-bluetooth-git libastal-cava-git libastal-hyprland-git libastal-mpris-git libastal-network-git libastal-notifd-git libastal-powerprofiles-git libastal-tray-git libastal-wireplumber-git linutil-git matugen-git mpvpaper opencode-bin pokemon-colorscripts-git spotify ttf-material-symbols-variable-git ttf-ms-win11-auto ttf-victor-mono visual-studio-code-bin visual-studio-code-insiders-bin wallust wazuh-agent webapp-manager wlrobs-hg zen-browser-bin
```

### Method 2: Using Package Lists

Save the package lists to files and install from them:

```bash
# Export current packages (for reference/backup)
pacman -Qqen > pkglist-official.txt   # Official repo packages
pacman -Qqem > pkglist-aur.txt        # AUR packages

# Install from package lists
sudo pacman -S --needed - < pkglist-official.txt
yay -S --needed - < pkglist-aur.txt
```

### Method 3: Clone This Repo's Package Lists

Package lists are included in this repository:

```bash
# Clone the repo
git clone https://github.com/xero/dotfiles.git ~/dotfiles
cd ~/dotfiles

# Install all packages
sudo pacman -S --needed - < pkglist-official.txt
yay -S --needed - < pkglist-aur.txt
```

### CachyOS-Specific Packages

This setup uses [CachyOS](https://cachyos.org/) packages. If on vanilla Arch, skip or substitute these:

```bash
# CachyOS packages (skip on vanilla Arch)
cachy-browser cachyos-fish-config cachyos-hello cachyos-hooks cachyos-kernel-manager cachyos-keyring cachyos-micro-settings cachyos-mirrorlist cachyos-packageinstaller cachyos-plymouth-theme cachyos-rate-mirrors cachyos-settings cachyos-v3-mirrorlist cachyos-v4-mirrorlist cachyos-wallpapers cachyos-zsh-config linux-cachyos linux-cachyos-headers
```

### Packages by Category

<details>
<summary><b>Core System</b></summary>

```
base base-devel linux-cachyos linux-cachyos-headers linux-firmware intel-ucode
mkinitcpio grub grub-hook efibootmgr efitools cryptsetup lvm2 btrfs-progs
device-mapper dosfstools e2fsprogs exfatprogs f2fs-tools jfsutils xfsprogs
nilfs-utils udftools mtools snapper timeshift btrfs-assistant
```

</details>

<details>
<summary><b>Hyprland & Wayland</b></summary>

```
hyprland aquamarine hyprcursor hyprgraphics hypridle hyprland-guiutils hyprlang
hyprlock hyprpolkitagent hyprsunset hyprtoolkit hyprutils hyprwayland-scanner
hyprwire niri xwayland-satellite uwsm xdg-desktop-portal-hyprland
xdg-desktop-portal-gtk xdg-desktop-portal-gnome cliphist grim slurp swappy
swww wlogout wlsunset
```

</details>

<details>
<summary><b>Desktop Environment</b></summary>

```
waybar rofi swaync thunar thunar-archive-plugin thunar-volman tumbler
ffmpegthumbnailer nwg-look nwg-displays kvantum qt5ct qt6ct polkit-kde-agent
gvfs gvfs-mtp xarchiver loupe okular mousepad
```

</details>

<details>
<summary><b>Terminals & Shells</b></summary>

```
kitty ghostty alacritty tmux zsh fish nushell starship atuin zoxide lsd
bash-completion nano nano-syntax-highlighting micro vim neovim
```

</details>

<details>
<summary><b>Development Tools</b></summary>

```
git git-lfs github-cli lazygit mercurial stow ripgrep plocate
cmake meson rust go python python-pip python-pipx typescript yarn
jdk17-openjdk jdk21-openjdk dbeaver docker docker-compose podman podman-compose
virt-manager libvirt qemu-desktop qemu-emulators-full swtpm distrobox
visual-studio-code-bin visual-studio-code-insiders-bin zed jetbrains-toolbox
opencode-bin eclipse-java-bin
```

</details>

<details>
<summary><b>Audio & Media</b></summary>

```
pipewire-alsa pipewire-pulse wireplumber pamixer playerctl pavucontrol
easyeffects calf lsp-plugins noisetorch mpv mpv-mpris mpvpaper vlc
obs-studio obs-vaapi obs-vkcapture v4l2loopback-dkms wlrobs-hg
cava spotify spotify-player yt-dlp
```

</details>

<details>
<summary><b>Graphics & Design</b></summary>

```
gimp inkscape freecad librecad libreoffice-fresh loupe okular zathura
zathura-pdf-poppler qpdf
```

</details>

<details>
<summary><b>Networking & Communication</b></summary>

```
networkmanager networkmanager-openvpn network-manager-applet blueman bluez
bluez-utils bluez-libs bluez-hid2hci kdeconnect localsend telegram-desktop
discord vesktop thunderbird weechat openssh openbsd-netcat wget iwd
wpa_supplicant modemmanager dhclient dnsmasq samba nfs-utils
```

</details>

<details>
<summary><b>Browsers</b></summary>

```
cachy-browser chromium zen-browser-bin
```

</details>

<details>
<summary><b>Fonts</b></summary>

```
ttf-jetbrains-mono ttf-jetbrains-mono-nerd ttf-meslo-nerd ttf-firacode-nerd
ttf-fira-code ttf-victor-mono otf-font-awesome ttf-material-symbols-variable-git
ttf-bitstream-vera ttf-dejavu ttf-droid ttf-liberation ttf-opensans
ttf-ms-win11-auto adobe-source-code-pro-fonts adobe-source-han-sans-cn-fonts
adobe-source-han-sans-jp-fonts adobe-source-han-sans-kr-fonts opendesktop-fonts
noto-fonts noto-fonts-cjk noto-fonts-emoji noto-color-emoji-fontconfig
awesome-terminal-fonts terminus-font
```

</details>

<details>
<summary><b>System Utilities</b></summary>

```
btop htop nvtop glances ncdu duf brightnessctl ddcutil cpupower
power-profiles-daemon upower ufw iptables-nft reflector pacman-contrib
rebuild-detector paru yay-bin octopi flatpak 7zip unrar unzip rsync pv
smartmontools hdparm lsscsi sg3_utils hwinfo hwdetect dmidecode lsb-release
```

</details>

<details>
<summary><b>Theming & Customization</b></summary>

```
wallust matugen-git aylurs-gtk-shell-git quickshell noctalia-shell
libastal-apps-git libastal-battery-git libastal-bluetooth-git libastal-cava-git
libastal-hyprland-git libastal-mpris-git libastal-network-git libastal-notifd-git
libastal-powerprofiles-git libastal-tray-git libastal-wireplumber-git
gtk-engine-murrine pokemon-colorscripts-git
```

</details>

<details>
<summary><b>ASUS ROG Laptop (Optional)</b></summary>

```
asusctl supergfxctl rog-control-center
```

</details>

<details>
<summary><b>Gaming & Wine</b></summary>

```
wine winetricks lib32-mesa lib32-vulkan-intel lib32-opencl-mesa
lib32-libva-intel-driver vulkan-intel vulkan-headers opencl-mesa
```

</details>

## Installer TUI

An interactive terminal installer is included to automate the full setup. It handles package installation, dotfile deployment, shell configuration, and theming in a single guided process.

### Quick Start

```bash
# Clone and run
git clone https://github.com/xero/dotfiles.git ~/dotfiles
cd ~/dotfiles/installer
make build
./dotfiles-installer
```

Or run directly with Go:

```bash
cd ~/dotfiles/installer
go run ./cmd/installer/
```

### Installation Modes

The TUI presents a menu with the following options:

| Mode              | Description                                                   |
| ----------------- | ------------------------------------------------------------- |
| **Full Install**  | Everything: packages, dotfiles, shell setup, theming          |
| **Minimal**       | Core packages + dotfiles only, skip shell plugins and theming |
| **Custom**        | Interactively select which stages to run                      |
| **Packages Only** | Install pacman and AUR packages from the package lists        |
| **Dotfiles Only** | Clone the repo and run `stow .` to create symlinks            |
| **Upgrade**       | Pull latest changes from the repo and re-run stow             |

### What It Does

The installer runs through up to 14 stages:

1. **Check Prerequisites** - Verify `git` and `base-devel` are available
2. **Install AUR Helper** - Install `yay` if no AUR helper is present
3. **Install Official Packages** - Install from `pkglist-official.txt` via pacman
4. **Install AUR Packages** - Install from `pkglist-aur.txt` via yay/paru
5. **Backup Configs** - Copy existing `~/.config` dirs to `~/.config-backup`
6. **Clone Dotfiles** - Clone the repo (or pull latest if already cloned)
7. **Install Oh-My-Zsh** - Install the OMZ framework
8. **Install Zsh Plugins** - F-Sy-H, autosuggestions, autocomplete, you-should-use
9. **Install Powerlevel10k** - Install the p10k Zsh theme
10. **Create Symlinks** - Run GNU Stow to link configs into place
11. **Install TPM** - Install Tmux Plugin Manager
12. **Set Default Shell** - Set Zsh as the login shell via `chsh`
13. **Apply Theme** - Run wallust to generate color scheme from wallpaper
14. **Cleanup** - Final summary

Stages that are already complete (e.g. packages already installed, OMZ already present) are automatically skipped.

### Keybindings

| Key               | Context   | Action                |
| ----------------- | --------- | --------------------- |
| `j`/`k` or arrows | Menu      | Navigate options      |
| `1`-`6`           | Menu      | Quick select option   |
| `Enter`           | Menu      | Confirm selection     |
| `r`               | Checklist | Retry failed stage    |
| `s`               | Checklist | Skip current stage    |
| `v`               | Any       | Toggle verbose output |
| `?`               | Any       | Toggle help overlay   |
| `r`               | Complete  | Reboot                |
| `q` / `Ctrl+C`    | Any       | Quit                  |

### CLI Flags

```
-v, -verbose       Show detailed command output
-dry-run           Simulate installation without making changes
-skip-backup       Skip backup of existing configs
-path <dir>        Path for dotfiles repository (default: ~/dotfiles)
-repo <url>        Dotfiles repository URL
-version           Show version
```

### Building

```bash
cd ~/dotfiles/installer
make build          # Compile binary
make install        # Build + install to /usr/local/bin/
make dev            # Run with --dry-run flag for testing
make clean          # Remove binary
```

Requires Go 1.22+.

---

## Installation

The installer TUI above is the recommended way to set up. For manual installation, follow the steps below.

### 1. Clone the Repository

```bash
git clone https://github.com/xero/dotfiles.git ~/dotfiles
cd ~/dotfiles
```

### 2. Backup Existing Configs

```bash
# Backup existing configs
mkdir -p ~/.config-backup
cp -r ~/.config/hypr ~/.config-backup/ 2>/dev/null
cp -r ~/.config/waybar ~/.config-backup/ 2>/dev/null
cp -r ~/.config/rofi ~/.config-backup/ 2>/dev/null
cp -r ~/.config/kitty ~/.config-backup/ 2>/dev/null
cp ~/.zshrc ~/.config-backup/ 2>/dev/null
```

### 3. Create Symlinks

```bash
# Using GNU Stow (recommended)
sudo pacman -S stow
cd ~/dotfiles
stow .

# Or manually symlink
ln -sf ~/dotfiles/.config/* ~/.config/
ln -sf ~/dotfiles/.zshrc ~/.zshrc
ln -sf ~/dotfiles/.zprofile ~/.zprofile
```

### 4. Set Zsh as Default Shell

```bash
chsh -s $(which zsh)
```

### 5. Install Oh-My-Zsh & Plugins

```bash
# Oh-My-Zsh
sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"

# Powerlevel10k theme
git clone --depth=1 https://github.com/romkatv/powerlevel10k.git ${ZSH_CUSTOM:-$HOME/.oh-my-zsh/custom}/themes/powerlevel10k

# Plugins
git clone https://github.com/z-shell/F-Sy-H.git ${ZSH_CUSTOM:-$HOME/.oh-my-zsh/custom}/plugins/F-Sy-H
git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-$HOME/.oh-my-zsh/custom}/plugins/zsh-autosuggestions
git clone https://github.com/marlonrichert/zsh-autocomplete.git ${ZSH_CUSTOM:-$HOME/.oh-my-zsh/custom}/plugins/zsh-autocomplete
git clone https://github.com/MichaelAquilina/zsh-you-should-use.git ${ZSH_CUSTOM:-$HOME/.oh-my-zsh/custom}/plugins/you-should-use
```

### 6. Install Tmux Plugin Manager

```bash
git clone https://github.com/tmux-plugins/tpm ~/.tmux/plugins/tpm
# Then press prefix + I (Ctrl-a + I) inside tmux to install plugins
```

### 7. Reboot

```bash
reboot
```

## File Structure

```
dotfiles/
├── .config/
│   ├── hypr/                    # Hyprland configuration
│   │   ├── hyprland.conf        # Main config (sources others)
│   │   ├── hypridle.conf        # Idle daemon config
│   │   ├── hyprlock.conf        # Lock screen config
│   │   ├── configs/
│   │   │   └── Keybinds.conf    # Default keybindings
│   │   ├── UserConfigs/         # User customization (safe to edit)
│   │   │   ├── ENVariables.conf # Environment variables
│   │   │   ├── Startup_Apps.conf# Autostart applications
│   │   │   ├── UserKeybinds.conf# Custom keybindings
│   │   │   ├── UserSettings.conf# Hyprland settings
│   │   │   └── WindowRules.conf # Window rules & tags
│   │   ├── scripts/             # Helper scripts
│   │   └── UserScripts/         # User scripts
│   ├── waybar/                  # Status bar
│   ├── rofi/                    # Application launcher
│   ├── kitty/                   # Terminal emulator
│   ├── swaync/                  # Notification center
│   ├── tmux/                    # Terminal multiplexer
│   └── fastfetch/               # System info display
├── .zshrc                       # Zsh configuration
├── .zprofile                    # Zsh profile
├── installer/                   # Interactive TUI installer (Go + Bubble Tea)
│   ├── cmd/installer/main.go    # Entry point
│   ├── internal/                # App logic, stages, UI, theming
│   └── Makefile                 # Build targets
├── pkglist-official.txt         # Official repo packages
├── pkglist-aur.txt              # AUR packages
└── README.md                    # This file
```

## Keybindings

### General

| Keybinding             | Action                           |
| ---------------------- | -------------------------------- |
| `SUPER + Return`       | Open terminal (Kitty)            |
| `SUPER + ALT + Return` | Open terminal with tmux attached |
| `SUPER + Q`            | Close active window              |
| `SUPER + D`            | Application launcher (Rofi)      |
| `SUPER + E`            | File manager (Thunar)            |
| `SUPER + B`            | Open default browser             |
| `CTRL + ALT + Delete`  | Exit Hyprland                    |

### Window Management

| Keybinding                   | Action          |
| ---------------------------- | --------------- |
| `SUPER + Arrow Keys`         | Move focus      |
| `SUPER + SHIFT + Arrow Keys` | Resize window   |
| `SUPER + CTRL + Arrow Keys`  | Move window     |
| `SUPER + ALT + Arrow Keys`   | Swap window     |
| `SUPER + SPACE`              | Toggle floating |
| `SUPER + SHIFT + F`          | Fullscreen      |
| `SUPER + CTRL + F`           | Fake fullscreen |

### Workspaces

| Keybinding            | Action                        |
| --------------------- | ----------------------------- |
| `SUPER + 1-0`         | Switch to workspace 1-10      |
| `SUPER + SHIFT + 1-0` | Move window to workspace 1-10 |
| `SUPER + Tab`         | Next workspace                |
| `SUPER + SHIFT + Tab` | Previous workspace            |
| `SUPER + U`           | Toggle special workspace      |
| `SUPER + SHIFT + U`   | Move to special workspace     |

### Utilities

| Keybinding          | Action                        |
| ------------------- | ----------------------------- |
| `SUPER + SHIFT + S` | Screenshot (area) with Swappy |
| `SUPER + Print`     | Screenshot (full)             |
| `SUPER + ALT + V`   | Clipboard manager             |
| `SUPER + H`         | Keybindings cheat sheet       |
| `SUPER + CTRL + L`  | Lock screen                   |
| `SUPER + CTRL + P`  | Power menu (wlogout)          |
| `SUPER + W`         | Wallpaper selector            |
| `SUPER + A`         | Window overview (Quickshell)  |

### Media Keys

| Keybinding              | Action          |
| ----------------------- | --------------- |
| `XF86AudioRaiseVolume`  | Volume up       |
| `XF86AudioLowerVolume`  | Volume down     |
| `XF86AudioMute`         | Toggle mute     |
| `XF86AudioMicMute`      | Toggle mic mute |
| `XF86MonBrightnessUp`   | Brightness up   |
| `XF86MonBrightnessDown` | Brightness down |

## Customization

### Safe Files to Edit

These files are designed for user customization and won't be overwritten:

- `.config/hypr/UserConfigs/UserKeybinds.conf` - Add your keybindings
- `.config/hypr/UserConfigs/UserSettings.conf` - Hyprland settings
- `.config/hypr/UserConfigs/Startup_Apps.conf` - Autostart apps
- `.config/hypr/UserConfigs/WindowRules.conf` - Window rules
- `.config/hypr/UserConfigs/01-UserDefaults.conf` - Default apps

### Changing Wallpaper

```bash
# Interactive selector
SUPER + W

# Or manually with swww
swww img /path/to/wallpaper.jpg --transition-type grow
```

### Changing Themes

- **Rofi Theme**: `SUPER + CTRL + R`
- **Waybar Style**: `SUPER + CTRL + B`
- **Waybar Layout**: `SUPER + ALT + B`
- **GTK Theme**: Use `nwg-look`
- **Qt Theme**: Use `qt5ct` / `qt6ct` and `kvantummanager`

### Monitor Configuration

Use `nwg-displays` for GUI configuration, or edit:

- `.config/hypr/monitors.conf`
- `.config/hypr/workspaces.conf`

## Troubleshooting

### Common Issues

#### Screen tearing or flickering

```bash
# Try disabling VRR in UserSettings.conf
vrr = 0
```

#### Waybar not showing

```bash
# Restart waybar
killall waybar && waybar &
```

#### Rofi not launching

```bash
# Kill existing instance first
pkill rofi; rofi -show drun
```

#### Wallpaper not loading

```bash
# Check if swww-daemon is running
pgrep swww-daemon || swww-daemon &
swww img ~/Pictures/wallpapers/your-wallpaper.jpg
```

#### Audio not working

```bash
# Restart PipeWire
systemctl --user restart pipewire pipewire-pulse wireplumber
```

#### NVIDIA users

Add these to `.config/hypr/UserConfigs/ENVariables.conf`:

```conf
env = LIBVA_DRIVER_NAME,nvidia
env = __GLX_VENDOR_LIBRARY_NAME,nvidia
env = NVD_BACKEND,direct
```

### Reset to Fresh State

```bash
# Remove initial boot marker to re-run setup
rm ~/.config/hypr/.initial_startup_done
# Then logout and login again
```

### Logs

```bash
# Hyprland logs
cat /tmp/hypr/$(ls -t /tmp/hypr/ | head -1)/hyprland.log

# Check for errors
journalctl --user -xe
```

## Environment Variables

Key environment variables configured for development:

| Variable                       | Value                             | Purpose                  |
| ------------------------------ | --------------------------------- | ------------------------ |
| `GDK_BACKEND`                  | wayland,x11,\*                    | GTK apps prefer Wayland  |
| `QT_QPA_PLATFORM`              | wayland;xcb                       | Qt apps prefer Wayland   |
| `ELECTRON_OZONE_PLATFORM_HINT` | auto                              | Electron apps on Wayland |
| `MOZ_ENABLE_WAYLAND`           | 1                                 | Firefox on Wayland       |
| `SSH_AUTH_SOCK`                | $XDG_RUNTIME_DIR/ssh-agent.socket | SSH agent for Git        |
| `_JAVA_AWT_WM_NONREPARENTING`  | 1                                 | Java GUI apps fix        |
| `EDITOR`                       | nvim                              | Default editor           |

## Credits

- [JaKooLit](https://github.com/JaKooLit) - Original Hyprland-Dots base
- [Hyprland](https://hyprland.org/) - The amazing Wayland compositor
- [vaxerski](https://github.com/vaxerski) - Hyprland creator

## License

These dotfiles are provided as-is. Feel free to use, modify, and distribute.

---

**Happy ricing!**

<div align="center">
  <img src="assets/footer.svg" width="100%" />
</div>
