package main

import (
	"fmt"
	"os"
	"proxy-grabber/helper"
	"proxy-grabber/input"
	"proxy-grabber/webscrape"
)

func main() {

	//filename with extension
	var fileName string
	var chunkSize int

	fileName = input.EnterFileName()
	chunkSize = input.EnterChunkSize()

	//current dir
	currentWd, _ := os.Getwd()
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
	activeProxies := helper.FindActiveProxies(chunkSize, proxies)
	err = helper.WriteProxiesToFile(activeProxies)

	if err != nil {
		panic(err.Error())
	}

}
