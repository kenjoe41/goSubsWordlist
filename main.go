package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"

	"github.com/kenjoe41/goSubsWordlist/ezUtils"
	"github.com/kenjoe41/goSubsWordlist/output"
)

func main() {
	
	// Print Header text
	output.Beautify()

	// Include the Root Domain names in words
	var includeRoot bool
	flag.BoolVar(&includeRoot, "iR", false, "Include root domain names in wordlist.")

	// Concurrency flag
	var concurrency int
	flag.IntVar(&concurrency, "t", runtime.NumCPU(), "Threads for concurrency. Default is current available logical CPUs available to this process.")

	flag.Parse()


	// This is divided up in the subroutine for loop, so a value below 2 is BS.
	if concurrency <= 1 {
		concurrency = 2
	} else {
		// We have 2 channels to share the concurrency with, let's reassure them that they'll have equal share.
		concurrency = concurrency * 2
	}

	// Create channels to use
	domains := make(chan string)
	subdomains := make(chan string)
	output := make(chan string)

	// Domain Input worker
	var domainsWG sync.WaitGroup
	for i := 0; i < concurrency/2; i++ {

		domainsWG.Add(1)

		go func() {
			for inDomain := range domains {
				inDomain = strings.TrimSpace(strings.ToLower(inDomain))

				domain := ezUtils.CleanDomain(inDomain)

				if domain == "" {
					// Log something but continue to next domain if available
					// log.Printf("Failed to get domain from: %s", domain)
					continue
				}

				subdomain := ezUtils.ExtractSubdomain(domain, includeRoot)

				if subdomain == "" {
					// Log something but continue to next domain if available
					// log.Printf("Failed to get subdomain for domain: %s", domain)
					continue
				}

				subdomains <- subdomain

			}
			domainsWG.Done()
		}()
	}

	var subdomainsWG sync.WaitGroup

	for i := 0; i < concurrency/2; i++ {

		subdomainsWG.Add(1)

		go func() {
			for inSubdomains := range subdomains {

				// Split the subdomain into separate words by the '.' char.
				// Returns slice of words.
				subWords := ezUtils.SplitSubToWords(inSubdomains)

				// Print to console for now
				for _, subword := range subWords {
					output <- subword
				}

			}
			subdomainsWG.Done()
		}()

	}

	// Close subdomains channel when done reading from domains chan.
	go func() {
		domainsWG.Wait()
		close(subdomains)
	}()

	var outputWG sync.WaitGroup
	outputWG.Add(1)
	go func() {
		for word := range output {
			fmt.Println(word)
		}
		outputWG.Done()
	}()

	// Close the Output Chan after subdomain worker is done.
	go func() {
		subdomainsWG.Wait()
		close(output)
	}()

	sc := bufio.NewScanner(os.Stdin)

	var inputItem bool
	for sc.Scan() {

		domains <- sc.Text()
		inputItem = true
	}

	if !inputItem {
		fmt.Fprintln(os.Stderr, "No domains or urls detected. Hint: cat domains.txt | goSubsWordlist")
		flag.Usage()
	}

	// Close domains chan
	close(domains)

	// check there were no errors reading stdin (unlikely)
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	// Wait until the output waitgroup is done
	outputWG.Wait()
}
