package main

import (
	"fmt"
	"os"
	"proxy-grabber/helper"
	"proxy-grabber/webscrape"
)

func main() {

	//current dir
	currentWd, _ := os.Getwd()
	//filename with extension
	fileName := "proxies.txt"

	//set Envs
	helper.SetEnvs(currentWd, fileName)

	_, err := helper.CheckFileExistsOrCreate()

	if err != nil {
		panic(err.Error())
	}

	err = helper.TruncateFile()

	if err != nil {
		message := fmt.Sprintf("Truncating File Failed %v", err.Error())
		panic(message)
	}

	proxies := webscrape.InitializeWebScrapeProxies()
	activeProxies := helper.FindActiveProxies(2, proxies)
	err = helper.WriteProxiesToFile(activeProxies)

	if err != nil {
		panic(err.Error())
	}

}
