package stages

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"dotfiles-installer/internal/config"
	"dotfiles-installer/internal/executor"
)

// CreateAllStages creates all installation stages
func CreateAllStages(cfg *config.Config, paths *config.Paths) *StageList {
	list := NewStageList()

	// 1. Prerequisites
	list.Add(createPrerequisitesStage())

	// 2. AUR Helper
	list.Add(createAURHelperStage())

	// 3. Official Packages
	list.Add(createOfficialPackagesStage(paths))

	// 4. AUR Packages
	list.Add(createAURPackagesStage(paths))

	// 5. Backup Configs
	if !cfg.SkipBackup {
		list.Add(createBackupStage(paths))
	}

	// 6. Clone Dotfiles
	list.Add(createCloneDotfilesStage(cfg, paths))

	// 7. GNU Stow
	list.Add(createStowStage(paths))

	// 8. Oh-My-Zsh
	list.Add(createOhMyZshStage(paths))

	// 9. Zsh Plugins
	list.Add(createZshPluginsStage(paths))

	// 10. Powerlevel10k
	list.Add(createPowerlevel10kStage(paths))

	// 11. TPM (Tmux Plugin Manager)
	list.Add(createTPMStage(paths))

	// 12. Default Shell
	list.Add(createDefaultShellStage())

	// 13. Wallust Theme
	list.Add(createWallustStage(paths))

	// 14. Cleanup
	list.Add(createCleanupStage())

	return list
}

// 1. Prerequisites
func createPrerequisitesStage() *Stage {
	stage := NewStage("prerequisites", "Check Prerequisites", "Verify git and base-devel are installed", "")
	stage.CanSkip = false

	stage.IsCompleteFunc = func() bool {
		return executor.CheckCommandExists("git") && executor.CheckCommandExists("makepkg")
	}

	stage.Action = func(ctx context.Context, progress chan<- ProgressUpdate, dryRun bool) error {
		progress <- ProgressUpdate{Message: "Checking for git...", Current: 1, Total: 2}
		if !executor.CheckCommandExists("git") {
			return fmt.Errorf("git not found - install with: sudo pacman -S git")
		}

		progress <- ProgressUpdate{Message: "Checking for base-devel...", Current: 2, Total: 2}
		if !executor.CheckCommandExists("makepkg") {
			return fmt.Errorf("base-devel not found - install with: sudo pacman -S base-devel")
		}

		return nil
	}

	return stage
}

// 2. AUR Helper
func createAURHelperStage() *Stage {
	stage := NewStage("aur-helper", "Install AUR Helper", "Install yay for AUR packages", "")
	stage.CanSkip = false

	stage.IsCompleteFunc = func() bool {
		return executor.CheckCommandExists("yay") || executor.CheckCommandExists("paru")
	}

	stage.Action = func(ctx context.Context, progress chan<- ProgressUpdate, dryRun bool) error {
		if executor.CheckCommandExists("yay") || executor.CheckCommandExists("paru") {
			progress <- ProgressUpdate{Message: "AUR helper already installed", Current: 1, Total: 1}
			return nil
		}

		if dryRun {
			progress <- ProgressUpdate{Message: "[DRY RUN] Would install yay-bin from AUR", Current: 1, Total: 1}
			return nil
		}

		progress <- ProgressUpdate{Message: "Cloning yay-bin...", Current: 1, Total: 3}

		tmpDir := "/tmp/yay-bin-install"
		os.RemoveAll(tmpDir)

		cmd := exec.CommandContext(ctx, "git", "clone", "https://aur.archlinux.org/yay-bin.git", tmpDir)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to clone yay-bin: %w", err)
		}

		progress <- ProgressUpdate{Message: "Building yay-bin...", Current: 2, Total: 3}

		cmd = exec.CommandContext(ctx, "makepkg", "-si", "--noconfirm")
		cmd.Dir = tmpDir
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to build yay-bin: %w", err)
		}

		progress <- ProgressUpdate{Message: "Cleaning up...", Current: 3, Total: 3}
		os.RemoveAll(tmpDir)

		return nil
	}

	return stage
}

