package helper

import (
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

	dirPath := os.Getenv(dirpathEnvName)

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err = os.Mkdir(dirPath, os.ModeAppend)

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
	f, err := os.OpenFile(fPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	defer f.Close()

	for _, proxy := range proxies {
		f.WriteString(proxy + "\n")
	}

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
