// internal/commands/docs_generator.go
package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

// DocsCommand handles documentation generation
func DocsCommand(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: devtool docs <type> [language] [path]")
	}

	docType := args[0]
	language := "golang"
	projectPath := "."

	if len(args) > 1 {
		language = args[1]
	}
	if len(args) > 2 {
		projectPath = args[2]
	}

	switch docType {
	case "readme":
		return generateREADME(projectPath, language)
	case "api":
		return generateAPIDocumentation(projectPath, language)
	case "adr":
		return generateADRSetup(projectPath)
	case "setup":
		return generateSetupGuide(projectPath, language)
	case "all":
		if err := generateREADME(projectPath, language); err != nil {
			return err
		}
		if err := generateADRSetup(projectPath); err != nil {
			return err
		}
		if err := generateSetupGuide(projectPath, language); err != nil {
			return err
		}
		return generateAPIDocumentation(projectPath, language)
	default:
		return fmt.Errorf("unknown doc type: %s (readme, api, adr, setup, all)", docType)
	}
}

func generateREADME(projectPath, language string) error {
	readmePath := filepath.Join(projectPath, "README.md")

	// Check if README exists
	if _, err := os.Stat(readmePath); err == nil {
		fmt.Printf("⚠️  README.md already exists at %s\n", readmePath)
		return nil
	}

	content := getREADMETemplate(language)

	Logf("Creating README.md at %s", readmePath)
	if err := os.WriteFile(readmePath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write README: %v", err)
	}

	fmt.Printf("✓ Generated README.md for %s\n", language)
	fmt.Printf("  Location: %s\n", readmePath)
	fmt.Println("  Next steps:")
	fmt.Println("    - Customize with your project details")
	fmt.Println("    - Add badges and links")
	fmt.Println("    - Keep it up-to-date")

	return nil
}

func generateAPIDocumentation(projectPath, language string) error {
	docsDir := filepath.Join(projectPath, "docs")
	if err := os.MkdirAll(docsDir, 0755); err != nil {
		return fmt.Errorf("failed to create docs directory: %v", err)
	}

	content := getAPIDocTemplate(language)
	apiDocPath := filepath.Join(docsDir, "API.md")

	Logf("Creating API documentation at %s", apiDocPath)
	if err := os.WriteFile(apiDocPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write API docs: %v", err)
	}

	fmt.Printf("✓ Generated API documentation for %s\n", language)
	fmt.Printf("  Location: %s\n", apiDocPath)
	fmt.Println("  Next steps:")
	fmt.Println("    - Document your endpoints/functions")
	fmt.Println("    - Add request/response examples")
	fmt.Println("    - Include authentication details")

	return nil
}

func generateADRSetup(projectPath string) error {
	adrDir := filepath.Join(projectPath, "docs", "adr")
	if err := os.MkdirAll(adrDir, 0755); err != nil {
		return fmt.Errorf("failed to create ADR directory: %v", err)
	}

	// Create ADR index
	indexPath := filepath.Join(adrDir, "README.md")
	indexContent := getADRIndexTemplate()

	if err := os.WriteFile(indexPath, []byte(indexContent), 0644); err != nil {
		return fmt.Errorf("failed to write ADR index: %v", err)
	}

	// Create example ADR
	examplePath := filepath.Join(adrDir, "0001-project-setup.md")
	exampleContent := getADRTemplate()

	if err := os.WriteFile(examplePath, []byte(exampleContent), 0644); err != nil {
		return fmt.Errorf("failed to write example ADR: %v", err)
	}

	fmt.Printf("✓ Generated ADR setup\n")
	fmt.Printf("  Location: %s\n", adrDir)
	fmt.Println("  Files created:")
	fmt.Println("    - docs/adr/README.md (index)")
	fmt.Println("    - docs/adr/0001-project-setup.md (template)")
	fmt.Println("  Next steps:")
	fmt.Println("    - Create new ADRs for important decisions")
	fmt.Println("    - Follow the template format")
	fmt.Println("    - Keep them concise and focused")

	return nil
}

func generateSetupGuide(projectPath, language string) error {
	docsDir := filepath.Join(projectPath, "docs")
	if err := os.MkdirAll(docsDir, 0755); err != nil {
		return fmt.Errorf("failed to create docs directory: %v", err)
	}

	content := getSetupGuideTemplate(language)
	setupPath := filepath.Join(docsDir, "SETUP.md")

	Logf("Creating setup guide at %s", setupPath)
	if err := os.WriteFile(setupPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write setup guide: %v", err)
	}

	fmt.Printf("✓ Generated setup guide for %s\n", language)
	fmt.Printf("  Location: %s\n", setupPath)
	fmt.Println("  Next steps:")
	fmt.Println("    - Add environment-specific instructions")
	fmt.Println("    - Include troubleshooting section")
	fmt.Println("    - Keep updated as dependencies change")

	return nil
}

