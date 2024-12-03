package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e.Error())
		panic(e)
	}
}

func safe(report []int) bool {
	var increasing = report[0] < report[len(report)-1]
	for i := 1; i < len(report); i++ {
		if increasing && (report[i] <= report[i-1] || report[i]-report[i-1] > 3) ||
			!increasing && (report[i] >= report[i-1] || report[i-1]-report[i] > 3) {
			return false
		}
	}
	return true
}

func main() {
	file, error := os.Open("input.txt")
	check(error)

	file_reader := bufio.NewScanner(file)

	safe_reports := 0

	for file_reader.Scan() {
		line := file_reader.Text()
		numbers := strings.Split(line, " ")
		var number_list []int
		for i := range numbers {
			number, error := strconv.Atoi(numbers[i])
			check(error)
			number_list = append(number_list, number)
		}

		if safe(number_list) {
			safe_reports += 1
		}
	}

	log.Println(safe_reports)
}
