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

var extractSubdomainTests = []extractSubdomainTest{
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
	// Initialize the fasttld extractor
	extract, err := fasttld.New(fasttld.SuffixListParams{})
	if err != nil {
		t.Fatalf("Failed to initialize fasttld: %v", err)
	}

	// Loop through test cases using subtests
	for _, test := range extractSubdomainTests {
		t.Run(test.url, func(t *testing.T) {
			subdomain := ExtractSubdomain(test.url, test.includeRootPtr, extract)
			if subdomain != test.expected {
				t.Errorf("For URL %q with includeRootPtr %v: expected %q, but got %q", test.url, test.includeRootPtr, test.expected, subdomain)
			}
		})
	}
}

func BenchmarkExtractSubdomain(b *testing.B) {
	// Initialize the fasttld extractor for benchmarks
	extract, err := fasttld.New(fasttld.SuffixListParams{})
	if err != nil {
		b.Fatalf("Failed to initialize fasttld: %v", err)
	}

	// Define a set of test URLs for benchmarking
	urls := []string{
		"a.b.c.d.e.example.com",
		"sub.example.com",
		"example.com",
		"255.255.255.255",
		"https://",
		"",
	}

	// Benchmark ExtractSubdomain with includeRootPtr = false
	b.Run("WithoutIncludeRootPtr", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, url := range urls {
				ExtractSubdomain(url, false, extract)
			}
		}
	})

	// Benchmark ExtractSubdomain with includeRootPtr = true
	b.Run("WithIncludeRootPtr", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, url := range urls {
				ExtractSubdomain(url, true, extract)
			}
		}
	})
}
