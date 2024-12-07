package main

import (
	"bufio"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e.Error())
		panic(e)
	}
}

func CharAt(x, y int, matrix [][]rune) rune {

	if x < 0 || x >= len(matrix) {
		return ' '
	}

	if y < 0 || y >= len(matrix[x]) {
		return ' '
	}

	return matrix[x][y]
}

func CheckXMAS(x, y, dx, dy int, matrix [][]rune) bool {
	return CharAt(x, y, matrix) == 'X' &&
		CharAt(x+dx, y+dy, matrix) == 'M' &&
		CharAt(x+(2*dx), y+(2*dy), matrix) == 'A' &&
		CharAt(x+(3*dx), y+(3*dy), matrix) == 'S'
}

func main() {
	input_file, error := os.Open("input.txt")
	check(error)

	file_reader := bufio.NewScanner(input_file)

	var matrix [][]rune
	for file_reader.Scan() {
		input_line := []rune(file_reader.Text())
		matrix = append(matrix, input_line)
	}

	count := 0

	for x, row := range matrix {
		for y := range row {
			possible_directions := [][]int{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
			for _, direction := range possible_directions {
				if CheckXMAS(x, y, direction[0], direction[1], matrix) {
					count++
				}
			}
		}
	}

	log.Println(count)
}
