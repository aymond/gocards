package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

// NewDeck creates a new deck
func NewDeck(deckName string) (deck Deck, err error) {
	deck = Deck{
		Name: deckName}
	err = nil
	return
}

// Initialize populates the given deck with cards.
// ToDo: Load a deck from a file.
func (d *Deck) Initialize(f string, t string) error {
	log.Printf("Creating Deck from type: %s", t)

	//Read the file
	data, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf("File contents: %s", data)

	switch t {
	case "s":
		// Standard deck
		log.Println("Standard Deck")
		// ToDo Load Standard Deck, where we need
		// to generate a card per suit.
		loadStandardDeck(*d, data)
	case "f":
		// Fixed deck
		log.Println("Fixed Deck")
	default:
		log.Fatalf("Unknown deck type: %q", t)

	}
	return nil
}

// Loads a standard deck. Takes a json of
// values and suits, and generates one value per suit
// and returns the deck.
// Input is an existing deck and json string.
// Output Deck
func loadStandardDeck(d Deck, data []byte) {
	fmt.Printf("%q", data)
	/* for i := 0; i < len(d.Values); i++ {
		for n := 0; n < len(d.Suits); n++ {
			d.add(d.Suits[n], d.Values[i])
		}
	} */
}

// Add a card to the deck by giving the suit and value as a string.
// The card object will be created and appended to the deck.
func (d *Deck) add(s string, v string) {
	card := Card{
		Suit:  s,
		Value: v}
	d.Cards = append(d.Cards, card)
	return
}

// Add a card object to the deck by passing in a card object.
// The card object will be appended to the deck.
func (d *Deck) addCard(c Card) {
	d.Cards = append(d.Cards, c)
	return
}

// Shuffle the deck.
func (d *Deck) Shuffle() {
	// Pick a random position in the deck and swap it.
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < len(d.Cards); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			d.Cards[r], d.Cards[i] = d.Cards[i], d.Cards[r]
		}
	}
	return
}

// Shuffle the deck.
func (gs *BattleKingsGameState) Shuffle(p Deck) {
	// Pick a random position in the deck and swap it.
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < len(p.Cards); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			p.Cards[r], p.Cards[i] = p.Cards[i], p.Cards[r]
		}
	}
	return
}

// Print out to stdout the number of cards specified in n
func (d *Deck) deal(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(d.Cards[i].Value + " of " + d.Cards[i].Suit)
	}
	return
}

// DealToPlayers the number of cards (n) to the number of players (p)
// TODO
func (d *Deck) dealToPlayers(n, p int) {
	fmt.Println(p)
	for i := 0; i < n; i++ {
		fmt.Println(d.Cards[i].Value + " of " + d.Cards[i].Suit)
	}
	return
}

// DealCard returns a card from the deck, removing it from the source deck.
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

// Debug prints the current state of the passed deck
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

// loadDeck reads a file that contains a deck of cards.
func loadDeck(f string) (Deck, error) {

	var deck Deck //:= Deck{}
	file, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatal(err)
		return deck, err
	}

	fmt.Printf("File contents: %s", file)

	err = json.Unmarshal(file, &deck)

	fmt.Println("Name: ", deck.Name)
	for i := 0; i < len(deck.Cards); i++ {
		fmt.Println("Suit: ", deck.Cards[i].Suit)
		fmt.Println("Value: ", deck.Cards[i].Value)
	}

	return deck, err
}
