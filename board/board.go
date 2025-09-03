package board

import (
	"fmt"
)

// bitboard representation
// a 1 in that bit signifies that a piece of that type exists there
type cBoard struct {
	WhitePawns uint64
	whiteKnights uint64
	whiteBishops uint64
	whiteRooks uint64
	whiteQueens uint64
	whiteKing uint64

	blackPawns uint64
	blackKnights uint64
	blackBishops uint64
	blackRooks uint64
	blackQueens uint64
	blackKing uint64
}

var ChessBoard cBoard
func PrintBitBoard(bitBoard uint64) {
	fmt.Println()
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			
			//sqaure index of the board
			square := rank * 8 + file // 0 - 64

			// empty is 0, filled is 1 
			bit := 0
			if getBit(bitBoard, square) == 0 {
				bit = 0
			} else {
				bit = 1
			}
			//also print the current rank 
			if file == 0 {
				fmt.Printf("%d ", 8 - rank)
			}
			// print bit state 
			fmt.Printf(" %d ", bit)
		}
		//print newline every rank
		fmt.Println()
	}

	// print the files of the bit board 
	fmt.Printf("\n   a  b  c  d  e  f  g  h \n\n")
}

// internal macros for getting/popping/setting bits

// returns the bit at the sqaure position given the bitboard and sqaure position
func getBit(bitBoard uint64, square int) (bit uint64) {

	// 1 << square will left shift one by 0 - 64
	// this will be a bitmask at the 0-64th bit
	// &'ing the bitboard at the bitmask will see if the bit is activated there
	return bitBoard & (1 << square) 
}