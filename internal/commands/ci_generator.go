// internal/commands/ci_generator.go
package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

// CICommand handles CI/CD pipeline generation
func CICommand(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: devtool ci <provider> <language> [path]")
	}

	provider := args[0]
	language := args[1]
	projectPath := "."
	if len(args) > 2 {
		projectPath = args[2]
	}

	switch provider {
	case "github-actions":
		return generateGitHubActions(projectPath, language)
	case "gitlab-ci":
		return generateGitLabCI(projectPath, language)
	default:
		return fmt.Errorf("unknown CI provider: %s (supported: github-actions, gitlab-ci)", provider)
	}
}

func generateGitHubActions(projectPath, language string) error {
	workflowDir := filepath.Join(projectPath, ".github", "workflows")
	if err := os.MkdirAll(workflowDir, 0755); err != nil {
		return fmt.Errorf("failed to create workflow directory: %v", err)
	}

	content := getGitHubActionsTemplate(language)
	filename := filepath.Join(workflowDir, "ci.yml")

	Logf("Creating GitHub Actions workflow at %s", filename)
	if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write workflow file: %v", err)
	}

	fmt.Printf("✓ Generated GitHub Actions workflow for %s\n", language)
	fmt.Printf("  Location: .github/workflows/ci.yml\n")
	fmt.Println("  Next steps:")
	fmt.Println("    - Customize the workflow as needed")
	fmt.Println("    - Push to GitHub to enable CI/CD")
	fmt.Println("    - View runs at: https://github.com/your-repo/actions")

	return nil
}

func generateGitLabCI(projectPath, language string) error {
	filename := filepath.Join(projectPath, ".gitlab-ci.yml")

	content := getGitLabCITemplate(language)

	Logf("Creating GitLab CI configuration at %s", filename)
	if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write GitLab CI file: %v", err)
	}

	fmt.Printf("✓ Generated GitLab CI configuration for %s\n", language)
	fmt.Printf("  Location: .gitlab-ci.yml\n")
	fmt.Println("  Next steps:")
	fmt.Println("    - Customize the pipeline as needed")
	fmt.Println("    - Push to GitLab to enable CI/CD")
	fmt.Println("    - View pipelines at: https://gitlab.com/your-repo/-/pipelines")

	return nil
}

func getGitHubActionsTemplate(language string) string {
	switch language {
	case "golang":
		return `name: Go CI

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main, develop ]

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
    - uses: actions/checkout@v3

    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.20'
        cache: true

    - name: Build
      run: go build -v ./cmd/devtool

    - name: Test
      run: go test -v -race -coverprofile=coverage.out ./...

    - name: Upload coverage
      uses: codecov/codecov-action@v3
      with:
        files: ./coverage.out

    - name: Lint
      uses: golangci/golangci-lint-action@v3
      with:
        version: latest
`

	case "spring":
		return `name: Spring Boot CI

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main, develop ]

jobs:
  build:
    runs-on: ubuntu-latest

    services:
      mysql:
        image: mysql:8.0
        env:
          MYSQL_DATABASE: testdb
          MYSQL_ROOT_PASSWORD: password
        options: >-
          --health-cmd="mysqladmin ping"
          --health-interval=10s
          --health-timeout=5s
          --health-retries=3

    steps:
    - uses: actions/checkout@v3

    - name: Set up JDK 17
      uses: actions/setup-java@v3
      with:
        java-version: '17'
        distribution: 'temurin'
        cache: maven

    - name: Build with Maven
      run: mvn clean verify

    - name: Upload coverage to Codecov
      uses: codecov/codecov-action@v3
      with:
        files: ./target/site/jacoco/jacoco.xml
`

	case "nextjs":
		return `name: Next.js CI

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main, develop ]

jobs:
  build-and-test:
    runs-on: ubuntu-latest

    strategy:
      matrix:
        node-version: [18.x, 20.x]

    steps:
    - uses: actions/checkout@v3

    - name: Setup Node.js ${{ matrix.node-version }}
      uses: actions/setup-node@v3
      with:
        node-version: ${{ matrix.node-version }}
        cache: 'npm'

    - name: Install dependencies
      run: npm ci

    - name: Lint
      run: npm run lint

    - name: Build
      run: npm run build

    - name: Test
      run: npm test -- --coverage

    - name: Upload coverage
      uses: codecov/codecov-action@v3
      with:
        files: ./coverage/lcov.info
`

	case "terraform":
		return `name: Terraform CI

on:
  push:
    branches: [ main, develop ]
    paths:
      - 'terraform/**'
  pull_request:
    branches: [ main, develop ]
    paths:
      - 'terraform/**'

jobs:
  terraform:
    runs-on: ubuntu-latest

    steps:
    - uses: actions/checkout@v3

    - name: Setup Terraform
      uses: hashicorp/setup-terraform@v2
      with:
        terraform_version: 1.6.0

    - name: Terraform Format Check
      run: terraform fmt -check -recursive ./terraform

    - name: Terraform Init
      run: cd terraform && terraform init -backend=false

    - name: Terraform Validate
      run: cd terraform && terraform validate

    - name: TFLint
      uses: terraform-linters/setup-tflint@v3

    - name: Run TFLint
      run: cd terraform && tflint --init && tflint --format compact
`

	default:
		return getGitHubActionsTemplate("golang")
	}
}

