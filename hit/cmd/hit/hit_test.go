package main

import (
"strings"
"testing"
"io"
)

type testEnv struct {
    stdout strings.Builder
    stderr strings.Builder
}

type parseArgsTest struct {
    name string
    args []string
    want config
}

func testRun(args ...string) (*testEnv, error) {  
    var tenv testEnv  
    err := run(&env{  
        args:   append([]string{"hit"}, args...), 
        stdout: &tenv.stdout, 
      stderr: &tenv.stderr,  
        dryRun: true,  
    })
    return &tenv, err 
}

func TestRunValidInput(t *testing.T) {
    t.Parallel()

    tenv, err := testRun("https://github.com/inancgumus")
    if err != nil {
        t.Fatalf("got %q;\nwant nil err", err)
    }
    if n := tenv.stdout.Len(); n == 0 { 
        t.Errorf("stdout = 0 bytes; want >0")
    }
    if n := tenv.stderr.Len(); n != 0 { 
        t.Errorf(
"stderr = %d bytes; want 0; stderr:\n%s",
            n, tenv.stderr.String(),
        )
    }
}

func TestRunInvalidInput(t *testing.T) {
    t.Parallel()

    tenv, err := testRun(
        "-c=2", "-n=1", "invalid-url",    
    )
    if err == nil {
        t.Fatalf("got nil; want err")
    }
    if n := tenv.stderr.Len(); n == 0 {  
        t.Error("stderr = 0 bytes; want >0")
    }
}

func TestParseArgsValidInput(t *testing.T) {
    t.Parallel()  

    for _, tt := range []parseArgsTest{
        {
            name: "all_flags",
            args: []string{"-n=10", "-c=5", "-rps=5", "http://test"},
            want: config{n: 10, c: 5, rps: 5, url: "http://test"},
        },
    } {
t.Run(tt.name, func(t *testing.T) {
            t.Parallel()  

            var got config
            if err := parseArgs(&got, tt.args, io.Discard); err != nil {
                t.Fatalf("parseArgs() error = %v, want no error", err)
            }
            if got != tt.want {
                t.Errorf("flags = %+v, want %+v", got, tt.want)
            }
        })
    }
}

func TestParseArgsInvalidInput(t *testing.T) {
    t.Parallel()

    for _, tt := range []parseArgsTest{
        {name: "n_syntax",   args: []string{"-n=ONE", "http://test"}},
        {name: "n_zero",     args: []string{"-n=0",   "http://test"}},
        {name: "n_negative", args: []string{"-n=-1",  "http://test"}},
    } {
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel()
           err := parseArgs(&config{}, tt.args, io.Discard)
            if err == nil {
                t.Fatal("parseArgs() = nil, want error")
            }
        })
    }
}
