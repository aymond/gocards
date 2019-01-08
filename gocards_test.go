package gocards

import "testing"

func TestNew(t *testing.T) {
	var err error
	d, err := newDeck("Test")
	if err != nil {
		t.Error(err)
	}
	if d.Name != "Test" {
		t.Error("Wrong Name. Was expecting Test but got ", d.Name)
	}
}

func TestInitialize(t *testing.T) {
	var err error
	var d Deck
	d.initialize("Blackjack")
	if err != nil {
		t.Error(err)
	}
	if d.Name != "Test" {
		t.Error("Wrong Name. Was expecting Test but got ", d.Name)
	}

}