// 3. Official Packages
func createOfficialPackagesStage(paths *config.Paths) *Stage {
	stage := NewStage("packages-official", "Install Official Packages", "Install packages from official repos", "")

	pkgListPath := filepath.Join(paths.Dotfiles, "pkglist-official.txt")

	stage.Action = func(ctx context.Context, progress chan<- ProgressUpdate, dryRun bool) error {
		packages, err := executor.ReadPackageList(pkgListPath)
		if err != nil {
			return fmt.Errorf("failed to read package list: %w", err)
		}

		stage.TotalItems = len(packages)
		progress <- ProgressUpdate{
			Message: fmt.Sprintf("Installing %d packages...", len(packages)),
			Total:   len(packages),
		}

		if dryRun {
			progress <- ProgressUpdate{
				Message: fmt.Sprintf("[DRY RUN] Would install %d packages", len(packages)),
				Current: len(packages),
				Total:   len(packages),
			}
			return nil
		}

		// Use pacman with --needed to skip already installed
		args := append([]string{"-S", "--needed", "--noconfirm"}, packages...)
		cmd := exec.CommandContext(ctx, "sudo", append([]string{"pacman"}, args...)...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("pacman failed: %w", err)
		}

		progress <- ProgressUpdate{
			Message: "All packages installed",
			Current: len(packages),
			Total:   len(packages),
		}

		return nil
	}

	return stage
}

// 4. AUR Packages
func createAURPackagesStage(paths *config.Paths) *Stage {
	stage := NewStage("packages-aur", "Install AUR Packages", "Install packages from AUR", "")

	pkgListPath := filepath.Join(paths.Dotfiles, "pkglist-aur.txt")

	stage.Action = func(ctx context.Context, progress chan<- ProgressUpdate, dryRun bool) error {
		packages, err := executor.ReadPackageList(pkgListPath)
		if err != nil {
			return fmt.Errorf("failed to read AUR package list: %w", err)
		}

		stage.TotalItems = len(packages)
		progress <- ProgressUpdate{
			Message: fmt.Sprintf("Installing %d AUR packages...", len(packages)),
			Total:   len(packages),
		}

		if dryRun {
			progress <- ProgressUpdate{
				Message: fmt.Sprintf("[DRY RUN] Would install %d AUR packages", len(packages)),
				Current: len(packages),
				Total:   len(packages),
			}
			return nil
		}

		// Use yay or paru
		helper := "yay"
		if executor.CheckCommandExists("paru") {
			helper = "paru"
		}

		args := append([]string{"-S", "--needed", "--noconfirm"}, packages...)
		cmd := exec.CommandContext(ctx, helper, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("%s failed: %w", helper, err)
		}

		return nil
	}

	return stage
}

// 5. Backup Configs
func createBackupStage(paths *config.Paths) *Stage {
	stage := NewStage("backup", "Backup Existing Configs", "Backup ~/.config to ~/.config-backup", "")

	stage.IsCompleteFunc = func() bool {
		return executor.DirExists(paths.ConfigBack)
	}

	stage.Action = func(ctx context.Context, progress chan<- ProgressUpdate, dryRun bool) error {
		if !executor.DirExists(paths.Config) {
			progress <- ProgressUpdate{Message: "No existing config to backup"}
			return nil
		}

		if dryRun {
			progress <- ProgressUpdate{Message: "[DRY RUN] Would backup ~/.config to ~/.config-backup"}
			return nil
		}

		progress <- ProgressUpdate{Message: "Creating backup..."}

		// Create backup directory
		if err := os.MkdirAll(paths.ConfigBack, 0755); err != nil {
			return fmt.Errorf("failed to create backup dir: %w", err)
		}

		// Copy important directories
		dirs := []string{"hypr", "waybar", "rofi", "kitty", "swaync", "tmux"}
		for _, dir := range dirs {
			src := filepath.Join(paths.Config, dir)
			if executor.DirExists(src) {
				dst := filepath.Join(paths.ConfigBack, dir)
				cmd := exec.CommandContext(ctx, "cp", "-r", src, dst)
				cmd.Run() // Ignore errors for non-existent dirs
			}
		}

		// Backup shell configs
		shellFiles := []string{".zshrc", ".bashrc", ".zprofile"}
		for _, f := range shellFiles {
			src := filepath.Join(paths.Home, f)
			if executor.FileExists(src) {
				dst := filepath.Join(paths.ConfigBack, f)
				exec.CommandContext(ctx, "cp", src, dst).Run()
			}
		}

		progress <- ProgressUpdate{Message: "Backup complete"}
		return nil
	}

	return stage
}

// 6. Clone Dotfiles
func createCloneDotfilesStage(cfg *config.Config, paths *config.Paths) *Stage {
	stage := NewStage("clone-dotfiles", "Clone Dotfiles", "Clone dotfiles repository", "")
	stage.CanSkip = false

	stage.IsCompleteFunc = func() bool {
		return executor.DirExists(filepath.Join(paths.Dotfiles, ".git"))
	}

	stage.Action = func(ctx context.Context, progress chan<- ProgressUpdate, dryRun bool) error {
		if executor.DirExists(filepath.Join(paths.Dotfiles, ".git")) {
			progress <- ProgressUpdate{Message: "Dotfiles already cloned, pulling latest..."}

			if dryRun {
				progress <- ProgressUpdate{Message: "[DRY RUN] Would pull latest changes"}
				return nil
			}

			cmd := exec.CommandContext(ctx, "git", "-C", paths.Dotfiles, "pull", "--rebase")
			if err := cmd.Run(); err != nil {
				// Not fatal, continue anyway
				progress <- ProgressUpdate{Message: "Pull failed, continuing with existing files"}
			}
			return nil
		}

		if dryRun {
			progress <- ProgressUpdate{Message: fmt.Sprintf("[DRY RUN] Would clone %s to %s", cfg.RepoURL, paths.Dotfiles)}
			return nil
		}

		progress <- ProgressUpdate{Message: "Cloning dotfiles repository..."}

		cmd := exec.CommandContext(ctx, "git", "clone", cfg.RepoURL, paths.Dotfiles)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to clone dotfiles: %w", err)
		}

		return nil
	}

	return stage
}

