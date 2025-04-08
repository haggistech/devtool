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
	
	fmt.Printf("Creating base Spring project in %s\n", projectPath)
	
	// For a proper Spring project, we'd typically use Spring Initializr
	// This is a simplified example that creates a basic structure
	
	// Create directories
	dirs := []string{
		"src/main/java/com/example/demo",
		"src/main/resources",
		"src/test/java/com/example/demo",
		"src/test/resources",
	}
	
	for _, dir := range dirs {
		err := os.MkdirAll(filepath.Join(projectPath, dir), 0755)
		if err != nil {
			return err
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
	return os.WriteFile(filepath.Join(projectPath, "pom.xml"), []byte(pomContent), 0644)
}
