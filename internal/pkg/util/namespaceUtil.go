package util

import (
	"bufio"
	"os"
	"strings"
)

// 获取命名空间
func GetNameSpace(rootPath string) (string, error) {
	file := rootPath + "go.mod"
	if CheckFileIsExist(file) {
		file, err := os.Open(file)
		if err != nil {
			return "", err
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		scanner.Scan()
		line := scanner.Text()
		line = strings.Replace(line, "module ", "", -1)

		return line, nil
	}

	return "", nil
}
