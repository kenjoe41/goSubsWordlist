package ezutils

import (
	"github.com/elliotwutingfeng/go-fasttld"
)

// ExtractSubdomain extracts the subdomain from a given URL.
// If includeRootPtr is true, the second-level domain will be included in the result.
func ExtractSubdomain(url string, includeRootPtr bool, extract *fasttld.FastTLD) string {
	result, err := extract.Extract(fasttld.URLParams{URL: url})
	if err != nil || result.HostType != fasttld.HostName {
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
