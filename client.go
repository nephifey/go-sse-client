package sseclient

import (
	"net/http"
	"sync"

	"github.com/nephifey/go-sse-client/internal"
)

// ConnectionStatus options.
const (
	Waiting ConnectionStatus = iota
	Opened
	Closed
)

// Event: Type alias.
type Event = internal.Event

// EventCallback: Type alias.
type EventCallback = internal.EventCallback

// ConnectionStatus: Different connection statuses for a current client and the underlying connection to the server-sent event stream.
type ConnectionStatus int

// String: ConnectionStatus Stringer implementation.
func (c ConnectionStatus) String() string {
	switch c {
	case Waiting:
		return "waiting"
	case Opened:
		return "opened"
	case Closed:
		return "closed"
	}

	return "unknown"
}

// Client: Server-sent event client, used for connecting and reading the stream.
type Client struct {
	Url    string
	Status ConnectionStatus
	mu     sync.Mutex
}

// Events: Blocking listener. Gets all events and returns as a slice.
func (c *Client) Events() ([]*Event, error) {
	var events []*Event
	err := c.Listen(func(event *Event) {
		events = append(events, event)
	})
	if err != nil {
		return nil, err
	}

	return events, nil
}

// Listen: Blocking listener.
func (c *Client) Listen(callback EventCallback) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Status = Waiting
	resp, err := http.Get(c.Url)
	if err != nil {
		c.Status = Closed
		return err
	}

	c.Status = Opened
	defer resp.Body.Close()

	parser := internal.NewParser(resp.Body)
	parser.Parse(callback)

	return nil
}

// ListenAsync: Non-blocking listener.
func (c *Client) ListenAsync(callback EventCallback) error {
	errCh := make(chan error, 1)

	go func(callback EventCallback) {
		c.mu.Lock()
		defer c.mu.Unlock()

		c.Status = Waiting
		resp, err := http.Get(c.Url)
		if err != nil {
			c.Status = Closed
			errCh <- err
			return
		}

		c.Status = Opened
		defer resp.Body.Close()

		parser := internal.NewParser(resp.Body)
		parser.Parse(callback)

		errCh <- nil
	}(callback)

	return <-errCh
}

// NewClient: Creates a new server-sent event client.
func NewClient(url string) *Client {
	return &Client{
		Url:    url,
		Status: Waiting,
	}
}
