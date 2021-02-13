package client

import (
	"github.com/RoughIndustries/go_streak/models"
	"net/http"
)

// PurchaseShippingLabel creates a new transaction object and purchases the shipping label for the provided rate.
func (c *Client) GetFilesInABox(boxKey string) (*models.FilesResult, error) {
	files := &models.FilesResult{}
	err := c.do(http.MethodGet, "/api/v1/boxes/"+boxKey+"/files", nil, &files.Files)
	return files, err
}

func (c *Client) GetSpecificFile(fileKey string) (*models.File, error) {
	file := &models.File{}
	err := c.do(http.MethodGet, "/api/v1/files/"+fileKey, nil, &file)
	return file, err
}

func (c *Client) GetFileContents(fileKey string) ([]byte, error) {
	var file []byte
	file, err := c.doBytes(http.MethodGet, "/api/v1/files/"+fileKey+"/contents", nil)
	return file, err
}
