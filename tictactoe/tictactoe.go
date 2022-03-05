package main

import "fmt"

func main() {

	var board = createBoard()
	fmt.Print(board)
}

func createBoard() [][]string {
	newBoard := [][]string{
		[]string{" ", " ", " "},
		[]string{" ", " ", " "},
		[]string{" ", " ", " "},
	}

	return newBoard
}

func printBoard(board *string) {

}
