package cmd

import (
	"fmt"

	"github.com/ayuxsec/cachex/pkg/config"
)

func buildHelpMessage(cfg *config.Config) string {
	fullTimeout := cfg.ScannerConfig.Client.DialTimeout +
		cfg.ScannerConfig.Client.HandshakeTimeout +
		cfg.ScannerConfig.Client.ResponseTimeout

	helpMessage := fmt.Sprintf(`cachex - Tool to detect cache poisoning vulnerabilities

USAGE:
  cachex [flags]

FLAGS:

INPUT:
  -u, --url                      URL to scan
  -l, --list                     Path to a file containing a list of URLs to scan

GENERAL:
  -m, --scan-mode                Scan mode: single or multi (default: %s)

CONCURRENCY:
  -t, --threads                  Number of concurrent scan workers (default: %d)

RATE LIMITING:
  -rl, --rate-limit-per-second   Max HTTP requests per second (default: %d)

HTTP CLIENT:
  -timeout, --request-timeout    Total request timeout in seconds (default: %.1f)
  -proxy, --proxy-url            Proxy URL to use for requests (default: %s)

PERSISTENCE CHECKER:
  -np, --no-chk-prst             Disable persistence checker (default: %v)
  -pr, --prst-requests           Requests sent to poison cache (default: %d)
  -pt, --prst-threads            Concurrent poisoning workers (default: %d)

OUTPUT:
  -j, --json                     Write JSONLines output
  -o, --output                   Path to output file (default: stdout)

PAYLOADS:
  -pcf, --payload-config-file    Path to payload config YAML file (default: %s)
`,
		cfg.ScannerConfig.ScanMode,
		cfg.ScannerConfig.Threads,
		cfg.ScannerConfig.Client.RateLimitRPS,
		fullTimeout,
		"None",
		!cfg.ScannerConfig.PersistenceCheckerArgs.Enabled,
		cfg.ScannerConfig.PersistenceCheckerArgs.NumRequestsToSend,
		cfg.ScannerConfig.PersistenceCheckerArgs.Threads,
		config.DefaultPayloadHeadersPath,
	)

	return helpMessage
}
