package client

import (
	"github.com/RoughIndustries/go_streak/models"
	"net/http"
)

// PurchaseShippingLabel creates a new transaction object and purchases the shipping label for the provided rate.
func (c *Client) GetAllPipelines() (*models.PipelinesResult, error) {
	pipelines := &models.PipelinesResult{}
	err := c.do(http.MethodGet, "/api/v1/pipelines", nil, &pipelines.Pipelines)
	return pipelines, err
}

func (c *Client) GetPipeline(pipelinKey string) (*models.Pipeline, error) {
	pipeline := &models.Pipeline{}
	err := c.do(http.MethodGet, "/api/v1/pipelines/"+pipelinKey, nil, &pipeline)
	return pipeline, err
}
