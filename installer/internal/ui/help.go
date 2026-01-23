package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// KeyBinding represents a keyboard shortcut
type KeyBinding struct {
	Key  string
	Desc string
}

// GlobalBindings are available in all states
var GlobalBindings = []KeyBinding{
	{"q", "Quit"},
	{"?", "Toggle help"},
	{"v", "Toggle verbose"},
}

// MenuBindings are available in menu state
var MenuBindings = []KeyBinding{
	{"j/↓", "Move down"},
	{"k/↑", "Move up"},
	{"Enter", "Select"},
	{"1-6", "Quick select"},
}

// ChecklistBindings are available during installation
var ChecklistBindings = []KeyBinding{
	{"r", "Retry stage"},
	{"s", "Skip stage"},
	{"Ctrl+C", "Cancel"},
}

// CompleteBindings are available on completion screen
var CompleteBindings = []KeyBinding{
	{"r", "Reboot now"},
	{"Enter", "Exit"},
}

// RenderHelpOverlay renders the help popup
func RenderHelpOverlay(width, height int, stateBindings []KeyBinding) string {
	// Combine bindings
	allBindings := append(stateBindings, GlobalBindings...)

	// Calculate box dimensions
	maxKeyLen := 0
	maxDescLen := 0
	for _, b := range allBindings {
		if len(b.Key) > maxKeyLen {
			maxKeyLen = len(b.Key)
		}
		if len(b.Desc) > maxDescLen {
			maxDescLen = len(b.Desc)
		}
	}

	boxWidth := maxKeyLen + maxDescLen + 10
	if boxWidth < 30 {
		boxWidth = 30
	}

	// Title
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(ColorNeonCyan).
		Padding(0, 1)

	// Key style
	keyStyle := lipgloss.NewStyle().
		Foreground(ColorNeonPink).
		Bold(true).
		Width(maxKeyLen + 2).
		Align(lipgloss.Right)

	// Description style
	descStyle := lipgloss.NewStyle().
		Foreground(ColorWhite).
		PaddingLeft(2)

	// Build content
	var lines []string
	lines = append(lines, "")
	lines = append(lines, titleStyle.Render("KEYBINDINGS"))
	lines = append(lines, "")

	for _, b := range allBindings {
		line := keyStyle.Render(b.Key) + descStyle.Render(b.Desc)
		lines = append(lines, line)
	}

	lines = append(lines, "")
	lines = append(lines, HelpStyle.Render("Press ? to close"))
	lines = append(lines, "")

	content := strings.Join(lines, "\n")

	// Box style
	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		BorderForeground(ColorNeonPink).
		Background(ColorDarkBg).
		Padding(0, 2).
		Width(boxWidth)

	box := boxStyle.Render(content)

	// Center the box
	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, box)
}

// RenderHelpHint renders a small help hint at the bottom
func RenderHelpHint() string {
	return HelpStyle.Render("Press ? for help • v for verbose • q to quit")
}

// RenderVerboseIndicator shows verbose mode status
func RenderVerboseIndicator(enabled bool) string {
	if enabled {
		return StatusRunning.Render("[VERBOSE ON]")
	}
	return HelpStyle.Render("[verbose off]")
}

// RenderDryRunIndicator shows dry-run mode
func RenderDryRunIndicator() string {
	return WarningStyle.Render("[DRY RUN MODE]")
}
