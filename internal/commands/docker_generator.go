// internal/commands/docker_generator.go
package commands

import (
	"os"
	"path/filepath"
)

// GenerateDockerFiles creates Dockerfile and docker-compose.yml
func GenerateDockerFiles(projectPath, projectType string) error {
	// Generate Dockerfile
	dockerfile := generateDockerfile(projectType)
	dockerPath := filepath.Join(projectPath, "Dockerfile")
	Logf("Creating Dockerfile at %s", dockerPath)
	if err := os.WriteFile(dockerPath, []byte(dockerfile), 0644); err != nil {
		return err
	}

	// Generate docker-compose.yml
	dockerCompose := generateDockerCompose(projectType)
	composePath := filepath.Join(projectPath, "docker-compose.yml")
	Logf("Creating docker-compose.yml at %s", composePath)
	if err := os.WriteFile(composePath, []byte(dockerCompose), 0644); err != nil {
		return err
	}

	// Generate .dockerignore
	dockerIgnore := generateDockerIgnore()
	ignorePath := filepath.Join(projectPath, ".dockerignore")
	Logf("Creating .dockerignore at %s", ignorePath)
	if err := os.WriteFile(ignorePath, []byte(dockerIgnore), 0644); err != nil {
		return err
	}

	return nil
}

func generateDockerfile(projectType string) string {
	switch projectType {
	case "golang":
		return `# Build stage
FROM golang:1.20-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/devtool

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
`

	case "spring":
		return `FROM eclipse-temurin:17-jdk-alpine

WORKDIR /app

COPY pom.xml .
RUN apk add --no-cache maven && mvn dependency:resolve

COPY . .

RUN mvn clean package -DskipTests

EXPOSE 8080

ENTRYPOINT ["java", "-jar", "target/*.jar"]
`

	case "nextjs":
		return `FROM node:18-alpine AS builder

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm ci

COPY . .
RUN npm run build

FROM node:18-alpine

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm ci --only=production

COPY --from=builder /app/.next ./.next
COPY --from=builder /app/public ./public

EXPOSE 3000

CMD ["npm", "start"]
`

	case "terraform":
		return `FROM hashicorp/terraform:latest

WORKDIR /terraform

COPY . .

ENTRYPOINT ["terraform"]
CMD ["--help"]
`

	default:
		return `FROM alpine:latest

WORKDIR /app

COPY . .

CMD ["/bin/sh"]
`
	}
}

func generateDockerCompose(projectType string) string {
	switch projectType {
	case "golang":
		return `version: '3.8'

services:
  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      - ENV=development
      - DB_HOST=db
      - DB_PORT=5432
    depends_on:
      - db
    volumes:
      - .:/app

  db:
    image: postgres:15-alpine
    environment:
      - POSTGRES_DB=myapp
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=postgres
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data

volumes:
  postgres_data:
`

	case "spring":
		return `version: '3.8'

services:
  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      - SPRING_DATASOURCE_URL=jdbc:mysql://db:3306/myapp
      - SPRING_DATASOURCE_USERNAME=root
      - SPRING_DATASOURCE_PASSWORD=password
      - SPRING_JPA_HIBERNATE_DDL_AUTO=update
    depends_on:
      - db
    volumes:
      - .:/app

  db:
    image: mysql:8.0
    environment:
      - MYSQL_DATABASE=myapp
      - MYSQL_ROOT_PASSWORD=password
    ports:
      - "3306:3306"
    volumes:
      - mysql_data:/var/lib/mysql

volumes:
  mysql_data:
`

	case "nextjs":
		return `version: '3.8'

services:
  app:
    build: .
    ports:
      - "3000:3000"
    environment:
      - NODE_ENV=development
      - NEXT_PUBLIC_API_URL=http://localhost:8000
    volumes:
      - .:/app
      - /app/node_modules
      - /app/.next
    command: npm run dev

  db:
    image: postgres:15-alpine
    environment:
      - POSTGRES_DB=myapp
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=postgres
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data

volumes:
  postgres_data:
`

	case "terraform":
		return `version: '3.8'

services:
  terraform:
    build: .
    working_dir: /terraform
    environment:
      - AWS_REGION=us-west-2
      - TF_VAR_environment=dev
    volumes:
      - .:/terraform
      - ~/.aws:/root/.aws:ro
    command: init
`

	default:
		return `version: '3.8'

services:
  app:
    build: .
    ports:
      - "8000:8000"
    environment:
      - ENV=development
`
	}
}

func generateDockerIgnore() string {
	return `.git
.gitignore
.env
.env.local
.vscode
.idea
.DS_Store
node_modules
dist
build
target
*.log
*.swp
*.swo
*~
.terraform
*.tfstate*
__pycache__
*.pyc
.pytest_cache
`
}
