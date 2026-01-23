package app

import (
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"

	"dotfiles-installer/internal/stages"
	"dotfiles-installer/internal/state"
	"dotfiles-installer/internal/ui"
)

// Update handles messages and updates the model
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Window resize
	case tea.WindowSizeMsg:
		m.WindowWidth = msg.Width
		m.WindowHeight = msg.Height
		// Recreate matrix rain with new dimensions
		if m.CurrentState == state.StateIntro {
			m.MatrixRain = ui.NewMatrixRain(msg.Width, msg.Height)
		}
		return m, nil

	// Keyboard input
	case tea.KeyMsg:
		return m.handleKeyPress(msg)

	// Animation ticks
	case IntroTickMsg:
		if m.CurrentState != state.StateIntro {
			return m, nil
		}
		m.IntroFrame++
		if m.IntroFrame >= m.IntroMaxFrames {
			m.CurrentState = state.StateMenu
			return m, nil
		}
		return m, tickIntro()

	case MatrixTickMsg:
		if m.CurrentState == state.StateIntro && m.MatrixRain != nil {
			m.MatrixRain.Tick()
			return m, tickMatrix()
		}
		return m, nil

	case SpinnerTickMsg:
		if m.CurrentState == state.StateChecklist {
			m.SpinnerFrame++
			return m, tickSpinner()
		}
		return m, nil

	// Stage execution
	case StageCompleteMsg:
		if msg.Index < len(m.StageList.Stages) {
			m.StageList.Stages[msg.Index].Status = stages.StatusComplete
		}

		// Move to next stage
		if m.StageList.Next() {
			return m, m.runCurrentStage()
		}

		// All done
		m.CurrentState = state.StateComplete
		return m, nil

	case StageErrorMsg:
		if msg.Index < len(m.StageList.Stages) {
			m.StageList.Stages[msg.Index].Status = stages.StatusFailed
			m.StageList.Stages[msg.Index].Error = msg.Error
		}
		// Stay on current stage, user can retry or skip
		return m, nil
	}

	return m, nil
}

// handleKeyPress handles keyboard input
func (m Model) handleKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	// Help overlay toggle (works in all states)
	if msg.String() == "?" {
		m.ShowHelp = !m.ShowHelp
		return m, nil
	}

	// If help is showing, any key closes it
	if m.ShowHelp {
		m.ShowHelp = false
		return m, nil
	}

	// Global keys
	switch msg.String() {
	case "ctrl+c", "q":
		if m.CurrentState == state.StateChecklist && m.StageCancel != nil {
			m.StageCancel()
		}
		return m, tea.Quit

	case "v":
		m.VerboseMode = !m.VerboseMode
		return m, nil
	}

	// State-specific keys
	switch m.CurrentState {
	case state.StateIntro:
		// Any key skips intro
		m.CurrentState = state.StateMenu
		return m, nil

	case state.StateMenu:
		return m.handleMenuKeys(msg)

	case state.StateChecklist:
		return m.handleChecklistKeys(msg)

	case state.StateComplete:
		return m.handleCompleteKeys(msg)
	}

	return m, nil
}

// handleMenuKeys handles menu navigation
func (m Model) handleMenuKeys(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "j", "down":
		m.MenuCursor++
		if m.MenuCursor >= len(m.MenuOptions) {
			m.MenuCursor = 0
		}

	case "k", "up":
		m.MenuCursor--
		if m.MenuCursor < 0 {
			m.MenuCursor = len(m.MenuOptions) - 1
		}

	case "g":
		// gg to go to top (simplified: just g goes to top)
		m.MenuCursor = 0

	case "G":
		// G to go to bottom
		m.MenuCursor = len(m.MenuOptions) - 1

	case "enter":
		selected := m.MenuOptions[m.MenuCursor]
		if selected == state.MenuQuit {
			return m, tea.Quit
		}
		if selected == state.MenuCustom {
			// TODO: Implement custom stage selection
			// For now, treat as full install
			return m, m.StartInstallation(state.MenuFullInstall)
		}
		return m, m.StartInstallation(selected)

	case "1", "2", "3", "4", "5", "6":
		idx := int(msg.String()[0] - '1')
		if idx >= 0 && idx < len(m.MenuOptions) {
			m.MenuCursor = idx
			selected := m.MenuOptions[idx]
			if selected == state.MenuQuit {
				return m, tea.Quit
			}
			return m, m.StartInstallation(selected)
		}
	}

	return m, nil
}

// handleChecklistKeys handles checklist navigation
func (m Model) handleChecklistKeys(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	current := m.StageList.Current()
	if current == nil {
		return m, nil
	}

	switch msg.String() {
	case "r":
		// Retry failed stage
		if current.Status == stages.StatusFailed && current.CanRetry {
			current.Reset()
			return m, m.runCurrentStage()
		}

	case "s":
		// Skip current stage
		if current.CanSkip && (current.Status == stages.StatusFailed || current.Status == stages.StatusRunning) {
			if m.StageCancel != nil {
				m.StageCancel()
			}
			current.Status = stages.StatusSkipped

			if m.StageList.Next() {
				return m, m.runCurrentStage()
			}
			m.CurrentState = state.StateComplete
		}
	}

	return m, nil
}

// handleCompleteKeys handles completion screen
func (m Model) handleCompleteKeys(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "r":
		// Reboot
		if !m.Config.DryRun {
			exec.Command("sudo", "reboot").Start()
		}
		return m, tea.Quit

	case "enter":
		return m, tea.Quit
	}

	return m, nil
}
