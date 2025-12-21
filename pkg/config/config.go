package config

// Config defines the configuration for the cache poisoning scanner
type Config struct {
	ScannerConfig ScannerConfig `yaml:"scanner"` // Scanner configuration
	PayloadConfig PayloadConfig `yaml:"payload"` // Payload configuration
}

type PayloadConfig struct {
	PayloadHeaders map[string]string `yaml:"payload_headers"` // Headers to be used for cache poisoning
}

// ScannerConfig defines the configuration for the scanner
type ScannerConfig struct {
	ScanMode               string                 `yaml:"scan_mode"`           // Mode of scanning (single or multi-header)
	Threads                int                    `yaml:"threads"`             // Number of threads to use for scanning
	RequestHeaders         map[string]string      `yaml:"request_headers"`     // Headers to be sent with the request
	Client                 ClientConfig           `yaml:"client"`              // Client configuration
	PersistenceCheckerArgs PersistenceCheckerArgs `yaml:"persistence_checker"` // Arguments for checking cache persistence
	LoggerConfig           LoggerConfig           `yaml:"logger"`              // Logger configuration
}

// ClientConfig defines the configuration for the HTTP client
type ClientConfig struct {
	DialTimeout      float64 `yaml:"dial_timeout"`          // Timeout for establishing the connection
	HandshakeTimeout float64 `yaml:"handshake_timeout"`     // Timeout for TLS handshake
	ResponseTimeout  float64 `yaml:"response_timeout"`      // Timeout for server response headers
	ProxyURL         string  `yaml:"proxy_url"`             // Proxy URL for the HTTP client (optional)
	RateLimitRPS     int     `yaml:"rate_limit_per_second"` // Max Requests to send per second (0 = unlimited)
}

type PersistenceCheckerArgs struct {
	Enabled           bool `yaml:"enabled"`              // Flag to enable persistence checking
	NumRequestsToSend int  `yaml:"num_requests_to_send"` // Number of requests to send for poisoning
	Threads           int  `yaml:"threads"`              // Number of threads to use for sending requests
}

type LoggerConfig struct {
	LogError     bool   `yaml:"log_error"`      // Flag to log errors to stderr
	LogMode      string `yaml:"log_mode"`       // Mode of logging (pretty or JSON)
	OutputFile   string `yaml:"output_file"`    // File to write logs to (optional)
	Debug        bool   `yaml:"debug"`          // Flag to enable debug logging
	SkipTenative bool   `yaml:"skip_tentative"` // Flag to skip stdout logging of tentative vulnerabilities
}
