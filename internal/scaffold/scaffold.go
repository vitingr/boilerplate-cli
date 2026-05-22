package scaffold

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/vitingr/boilerplate-cli/internal/config"
	"github.com/vitingr/boilerplate-cli/internal/git"
)

var (
	info    = color.New(color.FgCyan).PrintfFunc()
	success = color.New(color.FgGreen).PrintfFunc()
	warn    = color.New(color.FgYellow).PrintfFunc()
	fail    = color.New(color.FgRed).PrintfFunc()
)

func Run(cfg *config.ScaffoldConfig) error {
	info("⏳  Fetching boilerplate %s/%s/%s …\n", cfg.Language, cfg.Framework, cfg.Template)

	if err := cloneBoilerplate(cfg); err != nil {
		fail("✗ Clone failed: %v\n", err)
		return err
	}
	success("✔  Boilerplate copied to %s\n", cfg.OutputDir)

	if err := renameProject(cfg.OutputDir, cfg.ProjectName); err != nil {
		warn("⚠  Could not rename project references: %v\n", err)
	}

	if cfg.InitGit {
		info("⏳  Initialising git …\n")
		if err := git.Init(cfg.OutputDir); err != nil {
			warn("⚠  git init failed: %v\n", err)
		} else {
			_ = git.InitialCommit(cfg.OutputDir)
			success("✔  Git repository initialised\n")
		}
	}

	if cfg.InstallDeps {
		if err := installDeps(cfg); err != nil {
			warn("⚠  Dependency installation failed: %v\n", err)
		} else {
			success("✔  Dependencies installed\n")
		}
	}

	fmt.Println()
	success("🚀  Project %q ready at %s\n\n", cfg.ProjectName, cfg.OutputDir)
	info("  cd %s\n", cfg.OutputDir)
	fmt.Println()
	return nil
}

func cloneBoilerplate(cfg *config.ScaffoldConfig) error {
	tmp, err := os.MkdirTemp("", "bplt-*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmp)

	repoURL := config.RepoBaseURL + ".git"
	subdir := filepath.Join(cfg.Language, cfg.Framework, cfg.Template)

	cmds := [][]string{
		{"git", "clone", "--filter=blob:none", "--sparse", "--depth=1", repoURL, tmp},
		{"git", "-C", tmp, "sparse-checkout", "set", subdir},
	}

	for _, c := range cmds {
		cmd := exec.Command(c[0], c[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("%s: %w", strings.Join(c, " "), err)
		}
	}

	src := filepath.Join(tmp, subdir)
	if err := os.MkdirAll(filepath.Dir(cfg.OutputDir), 0o755); err != nil {
		return err
	}
	return os.Rename(src, cfg.OutputDir)
}

func renameProject(dir, name string) error {
	return filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		if !isTextFile(path) {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return nil // skip unreadable files
		}
		updated := strings.ReplaceAll(string(data), "boilerplate", name)
		updated = strings.ReplaceAll(updated, "BOILERPLATE", strings.ToUpper(name))
		if updated != string(data) {
			return os.WriteFile(path, []byte(updated), 0o644)
		}
		return nil
	})
}

func isTextFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	textExts := map[string]bool{
		".go": true, ".ts": true, ".js": true, ".json": true,
		".mod": true, ".sum": true, ".toml": true, ".yaml": true,
		".yml": true, ".env": true, ".md": true, ".txt": true,
		".py": true, ".sh": true, ".dockerfile": true,
	}
	return textExts[ext]
}

func installDeps(cfg *config.ScaffoldConfig) error {
	var cmd *exec.Cmd
	switch cfg.Language {
	case "node":
		cmd = exec.Command("npm", "install")
	case "golang":
		cmd = exec.Command("go", "mod", "tidy")
	case "python":
		cmd = exec.Command("pip", "install", "-r", "requirements.txt")
	default:
		return nil
	}
	cmd.Dir = cfg.OutputDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
