package board

import (
	"fmt"
)

// bitboard representation
// a 1 in that bit signifies that a piece of that type exists there
type cBoard struct {
	WhitePawns   uint64
	whiteKnights uint64
	whiteBishops uint64
	whiteRooks   uint64
	whiteQueens  uint64
	whiteKing    uint64

	blackPawns   uint64
	blackKnights uint64
	blackBishops uint64
	blackRooks   uint64
	blackQueens  uint64
	blackKing    uint64
}

// enum for board squares
type BitSquare int

const (
	A8 BitSquare = iota
	B8
	C8
	D8
	E8
	F8
	G8
	H8
	A7
	B7
	C7
	D7
	E7
	F7
	G7
	H7
	A6
	B6
	C6
	D6
	E6
	F6
	G6
	H6
	A5
	B5
	C5
	D5
	E5
	F5
	G5
	H5
	A4
	B4
	C4
	D4
	E4
	F4
	G4
	H4
	A3
	B3
	C3
	D3
	E3
	F3
	G3
	H3
	A2
	B2
	C2
	D2
	E2
	F2
	G2
	H2
	A1
	B1
	C1
	D1
	E1
	F1
	G1
	H1
)

// enum for sides of game
type Side int 

const (
	White Side = 0 
	Black 
)
var ChessBoard cBoard

// debugging tool
func PrintBitBoard(bitBoard uint64) {
	fmt.Println()
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {

			//sqaure index of the board
			square := rank*8 + file // 0 - 64

			// empty is 0, filled is 1
			bit := 0
			if getBit(bitBoard, square) == 0 {
				bit = 0
			} else {
				bit = 1
			}
			//also print the current rank
			if file == 0 {
				fmt.Printf("%d ", 8-rank)
			}
			// print bit state
			fmt.Printf(" %d ", bit)
		}
		//print newline every rank
		fmt.Println()
	}

	// print the files of the bit board
	fmt.Printf("\n   a  b  c  d  e  f  g  h \n\n")
	fmt.Printf(" Bitboard in decimal: %d\n", bitBoard)
}

// internal macros for getting/popping/setting bits

// returns the bit at the sqaure position given the bitboard and sqaure position
func getBit(bitBoard uint64, square int) (bit uint64) {

	// 1 << square will left shift one by 0 - 64
	// this will be a bitmask at the 0-64th bit
	// &'ing the bitboard at the bitmask will see if the bit is activated there
	return bitBoard & (1 << square)
}

// sets the bit at the human readable square to 1
func SetBit(bitBoard *uint64, square BitSquare) {
	// 1 << square will left shift one by 0 - 63rd bit
	// this will create a bitmask at the 0-63rd bit
	// or'ing the bitboard at the bitmask will set the bit at the square to 1
	*bitBoard |= (1 << square)
}

//removes a bit at at the human readable square/ sets it to 0
func PopBit(bitBoard *uint64, square BitSquare) {
	if getBit(*bitBoard, int(square)) != 0 { // nonzero value means bit was set there
		*bitBoard ^= (1 << square) //xor the bit
	} 
	// already 0, do nothing
}

// // generate pawn attacks
// func mask_pawn_attacks(square BitSquare, side Side) {
// 	// results of attacks bitboard 
// 	var attacks uint64 

// 	// piece bitboard 
// 	var bitBoard uint64 

// 	SetBit(&bitBoard, square)



// }