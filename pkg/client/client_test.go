package client

import (
	"testing"
)

func TestClientCreateOrder(t *testing.T) {
	tc := New("http://localhost:8080")
	tr := &CreateOrderRequest{
		Location: "90210",
		Item:     "sandwich",
	}

	resp, err := tc.CreateOrder(tr)
	if err != nil {
		t.Errorf("received error performing request: %v", err)
	}

	if resp.Error != "" {
		t.Errorf("received error from API: %v", err)
	}

	if resp.Order == nil {
		t.Errorf("order not created: %v", err)
	}
}
