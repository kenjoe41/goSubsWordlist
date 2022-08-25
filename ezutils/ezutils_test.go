package ezutils

import (
	"testing"

	"github.com/elliotwutingfeng/go-fasttld"
)

type extractSubdomainTest struct {
	url            string
	includeRootPtr bool
	expected       string
}

var extractSubdomainTests []extractSubdomainTest = []extractSubdomainTest{
	{url: "sub.example.com", includeRootPtr: false, expected: "sub"},
	{url: "sub.sub2.example.com", includeRootPtr: false, expected: "sub.sub2"},
	{url: "example.com", includeRootPtr: false, expected: ""},
	{url: "255.255.255.255", includeRootPtr: false, expected: ""},
	{url: "https://", includeRootPtr: false, expected: ""},
	{url: "", includeRootPtr: false, expected: ""},
	{url: "sub.example.com", includeRootPtr: true, expected: "sub.example"},
	{url: "sub.sub2.example.com", includeRootPtr: true, expected: "sub.sub2.example"},
	{url: "example.com", includeRootPtr: true, expected: "example"},
	{url: "255.255.255.255", includeRootPtr: true, expected: ""},
	{url: "https://", includeRootPtr: true, expected: ""},
	{url: "", includeRootPtr: true, expected: ""},
}

func TestExtractSubdomain(t *testing.T) {
	extract, err := fasttld.New(fasttld.SuffixListParams{})
	if err != nil {
		t.Errorf("%v", err)
	}
	for _, test := range extractSubdomainTests {
		subdomain := ExtractSubdomain(test.url, test.includeRootPtr, extract)
		if subdomain != test.expected {
			t.Errorf("Expected subdomain = %q but got %q", test.expected, subdomain)
		}
	}
}
