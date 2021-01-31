package tests

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/RoughIndustries/go_streak"
	"github.com/RoughIndustries/go_streak/client"
	"testing"
)

func TestClient(t *testing.T) {
	privateToken := ""
	if privateToken == "" {
		panic(errors.New("Please set PRIVATE_TOKEN"))
	}

	// create a Streak Client instance
	c := go_streak.NewClient(privateToken)
	// get tasks
	getTasks(c)
}

func getTasks(c *client.Client) {
	out, err := c.GetTasks()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Transaction:\n%s\n", dump(out))
}

func dump(v interface{}) string {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(data)
}
