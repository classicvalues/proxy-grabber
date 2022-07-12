package input

import (
	"fmt"
	"proxy-grabber/helper"
	"strings"
)

var fileName string
var chunkSize int = 50
var proxyType int = 4
var timeoutTime int = 3

const (
	Http   = 1
	Https  = 2
	Socks5 = 3
	All    = 4
)

func ChangeReverseProxyTimeout() {
	fmt.Println("Please enter the timeout of reverse proxy for checking the grabbed proxies, by default it's 3 seconds (if you wanna use default just enter please!)")
	fmt.Scanln(&timeoutTime)

	helper.ChangeReverseProxyTimeOut(timeoutTime)
}

func EnterProxyType() int {
	fmt.Println("Please enter the Proxy Type which you wanna webscrape: (Please just enter the number of proxy type!)")
	fmt.Println("1 -> Http Proxy")
	fmt.Println("2 -> Https Proxy")
	fmt.Println("3 -> Socks5 Proxy")
	fmt.Println("4 -> All types")
	fmt.Println("By default the program will webscrape all types so if you wanna use default , please enter 4 or just enter")

	fmt.Scanln(&proxyType)

	switch proxyType {
	case Http:
		return Http
	case Https:
		return Https
	case Socks5:
		return Socks5
	default:
		return All
	}

}

func EnterFileName() string {
	fmt.Println("Please enter the file name, by default it's proxies.txt (if you wanna use default just enter please!)")
	fmt.Scanln(&fileName)

	if len(fileName) == 0 {
		fileName = "proxies.txt"
	}

	if !strings.Contains(fileName, ".txt") {
		panic("please enter the fileName with .txt extension")
	}

	return fileName
}

func EnterChunkSize() int {
	fmt.Println("Please enter the chunkSize, by default it's 50 (if you wanna use default just enter please!)")
	fmt.Scanln(&chunkSize)

	if chunkSize < 0 {
		panic("minus number for chunk size is Invalid!")
	}

	return chunkSize
}
