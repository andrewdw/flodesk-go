package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestNewClient tests the NewClient function.
func TestNewClient(t *testing.T) {
	apiKey := "test-api-key"
	client := NewClient(apiKey)

	if client.apiKey != apiKey {
		t.Errorf("NewClient apiKey = %v, want %v", client.apiKey, apiKey)
	}

	if client.baseURL != baseURL {
		t.Errorf("NewClient baseURL = %v, want %v", client.baseURL, baseURL)
	}

	if client.httpClient == nil {
		t.Error("NewClient httpClient is nil")
	}
}

// TestClientOptions tests the WithBaseURL and WithHTTPClient options.
func TestClientOptions(t *testing.T) {
	customURL := "https://custom.api.com"
	customClient := &http.Client{}

	client := NewClient(
		"test-api-key",
		WithBaseURL(customURL),
		WithHTTPClient(customClient),
	)

	if client.baseURL != customURL {
		t.Errorf("WithBaseURL = %v, want %v", client.baseURL, customURL)
	}

	if client.httpClient != customClient {
		t.Errorf("WithHTTPClient not set correctly")
	}
}

// setupTestServer is a helper function to set up a test server and client.
func setupTestServer(t *testing.T, path string, response interface{}) (*httptest.Server, *Client) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != path {
			t.Errorf("Expected path %q, got %q", path, r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))

	client := NewClient("test-api-key", WithBaseURL(server.URL))
	return server, client
}
