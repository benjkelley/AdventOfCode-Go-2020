package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Finally turning this into a function since importing + splitting lines seems to be part of each challenge
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

func main() {

	lines := import_data()
	// since scanner.Scan() truncates blank lines at the end of files, cheating a little to add a blank line to denote
	// the end of the dataset
	lines = append(lines, "")

	var passports []map[string]string
	currentPassport := make(map[string]string)
	requiredFields := make([]string, 8)
	requiredFields[0] = "byr"
	requiredFields[1] = "iyr"
	requiredFields[2] = "eyr"
	requiredFields[3] = "hgt"
	requiredFields[4] = "hcl"
	requiredFields[5] = "ecl"
	requiredFields[6] = "pid"
	requiredFields[7] = "cid"

	for _, line := range lines {
		if line != "" {
			elements := strings.Split(line, " ")

			for _, item := range elements {
				keyValue := strings.Split(item, ":")
				currentPassport[keyValue[0]] = keyValue[1]
			}
		} else {
			passports = append(passports, currentPassport)
			currentPassport = make(map[string]string)
		}
	}

	validPassports := 0

	for _, passport := range passports {
		validPass := true
		for _, field := range requiredFields {
			_, exists := passport[field]
			if exists == false && field != "cid" {
				validPass = false
				break
			}
		}
		if validPass == true {
			validPassports++
		}
	}

	fmt.Printf("Valid Passports: %d\n", validPassports)
	fmt.Printf("Total Passports: %d\n", len(passports))
}
