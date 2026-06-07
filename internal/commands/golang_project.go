// internal/commands/golang_project.go
package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// CreateBaseGolangProject creates a new Golang project structure
func CreateBaseGolangProject(projectPath string) error {
	if projectPath == "" {
		projectPath = "."
	}

	if err := validateProjectPath(projectPath); err != nil {
		return err
	}

	fmt.Printf("Creating base Golang project in %s\n", projectPath)
	Logf("Project path validated: %s", projectPath)

	// Check if go.mod already exists
	if CheckFileExists(filepath.Join(projectPath, "go.mod")) {
		return fmt.Errorf("go.mod already exists at %s. This looks like an existing Go project", projectPath)
	}

	// Create directories
	dirs := []string{
		"cmd",
		"internal",
		"pkg",
		"api",
		"configs",
		"docs",
	}

	for _, dir := range dirs {
		fullPath := filepath.Join(projectPath, dir)
		Logf("Creating directory: %s", fullPath)
		err := os.MkdirAll(fullPath, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory %s: %v", dir, err)
		}
	}

	// Create main.go
	mainContent := `package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")
}
`
	err := os.WriteFile(filepath.Join(projectPath, "cmd", "main.go"), []byte(mainContent), 0644)
	if err != nil {
		return err
	}

	// Create go.mod
	modName := filepath.Base(projectPath)
	if modName == "." {
		modName = "example.com/myproject"
	} else {
		modName = "example.com/" + modName
	}

	modContent := fmt.Sprintf(`module %s

go 1.20
`, modName)

	err = os.WriteFile(filepath.Join(projectPath, "go.mod"), []byte(modContent), 0644)
	if err != nil {
		return err
	}

	// Create README.md
	readmeContent := fmt.Sprintf(`# %s

A Go project.

## Getting Started

### Prerequisites

- Go 1.20 or later

### Installation
`+"```"+`
go mod tidy
`+"```"+`

### Running

`+"```"+`
go run cmd/main.go
`+"```"+`

`, modName)

	err = os.WriteFile(filepath.Join(projectPath, "README.md"), []byte(readmeContent), 0644)
	if err != nil {
		return err
	}

	fmt.Println("✓ Golang project created successfully")
	fmt.Printf("Next steps:\n  cd %s\n  go mod tidy\n  go run cmd/main.go\n", projectPath)

	return nil
}

// validateProjectPath checks if the project path is valid
func validateProjectPath(path string) error {
	if path == "." {
		return nil
	}
	if len(path) == 0 {
		return fmt.Errorf("project path cannot be empty")
	}
	if strings.Contains(path, "..") {
		return fmt.Errorf("project path cannot contain '..'")
	}
	return nil
}
