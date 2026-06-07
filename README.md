# devtool

A powerful command-line utility for developers to quickly set up projects, integrate with services, and scaffold production-ready applications.

## Features

### 🚀 Interactive Project Wizard
- **`devtool init`** - Step-by-step guided project creation
- Prompts for project name, type, and configuration
- Automatic git initialization with initial commit
- Optional Docker, environment files, and pre-commit hooks setup

### 📦 Project Templates & Creation
- **`devtool list`** - Browse 8 built-in project templates
- **`devtool pull <template>`** - Clone template repositories
- **`devtool create base <type> project`** - Create base project structure
- Supports: Golang, Spring, Next.js, Terraform

### 🛠️ Development Environment
- **`devtool versions`** - Check installed tool versions
- Detects 20+ common development tools
- Shows version info with installation status
- Useful for verifying environment compatibility

### 🔄 CI/CD Pipelines
- **`devtool ci <provider> <language>`** - Generate CI/CD workflows
- GitHub Actions (`.github/workflows/ci.yml`)
- GitLab CI (`.gitlab-ci.yml`)
- Includes: linting, testing, building, coverage reports
- Languages: Golang, Spring, Next.js, Terraform

### 🧪 Testing Frameworks
- **`devtool test <language>`** - Setup testing frameworks
- Go: Built-in testing with benchmarks
- Spring: JUnit tests with Maven configuration
- Next.js: Jest with TypeScript support
- Python: pytest with fixtures
- Includes example tests and configuration files

### 📚 Documentation Generation
- **`devtool docs <type> [language]`** - Generate project documentation
- README: Project overview, features, installation, setup
- API: REST endpoint documentation with examples
- ADR: Architecture Decision Records structure
- Setup: Development, testing, deployment guides
- All: Generate complete documentation suite

### ⚙️ Configuration Management
- **`devtool config`** - Manage ~/.devtool.yaml
  - `show` - Display current configuration
  - `set <key> <value>` - Set individual values
  - `init` - Interactive setup wizard
  - `reset` - Delete configuration file
- Save API tokens, git defaults, and service credentials
- Auto-fallback to environment variables

### 🐳 Docker & Development Environment
Auto-generates for each project type:
- **Dockerfile** - Multi-stage, optimized builds
- **docker-compose.yml** - Local development with services
- **.dockerignore** - Optimized file exclusions

### 📝 Environment Files
- **`.env.example`** - Language/framework-specific templates
- Pre-configured variables for databases, APIs, logging
- Copy to `.env` for local development

### 🔧 Pre-commit Hooks
- **`.pre-commit-config.yaml`** - Auto-generated linting configuration
- Language-specific hooks:
  - **Go**: gofmt, vet, golangci-lint
  - **Java**: formatting, checkstyle
  - **JavaScript**: ESLint, Prettier
  - **Terraform**: fmt, validate, tflint

### 🔗 Service Integrations
- **GitHub** - Create repositories via `gh` CLI
- **Jira** - Fetch ticket information
- **Confluence** - Create pages programmatically

## Installation

### Download Binary
```bash
# Get the latest binary
wget https://github.com/yourusername/devtool/releases/download/latest/devtool
chmod +x devtool

# Install to system path
sudo mv devtool /usr/local/bin/
```

### Build from Source
```bash
git clone https://github.com/yourusername/devtool.git
cd devtool
go build -o devtool ./cmd/devtool
sudo mv devtool /usr/local/bin/
```

## Quick Start

### Create a New Project (Recommended)
```bash
devtool init
```

Follow the interactive prompts to:
1. Enter project name
2. Select project type
3. Choose optional features (Docker, .env, pre-commit hooks)
4. Review summary and confirm

**Example Output:**
```
🚀 Welcome to devtool project initializer!

Enter project name [my-project]: my-api
Select project type:
  1) golang
  2) spring
  3) nextjs
  4) terraform
Select option (1-4): 1

Project path [my-api]: my-api
Initialize git repository? [Y/n]: y
Create GitHub repository? [y/N]: n

Generate Docker files? [Y/n]: y
Generate .env.example? [Y/n]: y
Setup pre-commit hooks? [Y/n]: y

📋 Project Summary:
  Name:       my-api
  Type:       golang
  Path:       my-api
  Git:        true
  Docker:     true
  .env.example: true
  Pre-commit: true

Proceed with project creation? [Y/n]: y

✓ Golang project created successfully
✓ Generated .env.example
✓ Generated Dockerfile and docker-compose.yml
✓ Generated .pre-commit-config.yaml
✓ Git repository initialized

✨ Project 'my-api' created successfully at my-api

🚀 Next steps:
  cd my-api
  go mod tidy
  go run cmd/main.go
```

