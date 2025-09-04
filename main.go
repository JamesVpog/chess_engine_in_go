package main

import (

	"github.com/JamesVPog/chess_engine_in_go/board"
)



func main() {
	// set the e4, d4, and h7 bit
	// board.PrintBitBoard(board.Mask_pawn_attacks(board.E4, board.White))

	// var hgFile uint64 
	// hgFile = 0
	// // generate picture for file in comments
	// for rank := 0; rank < 8; rank++ {
	// 	for file := 0; file < 8; file++ {
	// 		square := rank * 8 + file
			 
	// 		if file > 1 { 
	// 			board.SetBit(&hgFile, board.BitSquare(square))
	// 		}
	// 	}
	// }

	
	// seeing all possible attacks of every pawn position 
	board.InitLeaperAttacks()
	for i :=0; i < 64; i++ {
		board.PrintBitBoard(board.PawnAttacks[board.Black][i])
	}

}
