// package memo provides a concurrency-unsafe
// memoization of a function of type Func.
package memo

import "sync"

// A Memo caches the results of calling a Func.
type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]result
}

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

// result models the return values of the functions
type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	var mu sync.Mutex
	return &Memo{f: f, mu: mu, cache: make(map[string]result)}
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()
	if !ok {
		res.value, res.err = memo.f(key)

		// Between the two critical sections, several goroutines
		// may race to compute f(key) and update the map
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	return res.value, res.err
}
