package cachex

import (
	"bufio"
	"os"
	"testing"

	"github.com/ayuxsec/cachex/pkg/config"
)

func TestRun(t *testing.T) {
	var urls []string
	file, _ := os.Open("../../tests/testdata/urls.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			urls = append(urls, scanner.Text())
		}
	}
	t.Logf("Scanning %d urls", len(urls))

	cfg := config.DefaultConfig()
	cachexScanner := Scanner{
		ScannerConfig: &cfg.ScannerConfig,
		PayloadConfig: &cfg.PayloadConfig,
		URLs:          urls,
	}
	cachexScanner.ScannerConfig.Client.RateLimitRPS = 1
	cachexScanner.ScannerConfig.Client.ProxyURL = "http://127.0.0.1:8080"

	scannerOutput, err := cachexScanner.Run()
	if err != nil {
		t.Log(err)
	}
	t.Log(scannerOutput)
}
