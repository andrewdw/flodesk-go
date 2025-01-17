# flodesk-go

A Go client library for the [Flodesk API](https://developers.flodesk.com/).

## Installation

```bash
go get github.com/andrewdw/flodesk-go
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/andrewdw/flodesk-go/client"
)

func main() {
    // Create a new client
    c := client.NewClient("your-api-key")

    // List subscribers with a default limit of 20
    subscribers, err := c.ListSubscribers(1, 20, "active", "")
    if err != nil {
        panic(err)
    }

    // Print subscribers
    for _, s := range subscribers.Data {
        fmt.Printf("Subscriber: %s <%s>\n", s.FirstName, s.Email)
    }

    // Create a new subscriber
    newSubscriber := &client.CreateSubscriberRequest{
        Email:     "test@example.com",
        FirstName: "Test",
        LastName:  "User",
    }

    subscriber, err := c.CreateOrUpdateSubscriber(newSubscriber)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Created subscriber: %s\n", subscriber.ID)
}
```

## Features

- Full support for Flodesk API v1
- Pagination support
- Custom HTTP client support
- Unit tested
- Strongly typed responses

## API Coverage

- Subscribers

  - List subscribers
  - Create or update subscriber
  - Get subscriber
  - Remove from segments
  - Add to segments
  - Unsubscribe from all

- Segments

  - List segments
  - Get segment

- Custom Fields

  - List custom fields
  - Create custom field
  - List all custom fields

- Webhooks
  - List webhooks
  - Create webhook
  - Delete webhook
  - Get webhook
  - Update webhook

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin feature/my-new-feature`)
5. Create a new Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
