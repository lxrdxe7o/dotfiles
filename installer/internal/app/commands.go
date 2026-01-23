package app

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// Message types

// IntroTickMsg is sent for intro animation frames
type IntroTickMsg struct{}

// MatrixTickMsg is sent for matrix rain animation
type MatrixTickMsg struct{}

// SpinnerTickMsg is sent for spinner animation
type SpinnerTickMsg struct{}

// StageCompleteMsg is sent when a stage completes successfully
type StageCompleteMsg struct {
	Index int
}

// StageErrorMsg is sent when a stage fails
type StageErrorMsg struct {
	Index int
	Error error
}

// StageProgressMsg is sent during stage execution
type StageProgressMsg struct {
	Index   int
	Message string
	Current int
	Total   int
}

// Animation tick commands

func tickIntro() tea.Cmd {
	return tea.Tick(time.Millisecond*33, func(t time.Time) tea.Msg {
		return IntroTickMsg{}
	})
}

func tickMatrix() tea.Cmd {
	return tea.Tick(time.Millisecond*50, func(t time.Time) tea.Msg {
		return MatrixTickMsg{}
	})
}

func tickSpinner() tea.Cmd {
	return tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
		return SpinnerTickMsg{}
	})
}
