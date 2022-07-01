package webscrape

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"proxy-grabber/helper"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

var proxies = make([]string, 0)
var wg sync.WaitGroup

func InitializeWebScrapeProxies() {

	wg.Add(5)

	go WebScrapeProxyListNet()
	go WebScrapeTheSpeedXGithub()
	go WebscrapeClarketmGithub()
	go WebScrapeJetkaiHttpGithub()
	go WebScrapeJetkaiHttpsGithub()

	wg.Wait()

	RunWriteProxiesToFile()
}

func WebScrapeProxyListNet() {

	fmt.Printf("grabbing proxies from %v ...\n", "ProxyListNet")

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

		proxy := fmt.Sprintf("%v:%v", ip, port)
		proxies = append(proxies, proxy)
	})

	fmt.Printf("grabbed from %v\n", "ProxyListNet")
	wg.Done()
}

func WebScrapeTheSpeedXGithub() {

	fmt.Printf("grabbing proxies from %v ...\n", "TheSpeedXGithub")

	res, err := http.Get("https://raw.githubusercontent.com/TheSpeedX/SOCKS-List/master/http.txt")

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

	fmt.Printf("grabbed from %v\n", "TheSpeedXGithub")
	wg.Done()

}

func WebscrapeClarketmGithub() {
	fmt.Printf("grabbing proxies from %v ...\n", "ClarketmGithub")

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

	fmt.Printf("grabbed from %v\n", "ClarketmGithub")
	wg.Done()
}

func WebScrapeJetkaiHttpGithub() {
	fmt.Printf("grabbing proxies from %v ...\n", "JetkaiHttpGithub")

	res, err := http.Get("https://raw.githubusercontent.com/jetkai/proxy-list/main/online-proxies/txt/proxies-http.txt")

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

	fmt.Printf("grabbed from %v\n", "JetkaiHttpGithub")
	wg.Done()
}

func WebScrapeJetkaiHttpsGithub() {
	fmt.Printf("grabbing proxies from %v ...\n", "JetkaiHttpsGithub")

	res, err := http.Get("https://raw.githubusercontent.com/jetkai/proxy-list/main/online-proxies/txt/proxies-https.txt")

	if err != nil {
		log.Fatalf("Error in WebScrapeJetkaiHttpsGithub info: %v\n", err.Error())
		return
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		message := fmt.Sprintf("Couldn't fetch site correctly, StatusCode: %v", res.StatusCode)
		log.Fatalf("Error in WebScrapeJetkaiHttpsGithub info: %v\n", message)
		return
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("Error in WebScrapeJetkaiHttpsGithub info: %v\n", err.Error())
		return
	}

	cbts := string(body)

	clarketmProxies := strings.Split(cbts, "\n")

	proxies = append(proxies, clarketmProxies...)

	fmt.Printf("grabbed from %v\n", "JetkaiHttpsGithub")
	wg.Done()
}

func RunWriteProxiesToFile() {

	fmt.Printf("proxies are writing into file\n")

	uniqueProxies := helper.RemoveDuplicateProxies(proxies)

	err := helper.WriteProxiesToFile(uniqueProxies)

	if err != nil {
		panic(err.Error())
	}

	proxies = []string{}

	fmt.Printf("proxies wrote in file!\n")
}
