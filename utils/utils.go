package utils

import (
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/joeguo/tldextract"
)

func CleanDomain(dmn string) string {

	if strings.HasPrefix(dmn, "http") {

		u, err := url.Parse(dmn)
		if err != nil {
			log.Fatal(err)
		}

		return u.Hostname()

	} else {
		// Check if maybe it has '/' in it
		if strings.Contains(dmn, "/") {
			// Use regex or equivalent to get domain from string
			domainRegex := regexp.MustCompile(`[\w]+[\w\-_~\.]+\.[a-zA-Z]+|$`)

			match := domainRegex.FindString(dmn)
			if match != "" {
				return match
			}
			// We have no match
			return ""

		} else {
			return dmn
		}
	}
}

func ExtractSubdomain(url string) string {

	cache := "/tmp/tldsub.cache"
	extract, _ := tldextract.New(cache, false)

	result := extract.Extract(url)
	if len(result.Sub) > 0 {

		return result.Sub

	} else {
		return ""
	}
	// Let's first remove this, dont remember why i returned the root instead.
	// else {

	// 	return result.Root
	// }
}

func SplitSubToWords(subdomain string) []string {
	subWords := strings.Split(subdomain, ".")
	return subWords
}
