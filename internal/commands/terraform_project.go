// internal/commands/terraform_project.go
package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

// CreateBaseTerraformProject creates a new Terraform project structure
func CreateBaseTerraformProject(projectPath string) error {
	if projectPath == "" {
		projectPath = "."
	}

	fmt.Printf("Creating base Terraform project in %s\n", projectPath)

	// Create directories for Terraform modules and environments
	dirs := []string{
		"modules",
		"environments/dev",
		"environments/prod",
	}

	for _, dir := range dirs {
		err := os.MkdirAll(filepath.Join(projectPath, dir), 0755)
		if err != nil {
			return err
		}
	}

	// Create main.tf in root
	mainTf := `# Main Terraform configuration

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.16"
    }
  }

  required_version = ">= 1.2.0"
}

provider "aws" {
  region = var.aws_region
}

# Example module usage
# module "vpc" {
#   source = "./modules/vpc"
#   # Add module parameters here
# }
`
	err := os.WriteFile(filepath.Join(projectPath, "main.tf"), []byte(mainTf), 0644)
	if err != nil {
		return err
	}

	// Create variables.tf
	varsTf := `# Variables

variable "aws_region" {
  description = "AWS region to deploy resources"
  type        = string
  default     = "us-west-2"
}

# Add more variables as needed
`
	err = os.WriteFile(filepath.Join(projectPath, "variables.tf"), []byte(varsTf), 0644)
	if err != nil {
		return err
	}

	// Create outputs.tf
	outputsTf := `# Outputs

# output "example_output" {
#   description = "Example output"
#   value       = module.example.output_value
# }
`
	err = os.WriteFile(filepath.Join(projectPath, "outputs.tf"), []byte(outputsTf), 0644)
	if err != nil {
		return err
	}

	// Create example module
	exampleModuleTf := `# Example module

variable "example_var" {
  description = "Example variable"
  type        = string
  default     = "default-value"
}

resource "aws_s3_bucket" "example" {
  bucket = var.example_var
}

output "bucket_name" {
  value = aws_s3_bucket.example.id
}
`
	err = os.MkdirAll(filepath.Join(projectPath, "modules/example"), 0755)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(projectPath, "modules/example/main.tf"), []byte(exampleModuleTf), 0644)
	if err != nil {
		return err
	}

	// Create dev environment config
	devTf := `# Development environment configuration

terraform {
  backend "s3" {
    # Configure your backend here
    # bucket = "my-terraform-state"
    # key    = "dev/terraform.tfstate"
    # region = "us-west-2"
  }
}

module "root" {
  source = "../.."
  
  aws_region = "us-west-2"
  # Add more variables as needed
}
`
	err = os.WriteFile(filepath.Join(projectPath, "environments/dev/main.tf"), []byte(devTf), 0644)
	if err != nil {
		return err
	}

	// Create prod environment config
	prodTf := `# Production environment configuration

terraform {
  backend "s3" {
    # Configure your backend here
    # bucket = "my-terraform-state"
    # key    = "prod/terraform.tfstate"
    # region = "us-west-2"
  }
}

module "root" {
  source = "../.."
  
  aws_region = "us-west-2"
  # Add more variables as needed
}
`
	err = os.WriteFile(filepath.Join(projectPath, "environments/prod/main.tf"), []byte(prodTf), 0644)
	if err != nil {
		return err
	}

	// Create .gitignore
	gitignore := `.terraform
*.tfstate
*.tfstate.backup
.terraform.lock.hcl
crash.log
`
	err = os.WriteFile(filepath.Join(projectPath, ".gitignore"), []byte(gitignore), 0644)
	if err != nil {
		return err
	}

	// Create README.md
	readmeMd := `# Terraform Project

This is a base Terraform project with a modular structure.

## Structure

- /modules/: Reusable Terraform modules
- /environments/: Environment-specific configurations
  - /dev/: Development environment
  - /prod/: Production environment

## Usage

### Initialize

cd environments/dev
terraform init


### Plan
` + "```" + `
terraform plan
` + "```" + `

### Apply
` + "```" + `
terraform apply
` + "```" + `
`
	return os.WriteFile(filepath.Join(projectPath, "README.md"), []byte(readmeMd), 0644)
}
