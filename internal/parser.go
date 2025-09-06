package internal

import (
	"bytes"
	"io"
	"regexp"
	"strconv"
	"strings"
)

// Parser: Server-sent event parser.
type Parser struct {
	stream io.ReadCloser
}

// Parse: Reads till the end of the stream and dispatches events to the callback.
// dispatchCb(event) receives an *Event.
func (p *Parser) Parse(dispatchCb func(*Event), retryCb func(int), lastEventIdCb func(string)) {
	crlfRegex := regexp.MustCompile(`\r\n|\r|\n`)

	var buf bytes.Buffer
	chunk := make([]byte, 1024)
	for {
		bytes, err := p.stream.Read(chunk)
		if err != nil && bytes == 0 {
			break
		}

		buf.Write(chunk[:bytes])

		str := buf.String()
		str = crlfRegex.ReplaceAllString(str, "\n")
		lines := strings.Split(str, "\n")
		if len(lines) > 0 {
			var id string
			var event string
			var data []string
			var offset int

			for i, line := range lines {
				ci := strings.Index(line, ":")

				if line == "" && len(data) > 0 {
					dispatchCb(NewEvent(id, event, data))
					event = ""
					data = nil
					offset = i
				} else if ci > 0 {
					split := strings.SplitN(line, ":", 2)
					field := split[0]
					value := split[1]

					switch field {
					case "data":
						data = append(data, value)
						break
					case "event":
						event = value
						break
					case "id":
						id = value
						lastEventIdCb(id)
						break
					case "retry":
						if ms, err := strconv.Atoi(value); err == nil {
							retryCb(ms)
						}
						break
					}
				}
			}

			buf.Reset()
			buf.WriteString(strings.Join(lines[offset:], "\r\n"))
		}
	}
}

// NewParser: Creates new server-sent event parser.
func NewParser(stream io.ReadCloser) *Parser {
	return &Parser{
		stream: stream,
	}
}
