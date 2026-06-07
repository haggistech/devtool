// internal/commands/config_cmd.go
package commands

import (
	"bufio"
	"fmt"
	"os"

	"github.com/yourusername/devtool/internal/config"
)

// ConfigCommand handles config-related operations
func ConfigCommand(args []string) error {
	if len(args) == 0 {
		return configShow()
	}

	switch args[0] {
	case "show":
		return configShow()
	case "set":
		if len(args) < 3 {
			return fmt.Errorf("usage: devtool config set <key> <value>")
		}
		return configSet(args[1], args[2])
	case "init":
		return configInit()
	case "reset":
		return configReset()
	default:
		return fmt.Errorf("unknown config command: %s", args[0])
	}
}

// configShow displays the current configuration
func configShow() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	fmt.Println("📋 Current Configuration:")
	fmt.Println()
	fmt.Printf("Default Language: %s\n", orDefault(cfg.DefaultLanguage, "not set"))
	fmt.Printf("Git Email:        %s\n", orDefault(cfg.DefaultGitEmail, "not set"))
	fmt.Printf("Git Name:         %s\n", orDefault(cfg.DefaultGitName, "not set"))
	fmt.Println()
	fmt.Printf("GitHub Token:     %s\n", maskSecret(cfg.GitHubToken))
	fmt.Printf("Jira Domain:      %s\n", orDefault(cfg.JiraDomain, "not set"))
	fmt.Printf("Jira Email:       %s\n", orDefault(cfg.JiraEmail, "not set"))
	fmt.Printf("Jira Token:       %s\n", maskSecret(cfg.JiraToken))
	fmt.Printf("Confluence URL:   %s\n", orDefault(cfg.ConfluenceURL, "not set"))
	fmt.Printf("Confluence User:  %s\n", orDefault(cfg.ConfluenceUser, "not set"))
	fmt.Printf("Confluence Token: %s\n", maskSecret(cfg.ConfluenceToken))
	fmt.Println()
	fmt.Printf("Config location: %s\n", config.GetConfigPath())

	return nil
}

// configSet sets a configuration value
func configSet(key, value string) error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	switch key {
	case "default_language":
		cfg.DefaultLanguage = value
	case "default_git_email":
		cfg.DefaultGitEmail = value
	case "default_git_name":
		cfg.DefaultGitName = value
	case "github_token":
		cfg.GitHubToken = value
	case "jira_domain":
		cfg.JiraDomain = value
	case "jira_email":
		cfg.JiraEmail = value
	case "jira_token":
		cfg.JiraToken = value
	case "confluence_url":
		cfg.ConfluenceURL = value
	case "confluence_user":
		cfg.ConfluenceUser = value
	case "confluence_token":
		cfg.ConfluenceToken = value
	default:
		return fmt.Errorf("unknown config key: %s", key)
	}

	if err := cfg.Save(); err != nil {
		return err
	}

	fmt.Printf("✓ Set %s = %s\n", key, maskSecret(value))
	return nil
}

// configInit starts an interactive configuration setup
func configInit() error {
	fmt.Println("🔧 Interactive Configuration Setup")
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	// Git configuration
	fmt.Println("Git Configuration:")
	name := promptUserWithReader(reader, "  Git name", cfg.DefaultGitName)
	cfg.DefaultGitName = name

	email := promptUserWithReader(reader, "  Git email", cfg.DefaultGitEmail)
	cfg.DefaultGitEmail = email

	// Default language
	fmt.Println()
	lang := promptChoiceWithReader(reader, "Default project language", []string{"golang", "spring", "nextjs", "terraform"})
	cfg.DefaultLanguage = lang

	// API credentials (optional)
	fmt.Println()
	if promptYesNoWithReader(reader, "Configure Jira credentials?", false) {
		cfg.JiraDomain = promptUserWithReader(reader, "  Jira domain (your-company.atlassian.net)", cfg.JiraDomain)
		cfg.JiraEmail = promptUserWithReader(reader, "  Jira email", cfg.JiraEmail)
		cfg.JiraToken = promptUserWithReader(reader, "  Jira API token", "")
	}

	fmt.Println()
	if promptYesNoWithReader(reader, "Configure Confluence credentials?", false) {
		cfg.ConfluenceURL = promptUserWithReader(reader, "  Confluence URL", cfg.ConfluenceURL)
		cfg.ConfluenceUser = promptUserWithReader(reader, "  Confluence user", cfg.ConfluenceUser)
		cfg.ConfluenceToken = promptUserWithReader(reader, "  Confluence token", "")
	}

	if err := cfg.Save(); err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("✓ Configuration saved!")
	fmt.Printf("Location: %s\n", config.GetConfigPath())

	return nil
}

// configReset removes the config file
func configReset() error {
	configPath := config.GetConfigPath()
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Println("No config file to reset")
		return nil
	}

	reader := bufio.NewReader(os.Stdin)
	if !promptYesNoWithReader(reader, "Delete config file?", false) {
		fmt.Println("Reset cancelled")
		return nil
	}

	if err := os.Remove(configPath); err != nil {
		return fmt.Errorf("failed to delete config: %v", err)
	}

	fmt.Println("✓ Config file deleted")
	return nil
}

func orDefault(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}

func maskSecret(value string) string {
	if value == "" {
		return "not set"
	}
	if len(value) <= 4 {
		return "****"
	}
	return "****" + value[len(value)-4:]
}
