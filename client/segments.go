package client

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// ListSegments returns a list of segments.
func (c *Client) ListSegments(page, perPage int) (*ListSegmentsResponse, error) {
	query := url.Values{}
	if page > 0 {
		query.Add("page", strconv.Itoa(page))
	}
	if perPage > 0 {
		query.Add("per_page", strconv.Itoa(perPage))
	}

	path := "/segments"
	if len(query) > 0 {
		path += "?" + query.Encode()
	}

	req, err := c.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var response ListSegmentsResponse
	err = c.do(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetSegment returns a segment by ID.
func (c *Client) GetSegment(id string) (*Segment, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("/segments/%s", id), nil)
	if err != nil {
		return nil, err
	}

	var response Segment
	err = c.do(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
