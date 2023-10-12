package main

import (
	"flag"
	"fmt"

	"github.com/taylormonacelli/fewpanel"
)

func main() {
	fewpanel.InitFlags()
	flag.Parse()

	verbose := fewpanel.GetVerbose()
	logFormat := fewpanel.GetLogFormat()

	fmt.Printf("Verbose: %v\n", verbose)
	fmt.Printf("Log Format: %s\n", logFormat)
}
