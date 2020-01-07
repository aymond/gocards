package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

// A Card has a Suit and Value.
type Card struct {
	Suit  string `json:"suit"`
	Value string `json:"value"`
}

// The BattleKingsGameState contains all game-state relevant information.
type BattleKingsGameState struct {
	gameTick           int     // Count of rounds played
	playerOneHand      Deck    // State of Player One's hand
	playerTwoHand      Deck    // State of Player Two's hand
	discardPile        Deck    // Cards that have been discarded because of Tactics cards
	variantDiscardPile Deck    // Variant 1: Tactics Cards discarded before game start.
	tacticsPile        Deck    // Tactics cards that have not yet been drawn.
	gamePile           Deck    // Game cards that have not yet been drawn.
	flags              [9]Flag // 9 Flags in the game
	won                int     // 0 = unresolved, 1 = Player1, 2 = Player2
	playerOnTurn       int     // 1 = Player 1, 2 = Player 2
	uuid               string
	ID                 int
	gameName           string
	createdAt          time.Time
}

// Deck can be loaded with multiple types of decks. e.g. a Standard deck or special deck
// ToDo: Remove Suits and Values/Ranks. Introduce a generator for different deck types.
type Deck struct {
	Name  string `json:"name"`
	Cards []Card `json:"cards"`
}

// Flag is what the players are battling for.
type Flag struct {
	playerOnePile Deck
	playerTwoPile Deck
	won           int // 0 = unresolved, 1 = Player1, 2 = Player2
}

// GameList used for listing games.
type GameList struct {
	uuid      string
	gameName  string
	createdAt time.Time
}

// Db stores the game state.
var Db *sql.DB

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

// moveCard moves a card (c) from one deck (d1) to another deck (d2).
func (gs *BattleKingsGameState) moveCard(d1 *Deck, d2 *Deck) {
	var c Card
	c, d1.Cards = d1.Cards[len(d1.Cards)-1], d1.Cards[:len(d1.Cards)-1]
	d2.Cards = append(d2.Cards, c)
	//return c
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

// Debug prints the current state of the passed deck
func (gs *BattleKingsGameState) Debug() {
	//s, _ := json.Marshal(gs.gamePile)
	fmt.Printf("DEBUG: %+v", gs)
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

// LoadDeck reads the deck
func (gs *BattleKingsGameState) LoadDeck(f string, d *Deck) error {
	fileData, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = json.Unmarshal([]byte(fileData), d)
	return nil
}

// SaveState writes the game state to the database.
func (gs *BattleKingsGameState) newGame() error {
	/* var err error
	Db, err = sql.Open("sqlite3", "dbname=game.db")
	if err != nil {
		log.Fatal(err)
	} */

	statement := "insert into games (uuid, game_name, created_at) values ($1, $2, $3)" // returning id, uuid, game_name, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmt.QueryRow(createUUID(), "BattleKings", time.Now()).Scan(&gs.ID, &gs.uuid, &gs.gameName, &gs.createdAt)
	return err
}

// NewGame creates a new game with a default state.
func (gs *BattleKingsGameState) NewGame(v1 bool) error {
	var err error
	// Load the Game Deck
	_ = gs.LoadDeck("configs/battlekings-deck.json", &gs.gamePile)
	_ = gs.LoadDeck("configs/battlekings-tactics.json", &gs.tacticsPile)

	// Shuffle the decks
	gs.Shuffle(gs.gamePile)
	gs.Shuffle(gs.tacticsPile)

	if v1 == true {
		// Move two Tactics cards to the Discard Pile
		gs.moveCard(&gs.tacticsPile, &gs.variantDiscardPile)
		gs.moveCard(&gs.tacticsPile, &gs.variantDiscardPile)
	}

	// Deal 7 cards to each player.
	for i := 1; i <= 7; i++ {
		gs.moveCard(&gs.gamePile, &gs.playerOneHand)
		gs.moveCard(&gs.gamePile, &gs.playerTwoHand)
	}

	// Choose random start player 1 or 2
	rand.Seed(time.Now().UnixNano())
	gs.playerOnTurn = rand.Intn(2) + 1

	/* Db, err = sql.Open("sqlite3", "dbname=game.db")
	if err != nil {
		log.Fatal(err)
	} */

	statement := "insert into games (uuid, game_name, created_at) values ($1, $2, $3)" // returning id, uuid, game_name, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmt.QueryRow(createUUID(), "BattleKings", time.Now()).Scan(&gs.ID, &gs.uuid, &gs.gameName, &gs.createdAt)
	return err
}

// ListGames returns all games.
func ListGames() (games []GameList, err error) {

	rows, err := Db.Query("SELECT uuid, game_name, created_at FROM games ORDER BY created_at DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		game := GameList{}
		err = rows.Scan(&game.uuid, &game.gameName, &game.createdAt)
		if err != nil {
			return
		}
		games = append(games, game)
	}
	rows.Close()
	return
}
