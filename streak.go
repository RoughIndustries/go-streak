package go_streak

import "github.com/RoughIndustries/go_streak/client"

const (
	APIVersion20170329 = "2020-12-30"
)

// NewClient creates a new Streak client.
func NewClient(privateToken string) *client.Client {
	return client.NewClient(privateToken)
}
