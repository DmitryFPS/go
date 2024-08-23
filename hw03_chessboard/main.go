package main

import "fmt"

func printChessboard(boardWidth, boardHeight *int) {
	for i := 0; i < *boardWidth; i++ {
		for j := 0; j < *boardHeight; j++ {
			if (i+j)%2 == 0 {
				fmt.Printf("\u0020")
			} else {
				fmt.Printf("\u0023")
			}
		}
		fmt.Println()
	}
}

func readerDataFromConsole() (*int, *int) {
	boardWidth := 8
	boardHeight := 8

	fmt.Printf("Введите длину шахмотной доски по горизонтали: ")
	_, boardWidthError := fmt.Scanln(&boardWidth)
	fmt.Printf("Введите длину шахмотной доски по вертикали: ")
	_, boardHeightError := fmt.Scanln(&boardHeight)

	if boardWidthError != nil {
		fmt.Println("BoardWidthError", &boardWidth)
	}

	if boardHeightError != nil {
		fmt.Println("BoardHeightError", &boardHeight)
	}

	return &boardWidth, &boardHeight
}

func main() {
	printChessboard(readerDataFromConsole())
}
