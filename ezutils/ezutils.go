package ezutils

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
	}
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
	}
	return dmn
}

func ExtractSubdomain(url string, includeRootPtr bool, extract *tldextract.TLDExtract) string {
	result := extract.Extract(url)
	if len(result.Sub) > 0 {
		if includeRootPtr {
			return result.Sub + "." + result.Root
		}
		return result.Sub
	}
	if includeRootPtr {
		return result.Root
	}
	return ""
}

func SplitSubToWords(subdomain string) []string {
	subWords := strings.Split(subdomain, ".")
	return subWords
}
