package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	p("gocards", version(), "started.")

	// Create a new Game State and store in the DB. This should later be handled by a /new handler.
	dbCreate()
	var gs BattleKingsGameState

	_ = gs.NewGame(true)
	//_ = gs.SaveState()
	//gs.Debug()

	r := http.NewServeMux()
	r.HandleFunc("/", homeHandler)
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Handlers

func homeHandler(w http.ResponseWriter, r *http.Request) {
	//_, err := session(w, r)
	gamelist, _ := ListGames()
	p(gamelist)
	// A very simple health check.
	w.WriteHeader(http.StatusOK)
	t, _ := template.ParseGlob("templates/*.html")
	t.ExecuteTemplate(w, "home", gamelist)
}
