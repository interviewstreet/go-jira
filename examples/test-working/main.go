package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	jira "github.com/interviewstreet/go-jira"
)

func main() {

	jiraURL := ""
	username := ""
	token := ""

	tp := jira.BasicAuthTransport{
		Username: username,
		Password: token,
	}

	client, err := jira.NewClient(tp.Client(), strings.TrimSpace(jiraURL))
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	ctx := context.TODO()

	webhooks, _, err := client.Organization.GetWebhookList(ctx)

	pretty, _ := json.MarshalIndent(webhooks, "", "  ")
	fmt.Println(string(pretty))

	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}
}
