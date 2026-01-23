package main

import (
	"os"

	"github.com/ayuxsec/cachex/internal/app/cachex/cmd"
	"github.com/ayuxsec/cachex/pkg/logger"
)

func main() {
	cmd.PrintBanner()
	if err := cmd.App().Run(os.Args); err != nil {
		logger.Errorf(err.Error())
	}
}
