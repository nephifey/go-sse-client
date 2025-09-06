package sseclient

import (
	"errors"
	"net/http"
	"sync"

	"github.com/nephifey/go-sse-client/internal"
)

// Default user agent.
const defaultUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/139.0.0.0 Safari/537.36"

// ConnectionStatus options.
const (
	Waiting ConnectionStatus = iota
	Opened
	Closed
)

// OnEventMissing represents an error identifying the OnEvent field is missing.
var OnEventMissing = errors.New("The OnEvent field is missing and it must be set")

// Event: Type alias.
type Event = internal.Event

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

// Client: Server-sent event client interface.
type Client interface {
	Events() ([]*Event, error)
	Listen() error
	ListenAsync() error
	Status() ConnectionStatus
}

// client: Server-sent event client, used for connecting and reading the stream.
type client struct {
	url         string
	lastEventId string
	reconnectMs int
	status      ConnectionStatus
	headers     map[string]string
	onOpen      func(*http.Response)
	onClose     func(*http.Response)
	onEvent     func(*Event)
	mu          sync.Mutex
}

// WithOnOpen: Adds onOpen callback.
func (c *client) WithOnOpen(onOpen func(*http.Response)) { c.onOpen = onOpen }

// WithOnClose: Adds onOpen callback.
func (c *client) WithOnClose(onClose func(*http.Response)) { c.onClose = onClose }

// WithOnEvent: Adds onEvent callback.
func (c *client) WithOnEvent(onEvent func(*Event)) { c.onEvent = onEvent }

// LastEventId: Get the last event id.
func (c *client) LastEventId() string { return c.lastEventId }

// ReconnectMs: Get the reconnect time in milliseconds.
func (c *client) ReconnectMs() int { return c.reconnectMs }

// Status: Get the status.
func (c *client) Status() ConnectionStatus { return c.status }

// Events: Blocking listener. Gets all events and returns as a slice.
func (c *client) Events() ([]*Event, error) {
	var events []*Event
	oOnEvent := c.onEvent
	c.WithOnEvent(func(event *Event) {
		events = append(events, event)
	})

	err := c.Listen(false)
	c.onEvent = oOnEvent
	if err != nil {
		return nil, err
	}

	return events, nil
}

// Listen: Blocking or non-blocking depending on async flag.
func (c *client) Listen(async bool) error {
	if nil == c.onEvent {
		return OnEventMissing
	}

	listen := func() error {
		c.mu.Lock()
		defer c.mu.Unlock()

		c.status = Waiting
		resp, err := c.get()

		if nil != c.onClose {
			defer c.onClose(resp)
		}

		if err != nil {
			c.status = Closed
			return err
		}

		if nil != c.onOpen {
			c.onOpen(resp)
		}

		c.status = Opened
		defer resp.Body.Close()

		parser := internal.NewParser(resp.Body)
		parser.Parse(c.onEvent, c.setReconnectMs, c.setLastEventId)

		// need to handle reconnection at some point

		return nil
	}

	if async {
		go listen()

		return nil
	}

	return listen()
}

// setReconnectMs: Set reconnect ms.
func (c *client) setReconnectMs(reconnectMs int) {
	c.reconnectMs = reconnectMs
}

// setLastEventId: Set last event id.
func (c *client) setLastEventId(lastEventId string) {
	c.lastEventId = lastEventId
}

// get: Custom get request implementation.
func (c *client) get() (*http.Response, error) {
	req, err := http.NewRequest("GET", c.url, nil)
	if err != nil {
		return nil, err
	}

	for n, v := range c.headers {
		req.Header.Set(n, v)
	}

	req.Header.Set("Accept", "text/event-stream")
	if c.lastEventId != "" {
		req.Header.Set("Last-Event-Id", c.lastEventId)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, http.ErrBodyNotAllowed
	}

	return resp, nil
}

// NewClient: Creates a new server-sent event client.
func NewClient(url string) *client {
	return &client{
		url:         url,
		reconnectMs: 3000,
		status:      Waiting,
		headers: map[string]string{
			"User-Agent": defaultUserAgent,
		},
	}
}

// NewClient: Custom headers if needed (e.g., authorization, key, other).
func NewClientWithCustomHeaders(url string, headers map[string]string) *client {
	return &client{
		url:         url,
		reconnectMs: 3000,
		status:      Waiting,
		headers:     headers,
	}
}
