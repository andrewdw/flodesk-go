package client

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// ListSubscribers returns a list of subscribers.
func (c *Client) ListSubscribers(page, perPage int, status, segmentID *string) (*ListSubscribersResponse, error) {
	query := url.Values{}
	if page > 0 {
		query.Add("page", strconv.Itoa(page))
	}
	if perPage > 0 {
		query.Add("per_page", strconv.Itoa(perPage))
	}
	if status != nil && *status != "" {
		query.Add("status", *status)
	}
	if segmentID != nil && *segmentID != "" {
		query.Add("segment_id", *segmentID)
	}

	path := "/subscribers"
	if len(query) > 0 {
		path += "?" + query.Encode()
	}

	req, err := c.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var response ListSubscribersResponse
	err = c.do(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// CreateOrUpdateSubscriber creates or updates a subscriber.
func (c *Client) CreateOrUpdateSubscriber(subscriber *CreateSubscriberRequest) (*Subscriber, error) {
	req, err := c.newRequest(http.MethodPost, "/subscribers", subscriber)
	if err != nil {
		return nil, err
	}

	var response Subscriber
	err = c.do(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetSubscriber returns a subscriber by ID or email.
func (c *Client) GetSubscriber(idOrEmail string) (*Subscriber, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("/subscribers/%s", idOrEmail), nil)
	if err != nil {
		return nil, err
	}

	var response Subscriber
	err = c.do(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) RemoveFromSegments(idOrEmail string, segmentIDs []string) (*Subscriber, error) {
	body := struct {
		SegmentIDs []string `json:"segment_ids"`
	}{
		SegmentIDs: segmentIDs,
	}

	req, err := c.newRequest(http.MethodDelete, fmt.Sprintf("/subscribers/%s/segments", idOrEmail), body)
	if err != nil {
		return nil, err
	}

	var response Subscriber
	err = c.do(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) AddToSegments(idOrEmail string, segmentIDs []string) (*Subscriber, error) {
	body := struct {
		SegmentIDs []string `json:"segment_ids"`
	}{
		SegmentIDs: segmentIDs,
	}

	req, err := c.newRequest(http.MethodPost, fmt.Sprintf("/subscribers/%s/segments", idOrEmail), body)
	if err != nil {
		return nil, err
	}

	var response Subscriber
	err = c.do(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UnsubscribeFromAll(idOrEmail string) (*Subscriber, error) {
	req, err := c.newRequest(http.MethodPost, fmt.Sprintf("/subscribers/%s/unsubscribe", idOrEmail), nil)
	if err != nil {
		return nil, err
	}

	var response Subscriber
	err = c.do(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
