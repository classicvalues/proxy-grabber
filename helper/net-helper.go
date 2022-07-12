package helper

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup
var activeProxies = make([]string, 0)
var totalProxies int
var client = http.Client{
	Timeout: 3 * time.Second,
}

func ChangeReverseProxyTimeOut(timeoutTime int) {
	client.Timeout = time.Duration(timeoutTime * int(time.Second))
}

func FindActiveProxies(chunkSize int, proxies []string) []string {

	totalProxies = len(proxies)

	chunkCount, chunkProxies := chunkProxies(proxies, chunkSize)

	if totalProxies%2 != 0 {
		chunkCount += 1
	}

	fmt.Printf("total proxies: %v\n", totalProxies)
	fmt.Printf("finding acive proxies , it may take a few minutes , please wait ...\n")

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
		proxyUrl := "http://" + proxy
		req, _ := http.NewRequest("GET", proxyUrl, nil)
		req.Host = "google.com"
		res, err := client.Do(req)

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
