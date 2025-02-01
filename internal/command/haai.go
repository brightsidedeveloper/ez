package command

import (
	"ez/internal/clone"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func Haai(name string) error {
	if name == "" {
		return fmt.Errorf("name is required")
	}
	baseDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Failed to get current working directory: %v\n", err)
	}

	repoDir := filepath.Join(baseDir, name)
	viteDir := filepath.Join(repoDir, "react-vite")
	nativeDir := filepath.Join(repoDir, "react-native")
	serverDir := filepath.Join(repoDir, "go-bsd")

	if err := clone.Repo(repoDir, "https://github.com/brightsidedeveloper/haai.git"); err != nil {
		return err
	}
	cmd := exec.Command("pnpm", "i")
	cmd.Dir = viteDir
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to install deps: %w", err)
	}
	cmd2 := exec.Command("npm", "i")
	cmd2.Dir = nativeDir
	if err := cmd2.Run(); err != nil {
		return fmt.Errorf("failed to install deps: %w", err)
	}
	cmd3 := exec.Command("go", "mod", "tidy")
	cmd3.Dir = serverDir

	if err := cmd3.Run(); err != nil {
		return fmt.Errorf("failed to tidy deps: %w", err)
	}
	return nil
}
