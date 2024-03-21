package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var plots []int
	// numWords := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		var nums []rune
		var str []rune
		// var letters []rune
		// add string one rune at a time and check if it is a digit or a word
		// add to nums if so

		for _, r := range line {
			// first line
			str = append(str, r)
			// str = ["hcpjssql4kjhbcqzkvr2fivebpllzqbkhg"]

			if unicode.IsDigit(r) {
				nums = append(nums, r)
				// nums = [4, 2, 5]
				str = []rune{}
			}




			// } else if unicode.IsLetter(r) {
			// 	letters = append(letters, r)
			// }

				switch {
				case strings.Contains(string(str), "zero"):
					nums = append(nums, 0)
					str = []rune{}
				case strings.Contains(string(str), "one"):
					nums = append(nums, 1)
					str = []rune{}
				case strings.Contains(string(str), "two"):
					nums = append(nums, 2)
					str = []rune{}
				case strings.Contains(string(str), "three"):
					nums = append(nums, 3)
					str = []rune{}
				case strings.Contains(string(str), "four"):
					nums = append(nums, 4)
					str = []rune{}
				case strings.Contains(string(str), "five"):
					nums = append(nums, 5)
					str = []rune{}
				case strings.Contains(string(str), "six"):
					nums = append(nums, 6)
					str = []rune{}
				case strings.Contains(string(str), "seven"):
					nums = append(nums, 7)
					str = []rune{}
				case strings.Contains(string(str), "eight"):
					nums = append(nums, 8)
					str = []rune{}
				case strings.Contains(string(str), "nine"):
					nums = append(nums, 9)
					str = []rune{}
				}
			}

			if len(nums) >= 2 {
				first, _ := strconv.Atoi(string(nums[0]))
				last, _ := strconv.Atoi(string(nums[len(nums)-1]))
				combined, _ := strconv.Atoi(fmt.Sprintf("%d%d", first, last))
				plots = append(plots, combined)
			} else if len(nums) == 1 {
				num, _ := strconv.Atoi(strings.Repeat(string(nums[0]), 2))
				plots = append(plots, num)
			}

			// fmt.Println(string(str))
			// fmt.Println(nums)
			str = []rune{}
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

// func checkSubstrings(str string, subs ...string) []string {
// 	var matches []string

// 	for _, sub := range subs {
// 		if strings.Contains(str, sub) {
// 			matches = append(matches, sub)
// 		}
// 	}

// 	return matches
// }

// func main() {
// 	file, err := os.Open("input.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)

// 	var plots []int

// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		var nums []int
// 		// var letters []rune
// 		// numWords := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

// 		for _, r := range line {
// 			if unicode.IsDigit(r) {
// 				nums = append(nums, int(r))
// 			}
// 			// if unicode.IsLetter(r) {
// 			// 	letters = append(letters, r)
// 			// }
// 		}

// 		// 	for _, word := range checkSubstrings(line, numWords...) {
// 		// 	switch {
// 		// 		case word == "zero":
// 		// 			nums = append(nums, 0)
// 		// 		case word == "one":
// 		// 			nums = append(nums, 1)
// 		// 		case word == "two":
// 		// 			nums = append(nums, 2)
// 		// 		case word == "three":
// 		// 			nums = append(nums, 3)
// 		// 		case word == "four":
// 		// 			nums = append(nums, 4)
// 		// 		case word == "five":
// 		// 			nums = append(nums, 5)
// 		// 		case word == "six":
// 		// 			nums = append(nums, 6)
// 		// 		case word == "seven":
// 		// 			nums = append(nums, 7)
// 		// 		case word == "eight":
// 		// 			nums = append(nums, 8)
// 		// 		case word == "nine":
// 		// 			nums = append(nums, 9)
// 		// 	}
// 		// }

// 		if len(nums) >= 2 {
// 			first := nums[0]
// 			last := nums[len(nums) - 1]
// 			combined, _ := strconv.Atoi(fmt.Sprintf("%d%d", first, last))
// 			plots = append(plots, combined)
// 		} else if len(nums) == 1 {
// 			num := nums[0] * 11
// 			plots = append(plots, num)
// 		}
// 	}

// 	err = scanner.Err()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	var total int
// 	for _, v := range plots {
// 		total += v
// 	}
// 	fmt.Println(plots)

// 	fmt.Println(total)
// }

// func checkSubstrings(str string, subs ...string) []string {
// 	var matches []string

// 	for _, sub := range subs {
// 		if strings.Contains(str, sub) {
// 			matches = append(matches, sub)
// 		}
// 	}

// 	return matches
// }
