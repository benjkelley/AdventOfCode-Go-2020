package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parse_range(highlow string) (int, int) {
	re := regexp.MustCompile("-")
	split := re.Split(highlow, 2)
	low, _ := strconv.Atoi(split[0])
	high, _ := strconv.Atoi(split[1])
	return low, high
}

func check_valid(low int, high int, letter string, pass string) bool {
	letterCount := 0
	for _, ch := range strings.Split(pass, "") {
		if ch == letter {
			letterCount++
		}
	}
	if letterCount >= low && letterCount <= high {
		return true
	} else {
		return false
	}
}

func main() {
	inputs, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(inputs)
	var lines []string
	validPassCount := 0

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for _, line := range lines {
		items := strings.Fields(line)

		// find the range in the first field and separate into ints
		low, high := parse_range(items[0])

		// find the letter we are searching for
		letter := items[1][:len(items[1])-1]

		valid := check_valid(low, high, letter, items[2])

		if valid == true {
			validPassCount++
		}
	}

	fmt.Printf("There are %d valid passwords", validPassCount)

}
