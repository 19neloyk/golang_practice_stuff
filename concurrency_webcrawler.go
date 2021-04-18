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
	// This implementation doesn't do either:
	pages := make( map[string]bool )
	pages[url] = false
	var mu sync.Mutex 
	printer := make (chan string)
	done := make (chan bool)
	go CrawlHelper(url, depth, fetcher, &pages, mu, printer, done)
	go func () {
		if (<-done) {
			close(printer)
		}
	}()
	
	for i:= range printer {
		fmt.Println(i)
	}
	
}

func CrawlHelper(url string, depth int, fetcher Fetcher, pages *map[string]bool, mu sync.Mutex, printer chan string, done chan bool) {
	mu.Lock()
	(*pages)[url] = true
	mu.Unlock()
	if depth <= 0 {
		done <- true
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		printer <- err.Error()
		done <- true
		return
	}
	printer <-fmt.Sprintf("found: %s %q\n", url, body)
	to_be_searched_count := 0
	subDone := make (chan bool)
	for _, u := range urls {
		mu.Lock()
		_, ok := (*pages)[u]
		mu.Unlock()
		if !ok {
			to_be_searched_count++
			mu.Lock()
			(*pages)[url] = false
			mu.Unlock()
			go CrawlHelper(u, depth-1, fetcher, pages, mu, printer, subDone)
		}
	}
	for i := 0 ; i < to_be_searched_count ; i ++ {
		<-subDone
	}
	done <- true
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
	time.Sleep(300*time.Millisecond)
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
