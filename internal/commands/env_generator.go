// internal/commands/env_generator.go
package commands

import (
	"os"
	"path/filepath"
)

// GenerateEnvFile creates a .env.example file for the project
func GenerateEnvFile(projectPath, projectType string) error {
	envPath := filepath.Join(projectPath, ".env.example")

	var envContent string
	switch projectType {
	case "golang":
		envContent = generateGoEnv()
	case "spring":
		envContent = generateSpringEnv()
	case "nextjs":
		envContent = generateNextJsEnv()
	case "terraform":
		envContent = generateTerraformEnv()
	default:
		envContent = generateGenericEnv()
	}

	Logf("Creating .env.example at %s", envPath)
	return os.WriteFile(envPath, []byte(envContent), 0644)
}

func generateGoEnv() string {
	return `# Go Application Environment Variables

# Server Configuration
SERVER_PORT=8080
SERVER_HOST=0.0.0.0
ENV=development

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_NAME=myapp
DB_USER=postgres
DB_PASSWORD=postgres

# Logging
LOG_LEVEL=info

# API Keys
API_KEY=your-api-key-here

# CORS
CORS_ALLOWED_ORIGINS=http://localhost:3000,http://localhost:8000
`
}

func generateSpringEnv() string {
	return `# Spring Boot Application Environment Variables

# Server Configuration
SERVER_PORT=8080
SERVER_SERVLET_CONTEXT_PATH=/api

# Database Configuration
SPRING_DATASOURCE_URL=jdbc:mysql://localhost:3306/myapp
SPRING_DATASOURCE_USERNAME=root
SPRING_DATASOURCE_PASSWORD=password
SPRING_JPA_HIBERNATE_DDL_AUTO=update

# Application Profile
SPRING_PROFILES_ACTIVE=dev

# Logging
LOGGING_LEVEL_ROOT=INFO
LOGGING_LEVEL_COM_EXAMPLE=DEBUG

# API Keys
API_KEY=your-api-key-here

# CORS
CORS_ALLOWED_ORIGINS=http://localhost:3000
`
}

func generateNextJsEnv() string {
	return `# Next.js Application Environment Variables

# API Configuration
NEXT_PUBLIC_API_URL=http://localhost:3000
API_SECRET_KEY=your-secret-key

# Database (if using)
DATABASE_URL=postgresql://user:password@localhost:5432/myapp

# Authentication
NEXTAUTH_SECRET=your-nextauth-secret
NEXTAUTH_URL=http://localhost:3000

# Third-party APIs
NEXT_PUBLIC_STRIPE_PUBLISHABLE_KEY=pk_test_xxx
STRIPE_SECRET_KEY=sk_test_xxx

# Environment
NODE_ENV=development
`
}

func generateTerraformEnv() string {
	return `# Terraform Environment Variables

# AWS Configuration
AWS_REGION=us-west-2
AWS_ACCESS_KEY_ID=your-access-key
AWS_SECRET_ACCESS_KEY=your-secret-key

# Terraform Settings
TF_VAR_environment=dev
TF_VAR_project_name=myapp
TF_VAR_availability_zones=us-west-2a,us-west-2b

# Tagging
TF_VAR_tags={"Environment":"dev","ManagedBy":"terraform"}
`
}

func generateGenericEnv() string {
	return `# Application Environment Variables

# Server Configuration
SERVER_PORT=8000
ENV=development

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_NAME=myapp
DB_USER=admin
DB_PASSWORD=password

# Logging
LOG_LEVEL=info
`
}
