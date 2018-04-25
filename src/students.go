package students

import (
	"github.com/crazcalm/flash-cards/cards"
)

//Students testing
type Students struct {
	Students []Student
}

//Shuffle randomizes the order of the students
func (s Students) Shuffle() {
	Shuffle(s)
}

//GetCards testing
func (s Students) GetCards() []flashcards.Card {
	var cards []flashcards.Card
	for _, item := range s.Students {
		cards = append(cards, flashcards.Card{Front: item.GetFront(), Back: item.GetBack(), Hint: item.GetHint(), Flipped: false})
	}
	return cards
}
