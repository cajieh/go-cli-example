package main

import (
	"fmt"
	"flag"
	"errors"
    "strconv"
)

type positiveIntValue int

func asPositiveIntValue(p *int) *positiveIntValue {
    return (*positiveIntValue)(p) 
}

func (n *positiveIntValue) String() string {
    return strconv.Itoa(int(*n)) 
}

func (n *positiveIntValue) Set(s string) error {
    v, err := strconv.ParseInt( 
        s,  
        0, 
        strconv.IntSize,  
    )
    if err != nil {
        return err
    }
    if v <= 0 {
        return errors.New("should be greater than zero")
    }
    *n = positiveIntValue(v) 

	return nil

}

type config struct {
	url string
	n   int
	c   int
	rps int
}

func parseArgs(c *config, args []string) error {
	fs :=flag.NewFlagSet("hit", flag.ContinueOnError)
	fs.Usage = func() {
		fmt.Fprintf(fs.Output(), "usage %s [options] url\n", fs.Name())
		fs.PrintDefaults()
	}
	fs.StringVar(
		&c.url,
		"url",
		"",
		"HTTP serev 'URL' (required)",
	)
    fs.Var(asPositiveIntValue(&c.n),
		"n",
		"Number of requests to perform",
	)
	fs.Var(asPositiveIntValue(&c.c),
		"c",
		"Concurrency level: Number of multiple requests to make at a time",
	)
	fs.Var(asPositiveIntValue(&c.rps),
		"rps",
		"Rate limit for requests per second (0 for no limit)",
	)
    
	if err := fs.Parse(args); err != nil {
		return err
	}
	c.url = fs.Arg(0)
	return nil
}