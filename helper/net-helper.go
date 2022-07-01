package helper

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var activeProxies = make([]string, 0)
var totalProxies int

func FindActiveProxies(chunkSize int, proxies []string) []string {

	totalProxies = len(proxies)

	chunkCount, chunkProxies := chunkProxies(proxies, chunkSize)

	wg.Add(chunkCount)

	for chunkNumber, proxiesOfEachChunk := range chunkProxies {
		go checkAndAddActiveProxies(chunkNumber, proxiesOfEachChunk)
	}

	wg.Wait()

	return activeProxies
}

func chunkProxies(proxies []string, chunkSize int) (int, [][]string) {
	var chunkProxies [][]string
	chunkCount := len(proxies) / chunkSize

	for i := 0; i < len(proxies); i += chunkSize {
		end := i + chunkSize

		if end > len(proxies) {
			end = len(proxies)
		}

		chunkProxies = append(chunkProxies, proxies[i:end])
	}

	return chunkCount, chunkProxies
}

func checkAndAddActiveProxies(chunkNumber int, proxies []string) {

	defer wg.Done()

	for _, proxy := range proxies {
		fmt.Printf("remained proxies to check: %v\n", totalProxies)
		proxyUrl := "http://" + proxy
		req, _ := http.NewRequest("GET", proxyUrl, nil)
		req.Host = "google.com"
		res, err := http.DefaultClient.Do(req)

		if err != nil {
			totalProxies -= 1
			continue
		}

		defer res.Body.Close()

		if res.StatusCode != 200 {
			totalProxies -= 1
			continue
		}

		totalProxies -= 1
		activeProxies = append(activeProxies, proxy)
	}

}
