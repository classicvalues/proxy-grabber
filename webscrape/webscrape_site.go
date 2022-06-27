package webscrape

import (
	"fmt"
	"net/http"

	"proxy-grabber/helper"

	"github.com/PuerkitoBio/goquery"
)

var proxies = make([]string, 0)

func WebScrapeProxyListNet(filePath string) {

	fmt.Printf("grabbing proxies from %v ...\n", "ProxyListNet")

	res, err := http.Get("https://free-proxy-list.net")

	if err != nil {
		panic(err.Error())
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		message := fmt.Sprintf("Couldn't fetch site correctly, StatusCode: %v", res.StatusCode)
		panic(message)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		panic(err.Error())
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

	err = helper.WriteProxiesToFile(proxies)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("grabbed & wrote from %v\n", "ProxyListNet")

}
