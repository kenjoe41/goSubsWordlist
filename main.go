package main

import (
	"flag"
	"log"

	"github.com/kenjoe41/goSubsWordlist/cli"
)

func main() {
	includeRoot, silent := parseFlags()

	if err := cli.Cli(includeRoot, silent); err != nil {
		log.Fatal(err)
	}
}

func parseFlags() (bool, bool) {
	var includeRoot, silent bool
	flag.BoolVar(&includeRoot, "iR", false, "Include root domain names in wordlist.")
	flag.BoolVar(&silent, "silent", false, "Don't print the banner.")
	flag.Parse()
	return includeRoot, silent
}
