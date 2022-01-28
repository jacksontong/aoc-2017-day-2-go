package matrix

// Take a matrix and calculate the checksum of the matrix
// itegrate through each row of the matrix
// find the only two numbers in each row where one evenly divides the other
// return the total of differences
func CalculateChecksum(matrix [][]int) int {
	differencesByRow := make([]int, len(matrix))

	// get the difference by row and push the result to differencesByRow slice
	for i, row := range matrix {
		min, max := findDividable(row)
		differencesByRow[i] = max / min
	}

	// calculate the checksum
	out := 0
	for _, n := range differencesByRow {
		out += n
	}

	return out
}

// Take a slice of int find the 2 numbers
// that one evenly divides the other
// the result is a whole number
func findDividable(in []int) (min int, max int) {
loop:
	for i, m := range in {
		for j, n := range in {
			if i == j {
				continue
			}
			if m%n == 0 {
				min, max = n, m
				break loop
			} else if n%m == 0 {
				min, max = m, n
				break loop
			}
		}
	}

	return min, max
}
