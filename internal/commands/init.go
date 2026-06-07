// internal/commands/init.go
package commands

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// InitProject starts an interactive wizard to create a new project
func InitProject() error {
	fmt.Println("🚀 Welcome to devtool project initializer!")
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)

	// Get project name
	projectName := promptUserWithReader(reader, "Enter project name", "my-project")
	if err := validateProjectName(projectName); err != nil {
		return fmt.Errorf("invalid project name: %v", err)
	}

	// Get project type
	projectType := promptChoiceWithReader(reader, "Select project type", []string{"golang", "spring", "nextjs", "terraform"})

	// Get project path
	projectPath := promptUserWithReader(reader, "Project path", projectName)
	if err := validateProjectPath(projectPath); err != nil {
		return fmt.Errorf("invalid project path: %v", err)
	}

	// Ask about git initialization
	initGit := promptYesNoWithReader(reader, "Initialize git repository?", true)

	// Ask about GitHub remote (only if git is enabled)
	var githubRepo string
	if initGit {
		createRemote := promptYesNoWithReader(reader, "Create GitHub repository?", false)
		if createRemote {
			githubRepo = promptUserWithReader(reader, "GitHub repository name", projectName)
		}
	}

	// Ask about additional features
	fmt.Println()
	addDocker := promptYesNoWithReader(reader, "Generate Docker files?", true)
	addEnvFile := promptYesNoWithReader(reader, "Generate .env.example?", true)
	addPrecommit := promptYesNoWithReader(reader, "Setup pre-commit hooks?", true)

	// Display summary
	fmt.Println()
	fmt.Println("📋 Project Summary:")
	fmt.Printf("  Name:       %s\n", projectName)
	fmt.Printf("  Type:       %s\n", projectType)
	fmt.Printf("  Path:       %s\n", projectPath)
	fmt.Printf("  Git:        %v\n", initGit)
	if githubRepo != "" {
		fmt.Printf("  GitHub:     %s\n", githubRepo)
	}
	fmt.Printf("  Docker:     %v\n", addDocker)
	fmt.Printf("  .env.example: %v\n", addEnvFile)
	fmt.Printf("  Pre-commit: %v\n", addPrecommit)
	fmt.Println()

	if !promptYesNoWithReader(reader, "Proceed with project creation?", true) {
		fmt.Println("❌ Project creation cancelled")
		return nil
	}

	// Create the project
	fmt.Println()
	var err error
	switch projectType {
	case "golang":
		err = CreateBaseGolangProject(projectPath)
	case "spring":
		err = CreateBaseSpringProject(projectPath)
	case "nextjs":
		err = CreateBaseNextJsProject(projectPath)
	case "terraform":
		err = CreateBaseTerraformProject(projectPath)
	}

	if err != nil {
		return err
	}

	// Check if project directory exists before generating files
	if _, err := os.Stat(projectPath); os.IsNotExist(err) {
		fmt.Printf("⚠️  Warning: Project directory not found at %s\n", projectPath)
		fmt.Println("Some generators were skipped. Please ensure the project creation completed successfully.")
		if addEnvFile || addDocker || addPrecommit {
			fmt.Println()
			fmt.Println("To generate files manually, run from your project directory:")
			if addEnvFile {
				fmt.Printf("  devtool env-generate %s\n", projectType)
			}
			if addDocker {
				fmt.Printf("  devtool docker-generate %s\n", projectType)
			}
			if addPrecommit {
				fmt.Printf("  devtool precommit-generate %s\n", projectType)
			}
		}
		return nil
	}

	// Generate additional files if requested
	if addEnvFile {
		if err := GenerateEnvFile(projectPath, projectType); err != nil {
			fmt.Printf("⚠️  Warning: Failed to generate .env.example: %v\n", err)
		} else {
			fmt.Println("✓ Generated .env.example")
		}
	}

	if addDocker {
		if err := GenerateDockerFiles(projectPath, projectType); err != nil {
			fmt.Printf("⚠️  Warning: Failed to generate Docker files: %v\n", err)
		} else {
			fmt.Println("✓ Generated Dockerfile and docker-compose.yml")
		}
	}

	if addPrecommit {
		if err := GeneratePrecommitConfig(projectPath, projectType); err != nil {
			fmt.Printf("⚠️  Warning: Failed to generate pre-commit config: %v\n", err)
		} else {
			fmt.Println("✓ Generated .pre-commit-config.yaml")
			fmt.Println("  Run 'pre-commit install' to setup hooks")
		}
	}

	// Initialize git if requested
	if initGit {
		if err := initializeGit(projectPath, projectName); err != nil {
			fmt.Printf("⚠️  Warning: Git initialization failed: %v\n", err)
		} else {
			fmt.Println("✓ Git repository initialized")
		}

		// Create GitHub remote if requested
		if githubRepo != "" {
			fmt.Printf("\n📝 To create the GitHub repository, run:\n")
			fmt.Printf("  cd %s\n", projectPath)
			fmt.Printf("  gh repo create %s --public --source=. --remote=origin --push\n", githubRepo)
		}
	}

	fmt.Println()
	fmt.Printf("✨ Project '%s' created successfully at %s\n", projectName, projectPath)
	fmt.Println()
	fmt.Println("🚀 Next steps:")
	fmt.Printf("  cd %s\n", projectPath)
	if projectType == "golang" {
		fmt.Println("  go mod tidy")
		fmt.Println("  go run cmd/main.go")
	} else if projectType == "spring" {
		fmt.Println("  mvn clean install")
		fmt.Println("  mvn spring-boot:run")
	} else if projectType == "nextjs" {
		fmt.Println("  npm install")
		fmt.Println("  npm run dev")
	} else if projectType == "terraform" {
		fmt.Println("  cd environments/dev")
		fmt.Println("  terraform init")
		fmt.Println("  terraform plan")
	}

	return nil
}

