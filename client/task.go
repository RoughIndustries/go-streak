package client

import (
	"github.com/RoughIndustries/go_streak/models"
	"net/http"
)

// PurchaseShippingLabel creates a new transaction object and purchases the shipping label for the provided rate.
func (c *Client) GetTasks() (*models.TaskResults, error) {
	output := &models.TaskResults{}
	err := c.do(http.MethodGet, "/api/v2/tasks?limit=500&direction=future", nil, output)
	return output, err
}
