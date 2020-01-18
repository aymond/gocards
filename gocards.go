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

// moveCard moves a card (c) from one deck (d1) to another deck (d2).
func (gs *BattleKingsGameState) moveCard(d1 *Deck, d2 *Deck) {
	var c Card
	c, d1.Cards = d1.Cards[len(d1.Cards)-1], d1.Cards[:len(d1.Cards)-1]
	d2.Cards = append(d2.Cards, c)
	//return c
}

// Debug prints the current state of the passed deck
func (gs *BattleKingsGameState) Debug() {
	//s, _ := json.Marshal(gs.gamePile)
	fmt.Printf("DEBUG: %+v", gs)
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
