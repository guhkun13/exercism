package blackjack

var cardValueMap = map[string]int{
	"ace":   11,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
}

const (
	stand string = "S"
	hit string = "H"
	split string = "P"
	win string = "W"
)


// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	// by default return 0 if the mapped value or key is not exist withing the map
	return cardValueMap[card]	
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {			
	var SumCard int = ParseCard(card1) + ParseCard(card2)
	
	if card1 == "ace" && card2 == "ace" {
		return split
	} else if ( SumCard == 21) && (ParseCard(dealerCard) < 10 )  {
		return win
	} else if SumCard >= 17 && SumCard <= 20 {
		return stand
	} else if SumCard >= 12 && SumCard <= 16 {
		if ParseCard(dealerCard) >= 7 {
			return hit
		}
		return stand
	} else if SumCard <= 11 {
		return hit
	}
	return stand
}
			