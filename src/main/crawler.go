package main

import (
	"fmt"
	"sync"
)

//Fetcher interface
type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	defer wg.Done()

	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		if err.Error() != "cached" {
			fmt.Println(err)
		}
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		wg.Add(1)
		go Crawl(u, depth-1, fetcher)
	}
	return
}

var wg sync.WaitGroup

// CrawlDemo demo
func CrawlDemo() {
	wg.Add(1)
	Crawl("https://golang.org/", 4, fetcher)
	wg.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	if cache.v[url] == nil {
		if res, ok := f[url]; ok {
			cache.v[url] = res
			return res.body, res.urls, nil
		}
		cache.v[url] = &fakeResult{}
		return "", nil, fmt.Errorf("not found: %s", url)
	}
	return "", nil, fmt.Errorf("cached")
}

// SafeCache is safe to use concurrently.
type SafeCache struct {
	mu sync.Mutex
	v  map[string]*fakeResult
}

var cache = SafeCache{v: make(map[string]*fakeResult)}

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