func getREADMETemplate(language string) string {
	switch language {
	case "golang":
		return `# Project Name

Short description of your Go project.

## Overview

Detailed explanation of what this project does and why it exists.

## Features

- Feature 1
- Feature 2
- Feature 3

## Prerequisites

- Go 1.20 or later
- Git
- [Other requirements]

## Installation

From source:

    git clone https://github.com/username/project.git
    cd project
    go mod download
    go build -o project ./cmd/project

Using Docker:

    docker build -t project .
    docker run project

## Quick Start

    # Basic usage
    ./project --help

    # Run with config
    ./project --config config.yml

    # Run tests
    go test -v ./...

## Configuration

Configuration can be provided via:
- Command-line flags
- Environment variables
- Configuration file (YAML/JSON)

See docs/SETUP.md for configuration details.

## Development

Setup development environment:

    go mod tidy
    go test -cover ./...
    golangci-lint run
    go fmt ./...

## API Reference

See docs/API.md for detailed API documentation.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit changes
4. Push to branch
5. Open a Pull Request

## Testing

    go test -v ./...
    go test -cover ./...
    go test -race ./...

## Deployment

See docs/SETUP.md for deployment instructions.

## License

MIT License

## Support

- Email: support@example.com
- Documentation: docs/
`

	case "spring":
		return `# Project Name

Spring Boot REST API.

## Overview

Brief description of the application.

## Features

- RESTful API endpoints
- Database integration
- Authentication & Authorization
- [Other features]

## Prerequisites

- Java 17+
- Maven 3.8+
- MySQL 8.0+
- Docker (optional)

## Installation

Clone and build:

    git clone https://github.com/username/project.git
    cd project
    mvn clean install

Using Docker Compose:

    docker-compose up

## Running

    mvn spring-boot:run

## Configuration

Application properties in src/main/resources/application.yml:

    spring:
      datasource:
        url: jdbc:mysql://localhost:3306/myapp
        username: root
        password: password
      jpa:
        hibernate:
          ddl-auto: update

## Testing

    mvn test
    mvn verify
    mvn jacoco:report

## API Endpoints

See docs/API.md for complete API documentation.

Example endpoints:
- GET /api/users - Get all users
- POST /api/users - Create user
- GET /api/users/{id} - Get user by ID
- PUT /api/users/{id} - Update user
- DELETE /api/users/{id} - Delete user

## Deployment

See docs/SETUP.md for deployment instructions.

## License

MIT License
`

	case "nextjs":
		return `# Project Name

Next.js fullstack application.

## Overview

Brief description of your Next.js project.

## Features

- Server-side rendering (SSR)
- Static generation (SSG)
- API routes
- TypeScript support
- Tailwind CSS

## Prerequisites

- Node.js 18+
- npm or yarn or pnpm
- Git

## Getting Started

Installation:

    git clone https://github.com/username/project.git
    cd project
    npm install

Development:

    npm run dev

Open http://localhost:3000

## Building

    npm run build
    npm start

## Testing

    npm test
    npm test -- --coverage

## Project Structure

    app/                  - App router
    components/           - Reusable components
    lib/                  - Utilities
    public/               - Static assets
    __tests__/            - Tests
    .env.example          - Environment template

## Environment Setup

    cp .env.example .env.local

Configure your variables in .env.local.

## API Routes

See docs/API.md for API documentation.

Example routes in app/api/:
- GET /api/users
- POST /api/users
- GET /api/users/[id]
- PUT /api/users/[id]
- DELETE /api/users/[id]

## Deployment

Docker:

    docker build -t project .
    docker run -p 3000:3000 project

See docs/SETUP.md for detailed deployment instructions.

## License

MIT License
`

	default:
		return getREADMETemplate("golang")
	}
}

func getAPIDocTemplate(language string) string {
	return `# API Documentation

## Overview

Description of your API.

## Base URL

    http://localhost:8080/api
    https://api.example.com

## Authentication

Authorization: Bearer <token>

## Error Handling

All errors follow this format:

    {
      "error": "error_code",
      "message": "Human readable message",
      "timestamp": "2024-01-01T00:00:00Z"
    }

## Endpoints

### Users

GET /users

Get all users

Query Parameters:
- page (optional): Page number (default: 1)
- limit (optional): Items per page (default: 20)

Response:

    {
      "data": [
        {
          "id": "123",
          "name": "John Doe",
          "email": "john@example.com",
          "created_at": "2024-01-01T00:00:00Z"
        }
      ],
      "pagination": {
        "page": 1,
        "limit": 20,
        "total": 100
      }
    }

GET /users/:id

Get user by ID

Response:

    {
      "id": "123",
      "name": "John Doe",
      "email": "john@example.com"
    }

POST /users

Create user

Request Body:

    {
      "name": "Jane Doe",
      "email": "jane@example.com",
      "password": "securepassword"
    }

Response (201):

    {
      "id": "124",
      "name": "Jane Doe",
      "email": "jane@example.com"
    }

PUT /users/:id

Update user

DELETE /users/:id

Delete user (204 No Content)

## Rate Limiting

- Rate limit: 1000 requests per hour
- Check X-RateLimit-Remaining header

## Versioning

API versions in URL: /api/v1/, /api/v2/, etc.

## Changelog

### v1.0.0 (2024-01-01)
- Initial release

### v1.1.0 (2024-02-01)
- Added pagination
- Performance improvements
`
}

