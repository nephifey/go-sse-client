package internal

import (
	"bytes"
	"io"
	"regexp"
	"strings"
)

// Parser: Server-sent event parser.
type Parser struct {
	Stream io.ReadCloser
}

// EventCallback: The callback func type for Parse.
type EventCallback func(*Event)

// Parse: Reads till the end of the stream and dispatches events to the callback.
// callback(event) receives an *Event.
func (p *Parser) Parse(callback EventCallback) {
	crlfRegex := regexp.MustCompile(`\r\n|\r|\n`)

	var buf bytes.Buffer
	chunk := make([]byte, 1024)
	for {
		bytes, err := p.Stream.Read(chunk)
		if err != nil && bytes == 0 {
			break
		}

		buf.Write(chunk[:bytes])

		var event string
		var data []string

		str := buf.String()
		str = crlfRegex.ReplaceAllString(str, "\n")
		lines := strings.Split(str, "\n")
		if len(lines) > 0 {
			var filteredLines []string

			for _, line := range lines {
				if line == "" && len(data) > 0 {
					callback(NewEvent(event, data))
					data = nil
				} else if strings.HasPrefix(line, "event:") {
					event = strings.TrimSpace(line[len("event:"):])
				} else if strings.HasPrefix(line, "data:") {
					data = append(data, strings.TrimSpace(line[len("data:"):]))
				} else {
					filteredLines = append(filteredLines, line)
				}
			}

			buf.Reset()
			buf.WriteString(strings.Join(filteredLines, "\n"))
		}
	}
}

// NewParser: Creates new server-sent event parser.
func NewParser(stream io.ReadCloser) *Parser {
	return &Parser{
		Stream: stream,
	}
}
