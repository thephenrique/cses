package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solution(lines []string) {
	n, err := strconv.Atoi(lines[0])
	if err != nil {
		log.Fatal(err)
	}

	rawList := strings.FieldsFunc(lines[1], func(r rune) bool {
		return r == ' '
	})

	numbers := make([]int, 0, len(rawList))
	for _, value := range rawList {
		parsedNumber, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}

		numbers = append(numbers, parsedNumber)
	}

	sort.Ints(numbers)

	for idx := range n {
		if idx+1 == numbers[idx] {
			continue
		}

		fmt.Println(idx + 1)
		break
	}
}

func main() {
	var buffer []byte
	var rawLines []string

	scanner := bufio.NewScanner(os.Stdin)
	// Increase buffer size to read long tokens from STDIN.
	scanner.Buffer(buffer, 1024*1024*10)

	for scanner.Scan() {
		line := scanner.Text()
		rawLines = append(rawLines, line)
	}

	if scanner.Err() != nil {
		log.Fatal((scanner.Err()))
	}

	solution(rawLines)
}
