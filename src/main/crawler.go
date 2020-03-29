package main

import (
	"fmt"
	"sync"
	"time"
)

//Fetcher interface
type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

//Crawler struct
type Crawler struct {
	scheduledUrls map[string]bool
	mu            sync.Mutex
}

func (c *Crawler) visitScheduled(url string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.scheduledUrls[url]; ok {
		return true
	}
	c.scheduledUrls[url] = true
	return false
}

func newCrawler() *Crawler {
	return &Crawler{scheduledUrls: make(map[string]bool)}
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (c *Crawler) Crawl(url string, depth int, fetcher Fetcher) {
	var wg sync.WaitGroup

	if c.visitScheduled(url) || depth <= 0 {
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
		go func(url string) {
			defer wg.Done()
			c.Crawl(url, depth-1, fetcher)
		}(u)
	}

	wg.Wait()
	return
}

// CrawlDemo demo
func CrawlDemo() {
	crawler := newCrawler()
	start := time.Now()
	crawler.Crawl("https://golang.org/", 4, fetcher)
	fmt.Println("Crawler elapsed:", time.Since(start))
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	// simulate real fetch
	time.Sleep(time.Millisecond * 300)
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
