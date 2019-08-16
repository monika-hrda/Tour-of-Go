package main

import (
	"fmt"
	"sync"
)

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	visited := make(map[string]bool)
	crawl(url, depth, fetcher, visited)
}

var mux = sync.Mutex{}

func crawl(url string, depth int, fetcher Fetcher, visited map[string]bool) {
	// We've reached max depth. Exit.
	if depth <= 0 {
		return
	}
	// Claim the url for this goroutine.
	// If someone else has already claimed it, we'll just exit.
	mux.Lock()
	if _, ok := visited[url]; ok {
		mux.Unlock()
		return
	}
	visited[url] = true
	mux.Unlock()

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	done := make(chan bool)
	// Spawn as many goroutines as there are links on the page.
	for _, u := range urls {
		go func(url string) {
			crawl(url, depth-1, fetcher, visited)
			done <- true
		}(u)
	}
	// Wait for all of them to finish.
	for range urls {
		<-done
	}
	// We're done crawling this page. Go onto the next one.
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
