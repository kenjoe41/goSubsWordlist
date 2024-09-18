package cli

import (
	"bufio"
	"context"
	"errors"
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

func Cli(includeRoot, silent bool) error {
	if !silent {
		output.Beautify()
	}

	concurrency := runtime.NumCPU()
	if concurrency < 2 {
		concurrency = 2
	}

	// Channels and WaitGroups
	domainsChan := make(chan string)
	subdomainsChan := make(chan string)
	wordsChan := make(chan string)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Domain input workers
	var domainsWG sync.WaitGroup
	for i := 0; i < concurrency; i++ {
		domainsWG.Add(1)

		go processDomains(ctx, &domainsWG, domainsChan, subdomainsChan, includeRoot)
	}

	// Subdomain processing workers
	var subdomainsWG sync.WaitGroup
	for i := 0; i < concurrency; i++ {
		subdomainsWG.Add(1)
		go processSubdomains(ctx, &subdomainsWG, subdomainsChan, wordsChan)
	}

	// Close subdomains channel when done reading from domains chan.
	go func() {
		domainsWG.Wait()
		close(subdomainsChan)
	}()

	// Output processor
	var outputWG sync.WaitGroup
	outputWG.Add(1)
	go func() {
		for word := range wordsChan {
			fmt.Println(word)
		}
		outputWG.Done()
	}()

	// Close the Words Chan after subdomain worker is done.
	go func() {
		subdomainsWG.Wait()
		close(wordsChan)
	}()

	// Read input from stdin
	if err := readStdin(domainsChan); err != nil {
		return err
	}
	close(domainsChan)

	// Wait for workers to complete
	outputWG.Wait()

	return nil
}

func processDomains(ctx context.Context, domainsWG *sync.WaitGroup, domains <-chan string, subdomains chan<- string, includeRoot bool) {
	extract, err := fasttld.New(fasttld.SuffixListParams{})
	if err != nil {
		log.Fatal(err) // unlikely
	}
	for {
		select {
		case <-ctx.Done():
			domainsWG.Done()
			return
		case domain, ok := <-domains:
			if !ok {
				domainsWG.Done()
				return
			}
			if subdomain := ezutils.ExtractSubdomain(domain, includeRoot, extract); subdomain != "" {
				subdomains <- subdomain
			}
		}
	}
}

func processSubdomains(ctx context.Context, subdomainsWG *sync.WaitGroup, subdomains <-chan string, words chan<- string) {

	for {
		select {
		case <-ctx.Done():
			subdomainsWG.Done()
			return
		case subdomain, ok := <-subdomains:
			if !ok {
				subdomainsWG.Done()
				return
			}
			subWords := strings.Split(subdomain, ".")
			for _, word := range subWords {
				words <- word
			}
		}
	}
}

func readStdin(domains chan<- string) error {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		return errors.New("no domains or URLs detected. Hint: cat domains.txt | goSubsWordlist")
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		domains <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
