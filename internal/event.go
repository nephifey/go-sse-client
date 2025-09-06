package internal

import "encoding/json"

// Event: Server-sent event.
type Event struct {
	Id   string   `json:"id"`
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

// NewEvent: Creates a new server-sent event. If name is empty then NewDefaultEvent is called.
func NewEvent(id string, name string, data []string) *Event {
	if name == "" {
		name = "message"
	}

	return &Event{
		Id:   id,
		Name: name,
		Data: data,
	}
}
