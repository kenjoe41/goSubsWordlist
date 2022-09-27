package main

import (
	"flag"
	"log"

	"github.com/kenjoe41/goSubsWordlist/cli"
)

func main() {
	// Include the Root Domain names in words
	var includeRoot bool
	flag.BoolVar(&includeRoot, "iR", false, "Include root domain names in wordlist.")

	// Silent flag, no print banner.
	var silent bool
	flag.BoolVar(&silent, "silent", false, "Don't print the banner.")

	flag.Parse()
	if err := cli.Cli(includeRoot, silent); err != nil {
		log.Fatal(err)
	}
}
