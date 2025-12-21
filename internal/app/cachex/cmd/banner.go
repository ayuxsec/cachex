package cmd

import (
	"fmt"
	"os"

	"github.com/ayuxsec/cachex/pkg/version"
)

var banner = `
                   __             
  _________ ______/ /_  ___  _  __
 / ___/ __ ` + "`" + `/ ___/ __ \/ _ \| |/_/
/ /__/ /_/ / /__/ / / /  __/>  <  
\___/\__,_/\___/_/ /_/\___/_/|_|  
                                  
`

func PrintBanner() {
	fmt.Fprint(os.Stderr, banner) // Print the ASCII banner
	fmt.Fprint(os.Stderr, "               "+version.String()+", with <3 by @ayuxsec\n\n")
}