### Use Project Templates
```bash
# List available templates
devtool list

# Clone a template
devtool pull golang-api my-api
cd my-api
```

### Manage Configuration
```bash
# View configuration
devtool config show

# Set individual values
devtool config set default_language golang
devtool config set default_git_name "John Doe"
devtool config set default_git_email "john@example.com"

# Interactive setup
devtool config init

# Reset configuration
devtool config reset
```

## Project Types

### Golang
- Project structure: `cmd/`, `internal/`, `pkg/`, `api/`, `configs/`, `docs/`
- go.mod file with proper naming
- README with setup instructions
- Docker multi-stage build
- Environment variables for server, database, logging
- Pre-commit: gofmt, vet, golangci-lint

### Spring Boot
- Maven project structure
- Application.java with @SpringBootApplication
- pom.xml with Spring Boot dependencies
- application.properties configuration
- Docker with Maven layer caching
- Environment variables for database, server, logging
- Pre-commit: Java formatting, checkstyle

### Next.js
- Next.js app directory structure
- package.json with dependencies
- Tailwind CSS support ready
- Docker multi-stage build for optimized size
- Environment variables for API URLs, auth
- Pre-commit: ESLint, Prettier

### Terraform
- Modular structure: `modules/`, `environments/dev/`, `environments/prod/`
- main.tf, variables.tf, outputs.tf
- Example S3 bucket module
- Backend configuration templates
- Docker container for Terraform
- Pre-commit: terraform fmt, validate, tflint

## Docker Support

Generated `docker-compose.yml` includes appropriate services:

**Golang/Next.js:**
```yaml
services:
  app:
    build: .
    ports: ["8080:8080"]
  db:
    image: postgres:15-alpine
    environment: [POSTGRES_DB=myapp]
```

**Spring Boot:**
```yaml
services:
  app:
    build: .
    ports: ["8080:8080"]
  db:
    image: mysql:8.0
    environment: [MYSQL_DATABASE=myapp]
```

**Usage:**
```bash
# Start services
docker-compose up

# Stop services
docker-compose down

# View logs
docker-compose logs -f app
```

## Environment Files

Auto-generated `.env.example` with framework-specific variables:

**Golang Example:**
```bash
SERVER_PORT=8080
DB_HOST=localhost
DB_USER=postgres
LOG_LEVEL=info
```

**Setup:**
```bash
cp .env.example .env
# Edit .env with your values
```

## Pre-commit Hooks

Automatically configured for code quality:

```bash
# Install hooks
pre-commit install

# Run manually
pre-commit run --all-files

# Bypass hooks (not recommended)
git commit --no-verify
```

## Service Integrations

### GitHub
```bash
devtool github create my-repo
# Uses 'gh' CLI (https://cli.github.com)
```

### Jira
Set environment variables:
```bash
export JIRA_DOMAIN=your-company.atlassian.net
export JIRA_EMAIL=your.email@company.com
export JIRA_API_TOKEN=your-api-token

# Fetch a ticket
devtool get_jira_ticket PROJ-123
devtool get_jira_ticket 123  # Uses JIRA_PROJECT prefix
```

### Confluence
Set environment variables:
```bash
export CONFLUENCE_URL=https://company.atlassian.net/wiki
export CONFLUENCE_USER=your.email@company.com
export CONFLUENCE_TOKEN=your-api-token
export CONFLUENCE_SPACE=DOCS  # Optional

# Create a page (interactive)
devtool confluence create page
```

## Configuration File

Configuration is stored in `~/.devtool.yaml`:

```yaml
default_language: golang
default_git_email: john@example.com
default_git_name: John Doe
github_token: ghp_xxxxx
jira_domain: company.atlassian.net
jira_email: user@company.com
jira_token: xxxxx
confluence_url: https://company.atlassian.net/wiki
confluence_user: user@company.com
confluence_token: xxxxx
```

## Flags

- `-v, --verbose` - Enable debug logging for troubleshooting

## Commands Reference

