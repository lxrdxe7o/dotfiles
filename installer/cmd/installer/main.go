package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"dotfiles-installer/internal/app"
	"dotfiles-installer/internal/config"
)

func main() {
	cfg := config.ParseFlags()

	if cfg.Version {
		fmt.Printf("dotfiles-installer v%s\n", config.Version)
		os.Exit(0)
	}

	// Create the application model
	model := app.NewModel(cfg)

	// Create and run the Bubble Tea program
	p := tea.NewProgram(
		model,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running installer: %v\n", err)
		os.Exit(1)
	}
}
