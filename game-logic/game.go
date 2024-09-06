package gamelogic

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"golang.org/x/exp/rand"
)

type Card struct {
	id        int
	name      string
	strength  int
	headology int
	magic     int
	luck      int
}

type Game struct {
	player1Cards []Card
	player2Cards []Card
}

func NewGame() *Game {
	return &Game{}
}

func (game *Game) ShuffleAndDeal(deck ...Card) {
	r := rand.New(rand.NewSource(uint64(time.Now().Unix())))
	var deckCopy []Card

	copy(deckCopy, deck)

	r.Shuffle(len(deckCopy), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

	for i, v := range deckCopy {
		if i%2 != 0 {
			game.player1Cards = append(game.player1Cards, v)
		} else {
			game.player2Cards = append(game.player2Cards, v)
		}
	}
}

func (game *Game) Battle(attribute string) string {

	p1 := reflect.ValueOf(game.player1Cards[0])
	p1Att := reflect.Indirect(p1).FieldByName(attribute).Int()
	p2 := reflect.ValueOf(game.player2Cards[0])
	p2Att := reflect.Indirect(p2).FieldByName(attribute).Int()

	if p1Att == p2Att {
		return "stalemate"
	} else if p1Att > p2Att {
		game.SwapCards("p1")
		return "player 1"
	} else {
		game.SwapCards("p2")
		return "player 2"
	}
}

func (game *Game) SwapCards(winner string) {
	var winnerSlice []Card
	var loserSlice []Card
	var winningCard Card
	var loosingCard Card

	if strings.Compare("p1", winner) == 0 {
		winnerSlice = make([]Card, len(game.player1Cards))
		loserSlice = make([]Card, len(game.player2Cards))
		copy(winnerSlice, game.player1Cards)
		copy(loserSlice, game.player2Cards)

	} else {
		copy(winnerSlice, game.player2Cards)
		copy(loserSlice, game.player1Cards)
	}

	fmt.Println(winnerSlice)
	fmt.Println(loserSlice)

	if len(winnerSlice) > 0 && len(loserSlice) > 0 {

		winningCard = winnerSlice[0]
		loosingCard = loserSlice[0]

		winnerSlice = winnerSlice[1:]
		loserSlice = loserSlice[1:]

		winnerSlice = append(winnerSlice, winningCard, loosingCard)

		if winner == "p1" {
			game.player1Cards = winnerSlice
			game.player2Cards = loserSlice
		} else {
			game.player2Cards = winnerSlice
			game.player1Cards = loserSlice
		}
	}
}
