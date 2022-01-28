package matrix

import (
	"testing"
)

func TestCalculateChecksum(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			`5 1 9 5
			7 5 3 
			2 4 6 8`,
			args{matrix: [][]int{
				{5, 1, 9, 5},
				{7, 5, 3},
				{2, 4, 6, 8},
			}},
			18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateChecksum(tt.args.matrix); got != tt.want {
				t.Errorf("CalculateChecksum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMinAndMax(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name    string
		args    args
		wantMin int
		wantMax int
	}{
		{
			"5 1 9 5",
			args{[]int{5, 1, 9, 5}},
			1,
			9,
		},
		{
			"2 4 6 8",
			args{[]int{2, 4, 6, 8}},
			2,
			8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMin, gotMax := findMinAndMax(tt.args.in)
			if gotMin != tt.wantMin {
				t.Errorf("findMinAndMax() gotMin = %v, want %v", gotMin, tt.wantMin)
			}
			if gotMax != tt.wantMax {
				t.Errorf("findMinAndMax() gotMax = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}
