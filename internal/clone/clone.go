package clone

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func Repo(path, repo string) error {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return fmt.Errorf("directory %s already exists", path)
	}
	cmd := exec.Command("git", "clone", "--depth", "1", repo, path)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to clone repo: %w", err)
	}
	gitDir := filepath.Join(path, ".git")
	if err := os.RemoveAll(gitDir); err != nil {
		return fmt.Errorf("failed to remove .git directory: %v", err)
	}
	return nil
}
