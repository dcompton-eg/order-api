package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Order struct {
	ID       string `json:"id"`
	Location string `json:"location"`
	Item     string `json:"item"`
}

type CreateOrderRequest struct {
	Location string `json:"location"`
	Item     string `json:"item"`
}

type CreateOrderResponse struct {
	Order *Order `json:"order"`
	Error string `json:"error"`
}

func New(b string) *Client {
	return &Client{
		baseURL: b,
	}
}

type Client struct {
	baseURL string
	http.Client
}

// CreateOrder submits an order via the api.
func (c *Client) CreateOrder(cor *CreateOrderRequest) (*CreateOrderResponse, error) {
	u, err := url.JoinPath(c.baseURL, "/orders")
	if err != nil {
		return nil, fmt.Errorf("building create order request url: %w", err)
	}

	bs, err := json.Marshal(cor)
	if err != nil {
		return nil, fmt.Errorf("marshaling %T: %w", cor, err)
	}

	req, err := http.NewRequest(http.MethodPost, u, bytes.NewBuffer(bs))
	if err != nil {
		return nil, fmt.Errorf("preparing create order http request: %w", err)
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("performing create order request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("performing create order request: got status code %v, expected %v", resp.StatusCode, http.StatusCreated)
	}

	rbs, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading create order response body: %w", err)
	}

	fmt.Println(string(rbs))

	var coresp CreateOrderResponse
	err = json.Unmarshal(rbs, &coresp)
	if err != nil {
		return nil, fmt.Errorf("unmarshaling create order response body: %w", err)
	}

	return &coresp, nil
}
