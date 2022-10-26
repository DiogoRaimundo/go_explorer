package main

import (
	"fmt"
	"sync"
)

type StringSet map[string]struct{}

var empty struct{}

type SafeStringSet struct {
	mu sync.Mutex
	v  StringSet
}

func (m *SafeStringSet) TryAdd(key string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, exists := m.v[key]
	if exists {
		return false
	}

	m.v[key] = empty
	return true
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	set := SafeStringSet{v: make(map[string]struct{})}
	var wg sync.WaitGroup

	var crawler func(string, int)
	crawler = func(url string, depth int) {
		defer wg.Done()

		// TODO: Fetch URLs in parallel.
		// TODO: Don't fetch the same URL twice.
		// This implementation doesn't do either:
		if depth <= 0 {
			return
		}

		if !set.TryAdd(url) {
			return
		}

		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("found: %s %q\n", url, body)
		for _, u := range urls {
			wg.Add(1)
			go crawler(u, depth-1)
		}
	}

	wg.Add(1)
	crawler(url, depth)
	wg.Wait()
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
