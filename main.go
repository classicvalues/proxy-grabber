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

	filePath, err := helper.CheckFileExistsOrCreate()

	if err != nil {
		panic(err.Error())
	}

	err = helper.TruncateFile()

	if err != nil {
		message := fmt.Sprintf("Truncating File Failed %v", err.Error())
		panic(message)
	}

	webscrape.WebScrapeProxyListNet(filePath)

}
