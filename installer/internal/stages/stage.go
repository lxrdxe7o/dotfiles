package stages

import (
	"context"
)

// StageStatus represents the current status of a stage
type StageStatus int

const (
	StatusPending StageStatus = iota
	StatusRunning
	StatusComplete
	StatusSkipped
	StatusFailed
)

func (s StageStatus) String() string {
	switch s {
	case StatusPending:
		return "Pending"
	case StatusRunning:
		return "Running"
	case StatusComplete:
		return "Complete"
	case StatusSkipped:
		return "Skipped"
	case StatusFailed:
		return "Failed"
	default:
		return "Unknown"
	}
}

// ProgressUpdate is sent during stage execution
type ProgressUpdate struct {
	Message  string
	Current  int
	Total    int
	ItemName string
}

// Stage represents an installation stage
type Stage struct {
	ID          string
	Name        string
	Description string
	Icon        string

	Status         StageStatus
	Progress       int    // 0-100
	CurrentItem    string // Current package/file being processed
	TotalItems     int
	ProcessedItems int

	// Behavior flags
	CanSkip  bool
	CanRetry bool
	Enabled  bool // Whether this stage will run

	// Idempotency check
	IsCompleteFunc func() bool

	// Execution function
	Action func(ctx context.Context, progress chan<- ProgressUpdate, dryRun bool) error

	// Error info
	Error  error
	Output []string
}

// NewStage creates a new stage with defaults
func NewStage(id, name, description, icon string) *Stage {
	return &Stage{
		ID:          id,
		Name:        name,
		Description: description,
		Icon:        icon,
		Status:      StatusPending,
		CanSkip:     true,
		CanRetry:    true,
		Enabled:     true,
		Output:      make([]string, 0),
	}
}

// IsComplete checks if the stage is already complete
func (s *Stage) IsComplete() bool {
	if s.IsCompleteFunc != nil {
		return s.IsCompleteFunc()
	}
	return false
}

// Reset resets the stage to pending state
func (s *Stage) Reset() {
	s.Status = StatusPending
	s.Progress = 0
	s.CurrentItem = ""
	s.ProcessedItems = 0
	s.Error = nil
	s.Output = make([]string, 0)
}

// AddOutput appends a line to stage output
func (s *Stage) AddOutput(line string) {
	s.Output = append(s.Output, line)
	// Keep only last 100 lines
	if len(s.Output) > 100 {
		s.Output = s.Output[len(s.Output)-100:]
	}
}

// LastOutput returns the last n lines of output
func (s *Stage) LastOutput(n int) []string {
	if n >= len(s.Output) {
		return s.Output
	}
	return s.Output[len(s.Output)-n:]
}

// StageList manages a list of stages
type StageList struct {
	Stages       []*Stage
	CurrentIndex int
}

// NewStageList creates a new stage list
func NewStageList() *StageList {
	return &StageList{
		Stages:       make([]*Stage, 0),
		CurrentIndex: 0,
	}
}

// Add adds a stage to the list
func (sl *StageList) Add(stage *Stage) {
	sl.Stages = append(sl.Stages, stage)
}

// Current returns the current stage
func (sl *StageList) Current() *Stage {
	if sl.CurrentIndex >= 0 && sl.CurrentIndex < len(sl.Stages) {
		return sl.Stages[sl.CurrentIndex]
	}
	return nil
}

// Next advances to the next stage
func (sl *StageList) Next() bool {
	sl.CurrentIndex++
	// Skip disabled stages
	for sl.CurrentIndex < len(sl.Stages) && !sl.Stages[sl.CurrentIndex].Enabled {
		sl.Stages[sl.CurrentIndex].Status = StatusSkipped
		sl.CurrentIndex++
	}
	return sl.CurrentIndex < len(sl.Stages)
}

// IsFinished returns true if all stages are done
func (sl *StageList) IsFinished() bool {
	return sl.CurrentIndex >= len(sl.Stages)
}

// HasFailed returns true if any stage failed
func (sl *StageList) HasFailed() bool {
	for _, s := range sl.Stages {
		if s.Status == StatusFailed {
			return true
		}
	}
	return false
}

// CompletedCount returns the number of completed stages
func (sl *StageList) CompletedCount() int {
	count := 0
	for _, s := range sl.Stages {
		if s.Status == StatusComplete || s.Status == StatusSkipped {
			count++
		}
	}
	return count
}

// EnabledCount returns the number of enabled stages
func (sl *StageList) EnabledCount() int {
	count := 0
	for _, s := range sl.Stages {
		if s.Enabled {
			count++
		}
	}
	return count
}
