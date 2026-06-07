// internal/commands/templates.go
package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type Template struct {
	Name        string
	Description string
	Type        string
	Repo        string
}

var BuiltInTemplates = []Template{
	{
		Name:        "golang-api",
		Description: "Go REST API with database",
		Type:        "golang",
		Repo:        "https://github.com/devtool-templates/golang-api",
	},
	{
		Name:        "golang-cli",
		Description: "Go CLI application",
		Type:        "golang",
		Repo:        "https://github.com/devtool-templates/golang-cli",
	},
	{
		Name:        "spring-boot-api",
		Description: "Spring Boot REST API",
		Type:        "spring",
		Repo:        "https://github.com/devtool-templates/spring-boot-api",
	},
	{
		Name:        "spring-boot-web",
		Description: "Spring Boot web application with Thymeleaf",
		Type:        "spring",
		Repo:        "https://github.com/devtool-templates/spring-boot-web",
	},
	{
		Name:        "nextjs-fullstack",
		Description: "Next.js fullstack app with database",
		Type:        "nextjs",
		Repo:        "https://github.com/devtool-templates/nextjs-fullstack",
	},
	{
		Name:        "nextjs-commerce",
		Description: "Next.js e-commerce template",
		Type:        "nextjs",
		Repo:        "https://github.com/devtool-templates/nextjs-commerce",
	},
	{
		Name:        "terraform-aws",
		Description: "Terraform AWS infrastructure",
		Type:        "terraform",
		Repo:        "https://github.com/devtool-templates/terraform-aws",
	},
	{
		Name:        "terraform-gcp",
		Description: "Terraform GCP infrastructure",
		Type:        "terraform",
		Repo:        "https://github.com/devtool-templates/terraform-gcp",
	},
}

// ListTemplates displays all available templates
func ListTemplates() error {
	fmt.Println("📚 Available Project Templates")
	fmt.Println()

	currentType := ""
	for _, tmpl := range BuiltInTemplates {
		if tmpl.Type != currentType {
			currentType = tmpl.Type
			fmt.Printf("\n%s:\n", toTitleCase(currentType))
		}
		fmt.Printf("  %-20s - %s\n", tmpl.Name, tmpl.Description)
		fmt.Printf("  %s\n", tmpl.Repo)
		fmt.Println()
	}

	fmt.Println("Usage: devtool pull <template-name>")
	return nil
}

// PullTemplate clones a template repository
func PullTemplate(name, projectPath string) error {
	var template *Template
	for i := range BuiltInTemplates {
		if BuiltInTemplates[i].Name == name {
			template = &BuiltInTemplates[i]
			break
		}
	}

	if template == nil {
		return fmt.Errorf("template not found: %s. Run 'devtool list' to see available templates", name)
	}

	fmt.Printf("📥 Pulling template: %s\n", name)
	fmt.Println(template.Description)
	fmt.Println()

	// Check if directory exists
	if _, err := os.Stat(projectPath); err == nil {
		return fmt.Errorf("directory already exists: %s", projectPath)
	}

	// Create parent directory
	parentDir := filepath.Dir(projectPath)
	if parentDir != "." {
		if err := os.MkdirAll(parentDir, 0755); err != nil {
			return fmt.Errorf("failed to create parent directory: %v", err)
		}
	}

	// Clone the repository
	Logf("Cloning from %s", template.Repo)
	fmt.Printf("⏳ Downloading template...\n")

	cmd := exec.Command("git", "clone", template.Repo, projectPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to clone template: %v", err)
	}

	// Remove .git directory
	gitDir := filepath.Join(projectPath, ".git")
	if err := os.RemoveAll(gitDir); err != nil {
		Logf("Warning: failed to remove .git directory: %v", err)
	}

	fmt.Println()
	fmt.Printf("✓ Template '%s' created at %s\n", name, projectPath)
	fmt.Println()
	fmt.Println("Next steps:")
	fmt.Printf("  cd %s\n", projectPath)
	fmt.Println("  # Review README.md for setup instructions")
	fmt.Println("  devtool config init  # Setup your preferences")

	return nil
}

func toTitleCase(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(s[0]-'a'+'A') + s[1:]
}
