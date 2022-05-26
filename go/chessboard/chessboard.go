package chessboard

import (
	"log"
)

/*
 * Optimize code, referring to https://exercism.org/tracks/go/exercises/chessboard/solutions/chris-argirop
*/

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	countRank := 0
	val, ok := cb[rank]; 
	
	if ok {		
		// log.Printf("found RANK = %s | value : %v \n", rank, val)
		for _, status := range val {
			if status {
				countRank += 1
			}
		}				
	}
	
	log.Printf("rank = %s | isFound = %v | countRank = %d \n \n", rank, ok, countRank)
	return countRank
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	// panic("Please implement CountInFile()")	
	counter := 0

	file--
	if file >= 0 && file < 8 {		
		for _, value := range cb {
			if value[file] {								
				counter++
			}
		}
	}
	
	log.Printf("file = %v | countFile = %d \n \n", file,  counter)		
	return counter
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	countAll := 0

	for _, v := range cb {
		countAll += len(v)		
	}
	log.Printf("cb = %v | countAll = %d \n \n", cb,  countAll)
	return countAll
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	// panic("Please implement CountOccupied()")

	countOccupied := 0
	for _, v := range cb {
		for _, i := range v {
			if i {
				countOccupied += 1
			}
		}
	}

	log.Printf("cb = %v | countOccupied = %d \n \n", cb,  countOccupied)	
	return countOccupied
}
