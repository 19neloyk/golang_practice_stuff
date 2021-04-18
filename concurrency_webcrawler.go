package main

import (
	"fmt"
	"time"
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
	fetched_urls := make(map[string]bool)
	fetched_urls [url] = true
	var mu sync.Mutex
	printer := make (chan string)
	go Crawl_Helper (url, depth, fetcher, &fetched_urls, mu, printer)
	//close(printer)
	for i := range printer {
		fmt.Println(i)
	}
	return
}

func Crawl_Helper(url string, depth int, fetcher Fetcher, fetched_urls *map[string]bool, mu sync.Mutex, printer chan string ){
	if depth <= 0 {
		(*fetched_urls)[url] = true
		return
	}
	_, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	mu.Lock()
	(*fetched_urls)[url] = true
	mu.Unlock()
	for _, u := range urls {
		mu.Lock()
		_, ok := (*fetched_urls)[u]
		mu.Unlock()
		//ok = false implies url has not been explored yet
		if !ok {
			mu.Lock()
			(*fetched_urls)[u] = false
			mu.Unlock()
			go Crawl_Helper (u, depth-1, fetcher, fetched_urls, mu, printer)
		}
	}
	fmt.Println(url)
	printer <-fmt.Sprintf("found: %s\n", url)
	return 
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
	time.Sleep(500 * time.Millisecond)
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
