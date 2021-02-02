package client

import (
	"github.com/RoughIndustries/go_streak/models"
	"net/http"
)

// PurchaseShippingLabel creates a new transaction object and purchases the shipping label for the provided rate.
func (c *Client) GetTasks() (*models.TaskResults, error) {
	future := &models.TaskResults{}
	err := c.do(http.MethodGet, "/api/v2/tasks?limit=500&direction=future", nil, future)
	past := &models.TaskResults{}
	err = c.do(http.MethodGet, "/api/v2/tasks?limit=500&direction=past", nil, past)
	output := &models.TaskResults{}
	output.Results = append(output.Results, future.Results...)
	output.Results = append(output.Results, past.Results...)
	return output, err
}

func (c *Client) GetTasksInBox(boxKey string) (*models.TaskResults, error) {
	output := &models.TaskResults{}
	err := c.do(http.MethodGet, "/api/v2/boxes/"+boxKey+"/tasks", nil, output)
	return output, err
}

func (c *Client) GetTask(taskKey string) (*models.Task, error) {
	output := &models.Task{}
	err := c.do(http.MethodGet, "/api/v2/tasks/"+taskKey, nil, output)
	return output, err
}
