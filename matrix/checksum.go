package matrix

// Take a matrix and calculate the checksum of the matrix
// itegrate through each row of the matrix
// get the difference of the largest and smallest
// return the total of differences
func CalculateChecksum(matrix [][]int) int {
	differencesByRow := make([]int, len(matrix))

	// get the difference by row and push the result to differencesByRow slice
	for i, row := range matrix {
		min, max := findMinAndMax(row)
		differencesByRow[i] = max - min
	}

	// calculate the checksum
	out := 0
	for _, n := range differencesByRow {
		out += n
	}

	return out
}

// Take a slice of int, find the minimum and maximum numbers
func findMinAndMax(in []int) (min int, max int) {
	min = in[0]
	max = in[0]
	for _, n := range in {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	return min, max
}
