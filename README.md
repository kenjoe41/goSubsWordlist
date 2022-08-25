# goSubsWordlist

Generate a wordlist from a list of already discovered subdomains.
This list can be used for further bruteforcing for more subdomains.

## Install

To install:

```shell
go install -v github.com/kenjoe41/goSubsWordlist@latest
```

## TODO

I plan to add:
    ..* `top N` flag. Outputs only most reoccuring word upto the Nth number like `-top 1000`.
        Might require in-memory tracking of word occurance, might not be efficient for xx-large huge subdomain lists.