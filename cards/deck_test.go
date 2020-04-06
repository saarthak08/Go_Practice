package main
import (
	"testing"
	"os"
)


func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected length of deck to be 16. But got %v",len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card `Ace of Spades` but got %v",d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card `Four of Clubs` but got %v",d[len(d)-1])
	}
}


func  TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")
	deck := newDeck()
	deck.saveToFile("_deckTesting")
	loadedDeck := newDeckFromFile("_deckTesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected length of deck to be 16. But got %v",len(loadedDeck))
	}

	os.Remove("_deckTesting")

}