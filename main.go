package main

import (
	"fmt"
	"sync"
	"web-parser/counter"
	"web-parser/scrapper"
)

const k = 2
const subStr = "Go"

func main() {
	var URLs = []string{
		"https://golang.org",
		"https://go.dev/",
		"https://ru.wikipedia.org/wiki/Go",
		"https://blog.skillfactory.ru/glossary/golang/",
	}
	activeProcesses := make(chan int, k)
	wg := &sync.WaitGroup{}

	counters := counter.NewCounters()

	for _, url := range URLs {
		wg.Add(1)
		activeProcesses <- 1
		go GetStrOccurrencesCount(url, counters, activeProcesses, wg)
	}

	wg.Wait()
	fmt.Println()

	totalCount := 0
	for URL, occurrencesCount := range counters.LoadAll() {
		fmt.Printf("Count for %s: %d\n", URL, occurrencesCount)
		totalCount += occurrencesCount
	}

	fmt.Printf("Total: %d", totalCount)
}

func GetStrOccurrencesCount(URL string, counters *counter.Counters, activeProcesses chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	req := scrapper.NewRequest(URL)

	fmt.Printf("Doing request to %s...\n", URL)
	err := req.GetResponse()
	if err != nil {
		return
	}

	occurrencesCount := req.CountRepeatedStrInBody(subStr)
	counters.Store(URL, occurrencesCount)
	fmt.Printf("%s done!\n", URL)

	<-activeProcesses
}
