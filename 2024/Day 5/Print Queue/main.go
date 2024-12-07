package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e.Error())
		panic(e)
	}
}

func main() {
	file, error := os.Open("input.txt")
	check(error)

	file_reader := bufio.NewScanner(file)

	rule_map := make(map[int]int)

	numbers_before_map := make(map[int][]int)
	numbers_after_map := make(map[int][]int)

	for file_reader.Scan() {
		line := file_reader.Text()

		if line == "" {
			break
		}

		numbers := strings.Split(line, "|")

		first_number, error := strconv.Atoi(numbers[0])
		check(error)

		second_number, error := strconv.Atoi(numbers[1])
		check(error)

		rule_map[first_number] = second_number
		numbers_before_map[first_number] = append(numbers_before_map[first_number], second_number)
		numbers_after_map[second_number] = append(numbers_after_map[second_number], first_number)
	}

	var number_lines [][]int
	for file_reader.Scan() {
		line := file_reader.Text()

		numbers := strings.Split(line, ",")

		number_list := []int{}
		for i := range numbers {
			number, error := strconv.Atoi(numbers[i])
			check(error)
			number_list = append(number_list, number)
		}
		number_lines = append(number_lines, number_list)
	}

	sum := 0

	for i := range number_lines {
		incorrect := false
		for j := range number_lines[i] {
			number := number_lines[i][j]

			should_break := false
			for k := range number_lines[i] {
				current_number := number_lines[i][k]
				if k < j {
					if !slices.Contains(numbers_after_map[number], current_number) {
						should_break = true
					}
				} else if k > j {
					if !slices.Contains(numbers_before_map[number], current_number) {
						should_break = true
					}
				}
			}

			if should_break {
				incorrect = true
				break
			}
		}

		if !incorrect {
			middle_index := len(number_lines[i]) / 2
			sum += number_lines[i][middle_index]
		}
	}

	log.Println(sum)
}
