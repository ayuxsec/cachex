package cmd

import (
	"fmt"
	"testing"

	"github.com/ayuxsec/cachex/pkg/config"
)

func TestBuildHelpMessage(t *testing.T) {
	cfg := config.DefaultConfig()
	helpMessage := buildHelpMessage(cfg)
	fmt.Print(helpMessage)
}
