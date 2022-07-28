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

var wg sync.WaitGroup
var mu sync.Mutex

var static_map = map[string]int{}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, c chan string, quit chan string) {
	// TODO: Don't fetch the same URL twice.
	/*wg.Wait()
	mu.Lock()
	if static_map[url] == 0 {
		static_map[url] = 1
	} else {
		return
	}
	mu.Unlock()
	wg.Done()*/

	// Hint you may use a map
	// This implementation doesn't do either
	if depth <= 0 {
		println(depth, " deeep")
		quit <- "finished loop"
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	// TODO: Fetch URLs in parallel.
	// Hint: find which object should be shared
	c <- "found: " + url + body + "\n"
	println("depth: ", depth)
	for _, u := range urls {
		println(depth)
		Crawl(u, depth-1, fetcher, c, quit)
	}
	return
}

func main() {
	c := make(chan string)
	quit := make(chan string)

	go Crawl("https://golang.org/", 4, fetcher, c, quit)

	for {
		select {
		case val := <-c:
			fmt.Println(val)
		case <-quit:
			fmt.Println("Quit ")
		}
	}

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
