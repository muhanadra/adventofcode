package main

import (
	"testing"
)

func Test_getSum(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"It returns expected result when there are no duplicates",
			args{
				[]string{"1", "2", "3", "4"},
			},
			0,
		},
		{
			"produces a sum of 3 (1 + 2)",
			args{
				[]string{"1", "1", "2", "2"},
			},
			3,
		},
		{
			"produces 4 because each digit (all 1) matches the next",
			args{
				[]string{"1", "1", "1", "1"},
			},
			4,
		},
		{
			"produces 9 because the only digit that matches the next one is the last digit, 9.",
			args{
				[]string{"9", "1", "2", "1", "2", "1", "2", "9"},
			},
			9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSum(tt.args.arr); got != tt.want {
				t.Errorf("getSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
