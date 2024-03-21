package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// open & read input file
// scan each line, add ints to temp slice
// take first and last elements of temp slice to create two digit number
// append new number to toSum slice
// sum all nums in toSum
// return value

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var plots []int

	for scanner.Scan() {
		line := scanner.Text()
		var nums []int
		numWords := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

		for _, r := range line {
			if unicode.IsDigit(r) {
				nums = append(nums, int(r))
			}

			wordsInLine := checkSubstrings(line, numWords...)

			for _, word := range wordsInLine {
				switch {
				case word == "zero":
					nums = append(nums, 0)
				case word == "one":
					nums = append(nums, 1)
				case word == "two":
					nums = append(nums, 2)
				case word == "three":
					nums = append(nums, 3)
				case word == "four":
					nums = append(nums, 4)
				case word == "five":
					nums = append(nums, 5)
				case word == "six":
					nums = append(nums, 6)
				case word == "seven":
					nums = append(nums, 7)
				case word == "eight":
					nums = append(nums, 8)
				case word == "nine":
					nums = append(nums, 9)
				}
			}
		}

		if len(nums) >= 2 {
			first := nums[0]
			last := nums[len(nums) - 1]
			combined := (first * 10) + last
			plots = append(plots, combined)
		} else if len(nums) == 1 {
			num := nums[0] * 11
			plots = append(plots, num)
		}
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println(err)
		return
	}

	var total int
	for _, v := range plots {
		total += v
	}

	fmt.Println(total)
}

func checkSubstrings(str string, subs ...string) []string {
	var matches []string

	for _, sub := range subs {
		if strings.Contains(str, sub) {
			matches = append(matches, sub)
		}
	}

	return matches
}
