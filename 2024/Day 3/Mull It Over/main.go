package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		log.Fatal(e.Error())
		panic(e)
	}
}

func main() {
	input_file, error := os.Open("input.txt")
	check(error)

	file_reader := bufio.NewScanner(input_file)

	var input_lines []string
	for file_reader.Scan() {
		input_lines = append(input_lines, file_reader.Text())
	}

	var matched_substrings []string
	for i := range input_lines {
		match_condition := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
		matched_substrings = append(matched_substrings, match_condition.FindAllString(input_lines[i], -1)...)
	}

	total := 0

	for i := range matched_substrings {
		substring := matched_substrings[i]
		match_condition := regexp.MustCompile(`\d+`)
		numbers := match_condition.FindAllString(substring, -1)

		first_number, error := strconv.Atoi(numbers[0])
		check(error)

		second_number, error := strconv.Atoi(numbers[1])
		check(error)

		total += first_number * second_number
	}

	log.Print(total)
}
