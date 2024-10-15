package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var cardValue int

	switch card {
	case "ace":
		cardValue = 11
	case "two":
		cardValue = 2
	case "three":
		cardValue = 3
	case "four":
		cardValue = 4
	case "five":
		cardValue = 5
	case "six":
		cardValue = 6
	case "seven":
		cardValue = 7
	case "eight":
		cardValue = 8
	case "nine":
		cardValue = 9
	case "ten":
		cardValue = 10
	case "jack":
		cardValue = 10
	case "queen":
		cardValue = 10
	case "king":
		cardValue = 10
	default:
		cardValue = 0
	}

	return cardValue
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var strategy string

	card1Value := ParseCard(card1)
	card2Value := ParseCard(card2)
	cardsSum := card1Value + card2Value

	dealerCardValue := ParseCard(dealerCard)

	switch {
	case card1 == "ace" && card2 == "ace":
		strategy = "P"
	case cardsSum == 21 &&
		dealerCard != "jack" && dealerCard != "queen" && dealerCard != "king" &&
		dealerCard != "ace" && dealerCard != "ten":
		strategy = "W"
	case cardsSum == 21:
		strategy = "S"
	case cardsSum >= 17 && cardsSum <= 20:
		strategy = "S"
	case cardsSum >= 12 && cardsSum <= 16 && dealerCardValue < 7:
		strategy = "S"
	case cardsSum >= 12 && cardsSum <= 16 && dealerCardValue >= 7:
		strategy = "H"
	case cardsSum <= 11:
		strategy = "H"
	}

	return strategy
}
