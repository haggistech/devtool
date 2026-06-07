// internal/config/config.go
package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

const ConfigFileName = ".devtool.yaml"

type Config struct {
	DefaultLanguage string `yaml:"default_language"`
	DefaultGitEmail string `yaml:"default_git_email"`
	DefaultGitName  string `yaml:"default_git_name"`
	GitHubToken     string `yaml:"github_token,omitempty"`
	JiraDomain      string `yaml:"jira_domain,omitempty"`
	JiraEmail       string `yaml:"jira_email,omitempty"`
	JiraToken       string `yaml:"jira_token,omitempty"`
	ConfluenceURL   string `yaml:"confluence_url,omitempty"`
	ConfluenceUser  string `yaml:"confluence_user,omitempty"`
	ConfluenceToken string `yaml:"confluence_token,omitempty"`
}

// Load reads the config file from home directory
func Load() (*Config, error) {
	configPath := GetConfigPath()
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return &Config{}, nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %v", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %v", err)
	}

	return &cfg, nil
}

// Save writes the config to home directory
func (c *Config) Save() error {
	configPath := GetConfigPath()
	data, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %v", err)
	}

	if err := os.WriteFile(configPath, data, 0600); err != nil {
		return fmt.Errorf("failed to write config: %v", err)
	}

	return nil
}

// GetConfigPath returns the path to the config file
func GetConfigPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ConfigFileName
	}
	return filepath.Join(home, ConfigFileName)
}

// GetOrDefault returns the config value or falls back to environment variable
func (c *Config) GetJiraDomain() string {
	if c.JiraDomain != "" {
		return c.JiraDomain
	}
	return os.Getenv("JIRA_DOMAIN")
}

// GetJiraEmail returns the config value or falls back to environment variable
func (c *Config) GetJiraEmail() string {
	if c.JiraEmail != "" {
		return c.JiraEmail
	}
	return os.Getenv("JIRA_EMAIL")
}

// GetJiraToken returns the config value or falls back to environment variable
func (c *Config) GetJiraToken() string {
	if c.JiraToken != "" {
		return c.JiraToken
	}
	return os.Getenv("JIRA_API_TOKEN")
}

// GetConfluenceURL returns the config value or falls back to environment variable
func (c *Config) GetConfluenceURL() string {
	if c.ConfluenceURL != "" {
		return c.ConfluenceURL
	}
	return os.Getenv("CONFLUENCE_URL")
}

// GetConfluenceUser returns the config value or falls back to environment variable
func (c *Config) GetConfluenceUser() string {
	if c.ConfluenceUser != "" {
		return c.ConfluenceUser
	}
	return os.Getenv("CONFLUENCE_USER")
}

// GetConfluenceToken returns the config value or falls back to environment variable
func (c *Config) GetConfluenceToken() string {
	if c.ConfluenceToken != "" {
		return c.ConfluenceToken
	}
	return os.Getenv("CONFLUENCE_TOKEN")
}
