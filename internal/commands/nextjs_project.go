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
	
	fmt.Printf("Creating base Next.js project in %s\n", projectPath)
	
	// For a proper Next.js project, we would use npm/npx
	// This example will execute the npx command
	cmd := exec.Command("npx", "create-next-app", projectPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	return cmd.Run()
}
