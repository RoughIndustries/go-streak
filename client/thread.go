package client

import (
	"github.com/RoughIndustries/go_streak/models"
	"net/http"
)

// PurchaseShippingLabel creates a new transaction object and purchases the shipping label for the provided rate.
func (c *Client) GetThreadsInBox(boxKey string) (*models.ThreadsResult, error) {
	output := &models.ThreadsResult{}
	err := c.do(http.MethodGet, "/api/v1/boxes/"+boxKey+"/threads", nil, &output.Threads)
	return output, err
}

func (c *Client) GetThread(threadKey string) (*models.Thread, error) {
	output := &models.Thread{}
	err := c.do(http.MethodGet, "/api/v1/threads/"+threadKey, nil, output)
	return output, err
}
