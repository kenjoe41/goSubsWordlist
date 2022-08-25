package ezutils

import (
	"strings"

	"github.com/elliotwutingfeng/go-fasttld"
)

func ExtractSubdomain(url string, includeRootPtr bool, extract *fasttld.FastTLD) string {
	result, _ := extract.Extract(fasttld.URLParams{URL: url})
	if len(result.SubDomain) > 0 {
		if includeRootPtr {
			return result.SubDomain + "." + result.Domain
		}
		return result.SubDomain
	}
	if includeRootPtr {
		return result.Domain
	}
	return ""
}

func SplitOnDot(subdomain string) []string {
	return strings.Split(subdomain, ".")
}
