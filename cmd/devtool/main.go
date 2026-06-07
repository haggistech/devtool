// cmd/devtool/main.go
package main

import (
	"fmt"
	"os"

	"github.com/yourusername/devtool/internal/commands"
)

var verbose bool

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	args := os.Args[1:]

	for i, arg := range args {
		if arg == "-v" || arg == "--verbose" {
			verbose = true
			args = append(args[:i], args[i+1:]...)
			break
		}
	}

	commands.SetVerbose(verbose)

	if len(args) == 0 {
		printUsage()
		os.Exit(1)
	}

	cmd := args[0]

	switch cmd {
	case "help", "-h", "--help":
		printUsage()
	case "init":
		if err := commands.InitProject(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "config":
		if err := commands.ConfigCommand(args[1:]); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "list":
		if err := commands.ListTemplates(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "versions":
		if err := commands.ListVersions(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "ci":
		if err := commands.CICommand(args[1:]); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "test":
		if err := commands.TestCommand(args[1:]); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "pull":
		handlePull(args[1:])
	case "github":
		handleGithub(args[1:])
	case "confluence":
		handleConfluence(args[1:])
	case "create":
		handleCreate(args[1:])
	case "get_jira_ticket":
		handleJira(args[1:])
	default:
		fmt.Printf("Unknown command: %s\n", cmd)
		printUsage()
		os.Exit(1)
	}
}

func handleGithub(args []string) {
	if len(args) < 2 || args[0] != "create" {
		fmt.Println("Usage: devtool github create <repo-name>")
		os.Exit(1)
	}
	if err := commands.HandleGithubCreate(args[1]); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func handleConfluence(args []string) {
	if len(args) < 2 || args[0] != "create" || args[1] != "page" {
		fmt.Println("Usage: devtool confluence create page")
		os.Exit(1)
	}
	if err := commands.HandleConfluenceCreatePage(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func handleCreate(args []string) {
	if len(args) < 3 || args[0] != "base" || args[2] != "project" {
		fmt.Println("Usage: devtool create base <type> project")
		fmt.Println("Types: golang, spring, nextjs, terraform")
		os.Exit(1)
	}

	projectType := args[1]
	var err error
	switch projectType {
	case "golang":
		err = commands.CreateBaseGolangProject(".")
	case "spring":
		err = commands.CreateBaseSpringProject(".")
	case "nextjs":
		err = commands.CreateBaseNextJsProject(".")
	case "terraform":
		err = commands.CreateBaseTerraformProject(".")
	default:
		fmt.Printf("Unknown project type: %s\n", projectType)
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func handleJira(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: devtool get_jira_ticket <ticket-number>")
		os.Exit(1)
	}
	if err := commands.GetJiraTicket(args[0]); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func handlePull(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: devtool pull <template-name> [project-path]")
		fmt.Println("Run 'devtool list' to see available templates")
		os.Exit(1)
	}

	projectPath := args[0]
	if len(args) > 1 {
		projectPath = args[1]
	}

	if err := commands.PullTemplate(args[0], projectPath); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("devtool - Developer utility for project setup and integrations")
	fmt.Println("\nUsage:")
	fmt.Println("  devtool <command> [options]")
	fmt.Println("\nCommands:")
	fmt.Println("  init                              Start interactive project wizard (recommended)")
	fmt.Println("  config [show|set|init|reset]      Manage configuration (~/.devtool.yaml)")
	fmt.Println("  list                              List available project templates")
	fmt.Println("  pull <template> [path]            Clone a template repository")
	fmt.Println("  versions                          Show installed tool versions")
	fmt.Println("  ci <provider> <language> [path]   Generate CI/CD pipeline")
	fmt.Println("  test <language> [path]            Setup testing framework")
	fmt.Println("  github create <repo>              Create a GitHub repository")
	fmt.Println("  confluence create page            Create a Confluence page")
	fmt.Println("  create base <type> project        Create a project template")
	fmt.Println("  get_jira_ticket <ticket>          Fetch a Jira ticket")
	fmt.Println("  help                              Show this help message")
	fmt.Println("\nFlags:")
	fmt.Println("  -v, --verbose                     Enable debug logging")
	fmt.Println("\nProject Types:")
	fmt.Println("  golang, spring, nextjs, terraform")
	fmt.Println("\nEnvironment Variables:")
	fmt.Println("  GitHub:     gh CLI must be installed")
	fmt.Println("  Jira:       JIRA_DOMAIN, JIRA_EMAIL, JIRA_API_TOKEN")
	fmt.Println("  Confluence: CONFLUENCE_URL, CONFLUENCE_USER, CONFLUENCE_TOKEN, CONFLUENCE_SPACE")
}
