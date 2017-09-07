package Solver

import (
	"fmt"
	"log"
	"time"
)

// Basic data structure for individual Sudoku cells
type Cell struct {
	row      int
	col      int
	value    int
	solution int
}

// The grid struct is a representation of the complete sudoku board
// its values
type Grid struct {
	cells         [9][9]Cell
	Solved        bool
	SolutionCount int
	firstSolve    bool
	possibilities []int
}

// new ganme grid for copying to once the puzzle is solved
var newGameGrid [9][9]int

// variable showing the number of solution a puzzle has
var SolutionCount int

// Time track function found at https://gobyexample.com/
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

// This method gets the suitable candiates for the solver to try when placing a number in a cell
// the method will check for all legal moves and then return a slice containing those values
func (grid *Grid) getCandidates(c Cell) []int {
	possibilities := make([]int, 0)
	possibilities = possibilities[:0]

	for i := 0; i < 9; i++ {
		if grid.cells[c.row][i].value != 0 {

			possibilities = append(possibilities, grid.cells[c.row][i].value)

		} else {

		}
	}
	for i := 0; i < 9; i++ {
		if grid.cells[i][c.col].value != 0 {
			possibilities = append(possibilities, grid.cells[i][c.col].value)
		} else {
		}
	}
	//check grid //
	var nrow = c.row
	var ncol = c.col
	var x1 = 3 * (nrow / 3)
	var y1 = 3 * (ncol / 3)
	var x2 = x1 + 2
	var y2 = y1 + 2

	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			if grid.cells[i][j].value != 0 {
				// fmt.Println("found in grid")
				possibilities = append(possibilities, grid.cells[i][j].value)
			}
		}
	}
	possd := make([]int, 0)
	possd = possd[:0]

	for i := 1; i <= 9; i++ {
		if !contains(possibilities, i) {
			possd = append(possd, i)
		}
	}
	return possd
}

// Simple contains function to see if a value v is already in a given slice
// returns Boolean indicating the condition is true or false
func contains(s []int, v int) bool {
	for _, a := range s {
		if a == v {
			return true
		}
	}
	return false
}

// The main solve method
// solve method takes a single Cell as a paramter and recurively gos through each cell in the Grid trying to solve the puzzle
//
func (grid *Grid) solve(c Cell) bool {
	// defer timeTrack(time.Now(), "Solving")
	if c.value == 10 {

		if !grid.firstSolve {
			fmt.Println("HELP")
			for i := 0; i < 9; i++ {
				for j := 0; j < 9; j++ {
					newGameGrid[i][j] = grid.cells[i][j].value

				}
			}
			grid.firstSolve = true
			grid.SolutionCount = 99999

		}
		grid.Solved = true

		grid.SolutionCount++
		if grid.SolutionCount == 1 {
			for i := 0; i < 9; i++ {
				for j := 0; j < 9; j++ {
					newGameGrid[i][j] = grid.cells[i][j].value

				}
			}
			printBoard(newGameGrid)

		}
		fmt.Println("Solution count = ", grid.SolutionCount)
		grid.printGrid()
		return true

	}

	if c.value != 0 {
		// fmt.Println("the value before goin in", c.value)
		return grid.solve(grid.getNextCell(c))

	}
	var candidates = grid.getCandidates(c)

	// loop to get the cycle through the slice of candidates
	for _, a := range candidates {

		grid.cells[c.row][c.col].value = a
		grid.solve(grid.getNextCell(c))

		if grid.SolutionCount == 100000 {
			return true
		}

		if grid.SolutionCount > 1 {
			return true
		}

		grid.cells[c.row][c.col].value = 0

	}

	return false
}

// This method will return the next cell in the Grid
// when the end is reached the return is a cell with a value 10 to indicate to the solve
// method it is at the end of the array
func (grid *Grid) getNextCell(c Cell) Cell {

	var test Cell
	test.value = 10

	var row = c.row
	var col = c.col

	col++

	if col > 8 {
		col = 0
		row++
	}

	if row > 8 {
		// fmt.Println("This is the end")
		return test
	}

	return grid.cells[row][col]

}

// This Exorted function is how and external file will solve a puzzle
// The parameters are a 2d int array and a boolean to indicate whether it is solving or genereating a puzzle
// returns a solved 2d array
func NewSolver(puzzle [9][9]int, gen bool) [9][9]int {

	grid := new(Grid)
	grid.SolutionCount = 0
	grid.Solved = false
	grid.firstSolve = gen

	var b [9][9]int

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			grid.cells[i][j].value = puzzle[i][j]
			grid.cells[i][j].row = i
			grid.cells[i][j].col = j

		}
	}

	grid.solve(grid.cells[0][0])
	SolutionCount = grid.SolutionCount

	if grid.Solved {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				b[i][j] = newGameGrid[i][j]

			}
		}

		return b

	}

	return b

}

// simple function to display the grid data structure to the console
// will display as a Sudoku Game
func (grid *Grid) printGrid() {

	for i := 0; i < 9; i++ {

		if i == 0 || i == 3 || i == 6 || i == 9 {
			fmt.Println(" -  - - -  -  - - -  -  - - -  -  -  -")
		}

		for j := 0; j < 9; j++ {
			if j == 3 || j == 6 {
				fmt.Print("|")
			}
			var v = grid.cells[i][j].value
			if v == 0 {
				fmt.Printf("    ")
			} else {
				fmt.Printf("  %d ", v)
			}

		}
		fmt.Println()
	}

}

// Simple method to print a 2d int array to the console
// Displays as a Sudoku puzzle
func printBoard(board [9][9]int) {

	for i := 0; i < 9; i++ {

		if i == 0 || i == 3 || i == 6 || i == 9 {
			fmt.Println(" -  - - -  -  - - -  -  - - -  -  -  -")
		}

		for j := 0; j < 9; j++ {
			if j == 3 || j == 6 {
				fmt.Print("|")
			}
			var v = board[i][j]
			if v == 0 {
				fmt.Printf("    ")
			} else {
				fmt.Printf("  %d ", v)
			}

		}
		fmt.Println()
	}

}
