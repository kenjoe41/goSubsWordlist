package cli

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"

	"github.com/elliotwutingfeng/go-fasttld"
	"github.com/kenjoe41/goSubsWordlist/ezutils"
	"github.com/kenjoe41/goSubsWordlist/output"
)

// Cli accepts a list of URLs, one URL per line, from stdin and generates a wordlist from all subdomains found in the list.
func Cli(includeRoot, silent bool) error {
	// Print Header text
	if !silent {
		output.Beautify()
	}

	// This is a CPU-bound task, increasing the threads beyond what's available will just make it slow so removed the flag option.
	concurrency := runtime.NumCPU()

	// This is divided up in the subroutine for loop, so a value below 2 is BS.
	if concurrency < 2 {
		concurrency = 2
	} else {
		// We have 2 channels to share the concurrency with, let's reassure them that they'll have equal share.
		concurrency *= 2
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
			extract, err := fasttld.New(fasttld.SuffixListParams{})
			if err != nil {
				log.Fatal(err) // unlikely
			}
			for domain := range domains {
				if domain == "" {
					// Log something but continue to next domain if available
					// log.Printf("Failed to get domain from: %s", domain)
					continue
				}
				subdomain := ezutils.ExtractSubdomain(domain, includeRoot, extract)

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
				subWords := strings.Split(inSubdomains, ".")

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

	// Check for stdin input
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		flag.Usage()
		return errors.New("No domains or urls detected. Hint: cat domains.txt | goSubsWordlist")
	}

	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		domains <- sc.Text()
	}

	// Close domains chan
	close(domains)

	// check there were no errors reading stdin (unlikely)
	if err := sc.Err(); err != nil {
		return err
	}

	// Wait until the output waitgroup is done
	outputWG.Wait()

	return nil
}
