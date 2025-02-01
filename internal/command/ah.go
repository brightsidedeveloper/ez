package command

import (
	"ez/internal/clone"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func Ah(name string) error {
	if name == "" {
		return fmt.Errorf("name is required")
	}
	baseDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Failed to get current working directory: %v\n", err)
	}

	repoDir := filepath.Join(baseDir, name)
	clientDir := filepath.Join(repoDir, "client")
	serverDir := filepath.Join(repoDir, "server")
	if err := clone.Repo(repoDir, "https://github.com/brightsidedeveloper/ah.git"); err != nil {
		return err
	}
	cmd := exec.Command("pnpm", "i")
	cmd.Dir = clientDir
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to install deps: %w", err)
	}
	cmd2 := exec.Command("go", "mod", "tidy")
	cmd2.Dir = serverDir
	if err := cmd2.Run(); err != nil {
		return fmt.Errorf("failed to tidy deps: %w", err)
	}
	return nil
}
