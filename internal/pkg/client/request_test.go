package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Hello, world!"))
	}))
	defer server.Close()

	clientConfig := DefaultConfig()
	rateLimitedClient := clientConfig.CreateNewClient()

	response, err := FetchResponse(server.URL, nil, rateLimitedClient)
	if err != nil {
		t.Errorf("FetchResponse returned an error: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, response.StatusCode)
	}

	if response.Body != "Hello, world!" {
		t.Errorf("Expected body 'Hello, world!', but got '%s'", response.Body)
	}
}

func TestSendRequest(t *testing.T) {

	clientConfig := DefaultConfig()
	clientConfig.RateLimitRPS = 1 // 1 request allowed per second
	rateLimitedClient := clientConfig.CreateNewClient()

	for range 10 { // send 10 request simulatneously to trriger 1 request per second rlimit
		err := SendRequest("http://127.0.0.1:8000", nil, rateLimitedClient)
		if err != nil {
			t.Errorf("SendRequest returned an error: %v", err)
		}
	}
}
