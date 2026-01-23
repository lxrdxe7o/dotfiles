package config

import (
	"flag"
	"os"
	"path/filepath"
)

type Config struct {
	Verbose      bool
	DryRun       bool
	SkipBackup   bool
	DotfilesPath string
	RepoURL      string
	Version      bool
}

var Version = "1.0.0"

func ParseFlags() *Config {
	cfg := &Config{}

	flag.BoolVar(&cfg.Verbose, "verbose", false, "Show detailed command output")
	flag.BoolVar(&cfg.Verbose, "v", false, "Show detailed command output (shorthand)")
	flag.BoolVar(&cfg.DryRun, "dry-run", false, "Simulate installation without making changes")
	flag.BoolVar(&cfg.SkipBackup, "skip-backup", false, "Skip backup of existing configs")
	flag.StringVar(&cfg.DotfilesPath, "path", defaultDotfilesPath(), "Path for dotfiles repository")
	flag.StringVar(&cfg.RepoURL, "repo", "https://github.com/xero/dotfiles.git", "Dotfiles repository URL")
	flag.BoolVar(&cfg.Version, "version", false, "Show version")

	flag.Parse()
	return cfg
}

func defaultDotfilesPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return "~/dotfiles"
	}
	return filepath.Join(home, "dotfiles")
}

// Paths returns common paths used during installation
type Paths struct {
	Home       string
	Dotfiles   string
	Config     string
	ConfigBack string
	OhMyZsh    string
	TpmDir     string
}

func NewPaths(dotfilesPath string) *Paths {
	home, _ := os.UserHomeDir()
	return &Paths{
		Home:       home,
		Dotfiles:   dotfilesPath,
		Config:     filepath.Join(home, ".config"),
		ConfigBack: filepath.Join(home, ".config-backup"),
		OhMyZsh:    filepath.Join(home, ".oh-my-zsh"),
		TpmDir:     filepath.Join(home, ".tmux", "plugins", "tpm"),
	}
}