// promptUserWithReader prompts for user input with a default value
func promptUserWithReader(reader *bufio.Reader, prompt, defaultValue string) string {
	fmt.Printf("%s [%s]: ", prompt, defaultValue)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "" {
		return defaultValue
	}
	return input
}

// promptChoiceWithReader prompts user to select from options
func promptChoiceWithReader(reader *bufio.Reader, prompt string, options []string) string {
	for {
		fmt.Printf("%s:\n", prompt)
		for i, option := range options {
			fmt.Printf("  %d) %s\n", i+1, option)
		}
		fmt.Print("Select option (1-" + fmt.Sprint(len(options)) + "): ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			if len(options) >= 1 {
				return options[0]
			}
		case "2":
			if len(options) >= 2 {
				return options[1]
			}
		case "3":
			if len(options) >= 3 {
				return options[2]
			}
		case "4":
			if len(options) >= 4 {
				return options[3]
			}
		}
		fmt.Println("❌ Invalid choice. Please try again.")
	}
}

// promptYesNoWithReader prompts for a yes/no question
func promptYesNoWithReader(reader *bufio.Reader, prompt string, defaultYes bool) bool {
	defaultStr := "y/N"
	if defaultYes {
		defaultStr = "Y/n"
	}

	fmt.Printf("%s [%s]: ", prompt, defaultStr)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))

	if input == "" {
		return defaultYes
	}

	return input == "y" || input == "yes"
}

// validateProjectName checks if project name is valid
func validateProjectName(name string) error {
	if len(name) == 0 {
		return fmt.Errorf("project name cannot be empty")
	}
	if len(name) > 100 {
		return fmt.Errorf("project name too long (max 100 characters)")
	}
	return nil
}

// initializeGit initializes a git repository in the project path
func initializeGit(projectPath, projectName string) error {
	Logf("Initializing git repository in %s", projectPath)

	// Check if git is installed
	_, err := exec.LookPath("git")
	if err != nil {
		return fmt.Errorf("git not found. Please install git: https://git-scm.com")
	}

	// Initialize repository
	cmd := exec.Command("git", "init")
	cmd.Dir = projectPath
	if err := cmd.Run(); err != nil {
		return err
	}

	// Configure git user if not already configured
	nameCmd := exec.Command("git", "config", "user.name")
	nameCmd.Dir = projectPath
	if err := nameCmd.Run(); err != nil {
		Logf("Git user.name not configured, skipping")
	}

	// Add all files
	addCmd := exec.Command("git", "add", ".")
	addCmd.Dir = projectPath
	if err := addCmd.Run(); err != nil {
		return fmt.Errorf("git add failed: %v", err)
	}

	// Create initial commit
	commitCmd := exec.Command("git", "commit", "-m", fmt.Sprintf("Initial commit: %s", projectName))
	commitCmd.Dir = projectPath
	if err := commitCmd.Run(); err != nil {
		Logf("Git commit failed (may need git user configuration): %v", err)
	}

	return nil
}