```bash
# Project creation
devtool init                          # Interactive wizard (recommended)
devtool create base golang project    # Basic project structure
devtool create base spring project
devtool create base nextjs project
devtool create base terraform project

# Templates
devtool list                          # List templates
devtool pull <template-name> [path]  # Clone template

# Configuration
devtool config show                   # Display config
devtool config set <key> <value>      # Set value
devtool config init                   # Setup wizard
devtool config reset                  # Delete config

# Environment Check
devtool versions                      # Show installed tool versions

# CI/CD & Testing
devtool ci github-actions golang      # GitHub Actions for Go
devtool ci gitlab-ci spring           # GitLab CI for Spring
devtool ci github-actions nextjs      # GitHub Actions for Next.js
devtool test golang                   # Setup Go testing
devtool test spring                   # Setup Spring testing
devtool test nextjs                   # Setup Jest for Next.js

# Documentation
devtool docs readme golang            # Generate README
devtool docs api spring               # Generate API docs
devtool docs adr                      # Setup ADRs
devtool docs setup nextjs             # Generate setup guide
devtool docs all golang               # Generate all docs

# Service Integration
devtool github create <repo>          # Create GitHub repo
devtool confluence create page        # Create Confluence page
devtool get_jira_ticket <ticket>      # Fetch Jira ticket

# Help
devtool help                          # Show help
devtool -v help                       # Verbose help
```

## Tool Version Detection

The `devtool versions` command detects and displays versions for:

**Languages & Runtimes:**
- Node.js, Python, Go, Java, Ruby, Rust, PHP

**Package Managers:**
- npm, pnpm, yarn, pip, Maven

**DevOps & Cloud:**
- Docker, Terraform, Packer, kubectl
- AWS CLI, gcloud

**Version Control & Tools:**
- Git, GitHub CLI (gh)

**Example Output:**
```bash
$ devtool versions

🔧 Development Tools Version Check

✅ Installed:
  Node.js              18.19.1
  npm                  10.2.4
  Go                   1.21.0
  Python               3.11.5
  Docker               25.0.0
  Git                  2.42.0
  Terraform            1.6.0
  kubectl              1.28.0

❌ Not Installed (3):
  Rust
  PHP
  Packer

Summary: 17/20 tools installed
```

Use this to verify your environment has the tools needed for a project before starting.

## CI/CD Pipeline Generation

The `devtool ci` command generates production-ready CI/CD pipelines for your projects.

### GitHub Actions
```bash
devtool ci github-actions golang
devtool ci github-actions spring
devtool ci github-actions nextjs
devtool ci github-actions terraform
```

**Generated file:** `.github/workflows/ci.yml`

**Includes:**
- Build and test stages
- Dependency caching
- Code linting
- Coverage reporting
- Multi-version testing (where applicable)

**Example workflow for Go:**
```yaml
name: Go CI
on: [push, pull_request]
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
      - run: go build ./...
      - run: go test -v ./...
      - uses: codecov/codecov-action@v3
```

### GitLab CI
```bash
devtool ci gitlab-ci golang
devtool ci gitlab-ci spring
devtool ci gitlab-ci nextjs
devtool ci gitlab-ci terraform
```

**Generated file:** `.gitlab-ci.yml`

**Includes:**
- Parallel job stages
- Docker image specifications
- Artifact management
- Coverage reports
- Backend configuration templates

## Testing Framework Setup

The `devtool test` command scaffolds testing frameworks and example tests.

### Go Testing
```bash
devtool test golang [path]
```

**Creates:**
- `internal/commands/example_test.go` - Example unit and benchmark tests
- Uses Go's built-in testing package

**Run tests:**
```bash
go test -v ./...          # Verbose
go test -cover ./...      # With coverage
go test -bench ./...      # Benchmarks
```

### Spring Boot Testing
```bash
devtool test spring [path]
```

**Creates:**
- `src/test/java/com/example/demo/ExampleTest.java` - JUnit 5 tests
- `src/test/resources/` directory for test configs
- Example test with assertions

**Add to pom.xml:**
```xml
<dependency>
  <groupId>org.springframework.boot</groupId>
  <artifactId>spring-boot-starter-test</artifactId>
  <scope>test</scope>
</dependency>
```

**Run tests:**
```bash
mvn test                  # Run all tests
mvn test -Dtest=*Test    # Run specific test
mvn jacoco:report         # Coverage report
```

### JavaScript/TypeScript Testing
```bash
devtool test nextjs [path]
```

**Creates:**
- `__tests__/example.test.tsx` - Jest test examples
- `jest.config.js` - Jest configuration with TypeScript support
- Coverage thresholds and rules

