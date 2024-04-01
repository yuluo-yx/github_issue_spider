package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func CountItems(path string) {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err.Error())
			return
		}
	}(file)

	scanner := bufio.NewScanner(file)

	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		count += strings.Count(line, "#")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("问题个数：", count)
}
