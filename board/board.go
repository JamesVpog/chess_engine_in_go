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

// enum for board squares in LERF (Little-Endian Rank-File) mapping opposite order
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
	White Side = iota
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
			if getBit(bitBoard, square) != 0 {
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
	return bitBoard & (uint64(1) << uint(square))
}

// sets the bit at the human readable square to 1
func SetBit(bitBoard *uint64, square BitSquare) {
	// 1 << square will left shift one by 0 - 63rd bit depending on the square
	// this will create a bitmask at the 0-63rd bit
	// or'ing the bitboard at the bitmask will set the bit at the square to 1
	*bitBoard |= (uint64(1) << uint(square))
}

//removes a bit at at the human readable square/ sets it to 0
func PopBit(bitBoard *uint64, square BitSquare) {
	if getBit(*bitBoard, int(square)) != 0 { // nonzero value means bit was set there
		*bitBoard ^=  (uint64(1) << uint(square)) //xor the bit
	} 
	// already 0, do nothing
}

// 2d array of uint64
var pawnAttacks [2][64]uint64

/*

	Not A File 

8  0  1  1  1  1  1  1  1 
7  0  1  1  1  1  1  1  1 
6  0  1  1  1  1  1  1  1 
5  0  1  1  1  1  1  1  1 
4  0  1  1  1  1  1  1  1 
3  0  1  1  1  1  1  1  1 
2  0  1  1  1  1  1  1  1 
1  0  1  1  1  1  1  1  1 

   a  b  c  d  e  f  g  h 
*/
const NotAFile uint64 = 18374403900871474942 

/*

	Not H File 

8  1  1  1  1  1  1  1  0 
7  1  1  1  1  1  1  1  0 
6  1  1  1  1  1  1  1  0 
5  1  1  1  1  1  1  1  0 
4  1  1  1  1  1  1  1  0 
3  1  1  1  1  1  1  1  0 
2  1  1  1  1  1  1  1  0 
1  1  1  1  1  1  1  1  0 

   a  b  c  d  e  f  g  h 
*/
	
const NotHFile uint64 = 9187201950435737471

/*

	Not HG File

8  1  1  1  1  1  1  0  0 
7  1  1  1  1  1  1  0  0 
6  1  1  1  1  1  1  0  0 
5  1  1  1  1  1  1  0  0 
4  1  1  1  1  1  1  0  0 
3  1  1  1  1  1  1  0  0 
2  1  1  1  1  1  1  0  0 
1  1  1  1  1  1  1  0  0 

   a  b  c  d  e  f  g  h 
*/
const NotHGFile uint64 = 4557430888798830399

/*
	Not AB File
8  0  0  1  1  1  1  1  1 
7  0  0  1  1  1  1  1  1 
6  0  0  1  1  1  1  1  1 
5  0  0  1  1  1  1  1  1 
4  0  0  1  1  1  1  1  1 
3  0  0  1  1  1  1  1  1 
2  0  0  1  1  1  1  1  1 
1  0  0  1  1  1  1  1  1 

   a  b  c  d  e  f  g  h 
*/

const NotABFile uint64 = 18229723555195321596


func Mask_pawn_attacks(square BitSquare, side Side) (attacks uint64) {
	// results of attacks bitboard 
	// var attacks uint64 

	// piece bitboard 
	var bitBoard uint64 

	// set piece on the bitboard 
	SetBit(&bitBoard, square)

	// generate the attacks from the piece
	switch side {
	case White:
		// white pawn attack on the diagonal top right
		// if we are on h file, we want to make sure we don't wrap around to A file 
		if ( (bitBoard >> 7) & NotAFile) != 0 { // top right attack &'d notAfile is 0, then we are trying to wrap around to A File
			attacks |= (bitBoard >> 7)
		}	

		// white pawn attack on the diagonal top left
		// if we are on a file, we want to make sure we don't wrap around to H file 
		if ( (bitBoard >> 9) & NotHFile) != 0 { // top left attack &'d notHFile is 0, then we are trying to wrap around to H File
			attacks |= (bitBoard >> 9)
		}
		
	case Black:
		// black pawn attack on the diagonal bottom right
		// if we are on a file, we want to make sure we don't wrap around to H file 
		if ( (bitBoard << 7) & NotHFile) != 0 { // TODO
			attacks |= (bitBoard << 7)
		}	

		// black pawn attack on the diagona bottom left
		// if we are on h file, we want to make sure we don't wrap around to a file 
		if ( (bitBoard << 9) & NotAFile) != 0 { // TODO
			attacks |= (bitBoard << 9)
		}
	}

	return attacks


}

