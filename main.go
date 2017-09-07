package main

import (
	"Sudoku-GO/solver"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// data structure for the main Sudoku Grid
// Done Boolean value is used for building the HTML table that
// displays the grid
type Sudoku struct {
	Done     bool
	Value    [9]int
	Solved   [9]bool
	Solution [9]int
	ID       [9]string
}

// Data structure to display the value for each cell i.e A1 to I9
type CellID struct {
	Row int
	Col int
}

// function main starts the application server. An empty puzzle is used to
// build the HTML game template - so that fromatting is correct
func main() {

	tmpl := template.Must(template.ParseFiles("template/GUI.html"))
	var b = Solver.GetEmptyPuzzle()

	sudoku := arrayToData(b)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, struct{ Sudokus [9]Sudoku }{sudoku})
	})

	//file server for handling static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/game/", gameHandler)
	http.HandleFunc("/killer/", killerHandler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

// Simple funtion to serve the favicon
func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/static/favicons/favicon.ico")
}

// the game handler will recieve information from the server with the requested difficulty level
// as a string
// The handler will then get a newly generated puzzle from the solver package and then
// solve it.
// The template will then be rendered appropriately
func gameHandler(w http.ResponseWriter, r *http.Request) {

	var optionSelected string
	r.ParseForm()
	fmt.Println(r.Form) // print information on server side.
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
		optionSelected = strings.Join(v, " ")
	}

	fmt.Println(optionSelected)
	var b = Solver.GetPuzzle(optionSelected)

	var solvedArray = Solver.NewSolver(b, true)
	sudoku := getUnsolved(b, solvedArray)

	t := template.Must(template.ParseFiles("template/easy_game.html"))
	t, _ = t.ParseFiles("template/easy_game.html")

	t.Execute(w, struct{ Sudokus [9]Sudoku }{sudoku})
}

// Function handler for the Killer template - needs updated
func killerHandler(w http.ResponseWriter, r *http.Request) {

	var b = Solver.GetPuzzle("")

	fmt.Println(b)
	var solvedArray = Solver.NewSolver(b, true)
	fmt.Println(solvedArray)

	// sudoku := arrayToData(solvedArray)
	sudoku := getUnsolved(b, solvedArray)

	t := template.Must(template.ParseFiles("template/killer.html"))
	t, _ = t.ParseFiles("template/killer.html")
	t.Execute(w, struct{ Sudokus [9]Sudoku }{sudoku})
}

func arrayToData(solved [9][9]int) [9]Sudoku {

	solvedSudoku := [9]Sudoku{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			solvedSudoku[i].Value[j] = solved[i][j]
			solvedSudoku[i].Solved[j] = true

			if i == 3 || i == 6 {
				solvedSudoku[i].Done = true
			} else {
				solvedSudoku[i].Done = false
			}

		}
	}

	return solvedSudoku
}

// This function is to set the variables in the Sudoku data Structure so that
// a Sudoku grid has both the value and solutions
// This can then be used appropriately for the Client side logic to control the game
// This function will also fill the CellId data structure with value coresponding to the Cell
// The method returns a Completed Sudoku array
func getUnsolved(unsolved [9][9]int, solved [9][9]int) [9]Sudoku {
	solvedSudoku := [9]Sudoku{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			if unsolved[i][j] != 0 {
				solvedSudoku[i].Value[j] = solved[i][j]

			} else {
				solvedSudoku[i].Solution[j] = solved[i][j]
			}
			var f string

			// switch statement to set correct Cell ID
			switch i {
			case 0:
				f = "A"
			case 1:
				f = "B"
			case 2:
				f = "C"
			case 3:
				f = "D"
			case 4:
				f = "E"
			case 5:
				f = "F"
			case 6:
				f = "G"
			case 7:
				f = "H"
			case 8:
				f = "I"

			}

			s := strconv.Itoa(j + 1)
			solvedSudoku[i].ID[j] = f + s

			// The done variable is used to identify in the template the boundaries of the minigrids for the
			// thicker border
			if i == 3 || i == 6 {
				solvedSudoku[i].Done = true
			} else {
				solvedSudoku[i].Done = false
			}

		}
	}
	return solvedSudoku

}
