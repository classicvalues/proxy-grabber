package input

import (
	"fmt"
	"strings"
)

var fileName string
var chunkSize int = 10

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
	fmt.Println("Please enter the chunkSize, by default it's 10 (if you wanna use default just enter please!)")
	fmt.Scanln(&chunkSize)

	if chunkSize < 0 {
		panic("minus number for chunk size is Invalid!")
	}

	return chunkSize
}
