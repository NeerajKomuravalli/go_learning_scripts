package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	defer safeCounter.waitGroup.Done()
	var crawler func(url string, depth int, fetcher Fetcher)
	crawler = func(url string, depth int, fetcher Fetcher) {
		defer safeCounter.waitGroup.Done()
		if depth <= 0 {
			return
		}
		if _, err := safeCounter.cacheUrl[url]; err == true {
			return
		}
		body, urls, err := fetcher.Fetch(url)
		safeCounter.mu.Lock()
		safeCounter.cacheUrl[url] += 1
		safeCounter.mu.Unlock()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)
		for _, u := range urls {
			safeCounter.waitGroup.Add(1)
			go crawler(u, depth-1, fetcher)
		}
		return
	}
	safeCounter.waitGroup.Add(1)
	go crawler(url, depth, fetcher)
	return
}

func main() {
	safeCounter.waitGroup.Add(1)
	go Crawl("https://golang.org/", 4, fetcher)
	safeCounter.waitGroup.Wait()
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

type SafeCounter struct {
	cacheUrl  map[string]int
	mu        sync.Mutex
	waitGroup sync.WaitGroup
}

var safeCounter SafeCounter = SafeCounter{cacheUrl: make(map[string]int)}

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
