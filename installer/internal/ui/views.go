package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"dotfiles-installer/internal/stages"
	"dotfiles-installer/internal/state"
)

// RenderMenu renders the main menu
func RenderMenu(options []state.MenuOption, cursor int, width, height int) string {
	var sb strings.Builder

	// Title
	title := RenderLogo()
	subtitle := RenderSubtitle()

	sb.WriteString(CenterText(title, width))
	sb.WriteString("\n")
	sb.WriteString(CenterText(subtitle, width))
	sb.WriteString("\n\n")

	// Menu box
	menuWidth := 60

	// Build menu items
	var items []string
	for i, opt := range options {
		icon := opt.Icon()
		label := opt.String()

		var itemStyle lipgloss.Style
		prefix := "  "

		if i == cursor {
			itemStyle = MenuSelectedStyle
			prefix = MenuCursorStyle.Render("▶ ")
		} else {
			itemStyle = MenuItemStyle
		}

		// Add number hint
		numHint := HelpStyle.Render(fmt.Sprintf("[%d] ", i+1))
		line := prefix + numHint + icon + " " + itemStyle.Render(label)
		items = append(items, line)

		// Add description for selected item
		if i == cursor {
			descStyle := lipgloss.NewStyle().
				Foreground(ColorLightGray).
				PaddingLeft(6).
				Width(menuWidth - 6)
			desc := descStyle.Render(opt.Description())
			items = append(items, desc)
		}
	}

	menuContent := strings.Join(items, "\n")

	menuBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ColorNeonPink).
		Padding(1, 2).
		Width(menuWidth)

	sb.WriteString(CenterText(menuBox.Render(menuContent), width))
	sb.WriteString("\n\n")

	// Help hint
	sb.WriteString(CenterText(RenderHelpHint(), width))

	return sb.String()
}

// RenderChecklist renders the installation progress
func RenderChecklist(stageList *stages.StageList, verbose bool, width, height int, spinnerFrame int) string {
	var sb strings.Builder

	// Header
	header := lipgloss.NewStyle().
		Bold(true).
		Foreground(ColorNeonCyan).
		Padding(0, 1).
		Render("INSTALLATION PROGRESS")

	progress := fmt.Sprintf("%d / %d", stageList.CompletedCount(), stageList.EnabledCount())
	progressStyle := lipgloss.NewStyle().Foreground(ColorNeonPink)

	headerLine := lipgloss.JoinHorizontal(lipgloss.Center,
		header,
		"  ",
		progressStyle.Render(progress),
	)

	sb.WriteString(CenterText(headerLine, width))
	sb.WriteString("\n\n")

	// Stages
	for i, stage := range stageList.Stages {
		if !stage.Enabled {
			continue
		}

		icon := getStatusIcon(stage.Status, spinnerFrame)
		name := stage.Name

		var lineStyle lipgloss.Style
		switch stage.Status {
		case stages.StatusComplete:
			lineStyle = StatusComplete
		case stages.StatusRunning:
			lineStyle = StatusRunning
		case stages.StatusFailed:
			lineStyle = StatusFailed
		case stages.StatusSkipped:
			lineStyle = StatusSkipped
		default:
			lineStyle = StatusPending
		}

		line := fmt.Sprintf("  %s %s", icon, lineStyle.Render(name))
		sb.WriteString(line)
		sb.WriteString("\n")

		// Progress bar for running stage
		if stage.Status == stages.StatusRunning && stage.TotalItems > 0 {
			progressPct := float64(stage.ProcessedItems) / float64(stage.TotalItems)
			bar := RenderProgressBar(progressPct, 40)
			count := fmt.Sprintf(" %d/%d", stage.ProcessedItems, stage.TotalItems)
			sb.WriteString("      " + bar + HelpStyle.Render(count))
			sb.WriteString("\n")

			if stage.CurrentItem != "" {
				sb.WriteString("      " + HelpStyle.Render("→ "+stage.CurrentItem))
				sb.WriteString("\n")
			}
		}

		// Error message
		if stage.Status == stages.StatusFailed && stage.Error != nil {
			errMsg := ErrorStyle.Render("      Error: " + stage.Error.Error())
			sb.WriteString(errMsg)
			sb.WriteString("\n")

			if stage.CanRetry {
				hint := HintStyle.Render("      Press 'r' to retry or 's' to skip")
				sb.WriteString(hint)
				sb.WriteString("\n")
			}
		}

		// Verbose output
		if verbose && i == stageList.CurrentIndex && len(stage.Output) > 0 {
			sb.WriteString("\n")
			output := stage.LastOutput(5)
			outputBox := VerboseBoxStyle.Width(width - 10).Render(strings.Join(output, "\n"))
			sb.WriteString("    " + outputBox)
			sb.WriteString("\n")
		}
	}

	sb.WriteString("\n")
	sb.WriteString(CenterText(RenderHelpHint(), width))

	return sb.String()
}

