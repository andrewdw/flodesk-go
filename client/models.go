package client

import "time"

// PaginationMeta represents the metadata for pagination.
type PaginationMeta struct {
	Page       int `json:"page"`
	TotalPages int `json:"total_pages"`
	PerPage    int `json:"per_page"`
	TotalItems int `json:"total_items"`
}

// SegmentMini represents a segment without subscriber counts.
type SegmentMini struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Subscriber represents a subscriber.
type Subscriber struct {
	ID             string                 `json:"id"`
	Status         string                 `json:"status"`
	Email          string                 `json:"email"`
	Source         string                 `json:"source"`
	FirstName      string                 `json:"first_name,omitempty"`
	LastName       string                 `json:"last_name,omitempty"`
	Segments       []SegmentMini          `json:"segments,omitempty"`
	CustomFields   map[string]interface{} `json:"custom_fields,omitempty"`
	OptinIP        string                 `json:"optin_ip,omitempty"`
	OptinTimestamp *time.Time             `json:"optin_timestamp,omitempty"`
	CreatedAt      *time.Time             `json:"created_at,omitempty"`
}

// ListSubscribersResponse represents a list of subscribers.
type ListSubscribersResponse struct {
	Meta PaginationMeta `json:"meta"`
	Data []Subscriber   `json:"data"`
}

// CreateSubscriberRequest represents a request to create or update a subscriber.
type CreateSubscriberRequest struct {
	ID             string                 `json:"id,omitempty"`
	Email          string                 `json:"email,omitempty"`
	FirstName      string                 `json:"first_name,omitempty"`
	LastName       string                 `json:"last_name,omitempty"`
	CustomFields   map[string]interface{} `json:"custom_fields,omitempty"`
	SegmentIDs     []string               `json:"segment_ids,omitempty"`
	DoubleOptIn    bool                   `json:"double_optin,omitempty"`
	OptinIP        string                 `json:"optin_ip,omitempty"`
	OptinTimestamp *time.Time             `json:"optin_timestamp,omitempty"`
}

// Segment represents a segment.
type Segment struct {
	ID                     string     `json:"id"`
	Name                   string     `json:"name"`
	TotalActiveSubscribers int        `json:"total_active_subscribers"`
	CreatedAt              *time.Time `json:"created_at,omitempty"`
}

// ListSegmentsResponse represents a list of segments.
type ListSegmentsResponse struct {
	Meta PaginationMeta `json:"meta"`
	Data []Segment      `json:"data"`
}

// CustomField represents a custom field.
type CustomField struct {
	Key   string `json:"key"`
	Label string `json:"label"`
}

// ListCustomFieldsResponse represents a list of custom fields.
type ListCustomFieldsResponse struct {
	Meta PaginationMeta `json:"meta"`
	Data []CustomField  `json:"data"`
}

// CreateCustomFieldRequest represents a request to create a custom field.
type CreateCustomFieldRequest struct {
	Label string `json:"label"`
}

// Webhook represents a webhook.
type Webhook struct {
	ID        string     `json:"id"`
	PostURL   string     `json:"post_url"`
	Events    []string   `json:"events"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// ListWebhooksResponse represents a list of webhooks.
type ListWebhooksResponse struct {
	Meta PaginationMeta `json:"meta"`
	Data []Webhook      `json:"data"`
}

// CreateWebhookRequest represents a request to create a webhook.
type CreateWebhookRequest struct {
	Name    string   `json:"name"`
	PostURL string   `json:"post_url"`
	Events  []string `json:"events"`
}

// UpdateWebhookRequest represents a request to update a webhook.
type UpdateWebhookRequest struct {
	Name    string   `json:"name,omitempty"`
	PostURL string   `json:"post_url,omitempty"`
	Events  []string `json:"events,omitempty"`
}
