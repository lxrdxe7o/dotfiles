package executor

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"

	tea "github.com/charmbracelet/bubbletea"
)

// OutputLine represents a line of command output
type OutputLine struct {
	StageIndex int
	Line       string
	IsError    bool
}

// CommandResult represents the result of a command execution
type CommandResult struct {
	StageIndex int
	Output     []string
	Error      error
}

// Executor handles async command execution
type Executor struct {
	dryRun bool
	mu     sync.Mutex
}

// NewExecutor creates a new executor
func NewExecutor(dryRun bool) *Executor {
	return &Executor{
		dryRun: dryRun,
	}
}

// RunCmd executes a command and returns a tea.Cmd that sends messages for output
func (e *Executor) RunCmd(ctx context.Context, stageIndex int, name string, args ...string) tea.Cmd {
	return func() tea.Msg {
		if e.dryRun {
			return CommandResult{
				StageIndex: stageIndex,
				Output:     []string{fmt.Sprintf("[DRY RUN] Would execute: %s %s", name, strings.Join(args, " "))},
				Error:      nil,
			}
		}

		cmd := exec.CommandContext(ctx, name, args...)
		output, err := e.runWithOutput(cmd)
		return CommandResult{
			StageIndex: stageIndex,
			Output:     output,
			Error:      err,
		}
	}
}

// RunCmdWithSudo executes a command with sudo
func (e *Executor) RunCmdWithSudo(ctx context.Context, stageIndex int, name string, args ...string) tea.Cmd {
	return func() tea.Msg {
		if e.dryRun {
			return CommandResult{
				StageIndex: stageIndex,
				Output:     []string{fmt.Sprintf("[DRY RUN] Would execute with sudo: %s %s", name, strings.Join(args, " "))},
				Error:      nil,
			}
		}

		// Build sudo command
		sudoArgs := append([]string{name}, args...)
		cmd := exec.CommandContext(ctx, "sudo", sudoArgs...)
		cmd.Stdin = os.Stdin // Allow sudo to prompt for password

		output, err := e.runWithOutput(cmd)
		return CommandResult{
			StageIndex: stageIndex,
			Output:     output,
			Error:      err,
		}
	}
}

// RunScript runs a shell script
func (e *Executor) RunScript(ctx context.Context, stageIndex int, script string) tea.Cmd {
	return func() tea.Msg {
		if e.dryRun {
			lines := strings.Split(script, "\n")
			output := []string{"[DRY RUN] Would execute script:"}
			for _, line := range lines {
				if strings.TrimSpace(line) != "" {
					output = append(output, "  "+line)
				}
			}
			return CommandResult{
				StageIndex: stageIndex,
				Output:     output,
				Error:      nil,
			}
		}

		cmd := exec.CommandContext(ctx, "bash", "-c", script)
		output, err := e.runWithOutput(cmd)
		return CommandResult{
			StageIndex: stageIndex,
			Output:     output,
			Error:      err,
		}
	}
}

// runWithOutput runs a command and captures output
func (e *Executor) runWithOutput(cmd *exec.Cmd) ([]string, error) {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to get stdout: %w", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to get stderr: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start command: %w", err)
	}

	var output []string
	var wg sync.WaitGroup

	readPipe := func(pipe io.Reader, prefix string) {
		defer wg.Done()
		scanner := bufio.NewScanner(pipe)
		// Increase buffer size for long lines
		buf := make([]byte, 0, 64*1024)
		scanner.Buffer(buf, 1024*1024)
		for scanner.Scan() {
			line := scanner.Text()
			e.mu.Lock()
			if prefix != "" {
				output = append(output, prefix+line)
			} else {
				output = append(output, line)
			}
			e.mu.Unlock()
		}
	}

	wg.Add(2)
	go readPipe(stdout, "")
	go readPipe(stderr, "")
	wg.Wait()

	err = cmd.Wait()
	return output, err
}

// CheckCommandExists checks if a command is available
func CheckCommandExists(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

// CheckPackageInstalled checks if a package is installed via pacman
func CheckPackageInstalled(pkg string) bool {
	cmd := exec.Command("pacman", "-Qi", pkg)
	return cmd.Run() == nil
}

// GetInstalledPackages returns a list of installed packages
func GetInstalledPackages() ([]string, error) {
	cmd := exec.Command("pacman", "-Qq")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	return lines, nil
}

// ReadPackageList reads packages from a file
func ReadPackageList(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var packages []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "#") {
			packages = append(packages, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return packages, nil
}

// FileExists checks if a file or directory exists
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// DirExists checks if a directory exists
func DirExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}
