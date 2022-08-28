package main

import (
	"flag"

	"github.com/kenjoe41/goSubsWordlist/cli"
)

func main() {
	// Include the Root Domain names in words
	var includeRoot bool
	flag.BoolVar(&includeRoot, "iR", false, "Include root domain names in wordlist.")
	// Concurrency flag
	// var concurrency int
	// flag.IntVar(&concurrency, "t", runtime.NumCPU(), "Threads for concurrency. Default is current available logical CPUs available to this process.")

	flag.Parse()
	cli.Cli(includeRoot)
}
