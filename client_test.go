package sseclient

import (
	"io"
	"net/http"
	"testing"
	"time"
)

var serverStarted = false
var client *Client

func createClient() *Client {
	if client != nil {
		return client
	}

	return NewClient("http://127.0.0.1:8080/stream")
}

func startLocalServer(t *testing.T) {
	if serverStarted {
		return
	}

	errCh := make(chan error, 1)

	go func() {
		http.HandleFunc("/stream", func(w http.ResponseWriter, _ *http.Request) {
			io.WriteString(w, "event: thing.beep\r\n")
			io.WriteString(w, "data: {\"id\":\"1\"}\r\n")
			io.WriteString(w, "\r\n")
			io.WriteString(w, "event: thing.bop\r\n")
			io.WriteString(w, "data: {\"id\":\"2\"}\r\n")
		})

		errCh <- http.ListenAndServe("127.0.0.1:8080", nil)
	}()

	select {
	case err := <-errCh:
		serverStarted = false
		t.Fatal("startLocalServer: Failed starting the local server stream", err)
	case <-time.After(1 * time.Second):
		serverStarted = true
		t.Log("startLocalServer: Started the local server stream")
	}
}

func TestClientEvents(t *testing.T) {
	startLocalServer(t)

	client := createClient()
	events, err := client.Events()
	if err != nil {
		t.Fatal("TestClientEvents: Failed listening to the local server stream", err)
		return
	}

	for _, event := range events {
		t.Log("TestClientEvents: ", event, client.Status)
	}
}

func TestClientListen(t *testing.T) {
	startLocalServer(t)

	client := createClient()
	err := client.Listen(func(event *Event) {
		t.Log("TestClientListen: ", event, client.Status)
	})
	if err != nil {
		t.Fatal("TestClientListen: Failed listening to the local server stream", err)
	}
}

func TestClientListenAsync(t *testing.T) {
	startLocalServer(t)

	client := createClient()
	err := client.ListenAsync(func(event *Event) {
		t.Log("TestClientListenAsync: ", event, client.Status)
	})
	if err != nil {
		t.Fatal("TestClientListenAsync: Failed listening to the local server stream", err)
	}

	// Pretend to do work.
	t.Log("TestClientListenAsync: Sleeping...")
	time.Sleep(5 * time.Second)
}
