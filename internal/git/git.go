package git

import (
	"fmt"
	"os"
	"os/exec"
)

// Init runs `git init` in dir.
func Init(dir string) error {
	return run(dir, "git", "init")
}

// InitialCommit stages everything and creates the first commit.
func InitialCommit(dir string) error {
	if err := run(dir, "git", "add", "."); err != nil {
		return err
	}
	return run(dir, "git", "commit", "-m", "chore: initial scaffold")
}

func run(dir string, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("command %q failed: %w", name+" "+fmt.Sprint(args), err)
	}
	return nil
}
