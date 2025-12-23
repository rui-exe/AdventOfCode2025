package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Open file
	file, err := os.Open("./input.txt")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	start := 50
	result := 0

	for scanner.Scan() {
		line := scanner.Text()
		n, _ := strconv.Atoi(line[1:])
		if line[0] == 'R' {
			start = (start + n) % 100
		} else {
			start = (start - n) % 100
		}
		if start == 0 {
			result++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	fmt.Println("result: ", result)

}
