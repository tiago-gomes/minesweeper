package game

import (
	"fmt"
	"math/rand"
)

// declares the board
var board = [][]int{}
var name string

func CreateBoard() {
	// generate the board
	board = [][]int{
		{rand.Intn(2), rand.Intn(2), rand.Intn(2), rand.Intn(2), rand.Intn(2)},
		{rand.Intn(2), rand.Intn(2), rand.Intn(2), rand.Intn(2), rand.Intn(2)},
		{rand.Intn(2), rand.Intn(2), rand.Intn(2), rand.Intn(2), rand.Intn(2)},
		{rand.Intn(2), rand.Intn(2), rand.Intn(2), rand.Intn(2), rand.Intn(2)},
		{rand.Intn(2), rand.Intn(2), rand.Intn(2), rand.Intn(2), rand.Intn(2)},
	}
}

func PrintBoard() {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			// Print each element with a space, without newline
			fmt.Print(board[i][j], " ")
		}
		// After printing each row, add a newline
		fmt.Println()
	}
}

func GetPlayerName() string {
	fmt.Print("Enter player name: ")
	fmt.Scanln(&name)
	return name
}

func getPlayerMove() (int, int) {

	var row int
	var column int

	for {
		fmt.Print("Enter row: ")
		fmt.Scanln(&row)

		if row < 0 || row > 5 {
			fmt.Println("Invalid row. Please enter a number between 0 and 5.")
			continue
		}

		fmt.Print("Enter column: ")
		fmt.Scanln(&column)

		if column < 0 || column > 5 {
			fmt.Println("Invalid column. Please enter a number between 0 and 5.")
			continue
		}

		if board[row][column] == 3 {
			fmt.Println("Invalid move. This position is already occupied.")
			continue
		}

		return row, column
	}
}

func SavePlay(row int, column int) {
	fmt.Printf("%s the coast is clear\n", name)
	// add play to board
	board[row][column] = 3
}

func IsWinner(row int, column int) bool {
	var isPlayerWinner bool = true
	// check if player has won
	for i := 0; i < len(board); i++ {
		// Iterate through the columns
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 0 {
				isPlayerWinner = false
			}
		}
	}
	return isPlayerWinner
}

func IsMine(row int, colum int) bool {
	return board[row][colum] == 1
}

func PlayGame() {
	for {
		// ask for move
		row, colum := getPlayerMove()

		// check if move is mine and we lost the game
		if IsMine(row, colum) {
			fmt.Printf("%s hit a mine!", name)
			return
		}

		// add play to board
		SavePlay(row, colum)

		// check if player won
		if IsWinner(row, colum) {
			fmt.Printf("%s is a winner!", name)
			return
		}
	}
}