// 7. GNU Stow
func createStowStage(paths *config.Paths) *Stage {
	stage := NewStage("stow", "Create Symlinks", "Run GNU Stow to create symlinks", "")
	stage.CanSkip = false

	stage.Action = func(ctx context.Context, progress chan<- ProgressUpdate, dryRun bool) error {
		if !executor.CheckCommandExists("stow") {
			return fmt.Errorf("stow not found - install with: sudo pacman -S stow")
		}

		if dryRun {
			progress <- ProgressUpdate{Message: "[DRY RUN] Would run: stow . in " + paths.Dotfiles}
			return nil
		}

		progress <- ProgressUpdate{Message: "Running stow..."}

		cmd := exec.CommandContext(ctx, "stow", ".")
		cmd.Dir = paths.Dotfiles

		output, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("stow failed: %s\n%w", string(output), err)
		}

		progress <- ProgressUpdate{Message: "Symlinks created"}
		return nil
	}

	return stage
}

// 8. Oh-My-Zsh
func createOhMyZshStage(paths *config.Paths) *Stage {
	stage := NewStage("ohmyzsh", "Install Oh-My-Zsh", "Install Oh-My-Zsh framework", "")

	stage.IsCompleteFunc = func() bool {
		return executor.DirExists(paths.OhMyZsh)
	}

	stage.Action = func(ctx context.Context, progress chan<- ProgressUpdate, dryRun bool) error {
		if executor.DirExists(paths.OhMyZsh) {
			progress <- ProgressUpdate{Message: "Oh-My-Zsh already installed"}
			return nil
		}

		if dryRun {
			progress <- ProgressUpdate{Message: "[DRY RUN] Would install Oh-My-Zsh"}
			return nil
		}

		progress <- ProgressUpdate{Message: "Installing Oh-My-Zsh..."}

		// Download and run installer non-interactively
		script := `sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" "" --unattended`
		cmd := exec.CommandContext(ctx, "bash", "-c", script)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to install Oh-My-Zsh: %w", err)
		}

		return nil
	}

	return stage
}

// 9. Zsh Plugins
func createZshPluginsStage(paths *config.Paths) *Stage {
	stage := NewStage("zsh-plugins", "Install Zsh Plugins", "Install syntax highlighting, autosuggestions, etc.", "")

	plugins := []struct {
		name string
		repo string
		path string
	}{
		{"F-Sy-H", "https://github.com/z-shell/F-Sy-H.git", "plugins/F-Sy-H"},
		{"zsh-autosuggestions", "https://github.com/zsh-users/zsh-autosuggestions", "plugins/zsh-autosuggestions"},
		{"zsh-autocomplete", "https://github.com/marlonrichert/zsh-autocomplete.git", "plugins/zsh-autocomplete"},
		{"you-should-use", "https://github.com/MichaelAquilina/zsh-you-should-use.git", "plugins/you-should-use"},
	}

	stage.TotalItems = len(plugins)

	stage.Action = func(ctx context.Context, progress chan<- ProgressUpdate, dryRun bool) error {
		customDir := filepath.Join(paths.OhMyZsh, "custom")

		for i, plugin := range plugins {
			pluginPath := filepath.Join(customDir, plugin.path)

			progress <- ProgressUpdate{
				Message:  fmt.Sprintf("Installing %s...", plugin.name),
				Current:  i + 1,
				Total:    len(plugins),
				ItemName: plugin.name,
			}

			if executor.DirExists(pluginPath) {
				continue // Already installed
			}

			if dryRun {
				continue
			}

			cmd := exec.CommandContext(ctx, "git", "clone", plugin.repo, pluginPath)
			if err := cmd.Run(); err != nil {
				return fmt.Errorf("failed to install %s: %w", plugin.name, err)
			}
		}

		return nil
	}

	return stage
}

