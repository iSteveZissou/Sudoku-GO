package Solver

import (
	"fmt"
	"math/rand"
	"time"
)

// Get empty puzzle
// parameter is a string - the user's choice of difficulty
// returns a new generated standard sudoku Puzzle
func GetPuzzle(i string) [9][9]int {

	switch i {
	case "Easy":
		fmt.Println("The user has selected ", i)
	case "Medium":
		fmt.Println("The user has selected ", i)
	case "Hard":
		fmt.Println("The user has selected ", i)

	}

	// loop to generate puzzle until one with a unique solution is found
	for {
		var b = GetNewStart()
		fmt.Print(b)
		var c = NewSolver(b, false)
		var d = removeCellsEasy(c)
		fmt.Print("here is d", d)

		NewSolver(d, true)
		if SolutionCount == 1 {
			return d
		}
	}

}

// get new puzzle
// returns an empty 2d int array
func GetEmptyPuzzle() [9][9]int {

	var b [9][9]int
	b[0] = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	b[1] = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	b[2] = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	b[3] = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	b[4] = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	b[5] = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	b[6] = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	b[7] = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	b[8] = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}

	return b
}

// This function gets the start of a puzzle by first randomly filling the first 3x3 mini-grid
// The first row will then be filled with the remaining numbers at random
// A 2d int array will be returned
func GetNewStart() [9][9]int {

	var b [9][9]int
	b[0] = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	b[1] = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	b[2] = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	b[3] = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	b[4] = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	b[5] = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	b[6] = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	b[7] = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	b[8] = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}

	var count = 0

	//found on go by example https://gobyexample.com/
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Print(r1.Intn(100), ",")

	list := r1.Perm(9)
	for i, _ := range list {
		list[i]++
	}
	var gridOne [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {

			gridOne[i][j] = list[count]
			count++
		}

	}
	var testGrid = getShuffledRows(gridOne[0])

	// fill in the first 3x3 grid
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {

			b[i][j] = gridOne[i][j]
		}
	}

	for j := 3; j < 9; j++ {

		b[0][j] = testGrid[j-3]

	}

	return b

}

func getShuffledRows(grid [3]int) [6]int {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	list2 := r1.Perm(9)
	for i, _ := range list2 {
		list2[i]++
	}

	var rowList [6]int
	var skip = false
	var counter = 0

	for i := 0; i < 9; i++ {
		skip = false
		for j := 0; j < 3; j++ {
			if list2[i] == grid[j] {

				skip = true
			}

		}
		if !skip {
			rowList[counter] = list2[i]
			counter++
		}
	}
	// fmt.Print(rowList)

	return rowList

}

func shuffleRow() {

}

func removeCellsEasy(board [9][9]int) [9][9]int {

	//generate 25 random numbers between one and nine for row index
	var N = 50

	var s [50][2]int
	var isFound = false
	// copy(board[i][:2], s))
	for i := 0; i < N; i++ {
		isFound = false
		// rand.Seed(time.Now().Unix())
		newCell := [2]int{Random(0, 9), Random(0, 9)}
		for _, h := range s {

			if h == newCell {
				i = i - 1
				isFound = true
				// fmt.Print("STOP")
				break

			}

		}
		// s = append(s, b)
		if !isFound {
			s[i] = newCell
			board[newCell[0]][newCell[1]] = 0
		}

	}
	// fmt.Print(s)

	return board

}

// Random function found on Go Cookbook
// Will return a ranom int between an minimum and maximum range
func Random(min, max int) int {

	return rand.Intn(max-min) + min
}
