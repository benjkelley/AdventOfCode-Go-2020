package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func check_year(year string, min_year int, max_year int) bool {

	year_int, err := strconv.Atoi(year)
	if err != nil {
		panic(err)
	}
	if year_int >= min_year && year_int <= max_year {
		return true
	} else {
		return false
	}
}

func check_height(height string) bool {

	if strings.HasSuffix(height, "cm") {
		elements := strings.Split(height, "cm")
		height_int, err := strconv.Atoi(elements[0])
		if err != nil {
			panic(err)
		}
		if height_int >= 150 && height_int <= 193 {
			return true
		} else {
			return false
		}

	} else if strings.HasSuffix(height, "in") {
		elements := strings.Split(height, "in")
		height_int, err := strconv.Atoi(elements[0])
		if err != nil {
			panic(err)
		}
		if height_int >= 59 && height_int <= 76 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func check_hair(color string) bool {

	if strings.HasPrefix(color, "#") {
		elements := strings.Split(color, "#")
		characters := strings.Split(elements[1], "")
		if len(characters) == 6 {
			for _, char := range characters {
				_, err := strconv.Atoi(char)
				if err != nil {
					if char != "a" && char != "b" && char != "c" && char != "d" && char != "e" && char != "f" {
						return false
					}
				}
			}
			return true
		}
	}
	return false

}

func check_eyes(color string) bool {

	valid_colors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, ec := range valid_colors {
		if color == ec {
			return true
		}
	}
	return false
}

func check_pid(pid string) bool {

	elements := strings.Split(pid, "")
	if len(elements) == 9 {
		for _, num := range elements {
			_, err := strconv.Atoi(num)
			if err != nil {
				return false
			}
		}
		return true
	}
	return false
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
			birth_valid := check_year(passport["byr"], 1920, 2002)
			issue_valid := check_year(passport["iyr"], 2010, 2020)
			expiration_valid := check_year(passport["eyr"], 2020, 2030)
			height_valid := check_height(passport["hgt"])
			color_valid := check_hair(passport["hcl"])
			eyes_valid := check_eyes(passport["ecl"])
			pid_valid := check_pid(passport["pid"])
			if birth_valid == true && issue_valid == true && expiration_valid == true && height_valid == true && color_valid == true && eyes_valid == true && pid_valid == true {
				validPassports++
			}
		}
	}

	fmt.Printf("Valid Passports: %d\n", validPassports)
	fmt.Printf("Total Passports: %d\n", len(passports))
}
