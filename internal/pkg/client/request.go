// Description: This file contains functions for making HTTP requests
// including fetching responses and sending requests without reading responses.
package client

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/ayuxsec/cachex/internal/pkg/logger"
)

// Response holds the response from the URL
type Response struct {
	StatusCode int                 `json:"StatusCode"`
	Headers    map[string][]string `json:"Headers"`
	Body       string              `json:"Body"`
	Location   string              `json:"Location"`
}

// FetchResponse sends a GET request and returns the response.
func FetchResponse(
	url string,
	requestHeaders map[string]string,
	r *RateLimitedClient,
) (*Response, error) {
	logger.Debugf("Fetching Response of %s", url)
	if r.Limiter != nil {
		if err := r.Limiter.Wait(context.Background()); err != nil {
			return nil, fmt.Errorf("rate limiter wait failed: %v", err)
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &Response{}, fmt.Errorf("error creating HTTP request: %v", err)
	}

	// Add headers to request
	for key, value := range requestHeaders {
		req.Header.Add(key, value)
	}

	// Send request
	resp, err := r.Client.Do(req)
	if err != nil {
		return &Response{}, fmt.Errorf("error sending HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &Response{}, fmt.Errorf("error reading response body: %v", err)
	}

	// Return response
	return &Response{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       string(respBody),
		Location:   resp.Header.Get("Location"),
	}, nil
}

// SendRequest sends a request but does not read or return the response.
func SendRequest(url string,
	requestHeaders map[string]string,
	r *RateLimitedClient,
) error {
	logger.Debugf("Sending GET request to %s", url)
	if r.Limiter != nil {
		if err := r.Limiter.Wait(context.Background()); err != nil {
			return fmt.Errorf("rate limiter wait failed: %v", err)
		}
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("error creating HTTP request: %v", err)
	}

	// Add headers to request
	for key, value := range requestHeaders {
		req.Header.Add(key, value)
	}

	// Send request (ignoring the response)
	_, err = r.Client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending HTTP request: %v", err)
	}

	return nil
}
