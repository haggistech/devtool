// internal/commands/jira.go
package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// JiraTicket represents the structure of a Jira ticket response
type JiraTicket struct {
	Key    string `json:"key"`
	Fields struct {
		Summary     string `json:"summary"`
		Description string `json:"description"`
		Status      struct {
			Name string `json:"name"`
		} `json:"status"`
		Priority struct {
			Name string `json:"name"`
		} `json:"priority"`
		Assignee struct {
			DisplayName string `json:"displayName"`
			EmailAddress string `json:"emailAddress"`
		} `json:"assignee"`
		Reporter struct {
			DisplayName string `json:"displayName"`
			EmailAddress string `json:"emailAddress"`
		} `json:"reporter"`
		Created    string `json:"created"`
		Updated    string `json:"updated"`
		DueDate    string `json:"duedate"`
		Resolution struct {
			Name string `json:"name"`
		} `json:"resolution"`
	} `json:"fields"`
}

// GetJiraTicket fetches a Jira ticket by its ticket number
func GetJiraTicket(ticketNumber string) error {
	if ticketNumber == "" {
		return fmt.Errorf("ticket number is required")
	}
	
	// Check for JIRA environment variables
	jiraDomain := os.Getenv("JIRA_DOMAIN")
	jiraEmail := os.Getenv("JIRA_EMAIL")
	jiraToken := os.Getenv("JIRA_API_TOKEN")
	
	if jiraDomain == "" || jiraEmail == "" || jiraToken == "" {
		fmt.Println("JIRA environment variables not fully set.")
		fmt.Println("Please set the following environment variables:")
		fmt.Println("  JIRA_DOMAIN - your Jira domain (e.g., your-company.atlassian.net)")
		fmt.Println("  JIRA_EMAIL - your Jira email address")
		fmt.Println("  JIRA_API_TOKEN - your Jira API token")
		return fmt.Errorf("missing required environment variables")
	}
	
	// Make sure ticket is properly formatted
	if !strings.Contains(ticketNumber, "-") {
		// Assume it's just the number part and add project prefix if available
		jiraProject := os.Getenv("JIRA_PROJECT")
		if jiraProject != "" {
			ticketNumber = jiraProject + "-" + ticketNumber
		} else {
			return fmt.Errorf("ticket number should include project prefix (e.g., PROJ-123)")
		}
	}
	
	fmt.Printf("Fetching Jira ticket: %s\n", ticketNumber)
	
	// Create the request
	url := fmt.Sprintf("https://%s/rest/api/2/issue/%s", jiraDomain, ticketNumber)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	
	// Set basic auth with API token
	req.SetBasicAuth(jiraEmail, jiraToken)
	req.Header.Set("Content-Type", "application/json")
	
	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	
	// Check response status
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to fetch ticket: %s - %s", resp.Status, string(body))
	}
	
	// Parse the response
	var ticket JiraTicket
	if err := json.NewDecoder(resp.Body).Decode(&ticket); err != nil {
		return err
	}
	
	// Display ticket information
	fmt.Printf("\n================== JIRA TICKET: %s ==================\n", ticket.Key)
	fmt.Printf("Summary:     %s\n", ticket.Fields.Summary)
	fmt.Printf("Status:      %s\n", ticket.Fields.Status.Name)
	fmt.Printf("Priority:    %s\n", ticket.Fields.Priority.Name)
	
	if ticket.Fields.Assignee.DisplayName != "" {
		fmt.Printf("Assignee:    %s (%s)\n", ticket.Fields.Assignee.DisplayName, ticket.Fields.Assignee.EmailAddress)
	} else {
		fmt.Printf("Assignee:    Unassigned\n")
	}
	
	fmt.Printf("Reporter:    %s\n", ticket.Fields.Reporter.DisplayName)
	fmt.Printf("Created:     %s\n", ticket.Fields.Created)
	fmt.Printf("Updated:     %s\n", ticket.Fields.Updated)
	
	if ticket.Fields.DueDate != "" {
		fmt.Printf("Due Date:    %s\n", ticket.Fields.DueDate)
	}
	
	if ticket.Fields.Resolution.Name != "" {
		fmt.Printf("Resolution:  %s\n", ticket.Fields.Resolution.Name)
	}
	
	fmt.Printf("\nDescription:\n%s\n", ticket.Fields.Description)
	fmt.Printf("\n============================================================\n")
	
	return nil
}