func getGitLabCITemplate(language string) string {
	switch language {
	case "golang":
		return `stages:
  - build
  - test
  - lint

variables:
  GO_VERSION: "1.20"

build:
  stage: build
  image: golang:${GO_VERSION}
  script:
    - go build -v ./cmd/devtool
  artifacts:
    paths:
      - devtool
    expire_in: 1 week

test:
  stage: test
  image: golang:${GO_VERSION}
  script:
    - go test -v -race -coverprofile=coverage.out ./...
  coverage: '/coverage:.+/'
  artifacts:
    reports:
      coverage_report:
        coverage_format: cobertura
        path: coverage.out

lint:
  stage: lint
  image: golangci/golangci-lint:latest
  script:
    - golangci-lint run --timeout 5m
`

	case "spring":
		return `stages:
  - build
  - test

variables:
  MAVEN_OPTS: "-Dmaven.repo.local=$CI_PROJECT_DIR/.m2/repository"

build:
  stage: build
  image: maven:3.8-openjdk-17
  script:
    - mvn clean package -DskipTests
  artifacts:
    paths:
      - target/
    expire_in: 1 week
  cache:
    paths:
      - .m2/repository

test:
  stage: test
  image: maven:3.8-openjdk-17
  script:
    - mvn clean verify
  coverage: '/Coverage: \d+\.\d+%/'
  artifacts:
    reports:
      junit: target/surefire-reports/TEST-*.xml
  cache:
    paths:
      - .m2/repository
`

	case "nextjs":
		return `stages:
  - install
  - lint
  - build
  - test

variables:
  NODE_VERSION: "18"

install:
  stage: install
  image: node:${NODE_VERSION}
  script:
    - npm ci
  cache:
    paths:
      - node_modules/
  artifacts:
    paths:
      - node_modules/
    expire_in: 1 hour

lint:
  stage: lint
  image: node:${NODE_VERSION}
  needs: ["install"]
  script:
    - npm run lint
  cache:
    paths:
      - node_modules/

build:
  stage: build
  image: node:${NODE_VERSION}
  needs: ["install"]
  script:
    - npm run build
  artifacts:
    paths:
      - .next/
    expire_in: 1 week
  cache:
    paths:
      - node_modules/

test:
  stage: test
  image: node:${NODE_VERSION}
  needs: ["install"]
  script:
    - npm test -- --coverage
  artifacts:
    reports:
      coverage_report:
        coverage_format: cobertura
        path: coverage/cobertura-coverage.xml
  cache:
    paths:
      - node_modules/
`

	case "terraform":
		return `stages:
  - validate
  - plan

variables:
  TERRAFORM_VERSION: "1.6"

validate:
  stage: validate
  image: hashicorp/terraform:${TERRAFORM_VERSION}
  script:
    - cd terraform
    - terraform init -backend=false
    - terraform validate
    - terraform fmt -check -recursive .

plan:
  stage: plan
  image: hashicorp/terraform:${TERRAFORM_VERSION}
  script:
    - cd terraform
    - terraform init
    - terraform plan -out=tfplan
  artifacts:
    paths:
      - terraform/tfplan
    expire_in: 7 days
  only:
    - merge_requests
`

	default:
		return getGitLabCITemplate("golang")
	}
}