func getADRIndexTemplate() string {
	return `# Architecture Decision Records

An Architecture Decision Record (ADR) is a document that captures an important architectural decision made along with its context and consequences.

## ADRs

- ADR-0001: Project Setup Decision

## Creating New ADRs

1. Copy the template from 0001-project-setup.md
2. Increment the number (0002, 0003, etc.)
3. Document your decision following the template
4. Add link to this README

## Format

Each ADR should include:
- Title: Short phrase describing the decision
- Status: Proposed, Accepted, Deprecated, Superseded
- Context: What is the issue we're addressing?
- Decision: What is the decision we're making?
- Consequences: What becomes easier or more difficult?
- Alternatives: What alternatives were considered?
`
}

func getADRTemplate() string {
	return `# ADR-0001: Project Setup Decision

Date: 2024-01-01

Status: Accepted

## Context

When starting a new project, we need to decide on the fundamental architecture and technology choices.

## Decision

We have decided to use [technology/framework] for this project because it provides:
- Better performance for our use case
- Easier to maintain and scale
- Good community support and documentation
- Aligns with team expertise

## Consequences

### Positive
- Faster development velocity
- Better code quality
- Easier onboarding for new team members
- Excellent ecosystem of tools

### Negative
- Requires team training
- Smaller community compared to alternatives
- Licensing considerations

## Alternatives Considered

### Alternative 1: [Technology]
- Pros: faster startup, larger community
- Cons: steeper learning curve, less suitable for our needs

### Alternative 2: [Technology]
- Pros: more mature, proven track record
- Cons: slower performance, more expensive

## References

- Documentation: https://example.com
- Decision Process: https://example.com
`
}

func getSetupGuideTemplate(language string) string {
	switch language {
	case "golang":
		return `# Setup Guide

Complete setup instructions for development, testing, and deployment.

## Development Setup

Prerequisites:
- Go 1.20 or later
- Git
- Your favorite IDE

### 1. Clone Repository

    git clone https://github.com/username/project.git
    cd project

### 2. Install Dependencies

    go mod download
    go mod tidy

### 3. Verify Installation

    go build -o project ./cmd/project
    ./project --version

## Running Locally

Development mode:

    go run ./cmd/project

With hot reload (using air):

    go install github.com/cosmtrek/air@latest
    air

Using Docker Compose:

    docker-compose up

## Testing

    go test -v ./...
    go test -cover ./...
    go test -race ./...
    go test -coverprofile=coverage.out ./...

## Database Setup

    createdb myapp
    ./project migrate up
    ./project seed

## Environment Configuration

1. Copy .env.example:

    cp .env.example .env

2. Edit .env with your settings:

    DATABASE_URL=postgres://user:pass@localhost/myapp
    PORT=8080
    LOG_LEVEL=info

## Pre-commit Setup

    pre-commit install
    pre-commit run --all-files

## Building for Production

    go build -ldflags="-s -w" -o project ./cmd/project

With version:

    VERSION=$(git describe --tags)
    go build -ldflags="-X main.Version=$VERSION" -o project ./cmd/project

## Deployment

Docker:

    docker build -t myapp:latest .
    docker run -p 8080:8080 --env-file .env myapp:latest

Kubernetes:

    kubectl apply -f k8s/

## Troubleshooting

"go: command not found"
- Install Go from https://golang.org/doc/install

"failed to dial database"
- Check DATABASE_URL and ensure database is running

Port already in use
- Change PORT in .env or use: lsof -i :8080

Module not found
- Run: go get github.com/username/module@latest && go mod tidy
`

	default:
		return `# Setup Guide

## Development Setup

1. Install prerequisites
2. Clone repository
3. Install dependencies
4. Configure environment variables
5. Run locally
6. Run tests

## Prerequisites

- Required tools and versions
- System requirements

## Installation

    git clone <repo>
    cd project
    npm install

## Running

    npm run dev

## Testing

    npm test

## Deployment

See deployment instructions above.

## Troubleshooting

Common issues and solutions.
`
	}
}
