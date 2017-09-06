package Solver

import (
	"reflect"
	"testing"
	"time"
)

func Test_timeTrack(t *testing.T) {
	type args struct {
		start time.Time
		name  string
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			timeTrack(tt.args.start, tt.args.name)
		})
	}
}

func TestGrid_getCandidates(t *testing.T) {
	type args struct {
		c Cell
	}
	tests := []struct {
		name string
		grid *Grid
		args args
		want []int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.grid.getCandidates(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Grid.getCandidates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_contains(t *testing.T) {
	type args struct {
		s []int
		v int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.s, tt.args.v); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrid_solve(t *testing.T) {
	type args struct {
		c Cell
	}
	tests := []struct {
		name string
		grid *Grid
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.grid.solve(tt.args.c); got != tt.want {
				t.Errorf("Grid.solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrid_getNextCell(t *testing.T) {
	type args struct {
		c Cell
	}
	tests := []struct {
		name string
		grid *Grid
		args args
		want Cell
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.grid.getNextCell(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Grid.getNextCell() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSolver(t *testing.T) {
	type args struct {
		puzzle [9][9]int
		gen    bool
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
			if got := NewSolver(tt.args.puzzle, tt.args.gen); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSolver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrid_printGrid(t *testing.T) {
	tests := []struct {
		name string
		grid *Grid
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.grid.printGrid()
		})
	}
}

func Test_printBoard(t *testing.T) {
	type args struct {
		board [9][9]int
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printBoard(tt.args.board)
		})
	}
}

func TestSolver(t *testing.T) {

	puzzle := GetPuzzle("Easy")
	grid := new(Grid)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			grid.cells[i][j].value = puzzle[i][j]
			grid.cells[i][j].row = i
			grid.cells[i][j].col = j

		}
	}
	grid.cells[0][0].value = 9

	var actual = grid.solve(grid.cells[0][0])

	if actual == false {
		t.Errorf("Test Failed ", actual)

	}
}

func TestSolverAccuracy(t *testing.T) {

	var puzzle [9][9]int
	puzzle[0] = [9]int{5, 3, 0, 0, 7, 0, 0, 0, 0}
	puzzle[1] = [9]int{6, 0, 0, 1, 9, 5, 0, 0, 0}
	puzzle[2] = [9]int{0, 9, 8, 0, 0, 0, 0, 6, 0}
	puzzle[3] = [9]int{8, 0, 0, 0, 6, 0, 0, 0, 3}
	puzzle[4] = [9]int{4, 0, 0, 8, 0, 3, 0, 0, 1}
	puzzle[5] = [9]int{7, 0, 0, 0, 2, 0, 0, 0, 6}
	puzzle[6] = [9]int{0, 6, 0, 0, 0, 0, 2, 8, 0}
	puzzle[7] = [9]int{0, 0, 0, 4, 1, 9, 0, 0, 5}
	puzzle[8] = [9]int{0, 0, 0, 0, 8, 0, 0, 7, 9}

	var expected [9][9]int
	expected[0] = [9]int{5, 3, 4, 6, 7, 8, 9, 1, 2}
	expected[1] = [9]int{6, 7, 2, 1, 9, 5, 3, 4, 8}
	expected[2] = [9]int{1, 9, 8, 3, 4, 2, 5, 6, 7}
	expected[3] = [9]int{8, 5, 9, 7, 6, 1, 4, 2, 3}
	expected[4] = [9]int{4, 2, 6, 8, 5, 3, 7, 9, 1}
	expected[5] = [9]int{7, 1, 3, 9, 2, 4, 8, 5, 6}
	expected[6] = [9]int{9, 6, 1, 5, 3, 7, 2, 8, 4}
	expected[7] = [9]int{2, 8, 7, 4, 1, 9, 6, 3, 5}
	expected[8] = [9]int{3, 4, 5, 2, 8, 6, 1, 7, 9}

	var actual = NewSolver(puzzle, false)

	var test = reflect.DeepEqual(actual, expected)

	if !test {
		t.Errorf("Test Failed got %v expected %v ", actual, expected)

	}
}
