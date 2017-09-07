package Solver

import (
	"reflect"
	"testing"
)

func TestGetPuzzle(t *testing.T) {
	type args struct {
		i string
	}
	tests := []struct {
		name string
		args args
		want [9][9]int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPuzzle(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPuzzle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEmptyPuzzle(t *testing.T) {
	var b = GetEmptyPuzzle()

	b[0][1] = 4

	want := GetEmptyPuzzle()

	if want != b {
		t.Errorf("GetNewStart() = %v, want %v", b, want)

	}

	b = GetEmptyPuzzle()
	if want != b {
		t.Errorf("GetNewStart() = %v, want %v", b, want)

	}

}

func TestGetNewStart(t *testing.T) {
	tests := []struct {
		name string
		want [9][9]int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNewStart(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNewStart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getShuffledRows(t *testing.T) {
	type args struct {
		grid [3]int
	}
	tests := []struct {
		name string
		args args
		want [6]int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getShuffledRows(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getShuffledRows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeCellsEasy(t *testing.T) {
	type args struct {
		board [9][9]int
	}
	tests := []struct {
		name string
		args args
		want [9][9]int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeCellsEasy(tt.args.board); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeCellsEasy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandom(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Random(tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("Random() = %v, want %v", got, tt.want)
			}
		})
	}
}
