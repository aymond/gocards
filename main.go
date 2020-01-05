package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	p("gocards", version(), "started.")

	// Create a new Game State and store in the DB. This should later be handled by a /new handler.
	dbCreate()
	var gs BattleKingsGameState

	_ = gs.NewGame(true)
	_ = gs.SaveState()
	//gs.Debug()

	r := http.NewServeMux()
	r.HandleFunc("/", homeHandler)
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Handlers

func homeHandler(w http.ResponseWriter, r *http.Request) {
	//_, err := session(w, r)

	// A very simple health check.
	w.WriteHeader(http.StatusOK)
	// w.Header().Set("Content-Type", "application/json")
	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, "<!DOCTYPE html><html><body>Hello, World!")
	io.WriteString(w, "</body></html>")
}
