// internal/commands/precommit_generator.go
package commands

import (
	"os"
	"path/filepath"
)

// GeneratePrecommitConfig creates .pre-commit-config.yaml
func GeneratePrecommitConfig(projectPath, projectType string) error {
	config := generatePrecommitConfig(projectType)
	configPath := filepath.Join(projectPath, ".pre-commit-config.yaml")

	Logf("Creating .pre-commit-config.yaml at %s", configPath)
	return os.WriteFile(configPath, []byte(config), 0644)
}

func generatePrecommitConfig(projectType string) string {
	switch projectType {
	case "golang":
		return `# Pre-commit configuration for Golang project
repos:
  # General
  - repo: https://github.com/pre-commit/pre-commit-hooks
    rev: v4.4.0
    hooks:
      - id: trailing-whitespace
      - id: end-of-file-fixer
      - id: check-yaml
      - id: check-added-large-files
      - id: check-merge-conflict

  # Go linting
  - repo: https://github.com/golangci/golangci-lint
    rev: v1.52.0
    hooks:
      - id: golangci-lint
        args: [--timeout=5m]

  # Go fmt
  - repo: https://github.com/dnephin/pre-commit-golang
    rev: v0.5.1
    hooks:
      - id: go-fmt
      - id: go-vet
`

	case "spring":
		return `# Pre-commit configuration for Spring project
repos:
  # General
  - repo: https://github.com/pre-commit/pre-commit-hooks
    rev: v4.4.0
    hooks:
      - id: trailing-whitespace
      - id: end-of-file-fixer
      - id: check-yaml
      - id: check-added-large-files
      - id: check-merge-conflict

  # Java formatting
  - repo: https://github.com/macisamuele/language-formatters-pre-commit-hooks
    rev: v2.9.0
    hooks:
      - id: pretty-format-java
        args: [--autofix]

  # Checkstyle
  - repo: local
    hooks:
      - id: maven-checkstyle
        name: maven-checkstyle
        entry: mvn checkstyle:check
        language: system
        types: [java]
        pass_filenames: false
        stages: [commit]
`

	case "nextjs":
		return `# Pre-commit configuration for Next.js project
repos:
  # General
  - repo: https://github.com/pre-commit/pre-commit-hooks
    rev: v4.4.0
    hooks:
      - id: trailing-whitespace
      - id: end-of-file-fixer
      - id: check-yaml
      - id: check-json
      - id: check-added-large-files
      - id: check-merge-conflict

  # ESLint
  - repo: https://github.com/pre-commit/mirrors-eslint
    rev: v8.40.0
    hooks:
      - id: eslint
        types: [javascript, jsx, typescript, tsx]
        args: [--fix]
        additional_dependencies:
          - eslint
          - eslint-config-next
          - eslint-plugin-react

  # Prettier
  - repo: https://github.com/pre-commit/mirrors-prettier
    rev: v3.0.0-alpha.9-for-vscode
    hooks:
      - id: prettier
        types_or: [javascript, jsx, typescript, tsx, json, yaml, markdown]
        args: [--write]
`

	case "terraform":
		return `# Pre-commit configuration for Terraform project
repos:
  # General
  - repo: https://github.com/pre-commit/pre-commit-hooks
    rev: v4.4.0
    hooks:
      - id: trailing-whitespace
      - id: end-of-file-fixer
      - id: check-yaml
      - id: check-added-large-files
      - id: check-merge-conflict

  # Terraform
  - repo: https://github.com/antonbabenko/pre-commit-terraform
    rev: v1.77.0
    hooks:
      - id: terraform_fmt
      - id: terraform_validate
      - id: terraform_tflint
        args:
          - --args=--only=terraform_unused_required_providers
      - id: terraform_docs
`

	default:
		return `# Pre-commit configuration
repos:
  - repo: https://github.com/pre-commit/pre-commit-hooks
    rev: v4.4.0
    hooks:
      - id: trailing-whitespace
      - id: end-of-file-fixer
      - id: check-yaml
      - id: check-added-large-files
      - id: check-merge-conflict
`
	}
}
