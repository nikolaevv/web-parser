package main

import (
	"fmt"
	"sync"
	"time"
	"web-parser/counters"
)

const k = 2

func main() {
	var URLs = []string{
		"https://golang.org",
		"https://go.dev/",
		"https://ru.wikipedia.org/wiki/Go",
		"https://blog.skillfactory.ru/glossary/golang/",
	}
	activeProcesses := make(chan int, k)
	wg := &sync.WaitGroup{}

	counters := counters.NewCounters()

	for _, url := range URLs {
		wg.Add(1)
		activeProcesses <- 1
		go GetStrOccurrencesCount(url, counters, activeProcesses, wg)
	}

	wg.Wait()

	totalCount := 0
	for URL, occurrencesCount := range counters.LoadAll() {
		fmt.Printf("Count for %s: %d\n", URL, occurrencesCount)
		totalCount += occurrencesCount
	}

	fmt.Printf("Total: %d", totalCount)
}

func GetStrOccurrencesCount(URL string, counters *counters.Counters, activeProcesses chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Doing big request to %s...\n", URL)
	time.Sleep(time.Second * 3)
	occurrencesCount := 0
	counters.Store(URL, occurrencesCount)
	<-activeProcesses
}
