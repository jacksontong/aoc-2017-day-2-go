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
			`5 9 2 8
			9 4 7 3
			3 8 6 5`,
			args{matrix: [][]int{
				{5, 9, 2, 8},
				{9, 4, 7, 3},
				{3, 8, 6, 5},
			}},
			9,
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

func Test_findDividable(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			"5 9 2 8",
			args{[]int{5, 9, 2, 8}},
			2,
			8,
		},
		{
			"9 4 7 3",
			args{[]int{9, 4, 7, 3}},
			3,
			9,
		},
		{
			"3 8 6 5",
			args{[]int{3, 8, 6, 5}},
			3,
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := findDividable(tt.args.in)
			if got != tt.want {
				t.Errorf("findDividable() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findDividable() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
