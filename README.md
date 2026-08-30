# goSubsWordlist

[![Go Reference](https://img.shields.io/badge/go-reference-blue?logo=go&logoColor=white&style=for-the-badge)](https://pkg.go.dev/github.com/kenjoe41/goSubsWordlist)
[![Go Report Card](https://goreportcard.com/badge/github.com/kenjoe41/goSubsWordlist?style=for-the-badge)](https://goreportcard.com/report/github.com/kenjoe41/goSubsWordlist)

[![GitHub license](https://img.shields.io/badge/LICENSE-MIT-GREEN?style=for-the-badge)](LICENSE)

Generate a wordlist from a list of already discovered subdomains, by splitting each one on its
`.` labels. Feed it subdomains you already found (e.g. from
[Roots](https://github.com/kenjoe41/Roots), `subfinder`, `amass`, etc.) and it prints the
individual words that make them up — useful as a permutation/bruteforce wordlist for finding
*more* subdomains.

## Install

```shell
go install -v github.com/kenjoe41/goSubsWordlist@latest
```

## Usage

goSubsWordlist reads subdomains from stdin, one per line, and prints words to stdout:

```shell
cat subdomains.txt | goSubsWordlist > words.txt
```

For `dev.api.example.com`, it prints `dev` and `api` — the labels below the registered domain.
The registered domain (`example.com`) is skipped by default; pass `-iR` to include it too:

```shell
cat subdomains.txt | goSubsWordlist -iR > words.txt
```

### Flags

| Flag      | Default | Description                                    |
|-----------|---------|-------------------------------------------------|
| `-iR`     | `false` | Include the root domain's own label in output.  |
| `-silent` | `false` | Suppress the startup banner.                    |

Output is deduplication-free and unsorted by design — pipe through `sort -u` if you need a
unique, sorted wordlist:

```shell
cat subdomains.txt | goSubsWordlist | sort -u > words.txt
```

## TODO

- `-top N` flag: output only the `N` most frequently occurring words (e.g. `-top 1000`).
  Requires in-memory word-occurrence tracking, which may not scale well for very large
  subdomain lists — needs benchmarking before landing.
