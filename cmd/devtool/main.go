// cmd/devtool/main.go
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/yourusername/devtool/internal/commands"
)

func main() {
	// Define subcommands
	githubCmd := flag.NewFlagSet("github", flag.ExitOnError)
	confluenceCmd := flag.NewFlagSet("confluence", flag.ExitOnError)
	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	getJiraCmd := flag.NewFlagSet("get_jira_ticket", flag.ExitOnError)
	
	// Define flags for subcommands
	githubCreateCmd := githubCmd.String("create", "", "Create a GitHub repository")
	confluenceCreatePageCmd := confluenceCmd.Bool("create", false, "Create a Confluence page")
	confluencePageCmd := confluenceCmd.Bool("page", false, "Create a Confluence page")
	createProjectType := createCmd.String("type", "", "Project type: golang, spring, nextjs, or terraform")
	createProjectPath := createCmd.String("path", ".", "Path where to create the project")
	
	// Check if no arguments provided
	if len(os.Args) < 2 {
		fmt.Println("Expected subcommands: github, confluence, create, get_jira_ticket")
		fmt.Println("Or use one of these direct commands:")
		fmt.Println("  github create <repo>")
		fmt.Println("  confluence create page")
		fmt.Println("  create base golang project")
		fmt.Println("  create base spring project")
		fmt.Println("  create base nextjs project")
		fmt.Println("  create base terraform project")
		fmt.Println("  get_jira_ticket <ticket-number>")
		os.Exit(1)
	}
	
	// Handle command syntax alternatives
	args := os.Args[1:]
	
	// Handle "github create <repo>"
	if args[0] == "github" && len(args) >= 3 && args[1] == "create" {
		commands.HandleGithubCreate(args[2])
		return
	}
	
	// Handle "confluence create page"
	if args[0] == "confluence" && len(args) >= 3 && args[1] == "create" && args[2] == "page" {
		commands.HandleConfluenceCreatePage()
		return
	}
	
	// Handle "create base <type> project"
	if args[0] == "create" && len(args) >= 4 && args[1] == "base" && args[3] == "project" {
		projectType := args[2]
		switch projectType {
		case "golang":
			commands.CreateBaseGolangProject(".")
		case "spring":
			commands.CreateBaseSpringProject(".")
		case "nextjs":
			commands.CreateBaseNextJsProject(".")
		case "terraform":
			commands.CreateBaseTerraformProject(".")
		default:
			fmt.Printf("Unknown project type: %s\n", projectType)
		}
		return
	}
	
	// Handle "get_jira_ticket <ticket-number>"
	if args[0] == "get_jira_ticket" && len(args) >= 2 {
		commands.GetJiraTicket(args[1])
		return
	}
	
	// Handle subcommands with flags
	switch args[0] {
	case "github":
		githubCmd.Parse(args[1:])
		if *githubCreateCmd != "" {
			commands.HandleGithubCreate(*githubCreateCmd)
		}
		
	case "confluence":
		confluenceCmd.Parse(args[1:])
		if *confluenceCreatePageCmd || *confluencePageCmd {
			commands.HandleConfluenceCreatePage()
		}
		
	case "create":
		createCmd.Parse(args[1:])
		if *createProjectType != "" {
			switch *createProjectType {
			case "golang":
				commands.CreateBaseGolangProject(*createProjectPath)
			case "spring":
				commands.CreateBaseSpringProject(*createProjectPath)
			case "nextjs":
				commands.CreateBaseNextJsProject(*createProjectPath)
			case "terraform":
				commands.CreateBaseTerraformProject(*createProjectPath)
			default:
				log.Fatalf("Unknown project type: %s", *createProjectType)
			}
		}
		
	case "get_jira_ticket":
		getJiraCmd.Parse(args[1:])
		if len(getJiraCmd.Args()) > 0 {
			commands.GetJiraTicket(getJiraCmd.Arg(0))
		} else {
			fmt.Println("Error: Ticket number is required")
			fmt.Println("Usage: devtool get_jira_ticket <ticket-number>")
		}
		
	default:
		fmt.Printf("Unknown command: %s\n", args[0])
		fmt.Println("Expected 'github', 'confluence', 'create', or 'get_jira_ticket'")
		os.Exit(1)
	}
}
