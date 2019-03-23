package gocards

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

// A Card has a Suit and Value
type Card struct {
	Suit  string
	Value string
}

// Deck can be loaded with multiple types of decks. e.g. a Standard deck or special deck
// ToDo: Remove Suits and Values/Ranks. Introduce a generator for different deck types.
type Deck struct {
	Name   string
	Cards  []Card
	Suits  []string
	Values []string
}

func main() {
	var deck2 Deck
	deck2.Initialize("Blackjack")
	deck2.Shuffle()
	deck2.dealToPlayers(5, 2)
	c := deck2.dealCard()
	p, _ := json.Marshal(c)
	fmt.Println("Dealt:" + string(p))
	player1Deck, e := NewDeck("Aymon")
	if e != nil {
		fmt.Println(e)
	}
	player1Deck.Debug()
	player1Deck.addCard(c)
	player1Deck.Debug()
	test, test2 := player1Deck.contains(c)
	fmt.Println("Card in Deck: ", test)
	fmt.Println("Position: ", test2)
}

func NewDeck(deckName string) (deck Deck, err error) {
	deck = Deck{
		Name: deckName}
	err = nil
	return
}

// Initialize returns an instance of a Deck. This method should return a game deck.
// Currently hardcoded a classic card deck.
func (d *Deck) Initialize(deckName string) error {
	fmt.Println("Creating Deck.")

	d.Name = deckName
	d.Suits = []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	d.Values = []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}

	for i := 0; i < len(d.Values); i++ {
		for n := 0; n < len(d.Suits); n++ {
			d.add(d.Suits[n], d.Values[i])
		}
	}
	return nil
}

// Add a card to the deck
func (d *Deck) add(s string, v string) {
	card := Card{
		Suit:  s,
		Value: v}
	d.Cards = append(d.Cards, card)
	return
}

// Add a card object to the deck
func (d *Deck) addCard(c Card) {
	d.Cards = append(d.Cards, c)
	return
}

func (d Deck) remove() Deck {
	return d
}

// Shuffle the deck
func (d *Deck) Shuffle() {
	// Pick a random position in the deck and swap it.
	for i := 1; i < len(d.Cards); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			d.Cards[r], d.Cards[i] = d.Cards[i], d.Cards[r]
		}
	}
	return
}

// Deal the number of cards specified in n
func (d *Deck) deal(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(d.Cards[i].Value + " of " + d.Cards[i].Suit)
	}
	return
}

// DealToPlayers the number of cards (n) to the number of players (p)
func (d *Deck) dealToPlayers(n, p int) {
	fmt.Println(p)
	for i := 0; i < n; i++ {
		fmt.Println(d.Cards[i].Value + " of " + d.Cards[i].Suit)
	}
	return
}

// DealCard returns a card from the deck, removing it from the source
// deck.
func (d *Deck) dealCard() Card {
	var c Card
	c, d.Cards = d.Cards[len(d.Cards)-1], d.Cards[:len(d.Cards)-1]
	return c
}

// contains locates a card in the deck and returning the cards position.
func (d *Deck) contains(c Card) (bool, int) {
	for i, card := range d.Cards {
		if card == c {
			return true, i
		}
	}
	return false, 0
}

// Debug prints the current state of the deck
func (d *Deck) Debug() {
	s, _ := json.Marshal(d)
	fmt.Println("DEBUG: " + string(s))
}

// GenerateDeck returns a slice of cards to be used in a deck
func GenerateDeck(d string) []Card {
	var c []Card
	switch d {
	case "hanikalone":
		fmt.Println("d")
		c = []Card{
			{Suit: "Pink", Value: "5"},
			{Suit: "Pink", Value: "5"},
			{Suit: "Pink", Value: "5"},
			{Suit: "Pink", Value: "5"},
			{Suit: "Pink", Value: "5"},
			{Suit: "Green", Value: "4"},
			{Suit: "Green", Value: "4"},
			{Suit: "Green", Value: "4"},
			{Suit: "Green", Value: "4"},
			{Suit: "Orange", Value: "3"},
			{Suit: "Orange", Value: "3"},
			{Suit: "Orange", Value: "3"},
			{Suit: "Blue", Value: "3"},
			{Suit: "Blue", Value: "3"},
			{Suit: "Blue", Value: "3"},
			{Suit: "Red", Value: "2"},
			{Suit: "Red", Value: "2"},
			{Suit: "Yellow", Value: "2"},
			{Suit: "Yellow", Value: "2"},
			{Suit: "Purple", Value: "2"},
			{Suit: "Purple", Value: "2"},
		}
	default:
		fmt.Println("Default")
	}
	return c
}
