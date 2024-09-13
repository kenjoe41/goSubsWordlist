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
	domains := make(chan string)
	subdomains := make(chan string)
	words := make(chan string)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	// Domain input workers
	wg.Add(concurrency / 2)
	for i := 0; i < concurrency/2; i++ {
		go processDomains(ctx, &wg, domains, subdomains, includeRoot)
	}

	// Subdomain processing workers
	wg.Add(concurrency / 2)
	for i := 0; i < concurrency/2; i++ {
		go processSubdomains(ctx, &wg, subdomains, words)
	}

	// Output processor
	go func() {
		for word := range words {
			fmt.Println(word)
		}
	}()

	// Read input from stdin
	if err := readStdin(domains); err != nil {
		return err
	}
	close(domains)

	// Wait for workers to complete
	wg.Wait()
	close(subdomains)
	close(words)

	return nil
}

func processDomains(ctx context.Context, wg *sync.WaitGroup, domains <-chan string, subdomains chan<- string, includeRoot bool) {
	defer wg.Done()
	extract, err := fasttld.New(fasttld.SuffixListParams{})
	if err != nil {
		log.Fatal(err) // unlikely
	}
	for {
		select {
		case <-ctx.Done():
			return
		case domain, ok := <-domains:
			if !ok {
				return
			}
			if subdomain := ezutils.ExtractSubdomain(domain, includeRoot, extract); subdomain != "" {
				subdomains <- subdomain
			}
		}
	}
}

func processSubdomains(ctx context.Context, wg *sync.WaitGroup, subdomains <-chan string, words chan<- string) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case subdomain, ok := <-subdomains:
			if !ok {
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
		return errors.New("No domains or URLs detected. Hint: cat domains.txt | goSubsWordlist")
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
