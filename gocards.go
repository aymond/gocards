package main

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
type Deck struct {
	Name   string
	Cards  []Card
	Suits  []string
	Values []string
}

func main() {

	var deck2 Deck
	deck2.initialize("Blackjack")
	deck2.shuffle()
	deck2.dealToPlayers(5, 2)
	c := deck2.dealCard()
	p, _ := json.Marshal(c)
	fmt.Println("Dealt:" + string(p))
}

func new(deckName string) (deck Deck) {
	var cards []Card

	deck = Deck{
		Name:   deckName,
		Suits:  []string{"Hearts", "Diamonds", "Clubs", "Spades"},
		Values: []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}}

	// Deck is defined, loop over the deck and associate each value to the suit.
	for i := 0; i < len(deck.Values); i++ {
		for n := 0; n < len(deck.Suits); n++ {
			card := Card{
				Suit:  deck.Suits[n],
				Value: deck.Values[i]}
			cards = append(cards, card)

		}
	}
	deck.Cards = cards
	return
}

// Initialize returns an instance of a Deck. This method should return a game deck.
// Currently hardcoded a classic card deck.
func (d *Deck) initialize(deckName string) {
	fmt.Println("Creating Deck.")

	d.Name = deckName
	d.Suits = []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	d.Values = []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}

	for i := 0; i < len(d.Values); i++ {
		for n := 0; n < len(d.Suits); n++ {
			d.add(d.Suits[n], d.Values[i])
		}
	}
	return
}

// Add a card to the deck
func (d *Deck) add(s string, v string) {
	card := Card{
		Suit:  s,
		Value: v}
	d.Cards = append(d.Cards, card)
	return
}

func (d Deck) remove() Deck {
	return d
}

// Shuffle the deck
func (d *Deck) shuffle() {
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
// deck, and returns the new size of the deck.
func (d *Deck) dealCard() Card {
	var c Card
	c, d.Cards = d.Cards[len(d.Cards)-1], d.Cards[:len(d.Cards)-1]
	return c
}
