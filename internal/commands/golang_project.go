// internal/commands/golang_project.go
package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

// CreateBaseGolangProject creates a new Golang project structure
func CreateBaseGolangProject(projectPath string) error {
	if projectPath == "" {
		projectPath = "."
	}

	fmt.Printf("Creating base Golang project in %s\n", projectPath)

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
		err := os.MkdirAll(filepath.Join(projectPath, dir), 0755)
		if err != nil {
			return err
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

	return os.WriteFile(filepath.Join(projectPath, "README.md"), []byte(readmeContent), 0644)
}
