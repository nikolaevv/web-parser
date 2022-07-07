package main

import (
	"log"
	"sync"
	"web-parser/counter"
	"web-parser/website"
)

const (
	processes = 5
	subStr    = "Go"
)

var (
	urls = []string{
		"https://golang.org",
		"https://go.dev/",
		"https://ru.wikipedia.org/wiki/Go",
		"http://golang-book.ru/",
		"https://golangify.com/",
		"https://tproger.ru/translations/golang-basics/",
		"https://goforum.info/",
	}
)

func main() {
	limit := make(chan int, processes)
	wg := &sync.WaitGroup{}
	c := counter.New()
	total := 0

	for _, url := range urls {
		wg.Add(1)
		limit <- 1
		go CountSubstrFromUrl(url, c, limit, wg)
	}

	wg.Wait()

	for url, count := range c.LoadAll() {
		log.Printf("Count for %s: %d\n", url, count)
		total += count
	}

	log.Printf("Total: %d\n", total)
}

func CountSubstrFromUrl(url string, c *counter.Counters, limit chan int, wg *sync.WaitGroup) {
	source := website.New(url)
	defer wg.Done()

	log.Printf("Doing request to %s...\n", url)
	if err := source.GetResponse(); err != nil {
		return
	}

	log.Printf("%s done!\n", url)
	c.Store(url, source.Count(subStr))

	<-limit
}
