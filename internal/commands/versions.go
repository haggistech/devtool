// internal/commands/versions.go
package commands

import (
	"fmt"
	"os/exec"
	"strings"
)

type ToolVersion struct {
	Name    string
	Version string
	Status  string // "installed", "not found"
}

var toolsToCheck = []struct {
	Name  string
	Cmd   string
	Args  []string
	Parse func(string) string
}{
	{
		Name: "Node.js",
		Cmd:  "node",
		Args: []string{"--version"},
		Parse: func(s string) string {
			return strings.TrimPrefix(strings.TrimSpace(s), "v")
		},
	},
	{
		Name: "npm",
		Cmd:  "npm",
		Args: []string{"--version"},
		Parse: func(s string) string {
			return strings.TrimSpace(s)
		},
	},
	{
		Name: "pnpm",
		Cmd:  "pnpm",
		Args: []string{"--version"},
		Parse: func(s string) string {
			return strings.TrimSpace(s)
		},
	},
	{
		Name: "yarn",
		Cmd:  "yarn",
		Args: []string{"--version"},
		Parse: func(s string) string {
			return strings.TrimSpace(s)
		},
	},
	{
		Name: "Go",
		Cmd:  "go",
		Args: []string{"version"},
		Parse: func(s string) string {
			parts := strings.Fields(strings.TrimSpace(s))
			if len(parts) >= 3 {
				return strings.TrimPrefix(parts[2], "go")
			}
			return strings.TrimSpace(s)
		},
	},
	{
		Name: "Java",
		Cmd:  "java",
		Args: []string{"-version"},
		Parse: func(s string) string {
			lines := strings.Split(s, "\n")
			if len(lines) > 0 {
				line := lines[0]
				if strings.Contains(line, "version") {
					parts := strings.Split(line, "\"")
					if len(parts) >= 2 {
						return parts[1]
					}
				}
			}
			return strings.TrimSpace(s)
		},
	},
	{
		Name: "Maven",
		Cmd:  "mvn",
		Args: []string{"--version"},
		Parse: func(s string) string {
			lines := strings.Split(s, "\n")
			if len(lines) > 0 {
				line := lines[0]
				parts := strings.Fields(line)
				if len(parts) >= 3 {
					return parts[2]
				}
			}
			return strings.TrimSpace(s)
		},
	},
	{
		Name: "Python",
		Cmd:  "python3",
		Args: []string{"--version"},
		Parse: func(s string) string {
			return strings.TrimPrefix(strings.TrimSpace(s), "Python ")
		},
	},
	{
		Name: "pip",
		Cmd:  "pip3",
		Args: []string{"--version"},
		Parse: func(s string) string {
			parts := strings.Fields(strings.TrimSpace(s))
			if len(parts) >= 2 {
				return parts[1]
			}
			return strings.TrimSpace(s)
		},
	},
	{
		Name: "Terraform",
		Cmd:  "terraform",
		Args: []string{"version"},
		Parse: func(s string) string {
			lines := strings.Split(s, "\n")
			if len(lines) > 0 {
				line := lines[0]
				parts := strings.Fields(line)
				if len(parts) >= 2 {
					return strings.TrimPrefix(parts[1], "v")
				}
			}
			return strings.TrimSpace(s)
		},
	},
	{
		Name: "Packer",
		Cmd:  "packer",
		Args: []string{"version"},
		Parse: func(s string) string {
			lines := strings.Split(s, "\n")
			if len(lines) > 0 {
				line := lines[0]
				parts := strings.Fields(line)
				if len(parts) >= 2 {
					return strings.TrimPrefix(parts[1], "v")
				}
			}
			return strings.TrimSpace(s)
		},
	},
	{
		Name: "Docker",
		Cmd:  "docker",
		Args: []string{"--version"},
		Parse: func(s string) string {
			parts := strings.Fields(strings.TrimSpace(s))
			if len(parts) >= 3 {
				return parts[2]
			}
			return strings.TrimSpace(s)
		},
	},
	{
		Name: "Git",
		Cmd:  "git",
		Args: []string{"--version"},
		Parse: func(s string) string {
			parts := strings.Fields(strings.TrimSpace(s))
			if len(parts) >= 3 {
				return parts[2]
			}
			return strings.TrimSpace(s)
		},
	},
	{
		Name: "gh (GitHub CLI)",
		Cmd:  "gh",
		Args: []string{"--version"},
		Parse: func(s string) string {
			lines := strings.Split(s, "\n")
			if len(lines) > 0 {
				parts := strings.Fields(lines[0])
				if len(parts) >= 3 {
					return parts[2]
				}
			}
			return strings.TrimSpace(s)
		},
	},
	{
		Name: "AWS CLI",
		Cmd:  "aws",
		Args: []string{"--version"},
		Parse: func(s string) string {
			parts := strings.Fields(strings.TrimSpace(s))
			if len(parts) > 0 {
				return parts[0]
			}
			return strings.TrimSpace(s)
		},
	},
	{
		Name: "gcloud",
		Cmd:  "gcloud",
		Args: []string{"--version"},
		Parse: func(s string) string {
			lines := strings.Split(s, "\n")
			if len(lines) > 0 {
				parts := strings.Fields(lines[0])
				if len(parts) >= 4 {
					return parts[3]
				}
			}
			return strings.TrimSpace(s)
		},
	},
	{
		Name: "kubectl",
		Cmd:  "kubectl",
		Args: []string{"version", "--client", "--short"},
		Parse: func(s string) string {
			parts := strings.Fields(strings.TrimSpace(s))
			if len(parts) >= 2 {
				return strings.TrimPrefix(parts[1], "v")
			}
			return strings.TrimSpace(s)
		},
	},
	{
		Name: "Rust",
		Cmd:  "rustc",
		Args: []string{"--version"},
		Parse: func(s string) string {
			parts := strings.Fields(strings.TrimSpace(s))
			if len(parts) >= 2 {
				return parts[1]
			}
			return strings.TrimSpace(s)
		},
	},
	{
		Name: "Ruby",
		Cmd:  "ruby",
		Args: []string{"--version"},
		Parse: func(s string) string {
			parts := strings.Fields(strings.TrimSpace(s))
			if len(parts) >= 2 {
				return parts[1]
			}
			return strings.TrimSpace(s)
		},
	},
	{
		Name: "PHP",
		Cmd:  "php",
		Args: []string{"--version"},
		Parse: func(s string) string {
			lines := strings.Split(s, "\n")
			if len(lines) > 0 {
				parts := strings.Fields(lines[0])
				if len(parts) >= 2 {
					return parts[1]
				}
			}
			return strings.TrimSpace(s)
		},
	},
}

