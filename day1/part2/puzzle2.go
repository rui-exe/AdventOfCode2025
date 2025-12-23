package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	start := 50
	result := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		dir := line[0]
		nStr := line[1:]
		n, _ := strconv.Atoi(nStr)
		result += n / 100

		remainder := n % 100
		if remainder == 0 {
			continue
		}

		if dir == 'R' {
			if start+remainder >= 100 {
				result++
			}

			start = (start + remainder) % 100

		} else {
			if start > 0 && start-remainder <= 0 {
				result++
			}

			start -= remainder
			if start < 0 {
				start += 100
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Final password:", result)
}
