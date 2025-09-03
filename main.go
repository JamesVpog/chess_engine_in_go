package main

import (
	"fmt"
	"github.com/JamesVPog/chess_engine_in_go/board"
)
func main() {
	fmt.Println("Hello chess engine world!")
	// setting the white pawns on the 2nd rank test
	board.ChessBoard.WhitePawns = 8192 + 4096 + 2048 + 1024 + 512 + 256
	board.PrintBitBoard(board.ChessBoard.WhitePawns)
}
