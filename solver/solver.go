package Solver

import (
	"fmt"
	"log"
	"time"
)

type Cell struct {
	row      int
	col      int
	value    int
	solution int
}

type Grid struct {
	cells         [9][9]Cell
	Solved        bool
	SolutionCount int
	firstSolve    bool
	possibilities []int
}

var newGameGrid [9][9]int
var SolutionCount int

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func (grid *Grid) getCandidates(c Cell) []int {
	possibilities := make([]int, 0)
	possibilities = possibilities[:0]
	// fmt.Print(possibilities)

	for i := 0; i < 9; i++ {
		if grid.cells[c.row][i].value != 0 {
			// fmt.Println("found in row")
			possibilities = append(possibilities, grid.cells[c.row][i].value)

		} else {
			// possibilities = append(possibilities, i)
		}
	}
	for i := 0; i < 9; i++ {
		if grid.cells[i][c.col].value != 0 {
			possibilities = append(possibilities, grid.cells[i][c.col].value)
		} else {
		}
	}
	//check grid // redo this!!!!
	// var v = c.value
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

	// fmt.Println(possd)
	return possd

}

func contains(s []int, v int) bool {
	for _, a := range s {
		if a == v {
			return true
		}
	}
	return false
}

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
			fmt.Println("FUCK")
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

	for _, a := range candidates {

		grid.cells[c.row][c.col].value = a
		grid.solve(grid.getNextCell(c))

		// if solved {
		// 	return true
		// }
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

func (grid *Grid) getNextCell(c Cell) Cell {

	var test Cell
	test.value = 10

	var row = c.row
	var col = c.col

	// fmt.Println("This goes in tooo", row, col)
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

//"newsolver"
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
