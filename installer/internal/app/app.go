package app

import (
	"context"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"dotfiles-installer/internal/config"
	"dotfiles-installer/internal/stages"
	"dotfiles-installer/internal/state"
	"dotfiles-installer/internal/ui"
)

// Model is the main application model
type Model struct {
	// Configuration
	Config *config.Config
	Paths  *config.Paths

	// State machine
	CurrentState state.AppState

	// Menu state
	MenuCursor  int
	MenuOptions []state.MenuOption

	// Installation state
	StageList    *stages.StageList
	StageContext context.Context
	StageCancel  context.CancelFunc

	// UI state
	VerboseMode  bool
	ShowHelp     bool
	WindowWidth  int
	WindowHeight int

	// Animation state
	IntroFrame   int
	IntroMaxFrames int
	SpinnerFrame int
	MatrixRain   *ui.MatrixRain

	// Timing
	StartTime time.Time
}

// NewModel creates a new application model
func NewModel(cfg *config.Config) *Model {
	paths := config.NewPaths(cfg.DotfilesPath)

	return &Model{
		Config:         cfg,
		Paths:          paths,
		CurrentState:   state.StateIntro,
		MenuCursor:     0,
		MenuOptions:    state.AllMenuOptions(),
		VerboseMode:    cfg.Verbose,
		IntroFrame:     0,
		IntroMaxFrames: 60, // ~2 seconds at 30fps
		SpinnerFrame:   0,
		StartTime:      time.Now(),
	}
}

// Init initializes the model
func (m Model) Init() tea.Cmd {
	return tea.Batch(
		tea.EnterAltScreen,
		tickIntro(),
		tickMatrix(),
	)
}

// View renders the UI
func (m Model) View() string {
	// Get terminal size
	width := m.WindowWidth
	height := m.WindowHeight
	if width == 0 {
		width = 80
	}
	if height == 0 {
		height = 24
	}

	var content string

	switch m.CurrentState {
	case state.StateIntro:
		content = ui.RenderIntro(m.MatrixRain, m.IntroFrame, m.IntroMaxFrames, width, height)

	case state.StateMenu:
		content = ui.RenderMenu(m.MenuOptions, m.MenuCursor, width, height)

	case state.StateChecklist:
		content = ui.RenderChecklist(m.StageList, m.VerboseMode, width, height, m.SpinnerFrame)

	case state.StateComplete:
		success := !m.StageList.HasFailed()
		content = ui.RenderComplete(success, m.StageList, width, height)
	}

	// Overlay help if shown
	if m.ShowHelp {
		var bindings []ui.KeyBinding
		switch m.CurrentState {
		case state.StateMenu:
			bindings = ui.MenuBindings
		case state.StateChecklist:
			bindings = ui.ChecklistBindings
		case state.StateComplete:
			bindings = ui.CompleteBindings
		}
		content = ui.RenderHelpOverlay(width, height, bindings)
	}

	// Add indicators
	if m.Config.DryRun {
		content = ui.RenderDryRunIndicator() + "\n" + content
	}

	return content
}

// StartInstallation initializes and starts the installation process
func (m *Model) StartInstallation(option state.MenuOption) tea.Cmd {
	// Create stages based on selected option
	m.StageList = stages.CreateAllStages(m.Config, m.Paths)

	// Configure stages based on option
	switch option {
	case state.MenuMinimal:
		// Disable shell plugins and theming
		for _, s := range m.StageList.Stages {
			if s.ID == "zsh-plugins" || s.ID == "p10k" || s.ID == "wallust" {
				s.Enabled = false
			}
		}

	case state.MenuPackagesOnly:
		// Only package stages
		for _, s := range m.StageList.Stages {
			if s.ID != "prerequisites" && s.ID != "aur-helper" &&
				s.ID != "packages-official" && s.ID != "packages-aur" && s.ID != "cleanup" {
				s.Enabled = false
			}
		}

	case state.MenuDotfilesOnly:
		// Only dotfiles stages
		for _, s := range m.StageList.Stages {
			if s.ID != "prerequisites" && s.ID != "clone-dotfiles" &&
				s.ID != "stow" && s.ID != "cleanup" {
				s.Enabled = false
			}
		}

	case state.MenuUpgrade:
		// Upgrade: Pull changes and sync with stow
		for _, s := range m.StageList.Stages {
			if s.ID != "prerequisites" && s.ID != "clone-dotfiles" &&
				s.ID != "stow" && s.ID != "cleanup" {
				s.Enabled = false
			}
			// Force clone-dotfiles to run even if repo exists (it will do git pull)
			if s.ID == "clone-dotfiles" {
				s.IsCompleteFunc = nil
			}
		}
	}

	// Skip stages that are already complete
	for _, s := range m.StageList.Stages {
		if s.Enabled && s.IsComplete() {
			s.Status = stages.StatusSkipped
		}
	}

	// Find first enabled stage
	for m.StageList.CurrentIndex < len(m.StageList.Stages) {
		if m.StageList.Stages[m.StageList.CurrentIndex].Enabled &&
			m.StageList.Stages[m.StageList.CurrentIndex].Status != stages.StatusSkipped {
			break
		}
		m.StageList.CurrentIndex++
	}

	// Create context for cancellation
	m.StageContext, m.StageCancel = context.WithCancel(context.Background())
	m.CurrentState = state.StateChecklist

	// Start first stage
	return m.runCurrentStage()
}

// runCurrentStage runs the current stage
func (m *Model) runCurrentStage() tea.Cmd {
	stage := m.StageList.Current()
	if stage == nil {
		m.CurrentState = state.StateComplete
		return nil
	}

	stage.Status = stages.StatusRunning

	return tea.Batch(
		tickSpinner(),
		m.executeStage(stage),
	)
}

// executeStage creates a command to execute a stage
func (m *Model) executeStage(stage *stages.Stage) tea.Cmd {
	return func() tea.Msg {
		progress := make(chan stages.ProgressUpdate, 10)
		done := make(chan error, 1)

		go func() {
			err := stage.Action(m.StageContext, progress, m.Config.DryRun)
			done <- err
			close(progress)
		}()

		// Collect progress updates
		for update := range progress {
			stage.CurrentItem = update.ItemName
			if update.Message != "" {
				stage.AddOutput(update.Message)
			}
			if update.Total > 0 {
				stage.TotalItems = update.Total
				stage.ProcessedItems = update.Current
			}
		}

		err := <-done
		if err != nil {
			return StageErrorMsg{Index: m.StageList.CurrentIndex, Error: err}
		}
		return StageCompleteMsg{Index: m.StageList.CurrentIndex}
	}
}