// ListVersions displays versions of installed development tools
func ListVersions() error {
	fmt.Println("🔧 Development Tools Version Check")
	fmt.Println()

	var tools []ToolVersion

	for _, tool := range toolsToCheck {
		version := checkToolVersion(tool.Cmd, tool.Args, tool.Parse)
		tools = append(tools, ToolVersion{
			Name:    tool.Name,
			Version: version,
			Status:  getStatus(version),
		})
	}

	// Display results grouped by status
	fmt.Println("✅ Installed:")
	for _, tool := range tools {
		if tool.Status == "installed" {
			fmt.Printf("  %-20s %s\n", tool.Name, tool.Version)
		}
	}

	notInstalled := 0
	for _, tool := range tools {
		if tool.Status == "not found" {
			notInstalled++
		}
	}

	if notInstalled > 0 {
		fmt.Println()
		fmt.Printf("❌ Not Installed (%d):\n", notInstalled)
		for _, tool := range tools {
			if tool.Status == "not found" {
				fmt.Printf("  %s\n", tool.Name)
			}
		}
	}

	fmt.Println()
	installed := len(tools) - notInstalled
	fmt.Printf("Summary: %d/%d tools installed\n", installed, len(tools))

	return nil
}

func checkToolVersion(cmd string, args []string, parse func(string) string) string {
	command := exec.Command(cmd, args...)
	output, err := command.CombinedOutput()
	if err != nil {
		Logf("Tool %s not found or error: %v", cmd, err)
		return "not installed"
	}

	outputStr := string(output)
	if outputStr == "" {
		return "installed (no version)"
	}

	return parse(outputStr)
}

func getStatus(version string) string {
	if version == "not installed" {
		return "not found"
	}
	return "installed"
}
