package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/kenjoe41/goSubsWordlist/utils"
)

func main() {

	// Create channels to use
	domains := make(chan string)
	subdomains := make(chan string)
	output := make(chan string)

	// Domain Input worker
	var domainsWG sync.WaitGroup
	domainsWG.Add(1)
	go func() {
		for inDomain := range domains {
			inDomain = strings.TrimSpace(strings.ToLower(inDomain))

			domain := utils.CleanDomain(inDomain)

			if domain == "" {
				// Log something but continue to next domain if available
				log.Printf("Failed to get domain from: %s", domain)
				continue
			}

			subdomain := utils.ExtractSubdomain(domain)

			if subdomain == "" {
				// Log something but continue to next domain if available
				log.Printf("Failed to get subdomain for domain: %s", domain)
				continue
			}

			subdomains <- subdomain

		}
		domainsWG.Done()
	}()

	var subdomainsWG sync.WaitGroup
	subdomainsWG.Add(1)
	go func() {
		for inSubdomains := range subdomains {

			// Split the subdomain into separate words by the '.' char.
			// Returns slice of words.
			subWords := utils.SplitSubToWords(inSubdomains)

			// Print to console for now
			for _, subword := range subWords {
				output <- subword
			}

		}
		subdomainsWG.Done()
	}()

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
	}()

	// Close the Output Chan after subdomain worker is done.
	go func() {
		subdomainsWG.Wait()
		close(output)
	}()

	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {

		domains <- sc.Text()

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
