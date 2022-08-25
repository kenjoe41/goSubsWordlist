package ezutils

import (
	"github.com/elliotwutingfeng/go-fasttld"
)

// ExtractSubdomain extracts the subdomain from a given url.
// If includeRootPtr is true, the second-level domain will be included
func ExtractSubdomain(url string, includeRootPtr bool, extract *fasttld.FastTLD) string {
	result, err := extract.Extract(fasttld.URLParams{URL: url})
	if err != nil {
		return ""
	}
	if result.HostType != fasttld.HostName {
		return ""
	}
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
