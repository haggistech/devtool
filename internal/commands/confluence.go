// internal/commands/confluence.go
package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type ConfluencePage struct {
	Type  string `json:"type"`
	Title string `json:"title"`
	Space struct {
		Key string `json:"key"`
	} `json:"space"`
	Body struct {
		Storage struct {
			Value          string `json:"value"`
			Representation string `json:"representation"`
		} `json:"storage"`
	} `json:"body"`
}

// HandleConfluenceCreatePage creates a new Confluence page
func HandleConfluenceCreatePage() error {
	confluenceURL := os.Getenv("CONFLUENCE_URL")
	confluenceUser := os.Getenv("CONFLUENCE_USER")
	confluenceToken := os.Getenv("CONFLUENCE_TOKEN")
	confluenceSpace := os.Getenv("CONFLUENCE_SPACE")

	if confluenceURL == "" || confluenceUser == "" || confluenceToken == "" {
		fmt.Println("Confluence environment variables not fully set.")
		fmt.Println("Please set the following environment variables:")
		fmt.Println("  CONFLUENCE_URL - your Confluence base URL (e.g., https://company.atlassian.net/wiki)")
		fmt.Println("  CONFLUENCE_USER - your Confluence email address")
		fmt.Println("  CONFLUENCE_TOKEN - your Confluence API token")
		fmt.Println("  CONFLUENCE_SPACE - the space key where to create the page (e.g., DOCS)")
		return fmt.Errorf("missing required environment variables")
	}

	if confluenceSpace == "" {
		confluenceSpace = "DOCS"
	}

	fmt.Print("Enter page title: ")
	var title string
	fmt.Scanln(&title)

	if strings.TrimSpace(title) == "" {
		return fmt.Errorf("page title cannot be empty")
	}

	fmt.Print("Enter page content (supports HTML): ")
	var content string
	fmt.Scanln(&content)

	if strings.TrimSpace(content) == "" {
		return fmt.Errorf("page content cannot be empty")
	}

	page := ConfluencePage{
		Type:  "page",
		Title: title,
	}
	page.Space.Key = confluenceSpace
	page.Body.Storage.Value = content
	page.Body.Storage.Representation = "storage"

	pageJSON, err := json.Marshal(page)
	if err != nil {
		return err
	}

	fmt.Printf("Creating Confluence page: %s\n", title)

	url := strings.TrimRight(confluenceURL, "/") + "/rest/api/content"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(pageJSON))
	if err != nil {
		return err
	}

	req.SetBasicAuth(confluenceUser, confluenceToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to create page: %s", resp.Status)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	pageID := result["id"]
	fmt.Printf("\n✓ Confluence page created successfully!\n")
	fmt.Printf("Page ID: %s\n", pageID)
	fmt.Printf("Title: %s\n", title)
	fmt.Printf("Space: %s\n", confluenceSpace)

	return nil
}
