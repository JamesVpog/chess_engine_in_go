package main

import (
	"github.com/JamesVPog/chess_engine_in_go/board"
)



func main() {
	board.SetBit(&board.ChessBoard.WhitePawns, board.H8)
	board.SetBit(&board.ChessBoard.WhitePawns, board.E4)
	board.SetBit(&board.ChessBoard.WhitePawns, board.D3)

	board.PrintBitBoard(board.ChessBoard.WhitePawns)

	board.PopBit(&board.ChessBoard.WhitePawns, board.E4)

	board.PrintBitBoard(board.ChessBoard.WhitePawns)

	board.PopBit(&board.ChessBoard.WhitePawns, board.E4)
	
	board.PrintBitBoard(board.ChessBoard.WhitePawns)

	board.PopBit(&board.ChessBoard.WhitePawns, board.E4)
	
	board.PrintBitBoard(board.ChessBoard.WhitePawns)
}
