// internal/commands/spring_project.go
package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

// CreateBaseSpringProject creates a new Spring project structure
func CreateBaseSpringProject(projectPath string) error {
	if projectPath == "" {
		projectPath = "."
	}

	if err := validateProjectPath(projectPath); err != nil {
		return err
	}

	fmt.Printf("Creating base Spring project in %s\n", projectPath)
	Logf("Project path validated: %s", projectPath)
	
	// Check if pom.xml already exists
	if CheckFileExists(filepath.Join(projectPath, "pom.xml")) {
		return fmt.Errorf("pom.xml already exists at %s. This looks like an existing Spring project", projectPath)
	}

	// Create directories
	dirs := []string{
		"src/main/java/com/example/demo",
		"src/main/resources",
		"src/test/java/com/example/demo",
		"src/test/resources",
	}

	for _, dir := range dirs {
		fullPath := filepath.Join(projectPath, dir)
		err := os.MkdirAll(fullPath, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory %s: %v", dir, err)
		}
	}
	
	// Create Application.java
	appContent := `package com.example.demo;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class Application {
    public static void main(String[] args) {
        SpringApplication.run(Application.class, args);
    }
}
`
	err := os.WriteFile(filepath.Join(projectPath, "src/main/java/com/example/demo/Application.java"), []byte(appContent), 0644)
	if err != nil {
		return err
	}
	
	// Create application.properties
	propContent := `# Application properties
server.port=8080
`
	err = os.WriteFile(filepath.Join(projectPath, "src/main/resources/application.properties"), []byte(propContent), 0644)
	if err != nil {
		return err
	}
	
	// Create pom.xml
	pomContent := `<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0" 
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 
                             https://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>
    
    <parent>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-parent</artifactId>
        <version>3.1.0</version>
        <relativePath/>
    </parent>
    
    <groupId>com.example</groupId>
    <artifactId>demo</artifactId>
    <version>0.0.1-SNAPSHOT</version>
    <name>demo</name>
    <description>Demo Spring Boot project</description>
    
    <properties>
        <java.version>17</java.version>
    </properties>
    
    <dependencies>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
        
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
        </dependency>
    </dependencies>
    
    <build>
        <plugins>
            <plugin>
                <groupId>org.springframework.boot</groupId>
                <artifactId>spring-boot-maven-plugin</artifactId>
            </plugin>
        </plugins>
    </build>
</project>
`
	err = os.WriteFile(filepath.Join(projectPath, "pom.xml"), []byte(pomContent), 0644)
	if err != nil {
		return err
	}

	fmt.Println("✓ Spring project created successfully")
	fmt.Printf("Next steps:\n  cd %s\n  mvn clean install\n  mvn spring-boot:run\n", projectPath)

	return nil
}
