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

	Logf("Validating repository name: %s", repoName)
	if err := validateRepoName(repoName); err != nil {
		return err
	}

	fmt.Printf("Creating GitHub repository: %s\n", repoName)
	Logf("Executing: gh repo create %s --public", repoName)

	cmd := exec.Command("gh", "repo", "create", repoName, "--public")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		if err.Error() == "executable file not found in $PATH" {
			return fmt.Errorf("gh CLI not found. Please install it: https://cli.github.com")
		}
		return err
	}

	Log("GitHub repository created successfully")
	return nil
}

// validateRepoName checks if the repository name is valid
func validateRepoName(name string) error {
	if len(name) == 0 {
		return fmt.Errorf("repository name cannot be empty")
	}
	if len(name) > 100 {
		return fmt.Errorf("repository name too long (max 100 characters)")
	}
	if !isValidRepoName(name) {
		return fmt.Errorf("repository name must contain only alphanumeric characters, hyphens, and underscores")
	}
	return nil
}

func isValidRepoName(name string) bool {
	for _, ch := range name {
		if !((ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') || ch == '-' || ch == '_') {
			return false
		}
	}
	return true
}
