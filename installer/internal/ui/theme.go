package ui

import "github.com/charmbracelet/lipgloss"

// Cyberpunk color palette
var (
	// Primary neon colors
	ColorNeonPink   = lipgloss.Color("#FF1493")
	ColorNeonCyan   = lipgloss.Color("#00FFFF")
	ColorNeonPurple = lipgloss.Color("#6A1B9A")
	ColorNeonGreen  = lipgloss.Color("#39FF14")
	ColorNeonOrange = lipgloss.Color("#FF6600")
	ColorNeonBlue   = lipgloss.Color("#00BFFF")

	// Background and text
	ColorDarkBg    = lipgloss.Color("#0D0D0D")
	ColorDarkGray  = lipgloss.Color("#1A1A2E")
	ColorMidGray   = lipgloss.Color("#2D2D44")
	ColorLightGray = lipgloss.Color("#4A4A6A")
	ColorWhite     = lipgloss.Color("#EAEAEA")
	ColorDimWhite  = lipgloss.Color("#888888")

	// Status colors
	ColorSuccess = lipgloss.Color("#00FF88")
	ColorWarning = lipgloss.Color("#FFD93D")
	ColorError   = lipgloss.Color("#FF3366")

	// Gradient stops for animations
	GradientPinkCyan = []lipgloss.Color{
		"#FF1493", "#FF00AA", "#CC00FF", "#9900FF", "#6600FF", "#00CCFF", "#00FFFF",
	}
	GradientGreenCyan = []lipgloss.Color{
		"#39FF14", "#00FF44", "#00FF88", "#00FFAA", "#00FFCC", "#00FFFF",
	}
)

// Style definitions
var (
	// Title and headers
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorNeonCyan).
			Padding(0, 1)

	SubtitleStyle = lipgloss.NewStyle().
			Foreground(ColorNeonPink).
			Italic(true)

	// Borders
	BorderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorNeonPink)

	DoubleBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.DoubleBorder()).
				BorderForeground(ColorNeonCyan)

	// Menu items
	MenuItemStyle = lipgloss.NewStyle().
			Foreground(ColorWhite).
			Padding(0, 2)

	MenuSelectedStyle = lipgloss.NewStyle().
				Foreground(ColorNeonCyan).
				Background(ColorMidGray).
				Bold(true).
				Padding(0, 2)

	MenuCursorStyle = lipgloss.NewStyle().
			Foreground(ColorNeonPink).
			Bold(true)

	// Progress and status
	ProgressBarFilled = lipgloss.NewStyle().
				Foreground(ColorNeonGreen)

	ProgressBarEmpty = lipgloss.NewStyle().
				Foreground(ColorMidGray)

	StatusPending = lipgloss.NewStyle().
			Foreground(ColorDimWhite)

	StatusRunning = lipgloss.NewStyle().
			Foreground(ColorNeonCyan).
			Bold(true)

	StatusComplete = lipgloss.NewStyle().
			Foreground(ColorSuccess)

	StatusFailed = lipgloss.NewStyle().
			Foreground(ColorError).
			Bold(true)

	StatusSkipped = lipgloss.NewStyle().
			Foreground(ColorLightGray).
			Strikethrough(true)

	// Help and hints
	HelpStyle = lipgloss.NewStyle().
			Foreground(ColorLightGray)

	HelpKeyStyle = lipgloss.NewStyle().
			Foreground(ColorNeonPink).
			Bold(true)

	HelpDescStyle = lipgloss.NewStyle().
			Foreground(ColorDimWhite)

	// Errors and warnings
	ErrorStyle = lipgloss.NewStyle().
			Foreground(ColorError).
			Bold(true)

	WarningStyle = lipgloss.NewStyle().
			Foreground(ColorWarning)

	HintStyle = lipgloss.NewStyle().
			Foreground(ColorNeonBlue).
			Italic(true)

	// Verbose output box
	VerboseBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorMidGray).
			Foreground(ColorDimWhite).
			Padding(0, 1)

	// Container styles
	CenterStyle = lipgloss.NewStyle().
			Align(lipgloss.Center)
)

// Spinner frames with cyberpunk aesthetic
var SpinnerFrames = []string{"▱▱▱▱▱", "▰▱▱▱▱", "▰▰▱▱▱", "▰▰▰▱▱", "▰▰▰▰▱", "▰▰▰▰▰", "▱▰▰▰▰", "▱▱▰▰▰", "▱▱▱▰▰", "▱▱▱▱▰"}

var DotSpinner = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

var CyberSpinner = []string{"◐", "◓", "◑", "◒"}

// Status icons
const (
	IconPending  = "○"
	IconRunning  = "◐"
	IconComplete = "✓"
	IconFailed   = "✗"
	IconSkipped  = "○"
	IconArrow    = "▶"
	IconDot      = "•"
)

// RenderGradientText applies a gradient effect to text
func RenderGradientText(text string, colors []lipgloss.Color) string {
	if len(colors) == 0 || len(text) == 0 {
		return text
	}

	result := ""
	for i, char := range text {
		colorIdx := (i * len(colors)) / len(text)
		if colorIdx >= len(colors) {
			colorIdx = len(colors) - 1
		}
		style := lipgloss.NewStyle().Foreground(colors[colorIdx])
		result += style.Render(string(char))
	}
	return result
}
