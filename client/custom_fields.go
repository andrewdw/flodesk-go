package client

import (
	"net/http"
	"net/url"
	"strconv"
)

func (c *Client) ListCustomFields(page, perPage int) (*ListCustomFieldsResponse, error) {
	query := url.Values{}
	if page > 0 {
		query.Add("page", strconv.Itoa(page))
	}
	if perPage > 0 {
		query.Add("per_page", strconv.Itoa(perPage))
	}

	path := "/custom-fields"
	if len(query) > 0 {
		path += "?" + query.Encode()
	}

	req, err := c.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var response ListCustomFieldsResponse
	err = c.do(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) CreateCustomField(field *CreateCustomFieldRequest) (*CustomField, error) {
	req, err := c.newRequest(http.MethodPost, "/custom-fields", field)
	if err != nil {
		return nil, err
	}

	var response CustomField
	err = c.do(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) ListAllCustomFields() ([]CustomField, error) {
	req, err := c.newRequest(http.MethodGet, "/custom-fields/all", nil)
	if err != nil {
		return nil, err
	}

	var response []CustomField
	err = c.do(req, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}