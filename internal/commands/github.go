// internal/commands/github.go
package commands

import (
	"fmt"
	"os"
	"os/exec"
)

// HandleGithubCreate creates a new GitHub repository
func HandleGithubCreate(repoName string) error {
	if repoName == "" {
		return fmt.Errorf("repository name is required")
	}
	
	fmt.Printf("Creating GitHub repository: %s\n", repoName)
	
	// You'll need to use GitHub API or gh CLI tool here
	// This example uses gh CLI which should be installed
	cmd := exec.Command("gh", "repo", "create", repoName, "--public")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	return cmd.Run()
}
