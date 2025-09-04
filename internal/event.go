package internal

import "encoding/json"

// Event: Server-sent event.
type Event struct {
	Name string   `json:"name"`
	Data []string `json:"data"`
}

// String: Event Stringer implementation.
func (e *Event) String() string {
	bytes, err := json.Marshal(e)
	if err != nil {
		return ""
	}

	return string(bytes)
}

// NewDefaultEvent: Creates a new server-sent event with "message" as the event name.
func NewDefaultEvent(data []string) *Event {
	return &Event{
		Name: "message",
		Data: data,
	}
}

// NewEvent: Creates a new server-sent event. If name is empty then NewDefaultEvent is called.
func NewEvent(name string, data []string) *Event {
	if name == "" {
		return NewDefaultEvent(data)
	}

	return &Event{
		Name: name,
		Data: data,
	}
}