// RenderComplete renders the completion screen
func RenderComplete(success bool, stageList *stages.StageList, width, height int) string {
	var sb strings.Builder

	if success {
		sb.WriteString(CenterText(RenderSuccess(), width))
	} else {
		sb.WriteString(CenterText(RenderFailed(), width))
	}

	sb.WriteString("\n\n")

	// Summary
	summaryStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ColorNeonCyan).
		Padding(1, 2)

	var summaryLines []string
	summaryLines = append(summaryLines, TitleStyle.Render("INSTALLATION SUMMARY"))
	summaryLines = append(summaryLines, "")

	completed := 0
	skipped := 0
	failed := 0

	for _, stage := range stageList.Stages {
		if !stage.Enabled {
			continue
		}
		switch stage.Status {
		case stages.StatusComplete:
			completed++
		case stages.StatusSkipped:
			skipped++
		case stages.StatusFailed:
			failed++
		}
	}

	summaryLines = append(summaryLines, StatusComplete.Render(fmt.Sprintf("  ✓ Completed: %d", completed)))
	if skipped > 0 {
		summaryLines = append(summaryLines, StatusSkipped.Render(fmt.Sprintf("  ○ Skipped: %d", skipped)))
	}
	if failed > 0 {
		summaryLines = append(summaryLines, StatusFailed.Render(fmt.Sprintf("  ✗ Failed: %d", failed)))
	}

	summaryLines = append(summaryLines, "")

	if success {
		summaryLines = append(summaryLines, HelpStyle.Render("  A reboot is recommended to apply all changes."))
		summaryLines = append(summaryLines, "")
		summaryLines = append(summaryLines, HelpKeyStyle.Render("  r")+" - Reboot now")
		summaryLines = append(summaryLines, HelpKeyStyle.Render("  Enter")+" - Exit")
	} else {
		summaryLines = append(summaryLines, ErrorStyle.Render("  Some stages failed. Check the errors above."))
		summaryLines = append(summaryLines, "")
		summaryLines = append(summaryLines, HelpKeyStyle.Render("  Enter")+" - Exit")
	}

	summary := summaryStyle.Render(strings.Join(summaryLines, "\n"))
	sb.WriteString(CenterText(summary, width))

	return sb.String()
}

// RenderIntro renders the intro animation
func RenderIntro(matrix *MatrixRain, frame int, maxFrames int, width, height int) string {
	// Animated logo reveal
	logo := RenderLogoAnimated(frame, maxFrames)
	subtitle := ""

	// Show subtitle after logo is revealed
	if frame > maxFrames/2 {
		subtitle = RenderSubtitle()
	}

	overlay := logo
	if subtitle != "" {
		overlay = logo + "\n\n" + subtitle
	}

	// Render matrix with logo overlay
	if matrix != nil {
		return matrix.RenderWithOverlay(overlay, width, height)
	}

	return CenterText(overlay, width)
}

// Helper functions

func getStatusIcon(status stages.StageStatus, frame int) string {
	switch status {
	case stages.StatusComplete:
		return StatusComplete.Render(IconComplete)
	case stages.StatusRunning:
		// Animated spinner
		spinnerIdx := frame % len(CyberSpinner)
		return StatusRunning.Render(CyberSpinner[spinnerIdx])
	case stages.StatusFailed:
		return StatusFailed.Render(IconFailed)
	case stages.StatusSkipped:
		return StatusSkipped.Render(IconSkipped)
	default:
		return StatusPending.Render(IconPending)
	}
}

// RenderProgressBar renders a progress bar
func RenderProgressBar(percent float64, width int) string {
	filled := int(percent * float64(width))
	if filled > width {
		filled = width
	}
	empty := width - filled

	filledStr := ProgressBarFilled.Render(strings.Repeat("█", filled))
	emptyStr := ProgressBarEmpty.Render(strings.Repeat("░", empty))

	return filledStr + emptyStr
}
