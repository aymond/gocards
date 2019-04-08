package gocards

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	var err error
	d, err := NewDeck("Test")
	if err != nil {
		t.Error(err)
	}
	if d.Name != "Test" {
		t.Error("Wrong Name. Was expecting Test but got ", d.Name)
	}
}

/* func TestInitialize(t *testing.T) {
	var err error
	var d Deck
	d.Initialize("Blackjack")
	if err != nil {
		t.Error(err)
	}
	if d.Name != "Blackjack" {
		t.Error("Wrong Name. Was expecting Blackjack but got ", d.Name)
	}

} */

/* func TestDealToPlayerApp(t *testing.T) {
	var deck2 Deck
	deck2.Initialize("Blackjack")
	deck2.Shuffle()
	deck2.dealToPlayers(5, 2)
	c := deck2.dealCard()
	p, _ := json.Marshal(c)
	fmt.Println("Dealt:" + string(p))

	player1Deck, e := NewDeck("Aymon")
	if e != nil {
		t.Error("Player1 Deck Failed:  ", e)
	}
	player1Deck.addCard(c)
	test, test2 := player1Deck.contains(c)
	fmt.Println("Card in Deck: ", test)
	fmt.Println("Position: ", test2)

} */

func TestGame(t *testing.T) {
	testdeck := GenerateDeck("hanikalone")
	fmt.Println(testdeck)
	// TODO
}

/* func TestRemove(t *testing.T) {
	var deck Deck

	 := []Card{
		{Suit: "Pink", Value: "5"},
		{Suit: "Pink", Value: "5"},
		{Suit: "Pink", Value: "5"},
		{Suit: "Pink", Value: "5"},
		{Suit: "Pink", Value: "5"},
		{Suit: "Green", Value: "4"},
	}

} */
