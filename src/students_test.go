package students

import (
	"testing"
)

func TestGetCards(t *testing.T) {
	//Students
	ss := getTestData()

	//Gets the cards
	cards := ss.GetCards()

	//Ensures that the slice of cards exists
	if len(cards) <= 0 {
		t.Fail()
	}
}
