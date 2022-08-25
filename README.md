# goSubsWordlist

[![Go Reference](https://img.shields.io/badge/go-reference-blue?logo=go&logoColor=white&style=for-the-badge)](https://pkg.go.dev/github.com/kenjoe41/goSubsWordlist)
[![Go Report Card](https://goreportcard.com/badge/github.com/kenjoe41/goSubsWordlist?style=for-the-badge)](https://goreportcard.com/report/github.com/kenjoe41/goSubsWordlist)

[![GitHub license](https://img.shields.io/badge/LICENSE-MIT-GREEN?style=for-the-badge)](LICENSE)

Generate a wordlist from a list of already discovered subdomains.
This list can be used for further bruteforcing for more subdomains.

## Install

To install:

```shell
go install -v github.com/kenjoe41/goSubsWordlist@latest
```

## TODO

I plan to add:
    ..* `top N` flag. Outputs only most reoccurring word up to the Nth number like `-top 1000`.
        Might require in-memory tracking of word occurrence, might not be efficient for xx-large huge subdomain lists.
