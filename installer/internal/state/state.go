package state

// AppState represents the current state of the application
type AppState int

const (
	StateIntro     AppState = iota // Animated ASCII intro with matrix rain
	StateMenu                      // Installation options menu
	StateChecklist                 // Multi-stage installation progress
	StateComplete                  // Success/reboot prompt
)

func (s AppState) String() string {
	switch s {
	case StateIntro:
		return "Intro"
	case StateMenu:
		return "Menu"
	case StateChecklist:
		return "Checklist"
	case StateComplete:
		return "Complete"
	default:
		return "Unknown"
	}
}

// MenuOption represents an installation option
type MenuOption int

const (
	MenuFullInstall MenuOption = iota
	MenuMinimal
	MenuCustom
	MenuPackagesOnly
	MenuDotfilesOnly
	MenuUpgrade
	MenuQuit
)

func (m MenuOption) String() string {
	switch m {
	case MenuFullInstall:
		return "Full Install"
	case MenuMinimal:
		return "Minimal (essentials only)"
	case MenuCustom:
		return "Custom (choose stages)"
	case MenuPackagesOnly:
		return "Packages Only"
	case MenuDotfilesOnly:
		return "Dotfiles Only (stow)"
	case MenuUpgrade:
		return "Upgrade (git pull + stow)"
	case MenuQuit:
		return "Quit"
	default:
		return "Unknown"
	}
}

func (m MenuOption) Description() string {
	switch m {
	case MenuFullInstall:
		return "Install everything: packages, dotfiles, shell setup, theming"
	case MenuMinimal:
		return "Core packages + dotfiles only, skip shell plugins and theming"
	case MenuCustom:
		return "Interactively select which stages to run"
	case MenuPackagesOnly:
		return "Install pacman and AUR packages only"
	case MenuDotfilesOnly:
		return "Clone repo and run stow only"
	case MenuUpgrade:
		return "Pull latest changes and re-run stow"
	case MenuQuit:
		return "Exit the installer"
	default:
		return ""
	}
}

func (m MenuOption) Icon() string {
	switch m {
	case MenuFullInstall:
		return ""
	case MenuMinimal:
		return ""
	case MenuCustom:
		return ""
	case MenuPackagesOnly:
		return ""
	case MenuDotfilesOnly:
		return ""
	case MenuUpgrade:
		return ""
	case MenuQuit:
		return ""
	default:
		return ""
	}
}

// AllMenuOptions returns all available menu options
func AllMenuOptions() []MenuOption {
	return []MenuOption{
		MenuFullInstall,
		MenuMinimal,
		MenuCustom,
		MenuPackagesOnly,
		MenuDotfilesOnly,
		MenuUpgrade,
		MenuQuit,
	}
}
