package main

import (
	"fmt"
	"os"
	"strings"
)

func readFileByDelimiter(fileName string, delimiter string) []string {
	path := fmt.Sprintf("/inputs/%v", fileName)
	data, err := os.ReadFile(path)
	fmt.Println(string(data))
	splitData := strings.Split(string(data), delimiter)
	if err != nil {
		panic(err)
	}

	return splitData
}
