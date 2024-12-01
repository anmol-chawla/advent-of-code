package main

import (
	"bufio"
	"log"
	"math"
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

	differences := []float64{}

	for i := range first_line {
		num_one := first_line[i]
		num_two := second_line[i]

		difference := math.Abs(float64(num_one - num_two))
		differences = append(differences, difference)
	}

	var final_sum int
	for i := range differences {
		final_sum += int(differences[i])
	}

	log.Println(int(final_sum))
}
