package client

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// ListWebhooks returns a list of webhooks.
func (c *Client) ListWebhooks(page, perPage int) (*ListWebhooksResponse, error) {
	query := url.Values{}
	if page > 0 {
		query.Add("page", strconv.Itoa(page))
	}
	if perPage > 0 {
		query.Add("per_page", strconv.Itoa(perPage))
	}

	path := "/webhooks"
	if len(query) > 0 {
		path += "?" + query.Encode()
	}

	req, err := c.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var response ListWebhooksResponse
	err = c.do(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// CreateWebhook creates a webhook.
func (c *Client) CreateWebhook(webhook *CreateWebhookRequest) (*Webhook, error) {
	req, err := c.newRequest(http.MethodPost, "/webhooks", webhook)
	if err != nil {
		return nil, err
	}

	var response Webhook
	err = c.do(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteWebhook deletes a webhook by ID.
func (c *Client) DeleteWebhook(id string) error {
	req, err := c.newRequest(http.MethodDelete, fmt.Sprintf("/webhooks/%s", id), nil)
	if err != nil {
		return err
	}

	return c.do(req, nil)
}

// GetWebhook returns a webhook by ID.
func (c *Client) GetWebhook(id string) (*Webhook, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("/webhooks/%s", id), nil)
	if err != nil {
		return nil, err
	}

	var response Webhook
	err = c.do(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdateWebhook updates a webhook by ID.
func (c *Client) UpdateWebhook(id string, webhook *UpdateWebhookRequest) (*Webhook, error) {
	req, err := c.newRequest(http.MethodPut, fmt.Sprintf("/webhooks/%s", id), webhook)
	if err != nil {
		return nil, err
	}

	var response Webhook
	err = c.do(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
