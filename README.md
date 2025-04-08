# devtool

A command-line utility for developers to quickly set up projects and integrate with various services.

## Features

- Create GitHub repositories
- Create Confluence pages
- Initialize project templates for:
  - Golang
  - Spring
  - Next.js
  - Terraform
- Fetch Jira ticket information

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/devtool.git
cd devtool

# Build the tool
go build -o devtool ./cmd/devtool

# Install to your $GOPATH/bin (optional)
go install ./cmd/devtool
```

## Usage

### GitHub Integration

```bash
# Create a GitHub repository
./devtool github create my-new-repo
```

### Confluence Integration

```bash
# Create a Confluence page
./devtool confluence create page
```

### Project Templates

```bash
# Create a Golang project
./devtool create base golang project

# Create a Spring project
./devtool create base spring project

# Create a Next.js project
./devtool create base nextjs project

# Create a Terraform project
./devtool create base terraform project
```

### Jira Integration

```bash
# Set up Jira environment variables (recommended to add to your .bashrc or .zshrc)
export JIRA_DOMAIN="your-company.atlassian.net"
export JIRA_EMAIL="your.email@company.com"
export JIRA_API_TOKEN="your-jira-api-token"
export JIRA_PROJECT="PROJ"  # Default project prefix (optional)

# Fetch a Jira ticket
./devtool get_jira_ticket PROJ-123

# Or with just the number (uses JIRA_PROJECT prefix)
./devtool get_jira_ticket 123
```

## Project Structure

```
.
├── api/            # REST API and service definitions
├── cmd/            # Main applications
│   └── devtool/    # Main CLI application
├── configs/        # Configuration files
├── docs/           # Documentation files
├── internal/       # Private application code
│   └── commands/   # Command implementations
└── pkg/            # Public libraries that can be used by external applications
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.