**Install dependencies:**
```bash
npm install --save-dev jest @jest/globals ts-jest @types/jest
```

**Run tests:**
```bash
npm test                  # Run all tests
npm test -- --watch      # Watch mode
npm test -- --coverage   # With coverage report
```

### Python Testing
```bash
devtool test python [path]
```

**Creates:**
- `tests/test_example.py` - pytest test examples with fixtures
- `pytest.ini` - pytest configuration
- Markers for test categorization

**Install dependencies:**
```bash
pip install pytest pytest-cov
```

**Run tests:**
```bash
pytest                    # Run all tests
pytest -v               # Verbose
pytest --cov            # With coverage
pytest -m slow          # Run marked tests
`

## Documentation Generation

The `devtool docs` command scaffolds professional project documentation.

### README Templates
```bash
devtool docs readme golang
devtool docs readme spring
devtool docs readme nextjs
```

**Generates:** `README.md` with:
- Project overview and features
- Installation instructions (source, Docker)
- Quick start guide
- Configuration details
- Development setup
- Testing instructions
- Deployment guidelines
- Contributing guidelines
- Troubleshooting section
- Support information

### API Documentation
```bash
devtool docs api [language]
```

**Generates:** `docs/API.md` with:
- Base URL and endpoints
- Authentication details
- Error handling format
- Example GET/POST/PUT/DELETE endpoints
- Request/response examples
- Rate limiting information
- Versioning strategy
- Changelog template

### Architecture Decision Records
```bash
devtool docs adr
```

**Generates:**
- `docs/adr/README.md` - ADR index and guidelines
- `docs/adr/0001-project-setup.md` - Template ADR

Each ADR documents:
- Date and status
- Context and problem
- Decision made
- Consequences (positive/negative)
- Alternatives considered
- References

### Setup Guides
```bash
devtool docs setup golang
devtool docs setup spring
devtool docs setup nextjs
```

**Generates:** `docs/SETUP.md` with:
- Prerequisites and installation
- Development environment setup
- Running locally (dev mode, with hot reload)
- Testing procedures
- Database setup (if applicable)
- Environment configuration
- Building for production
- Deployment instructions
- Troubleshooting guide
- Performance monitoring

### Generate All Documentation
```bash
devtool docs all golang
```

Generates all documentation at once:
- README.md
- docs/API.md
- docs/SETUP.md
- docs/adr/README.md
- docs/adr/0001-project-setup.md

## Documentation Best Practices

1. **Keep it current** - Update docs when code changes
2. **Be specific** - Use actual commands, not generic instructions
3. **Include examples** - Show real requests/responses for APIs
4. **Document why** - Explain decisions, not just what was chosen
5. **Add troubleshooting** - Include common issues and solutions
6. **Use consistent formatting** - Follow the generated templates
7. **Link between docs** - Cross-reference related sections

## Customization

The generated templates are starting points. Customize them with:
- Your actual API endpoints and responses
- Specific database and service configurations
- Team-specific development practices
- Actual deployment targets and procedures
- Real troubleshooting from your experience
- Project-specific features and capabilities
```

## Project Structure

```
.
├── cmd/devtool/          # CLI application entry point
├── internal/
│   ├── commands/         # Command implementations
│   └── config/           # Configuration management
├── pkg/                  # Public libraries (future)
├── api/                  # API definitions (future)
├── go.mod, go.sum        # Dependencies
└── README.md
```

## Requirements

- **Go 1.20+** - For building from source
- **Git** - For repository initialization and template cloning
- **gh CLI** - For GitHub integration (optional)
- **Docker/Docker Compose** - For containerized development (optional)
- **pre-commit** - For git hooks setup (optional)

## Troubleshooting

### "gh CLI not found"
Install GitHub CLI: https://cli.github.com

### "git not found"
Install Git: https://git-scm.com

### "Failed to clone template"
Ensure you have git installed and GitHub access.

### Missing Jira credentials
```bash
export JIRA_DOMAIN=company.atlassian.net
export JIRA_EMAIL=email@company.com
export JIRA_API_TOKEN=your-token
```

Or use:
```bash
devtool config init  # Interactive setup
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - See LICENSE file for details

## Roadmap

- [ ] Shell completion (bash/zsh)
- [ ] Plugin system for custom templates
- [ ] GitLab integration
- [ ] Database migration templates
- [ ] Monorepo support
- [ ] Team configuration sharing
