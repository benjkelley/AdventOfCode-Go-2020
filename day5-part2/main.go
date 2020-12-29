package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func import_data() []string {
	inputs, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(inputs)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func find_seat(mapper string) int {
	elements := strings.Split(mapper, "")
	low_row := 0
	high_row := 127
	low_col := 0
	high_col := 7
	for _, letter := range elements {
		if letter == "F" {
			range_vals := high_row - low_row
			midpoint := range_vals / 2
			high_row = (high_row - midpoint) - 1
		} else if letter == "B" {
			range_vals := high_row - low_row
			midpoint := range_vals / 2
			low_row = low_row + midpoint + 1
		} else if letter == "L" {
			range_vals := high_col - low_col
			midpoint := range_vals / 2
			high_col = (high_col - midpoint) - 1
		} else if letter == "R" {
			range_vals := high_col - low_col
			midpoint := range_vals / 2
			low_col = low_col + midpoint + 1
		}
	}
	if (low_col == high_col) && (low_row == high_row) {
		return ((low_row * 8) + low_col)
	}
	return 0
}

func main() {

	var seat_ids []int
	lines := import_data()
	for _, line := range lines {
		seat_ids = append(seat_ids, find_seat(line))
	}

	sort.Ints(seat_ids)
	for i := 0; i < 960; i++ {
		base := i + 38
		if base != seat_ids[i] {
			fmt.Println(base)
			break
		}
	}

}
