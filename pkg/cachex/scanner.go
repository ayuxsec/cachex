package cachex

import (
	"time"

	"github.com/ayuxsec/cachex/internal/pkg/client"
	"github.com/ayuxsec/cachex/pkg/logger"
	"github.com/ayuxsec/cachex/internal/scanner"
	"github.com/ayuxsec/cachex/pkg/config"
)

// Scanner orchestrates the batch scanning of URLs for cache-related vulnerabilities.
type Scanner struct {
	URLs          []string              // List of target URLs to scan
	ScannerConfig *config.ScannerConfig // Configuration for scanner behavior and logging
	PayloadConfig *config.PayloadConfig // Configuration for payload/header injection
}

// Run initializes and runs the internal scanner with the provided configuration.
// It returns a slice of ScannerOutput and a slice of errors encountered during scanning.
func (s *Scanner) Run() ([]scanner.ScannerOutput, []error) {
	if err := s.Validate(); err != nil {
		logger.Errorf("failed to validate args: %v", err)
		return nil, []error{err}
	}
	// Prepare the internal scanner with the configured arguments
	internalScanner := scanner.ScannerArgs{
		ScanMode:       mapStrScanMode(s.ScannerConfig.ScanMode),
		RequestHeaders: s.ScannerConfig.RequestHeaders,
		PayloadHeaders: s.PayloadConfig.PayloadHeaders,

		// Configuration for persistence checking (e.g., whether headers persist in cache)
		PersistenceCheckerArgs: &scanner.PersistenceCheckerArgs{
			DoCheck:           s.ScannerConfig.PersistenceCheckerArgs.Enabled,
			NumRequestsToSend: s.ScannerConfig.PersistenceCheckerArgs.NumRequestsToSend,
			NumThreads:        s.ScannerConfig.PersistenceCheckerArgs.Threads,
		},

		// Logger configuration: what to log, where, and how
		LoggerArgs: scanner.LoggerArgs{
			LogError: s.ScannerConfig.LoggerConfig.LogError,
			LogMode:  mapStrLogMode(s.ScannerConfig.LoggerConfig.LogMode),
			//LogTarget:    mapStrLogTarget(s.ScannerConfig.LoggerConfig.LogTarget),
			SkipTenative: s.ScannerConfig.LoggerConfig.SkipTenative,
		},
	}

	client := client.Config{
		DialTimeout:           time.Duration(s.ScannerConfig.Client.DialTimeout) * time.Second,
		HandshakeTimeout:      time.Duration(s.ScannerConfig.Client.HandshakeTimeout) * time.Second,
		ResponseHeaderTimeout: time.Duration(s.ScannerConfig.Client.ResponseTimeout) * time.Second,
		ProxyURL:              s.ScannerConfig.Client.ProxyURL,
	}.CreateNewClient()

	internalScanner.Client = client

	if !s.ScannerConfig.LoggerConfig.Debug {
		logger.DisableDebug = true
	}

	if s.ScannerConfig.LoggerConfig.OutputFile != "" {
		internalScanner.LoggerArgs.LogTarget = mapStrLogTarget("both")
		internalScanner.LoggerArgs.OutputFile = s.ScannerConfig.LoggerConfig.OutputFile
	} else {
		internalScanner.LoggerArgs.LogTarget = mapStrLogTarget("stdout")
	}

	// Run the batch scan with the prepared arguments
	return internalScanner.RunBatchScan(s.URLs, s.ScannerConfig.Threads)
}
