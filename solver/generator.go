package Solver

import (
	"fmt"
	"math/rand"
	"time"
)

// Get empty puzzle
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
	// return d

}

// get new puzzle
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
	//

	//found on go by example
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
		// b[1][j] = gridTwo[j-3]
		// b[2][j] = gridThree[j-3]
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
	// s1 := rand.NewSource(time.Now().UnixNano())
	// r1 := rand.New(s1)

	// var test = r1.Intn(9 - 1 ) + 1

	//generate 25 randoms for column index

	// b := []int{random(0, 9), random(0, 9)}
	// fmt.Print(b)
	var N = 50

	var s [50][2]int
	var isFound = false
	// copy(board[i][:2], s))
	for i := 0; i < N; i++ {
		isFound = false
		// rand.Seed(time.Now().Unix())
		newCell := [2]int{Random(0, 9), Random(0, 9)}
		for _, h := range s {
			// fmt.Print(newCell)
			// fmt.Print(h)

			if h == newCell {
				i = i - 1
				isFound = true
				// fmt.Print("STOP")
				break

			}
			// for _, cell := range h {

			// 	fmt.Print(h, " here")
			// 	fmt.Print(cell)
			// }

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

//go cookbook
func Random(min, max int) int {

	return rand.Intn(max-min) + min
}
