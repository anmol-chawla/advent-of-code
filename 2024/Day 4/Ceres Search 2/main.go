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

func CheckCrossMAS(x, y int, matrix [][]rune) bool {
	return CharAt(x, y, matrix) == 'A' &&
		(CharAt(x-1, y-1, matrix) == 'M' && CharAt(x-1, y+1, matrix) == 'M' && CharAt(x+1, y-1, matrix) == 'S' && CharAt(x+1, y+1, matrix) == 'S' || // Two M's above A
			CharAt(x+1, y-1, matrix) == 'M' && CharAt(x+1, y+1, matrix) == 'M' && CharAt(x-1, y-1, matrix) == 'S' && CharAt(x-1, y+1, matrix) == 'S' || // Two M's below A
			CharAt(x-1, y-1, matrix) == 'M' && CharAt(x+1, y-1, matrix) == 'M' && CharAt(x-1, y+1, matrix) == 'S' && CharAt(x+1, y+1, matrix) == 'S' || // Two M's left of A
			CharAt(x-1, y+1, matrix) == 'M' && CharAt(x+1, y+1, matrix) == 'M' && CharAt(x-1, y-1, matrix) == 'S' && CharAt(x+1, y-1, matrix) == 'S') // Two M's right of A
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
			if CheckCrossMAS(x, y, matrix) {
				count++
			}
		}
	}

	log.Println(count)
}
