package client

import (
	"testing"
	"time"
)

// TestListSubscribers tests the ListSubscribers function.
func TestListSubscribers(t *testing.T) {
	status := "active"
	segmentID := "123"

	expected := &ListSubscribersResponse{
		Meta: PaginationMeta{
			Page:       1,
			TotalPages: 2,
			PerPage:    20,
			TotalItems: 30,
		},
		Data: []Subscriber{
			{
				ID:     "test-id",
				Email:  "test@example.com",
				Status: "active",
			},
		},
	}

	server, client := setupTestServer(t, "/subscribers", expected)
	defer server.Close()

	response, err := client.ListSubscribers(1, 20, &status, &segmentID)
	if err != nil {
		t.Fatalf("ListSubscribers returned error: %v", err)
	}

	if response.Meta.Page != expected.Meta.Page {
		t.Errorf("ListSubscribers page = %v, want %v", response.Meta.Page, expected.Meta.Page)
	}

	if len(response.Data) != len(expected.Data) {
		t.Errorf("ListSubscribers data length = %v, want %v", len(response.Data), len(expected.Data))
	}
}

// TestCreateOrUpdateSubscriber tests the CreateOrUpdateSubscriber function.
func TestCreateOrUpdateSubscriber(t *testing.T) {
	now := time.Now()
	expected := &Subscriber{
		ID:        "test-id",
		Email:     "test@example.com",
		Status:    "active",
		FirstName: "Test",
		LastName:  "User",
		CreatedAt: &now,
	}

	server, client := setupTestServer(t, "/subscribers", expected)
	defer server.Close()

	request := &CreateSubscriberRequest{
		Email:     "test@example.com",
		FirstName: "Test",
		LastName:  "User",
	}

	response, err := client.CreateOrUpdateSubscriber(request)
	if err != nil {
		t.Fatalf("CreateOrUpdateSubscriber returned error: %v", err)
	}

	if response.Email != expected.Email {
		t.Errorf("CreateOrUpdateSubscriber email = %v, want %v", response.Email, expected.Email)
	}
}

// TestGetSubscriber tests the GetSubscriber function.
func TestGetSubscriber(t *testing.T) {
	expected := &Subscriber{
		ID:     "test-id",
		Email:  "test@example.com",
		Status: "active",
	}

	server, client := setupTestServer(t, "/subscribers/test-id", expected)
	defer server.Close()

	response, err := client.GetSubscriber("test-id")
	if err != nil {
		t.Fatalf("GetSubscriber returned error: %v", err)
	}

	if response.ID != expected.ID {
		t.Errorf("GetSubscriber ID = %v, want %v", response.ID, expected.ID)
	}
}
