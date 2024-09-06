package gamelogic

import (
	"fmt"
	"reflect"
	"testing"
)

func TestShuffleAndDeal(t *testing.T) {
	t.Run("check cards shuffle", func(t *testing.T) {
		game := NewGame()

		game.ShuffleAndDeal(InitialDeck...)

		p1 := game.player1Cards
		p2 := game.player2Cards

		if reflect.DeepEqual(p1, p2) {
			t.Errorf("got %v and %v hands should be different", p1, p2)
		}
	})

	t.Run("check each player has the same number of cards", func(t *testing.T) {
		game := NewGame()

		game.ShuffleAndDeal(InitialDeck...)

		p1 := len(game.player1Cards)
		p2 := len(game.player2Cards)

		if p1 != p2 {
			t.Errorf("got %v and %v hands should be same length", p1, p2)
		}
	})
}

func TestBattle(t *testing.T) {
	t.Run("check player 1 wins", func(t *testing.T) {
		game := NewGame()

		game.player1Cards = InitialDeck[0:2]
		game.player2Cards = InitialDeck[2:]

		fmt.Println("p1:", game.player1Cards)
		fmt.Println("p1:", game.player2Cards)

		want := "player 1"
		got := game.Battle("headology")

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

		wantP1Len := len(game.player1Cards)
		wantP2Len := len(game.player2Cards)

		if wantP1Len != 3 && wantP2Len != 1 {
			t.Errorf("P1 length %d want %d P2 length %d want %d", wantP1Len, 3, wantP2Len, 1)
		}
	})
}
