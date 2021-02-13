package client

import (
	"github.com/RoughIndustries/go_streak/models"
	"net/http"
)

// PurchaseShippingLabel creates a new transaction object and purchases the shipping label for the provided rate.
func (c *Client) GetMeetingsInBox(boxKey string) (*models.MeetingResults, error) {
	output := &models.MeetingResults{}
	err := c.do(http.MethodGet, "/api/v2/boxes/"+boxKey+"/meetings", nil, &output)
	return output, err
}
