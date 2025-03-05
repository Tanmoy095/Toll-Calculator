package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Tanmoy095/Toll-Calculator.git/types"
)

type Client struct {
	Endpoint string
}

func NewClient(endpoint string) *Client {
	return &Client{Endpoint: endpoint}
}

func (c *Client) AggregateInvoice(distance types.Distance) error {
	b, err := json.Marshal(distance)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", c.Endpoint, bytes.NewReader(b))
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err

	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("the service re-posted with non 200 status code: %d", resp.StatusCode)
	}
	return nil
}
