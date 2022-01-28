package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jacksontong/aoc-2017-day-2-go/matrix"
)

func main() {
	csv := flag.String("csv", "input.csv", "The input csv file")
	flag.Parse()

	input := readCSVFile(*csv)
	answer := matrix.CalculateChecksum(input)

	fmt.Printf("Answer: %d\n", answer)
}

// Read the csv input file and returns a 2d slice
func readCSVFile(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("unable to read input file "+path, err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("unable to parse file as CSV for "+path, err)
	}

	out := make([][]int, len(records))

	for i, row := range records {
		out[i] = make([]int, len(row))
		for j, str := range row {
			n, err := strconv.Atoi(str)
			if err != nil {
				out[i][j] = 0
			} else {
				out[i][j] = n
			}
		}
	}

	return out
}
