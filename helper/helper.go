package helper

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	dirpathEnvName = "dirPath"
	fNameEnvName   = "fName"
	fPath          = ""
)

func SetEnvs(dirPath string, fName string) {
	os.Setenv(dirpathEnvName, dirPath)
	os.Setenv(fNameEnvName, fName)
	fPath = filepath.Join(dirPath, fName)
}

func CheckFileExistsOrCreate() (string, error) {

	if _, err := os.Stat(fPath); os.IsNotExist(err) {
		_, err = os.Create(fPath)

		if err != nil {
			return "", err
		}
	}

	return fPath, nil
}

func TruncateFile() error {
	if err := os.Truncate(fPath, 0); err != nil {
		return err
	}
	return nil
}

func WriteProxiesToFile(proxies []string) error {
	fmt.Printf("proxies are writing into file\n")

	f, err := os.OpenFile(fPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	defer f.Close()

	for _, proxy := range proxies {
		f.WriteString(proxy + "\n")
	}

	fmt.Printf("proxies wrote in file!\n")

	return nil
}

func RemoveDuplicateProxies(proxies []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range proxies {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
