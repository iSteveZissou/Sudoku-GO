package main

import (
	"Sudoku-GO/solver"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type Todo struct {
	Task string
	Done bool
}

type Sudoku struct {
	Done     bool
	Value    [9]int
	Solved   [9]bool
	Solution [9]int
	ID       [9]string
}
type CellID struct {
	Row int
	Col int
}

var n int

var f Solver.NewDataStruct

func main() {
	tmpl := template.Must(template.ParseFiles("template/GUI.html"))

	var b = Solver.GetEmptyPuzzle()

	sudoku := arrayToData(b)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, struct{ Sudokus [9]Sudoku }{sudoku})
	})

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/game/", gameHandler)
	http.HandleFunc("/killer/", killerHandler)
	http.ListenAndServe(":8080", nil)
}
func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/static/favicons/favicon.ico")
}

func gameHandler(w http.ResponseWriter, r *http.Request) {

	var optionSelected string
	r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
	// attention: If you do not call ParseForm method, the following data can not be obtained form
	fmt.Println(r.Form) // print information on server side.
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
		optionSelected = strings.Join(v, " ")
	}

	fmt.Println(optionSelected)
	var b = Solver.GetPuzzle(optionSelected)

	// sudoku := arrayToData(b)

	var solvedArray = Solver.NewSolver(b)
	sudoku := getUnsolved(b, solvedArray)
	// fmt.Print("This is reloading the page")

	// var solved = newSolver(b)

	t := template.Must(template.ParseFiles("template/easy_game.html"))
	t, _ = t.ParseFiles("template/easy_game.html")

	t.Execute(w, struct{ Sudokus [9]Sudoku }{sudoku})
}

func killerHandler(w http.ResponseWriter, r *http.Request) {

	var b = Solver.GetPuzzle("")

	fmt.Println(b)
	var solvedArray = Solver.NewSolver(b)
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

			if i == 3 || i == 6 {
				solvedSudoku[i].Done = true
			} else {
				solvedSudoku[i].Done = false
			}

		}
	}
	return solvedSudoku

}
