package hit

import (
	"net/http"
	"fmt"
)

// SendFunc is a type of function that sends an
// [http.Request] and returns a [Result].
type SendFunc func(*http.Request) Result

// Options defines the options for sending requests.
// Uses default options for unset options.
type Options struct {
    // Concurrency is the number of concurrent requests to send.
    // Default: 1
    Concurrency int

    // RPS is the requests to send persecond.
    // Default: 0 (no rate limiting)
    RPS int

    // Send processes requests.
    // Default: Uses [Send].
    Send SendFunc
}

// Defaults returns the default [Options].
func Defaults() Options {
    return withDefaults(Options{})  
}

func withDefaults(o Options) Options {  
    if o.Concurrency == 0 {
        o.Concurrency = 1
  }
    if o.Send == nil {  
        o.Send = func(r *http.Request) Result {  
            return Send(http.DefaultClient, r)   
        }   
    }
    return o
}

// SendN sends N requests using [Send].
// It returns a single-use [Results] iterator that
// pushes a [Result] for each [http.Request] sent.
func SendN(
    n int, req *http.Request, opts Options,
) (Results, error) {
    opts = withDefaults(opts) 
    if n <= 0 {
        return nil, fmt.Errorf("n must be positive: got %d", n)
    }
    // other checks are omitted for brevity  

    return func(yield func(Result) bool) {
        for range n {
            if !yield(opts.Send(req)) {  
                return
            }
        }
    }, nil
}
