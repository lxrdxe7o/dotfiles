package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Main logo - cyberpunk style
const Logo = `
██████╗  ██████╗ ████████╗███████╗██╗██╗     ███████╗███████╗
██╔══██╗██╔═══██╗╚══██╔══╝██╔════╝██║██║     ██╔════╝██╔════╝
██║  ██║██║   ██║   ██║   █████╗  ██║██║     █████╗  ███████╗
██║  ██║██║   ██║   ██║   ██╔══╝  ██║██║     ██╔══╝  ╚════██║
██████╔╝╚██████╔╝   ██║   ██║     ██║███████╗███████╗███████║
╚═════╝  ╚═════╝    ╚═╝   ╚═╝     ╚═╝╚══════╝╚══════╝╚══════╝`

const LogoSmall = `
┳┓┏┓┏┳┓┏┓┳┓  ┏┓┏┓
┃┃┃┃ ┃ ┣ ┃┃  ┣ ┗┓
┻┛┗┛ ┻ ┻ ┻┗┛•┗┛┗┛`

const Subtitle = `
░▒▓█ ARCH LINUX + HYPRLAND DOTFILES █▓▒░`

const SubtitleBox = `
        ╔═══════════════════════════════╗
        ║   C Y B E R P U N K   2 0 7 7   ║
        ╚═══════════════════════════════╝`

const InstallerText = `
    ██╗███╗   ██╗███████╗████████╗ █████╗ ██╗     ██╗     ███████╗██████╗
    ██║████╗  ██║██╔════╝╚══██╔══╝██╔══██╗██║     ██║     ██╔════╝██╔══██╗
    ██║██╔██╗ ██║███████╗   ██║   ███████║██║     ██║     █████╗  ██████╔╝
    ██║██║╚██╗██║╚════██║   ██║   ██╔══██║██║     ██║     ██╔══╝  ██╔══██╗
    ██║██║ ╚████║███████║   ██║   ██║  ██║███████╗███████╗███████╗██║  ██║
    ╚═╝╚═╝  ╚═══╝╚══════╝   ╚═╝   ╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝╚═╝  ╚═╝`

const SuccessArt = `
    ╔═══════════════════════════════════════════╗
    ║                                           ║
    ║   ██████╗  ██████╗ ███╗   ██╗███████╗██╗  ║
    ║   ██╔══██╗██╔═══██╗████╗  ██║██╔════╝██║  ║
    ║   ██║  ██║██║   ██║██╔██╗ ██║█████╗  ██║  ║
    ║   ██║  ██║██║   ██║██║╚██╗██║██╔══╝  ╚═╝  ║
    ║   ██████╔╝╚██████╔╝██║ ╚████║███████╗██╗  ║
    ║   ╚═════╝  ╚═════╝ ╚═╝  ╚═══╝╚══════╝╚═╝  ║
    ║                                           ║
    ╚═══════════════════════════════════════════╝`

const FailedArt = `
    ╔═══════════════════════════════════════╗
    ║                                       ║
    ║   ███████╗ █████╗ ██╗██╗     ███████╗ ║
    ║   ██╔════╝██╔══██╗██║██║     ██╔════╝ ║
    ║   █████╗  ███████║██║██║     █████╗   ║
    ║   ██╔══╝  ██╔══██║██║██║     ██╔══╝   ║
    ║   ██║     ██║  ██║██║███████╗███████╗ ║
    ║   ╚═╝     ╚═╝  ╚═╝╚═╝╚══════╝╚══════╝ ║
    ║                                       ║
    ╚═══════════════════════════════════════╝`

// RenderLogo renders the main logo with gradient colors
func RenderLogo() string {
	lines := strings.Split(Logo, "\n")
	var result []string

	for i, line := range lines {
		if len(line) == 0 {
			result = append(result, "")
			continue
		}
		// Gradient from pink to cyan across lines
		progress := float64(i) / float64(len(lines))
		var color lipgloss.Color
		if progress < 0.5 {
			color = ColorNeonPink
		} else {
			color = ColorNeonCyan
		}
		style := lipgloss.NewStyle().Foreground(color).Bold(true)
		result = append(result, style.Render(line))
	}

	return strings.Join(result, "\n")
}

// RenderLogoAnimated renders logo with animation frame (for intro)
func RenderLogoAnimated(frame int, maxFrames int) string {
	lines := strings.Split(Logo, "\n")
	var result []string

	// Calculate how many lines to show based on frame
	linesToShow := (len(lines) * frame) / maxFrames
	if linesToShow > len(lines) {
		linesToShow = len(lines)
	}

	for i, line := range lines {
		if i >= linesToShow {
			// Not yet revealed
			result = append(result, strings.Repeat(" ", len(line)))
			continue
		}
		if len(line) == 0 {
			result = append(result, "")
			continue
		}
		// Gradient effect
		progress := float64(i) / float64(len(lines))
		colorIdx := int(progress * float64(len(GradientPinkCyan)-1))
		style := lipgloss.NewStyle().Foreground(GradientPinkCyan[colorIdx]).Bold(true)
		result = append(result, style.Render(line))
	}

	return strings.Join(result, "\n")
}

// RenderSubtitle renders the subtitle with neon effect
func RenderSubtitle() string {
	return lipgloss.NewStyle().
		Foreground(ColorNeonPink).
		Bold(true).
		Render(Subtitle)
}

// RenderSubtitleBox renders the cyberpunk box
func RenderSubtitleBox() string {
	return lipgloss.NewStyle().
		Foreground(ColorNeonCyan).
		Render(SubtitleBox)
}

// RenderInstallerText renders the INSTALLER text
func RenderInstallerText() string {
	return lipgloss.NewStyle().
		Foreground(ColorNeonPurple).
		Render(InstallerText)
}

// RenderSuccess renders success ASCII art
func RenderSuccess() string {
	return lipgloss.NewStyle().
		Foreground(ColorSuccess).
		Bold(true).
		Render(SuccessArt)
}

// RenderFailed renders failed ASCII art
func RenderFailed() string {
	return lipgloss.NewStyle().
		Foreground(ColorError).
		Bold(true).
		Render(FailedArt)
}

// CenterText centers text within a given width
func CenterText(text string, width int) string {
	lines := strings.Split(text, "\n")
	var result []string
	for _, line := range lines {
		padding := (width - lipgloss.Width(line)) / 2
		if padding < 0 {
			padding = 0
		}
		result = append(result, strings.Repeat(" ", padding)+line)
	}
	return strings.Join(result, "\n")
}
