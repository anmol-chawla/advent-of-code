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

func orderedAppend(number_list []int, number int) []int {
	list_length := len(number_list)

	if list_length == 0 {
		return append(number_list, number)
	}

	left := 0
	right := list_length

	for left < right {
		middle := int(uint(right+left) >> 1)
		if number_list[middle] <= number {
			left = middle + 1
		} else {
			right = middle
		}
	}
	return slices.Insert(number_list, left, number)
}

func main() {
	file, error := os.Open("input.txt")
	check(error)

	var file_reader *bufio.Scanner = bufio.NewScanner(file)

	var first_line []int
	var second_line []int

	occurences_map := make(map[int]int)

	for file_reader.Scan() {
		numbers := strings.Split(file_reader.Text(), "   ")
		first_number, error := strconv.Atoi(strings.Replace(numbers[0], " ", "", -1))
		check(error)

		second_number, error := strconv.Atoi(strings.Replace(numbers[1], " ", "", -1))
		check(error)

		first_line = orderedAppend(first_line, first_number)
		second_line = orderedAppend(second_line, second_number)
		occurences_map[second_number] += 1
	}

	similarity_score := 0

	for i := range first_line {
		second_line_repeats := occurences_map[first_line[i]]
		similarity_score += first_line[i] * second_line_repeats
	}

	log.Println(similarity_score)
}
