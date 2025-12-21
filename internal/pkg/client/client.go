// Description: This file contains the client package which is responsible for
// creating a new HTTP client with custom transport settings.
package client

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/time/rate"
)

// Config holds configuration settings for the HTTP client.
type Config struct {
	DialTimeout           time.Duration // Timeout for establishing the connection
	HandshakeTimeout      time.Duration // Timeout for TLS handshake
	ResponseHeaderTimeout time.Duration // Timeout for server response headers
	ProxyURL              string        // Proxy URL for the HTTP client (optional)
	RateLimitRPS          int           // max requests per second to send (0 = unlimited)
}

type RateLimitedClient struct {
	Client  *http.Client
	Limiter *rate.Limiter // limiter to limit requests
}

// DefaultConfig returns a default configuration for the application
func DefaultConfig() Config {
	config := Config{
		DialTimeout:           5 * time.Second,
		HandshakeTimeout:      5 * time.Second,
		ResponseHeaderTimeout: 5 * time.Second,
		RateLimitRPS:          0,
	}
	return config
}

// New creates a new HTTP RateLimited client with custom transport settings.
func (c Config) CreateNewClient() *RateLimitedClient {
	// Custom transport settings for the HTTP client.
	CustomTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // Ignore invalid SSL certificates
		DialContext: (&net.Dialer{
			Timeout: c.DialTimeout, // Timeout for establishing the connection
		}).DialContext,
		TLSHandshakeTimeout:   c.HandshakeTimeout,      // Timeout for TLS handshake
		ResponseHeaderTimeout: c.ResponseHeaderTimeout, // Timeout for server response headers
	}

	if c.ProxyURL != "" {
		proxyURL, err := url.Parse(c.ProxyURL)
		if err == nil {
			CustomTransport.Proxy = http.ProxyURL(proxyURL)
		}
	}

	httpClient := &http.Client{Transport: CustomTransport,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse // Prevents following redirects
		}}

	var limiter *rate.Limiter
	if c.RateLimitRPS > 0 {
		limiter = rate.NewLimiter(rate.Limit(c.RateLimitRPS), c.RateLimitRPS)
	}
	return &RateLimitedClient{
		Client:  httpClient,
		Limiter: limiter,
	}
}
