package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	baseURL = "https://api.flodesk.com/v1"
)

// Client represents the Flodesk API client.
type Client struct {
	apiKey     string
	httpClient *http.Client
	baseURL    string
}

// ClientOption is a functional option for the Client.
type ClientOption func(*Client)

// WithBaseURL sets the base URL for the client.
func WithBaseURL(url string) ClientOption {
	return func(c *Client) {
		c.baseURL = url
	}
}

// WithHTTPClient sets the HTTP client to use for requests.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

func NewClient(apiKey string, opts ...ClientOption) *Client {
	c := &Client{
		apiKey:     apiKey,
		httpClient: http.DefaultClient,
		baseURL:    baseURL,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// newRequest creates a new HTTP request.
func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	auth := base64.StdEncoding.EncodeToString([]byte(c.apiKey + ":"))
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", auth))
	req.Header.Set("User-Agent", "flodesk-go")

	return req, nil
}

// do performs the HTTP request and unmarshals the response into the provided interface.
func (c *Client) do(req *http.Request, v interface{}) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	if v != nil && resp.StatusCode != http.StatusNoContent {
		if err = json.NewDecoder(resp.Body).Decode(v); err != nil {
			return err
		}
	}

	return nil
}