// 10. Powerlevel10k
func createPowerlevel10kStage(paths *config.Paths) *Stage {
	stage := NewStage("p10k", "Install Powerlevel10k", "Install Powerlevel10k theme", "")

	p10kPath := filepath.Join(paths.OhMyZsh, "custom", "themes", "powerlevel10k")

	stage.IsCompleteFunc = func() bool {
		return executor.DirExists(p10kPath)
	}

	stage.Action = func(ctx context.Context, progress chan<- ProgressUpdate, dryRun bool) error {
		if executor.DirExists(p10kPath) {
			progress <- ProgressUpdate{Message: "Powerlevel10k already installed"}
			return nil
		}

		if dryRun {
			progress <- ProgressUpdate{Message: "[DRY RUN] Would install Powerlevel10k"}
			return nil
		}

		progress <- ProgressUpdate{Message: "Installing Powerlevel10k..."}

		cmd := exec.CommandContext(ctx, "git", "clone", "--depth=1",
			"https://github.com/romkatv/powerlevel10k.git", p10kPath)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to install Powerlevel10k: %w", err)
		}

		return nil
	}

	return stage
}

// 11. TPM
func createTPMStage(paths *config.Paths) *Stage {
	stage := NewStage("tpm", "Install TPM", "Install Tmux Plugin Manager", "")

	stage.IsCompleteFunc = func() bool {
		return executor.DirExists(paths.TpmDir)
	}

	stage.Action = func(ctx context.Context, progress chan<- ProgressUpdate, dryRun bool) error {
		if executor.DirExists(paths.TpmDir) {
			progress <- ProgressUpdate{Message: "TPM already installed"}
			return nil
		}

		if dryRun {
			progress <- ProgressUpdate{Message: "[DRY RUN] Would install TPM"}
			return nil
		}

		progress <- ProgressUpdate{Message: "Installing TPM..."}

		cmd := exec.CommandContext(ctx, "git", "clone",
			"https://github.com/tmux-plugins/tpm", paths.TpmDir)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to install TPM: %w", err)
		}

		return nil
	}

	return stage
}

// 12. Default Shell
func createDefaultShellStage() *Stage {
	stage := NewStage("default-shell", "Set Default Shell", "Set Zsh as default shell", "")

	stage.IsCompleteFunc = func() bool {
		shell := os.Getenv("SHELL")
		return strings.HasSuffix(shell, "zsh")
	}

	stage.Action = func(ctx context.Context, progress chan<- ProgressUpdate, dryRun bool) error {
		shell := os.Getenv("SHELL")
		if strings.HasSuffix(shell, "zsh") {
			progress <- ProgressUpdate{Message: "Zsh is already the default shell"}
			return nil
		}

		if dryRun {
			progress <- ProgressUpdate{Message: "[DRY RUN] Would set Zsh as default shell"}
			return nil
		}

		progress <- ProgressUpdate{Message: "Setting Zsh as default shell..."}

		zshPath, err := exec.LookPath("zsh")
		if err != nil {
			return fmt.Errorf("zsh not found: %w", err)
		}

		cmd := exec.CommandContext(ctx, "chsh", "-s", zshPath)
		cmd.Stdin = os.Stdin
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to set default shell: %w", err)
		}

		return nil
	}

	return stage
}

// 13. Wallust Theme
func createWallustStage(paths *config.Paths) *Stage {
	stage := NewStage("wallust", "Apply Theme", "Run initial Wallust theming", "")

	stage.Action = func(ctx context.Context, progress chan<- ProgressUpdate, dryRun bool) error {
		if !executor.CheckCommandExists("wallust") {
			progress <- ProgressUpdate{Message: "Wallust not installed, skipping theming"}
			return nil
		}

		wallpaperPath := filepath.Join(paths.Config, "hypr", "wallpaper_effects", ".wallpaper_current")
		if !executor.FileExists(wallpaperPath) {
			progress <- ProgressUpdate{Message: "No wallpaper set, skipping theming"}
			return nil
		}

		if dryRun {
			progress <- ProgressUpdate{Message: "[DRY RUN] Would run wallust theming"}
			return nil
		}

		progress <- ProgressUpdate{Message: "Applying theme..."}

		cmd := exec.CommandContext(ctx, "wallust", "run", wallpaperPath)
		if err := cmd.Run(); err != nil {
			// Not fatal
			progress <- ProgressUpdate{Message: "Wallust theming failed, continuing..."}
		}

		return nil
	}

	return stage
}

// 14. Cleanup
func createCleanupStage() *Stage {
	stage := NewStage("cleanup", "Cleanup", "Final cleanup and summary", "")
	stage.CanSkip = false

	stage.Action = func(ctx context.Context, progress chan<- ProgressUpdate, dryRun bool) error {
		progress <- ProgressUpdate{Message: "Installation complete!"}

		if dryRun {
			progress <- ProgressUpdate{Message: "[DRY RUN] No changes were made"}
		}

		return nil
	}

	return stage
}
