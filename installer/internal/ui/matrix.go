package ui

import (
	"math/rand"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// MatrixRain creates a matrix-style rain effect
type MatrixRain struct {
	Width   int
	Height  int
	Columns [][]rune
	Drops   []int
	Speeds  []int
	Chars   []rune
}

// NewMatrixRain creates a new matrix rain effect
func NewMatrixRain(width, height int) *MatrixRain {
	// Katakana + numbers + symbols for cyberpunk feel
	chars := []rune("アイウエオカキクケコサシスセソタチツテトナニヌネノハヒフヘホマミムメモヤユヨラリルレロワヲン0123456789ABCDEF@#$%&*<>[]{}|/\\")

	m := &MatrixRain{
		Width:   width,
		Height:  height,
		Columns: make([][]rune, width),
		Drops:   make([]int, width),
		Speeds:  make([]int, width),
		Chars:   chars,
	}

	for i := range m.Columns {
		m.Columns[i] = make([]rune, height)
		m.Drops[i] = rand.Intn(height)
		m.Speeds[i] = rand.Intn(2) + 1
		for j := range m.Columns[i] {
			m.Columns[i][j] = chars[rand.Intn(len(chars))]
		}
	}

	return m
}

// Tick advances the animation by one frame
func (m *MatrixRain) Tick() {
	for i := range m.Drops {
		m.Drops[i] += m.Speeds[i]
		if m.Drops[i] >= m.Height+10 {
			m.Drops[i] = -rand.Intn(m.Height/2) - 5
			m.Speeds[i] = rand.Intn(2) + 1
		}
		// Randomize some characters for flickering effect
		if rand.Float32() < 0.05 {
			j := rand.Intn(m.Height)
			m.Columns[i][j] = m.Chars[rand.Intn(len(m.Chars))]
		}
	}
}

// Render returns the matrix rain as a styled string
func (m *MatrixRain) Render() string {
	var sb strings.Builder

	// Green color shades from dark to bright
	greenShades := []lipgloss.Color{
		"#001100", "#002200", "#003300", "#004400",
		"#006600", "#008800", "#00AA00", "#00CC00", "#00FF00",
	}

	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			dist := m.Drops[x] - y
			if dist < 0 {
				dist = -dist + m.Height // Wrap around effect
			}

			var style lipgloss.Style
			if dist == 0 {
				// Leading character is bright white
				style = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Bold(true)
			} else if dist < len(greenShades) {
				// Trail fades out
				idx := len(greenShades) - 1 - dist
				if idx < 0 {
					idx = 0
				}
				style = lipgloss.NewStyle().Foreground(greenShades[idx])
			} else {
				// Very dark/invisible
				style = lipgloss.NewStyle().Foreground(lipgloss.Color("#000500"))
			}

			sb.WriteString(style.Render(string(m.Columns[x][y])))
		}
		if y < m.Height-1 {
			sb.WriteString("\n")
		}
	}

	return sb.String()
}

// RenderWithOverlay renders matrix rain with content overlaid on top
func (m *MatrixRain) RenderWithOverlay(overlay string, width, height int) string {
	rain := m.Render()
	rainLines := strings.Split(rain, "\n")

	overlayLines := strings.Split(overlay, "\n")

	// Calculate centering for overlay
	startY := (height - len(overlayLines)) / 2
	if startY < 0 {
		startY = 0
	}

	var result []string
	for i := 0; i < len(rainLines) && i < height; i++ {
		overlayIdx := i - startY
		if overlayIdx >= 0 && overlayIdx < len(overlayLines) {
			line := overlayLines[overlayIdx]
			// Center the overlay line
			padding := (width - lipgloss.Width(line)) / 2
			if padding < 0 {
				padding = 0
			}
			// Overlay replaces rain in the center
			result = append(result, strings.Repeat(" ", padding)+line)
		} else {
			result = append(result, rainLines[i])
		}
	}

	return strings.Join(result, "\n")
}

// SimpleMatrixLine generates a single line of matrix characters
func SimpleMatrixLine(width int, brightness float64) string {
	chars := []rune("アイウエオカキクケコ0123456789ABCDEF")
	var sb strings.Builder

	for i := 0; i < width; i++ {
		if rand.Float64() < 0.3 { // 30% density
			char := chars[rand.Intn(len(chars))]
			var color lipgloss.Color
			if brightness > 0.7 {
				color = "#00FF00"
			} else if brightness > 0.4 {
				color = "#00AA00"
			} else {
				color = "#004400"
			}
			style := lipgloss.NewStyle().Foreground(color)
			sb.WriteString(style.Render(string(char)))
		} else {
			sb.WriteString(" ")
		}
	}

	return sb.String()
}
