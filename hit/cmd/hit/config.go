package main

import (
	"flag"
)


type config struct {
	url string
	n   int
	c   int
	rps int
}

func parseArgs(c *config, args []string) error {
	fs :=flag.NewFlagSet("hit", flag.ContinueOnError)
	fs.StringVar(
		&c.url,
		"url",
		"",
		"HTTP serev 'URL' (required)",
	)
    fs.IntVar(
		&c.n,
		"n",
		1,
		"Number of requests to perform",
	)
	fs.IntVar(
		&c.c,
		"c",
		1,
		"Concurrency level: Number of multiple requests to make at a time",
	)
	fs.IntVar(
		&c.rps,
		"rps",
		0,
		"Rate limit for requests per second (0 for no limit)",
	)

	return fs.Parse(args)
}
