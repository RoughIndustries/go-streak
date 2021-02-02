package client

import (
	"github.com/RoughIndustries/go_streak/models"
	"net/http"
)

// PurchaseShippingLabel creates a new transaction object and purchases the shipping label for the provided rate.
func (c *Client) GetAllBoxesInPipelines(pipelineKey string) (*models.BoxesResult, error) {
	boxes := &models.BoxesResult{}
	err := c.do(http.MethodGet, "/api/v1/pipelines/"+pipelineKey+"/boxes", nil, &boxes.Boxes)
	return boxes, err
}

func (c *Client) GetBox(boxKey string) (*models.Box, error) {
	box := &models.Box{}
	err := c.do(http.MethodGet, "/api/v1/boxes/"+boxKey, nil, &box)
	return box, err
}
