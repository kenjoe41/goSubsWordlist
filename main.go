package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kenjoe41/goSubsWordlist/utils"
)

func main() {

	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		scDomain := strings.TrimSpace(strings.ToLower(sc.Text()))

		domain := utils.CleanDomain(scDomain)

		// fmt.Println(domain)

		if domain == "" {
			// Log something but continue to next domain if available
			log.Printf("Failed to get domain from: %s", scDomain)
			continue
		}

		subdomain := utils.ExtractSubdomain(domain)

		if subdomain == "" {
			// Log something but continue to next domain if available
			log.Printf("Failed to get subdomain for domain: %s", domain)
			continue
		}

		// Split the subdomain into separate words by the '.' char.
		// Returns slice of words.
		subWords := utils.SplitSubToWords(subdomain)
		// fmt.Println(subWords)

		// Print to console for now
		for _, subword := range subWords {
			fmt.Println(subword)
		}

		// check there were no errors reading stdin (unlikely)
		if err := sc.Err(); err != nil {
			log.Fatal(err)
		}
	}

}
