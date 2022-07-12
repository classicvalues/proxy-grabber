package webscrape

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"proxy-grabber/helper"
	"strings"
	"sync"

	"proxy-grabber/input"

	"github.com/PuerkitoBio/goquery"
)

var proxies = make([]string, 0)
var wg sync.WaitGroup

func InitializeWebScrapeProxies(proxyType int) []string {

	switch proxyType {
	case input.Http:
		result := initializeWebscrapeHttpProxies()
		proxies = []string{}
		return result
	case input.Https:
		result := initializeWebscrapeHttpsProxies()
		proxies = []string{}
		return result
	case input.Socks5:
		result := initializeWebscrapeSocks5Proxies()
		proxies = []string{}
		return result
	default:
		result := initializeWebscrapeAllProxies()
		proxies = []string{}
		return result
	}

}

func initializeWebscrapeHttpProxies() []string {

	wg.Add(4)

	go webScrapeProxyListNet(input.Http)
	go webScrapeTheSpeedXGithub(input.Http)
	go webscrapeClarketmGithub()
	go webScrapeJetkaiGithub(input.Http)

	wg.Wait()

	fmt.Printf("Removing duplicate http proxies\n")

	uniqueProxies := helper.RemoveDuplicateProxies(proxies)

	fmt.Printf("Duplicate http proxies removed\n")

	return uniqueProxies
}

func initializeWebscrapeHttpsProxies() []string {

	wg.Add(2)

	go webScrapeProxyListNet(input.Https)
	go webScrapeJetkaiGithub(input.Https)

	wg.Wait()

	fmt.Printf("Removing duplicate https proxies\n")

	uniqueProxies := helper.RemoveDuplicateProxies(proxies)

	fmt.Printf("Duplicate https proxies removed\n")

	return uniqueProxies
}

func initializeWebscrapeSocks5Proxies() []string {

	wg.Add(2)

	go webScrapeTheSpeedXGithub(input.Socks5)
	go webScrapeJetkaiGithub(input.Socks5)

	wg.Wait()

	fmt.Printf("Removing duplicate socks5 proxies\n")

	uniqueProxies := helper.RemoveDuplicateProxies(proxies)

	fmt.Printf("Duplicate socks5 proxies removed\n")

	return uniqueProxies
}

func initializeWebscrapeAllProxies() []string {

	var allTypeOfProxies []string

	httpProxies := initializeWebscrapeHttpProxies()
	proxies = []string{}
	httpsProxies := initializeWebscrapeHttpsProxies()
	proxies = []string{}
	socks5Proxies := initializeWebscrapeSocks5Proxies()

	allTypeOfProxies = append(allTypeOfProxies, httpProxies...)
	allTypeOfProxies = append(allTypeOfProxies, httpsProxies...)
	allTypeOfProxies = append(allTypeOfProxies, socks5Proxies...)

	return allTypeOfProxies
}

func webScrapeProxyListNet(proxyType int) {
	defer wg.Done()

	res, err := http.Get("https://free-proxy-list.net")

	if err != nil {
		log.Fatalf("Error in WebScrapeProxyListNet info: %v\n", err.Error())
		return
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		message := fmt.Sprintf("Couldn't fetch site correctly, StatusCode: %v", res.StatusCode)
		log.Fatalf("Error in WebScrapeProxyListNet info: %v\n", message)
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatalf("Error in WebScrapeProxyListNet info: %v\n", err.Error())
		return
	}

	doc.Find("#list table tr").Each(func(i int, s1 *goquery.Selection) {
		nodes := s1.Find("td").Nodes

		if len(nodes) == 0 {
			return
		}

		ip := nodes[0].FirstChild.Data
		port := nodes[1].FirstChild.Data
		isHttps := nodes[6].FirstChild.Data

		if proxyType == input.Https && isHttps == "no" {
			return
		}

		proxy := fmt.Sprintf("%v:%v", ip, port)
		proxies = append(proxies, proxy)
	})

}

func webScrapeTheSpeedXGithub(proxyType int) {
	defer wg.Done()

	var siteUrl string

	switch proxyType {
	case input.Http:
		siteUrl = "https://raw.githubusercontent.com/TheSpeedX/SOCKS-List/master/http.txt"
	case input.Socks5:
		siteUrl = "https://raw.githubusercontent.com/TheSpeedX/PROXY-List/master/socks5.txt"
	}

	res, err := http.Get(siteUrl)

	if err != nil {
		log.Fatalf("Error in WebScrapeTheSpeedXGithub info: %v\n", err.Error())
		return
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		message := fmt.Sprintf("Couldn't fetch site correctly, StatusCode: %v", res.StatusCode)
		log.Fatalf("Error in WebScrapeTheSpeedXGithub info: %v\n", message)
		return
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("Error in WebScrapeTheSpeedXGithub info: %v\n", err.Error())
		return
	}

	cbts := string(body)

	speedxProxies := strings.Split(cbts, "\n")

	proxies = append(proxies, speedxProxies...)

}

func webscrapeClarketmGithub() {
	defer wg.Done()

	res, err := http.Get("https://raw.githubusercontent.com/Clarketm/proxy-list/master/proxy-list-raw.txt")

	if err != nil {
		log.Fatalf("Error in WebscrapeClarketmGithub info: %v\n", err.Error())
		return
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		message := fmt.Sprintf("Couldn't fetch site correctly, StatusCode: %v", res.StatusCode)
		log.Fatalf("Error in WebscrapeClarketmGithub info: %v\n", message)
		return
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("Error in WebscrapeClarketmGithub info: %v\n", err.Error())
		return
	}

	cbts := string(body)

	clarketmProxies := strings.Split(cbts, "\n")

	proxies = append(proxies, clarketmProxies...)

}

func webScrapeJetkaiGithub(proxyType int) {
	defer wg.Done()

	var siteUrl string

	switch proxyType {
	case input.Http:
		siteUrl = "https://raw.githubusercontent.com/jetkai/proxy-list/main/online-proxies/txt/proxies-http.txt"
	case input.Https:
		siteUrl = "https://raw.githubusercontent.com/jetkai/proxy-list/main/online-proxies/txt/proxies-https.txt"
	case input.Socks5:
		siteUrl = "https://raw.githubusercontent.com/jetkai/proxy-list/main/online-proxies/txt/proxies-socks5.txt"
	}

	res, err := http.Get(siteUrl)

	if err != nil {
		log.Fatalf("Error in WebScrapeJetkaiHttpGithub info: %v\n", err.Error())
		return
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		message := fmt.Sprintf("Couldn't fetch site correctly, StatusCode: %v", res.StatusCode)
		log.Fatalf("Error in WebScrapeJetkaiHttpGithub info: %v\n", message)
		return
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("Error in WebScrapeJetkaiHttpGithub info: %v\n", err.Error())
		return
	}

	cbts := string(body)

	clarketmProxies := strings.Split(cbts, "\n")

	proxies = append(proxies, clarketmProxies...)
}
