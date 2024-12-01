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

	var file_reader *bufio.Scanner = bufio.NewScanner(file)

	var first_line []int
	var second_line []int
	for file_reader.Scan() {
		numbers := strings.Split(file_reader.Text(), "   ")
		first_number, error := strconv.Atoi(strings.Replace(numbers[0], " ", "", -1))
		check(error)

		second_number, error := strconv.Atoi(strings.Replace(numbers[1], " ", "", -1))
		check(error)

		first_line = append(first_line, first_number)
		second_line = append(second_line, second_number)
	}

	slices.Sort(first_line)
	slices.Sort(second_line)

	occurences_map := make(map[int]int)

	for i := range second_line {
		occurences_map[second_line[i]] += 1
	}

	similarity_score := 0

	for i := range first_line {
		second_line_repeats := occurences_map[first_line[i]]

		similarity_score += first_line[i] * second_line_repeats
	}

	log.Println(similarity_score)
}
