// internal/commands/nextjs_project.go
package commands

import (
	"fmt"
	"os"
	"os/exec"
)

// CreateBaseNextJsProject creates a new Next.js project
func CreateBaseNextJsProject(projectPath string) error {
	if projectPath == "" {
		projectPath = "."
	}

	if err := validateProjectPath(projectPath); err != nil {
		return err
	}

	fmt.Printf("Creating base Next.js project in %s\n", projectPath)
	fmt.Println("⏳ This may take a few minutes...")
	Logf("Executing: npx create-next-app %s", projectPath)

	cmd := exec.Command("npx", "create-next-app", projectPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		if err.Error() == "executable file not found in $PATH" {
			return fmt.Errorf("npx not found. Please install Node.js: https://nodejs.org")
		}
		return err
	}

	fmt.Println("✓ Next.js project created successfully")
	return nil
}
