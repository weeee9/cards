package blackjack

import (
	"fmt"

	"github.com/weeee9/cards/deck"
)

type turn uint

const (
	gameOver turn = iota
	playerTurn
	dealerTurn
)

var defaultDeckOpts = []deck.DeckOption{
	deck.Decks(3), deck.Shuffle,
}

func (t turn) String() string {
	switch t {
	case playerTurn:
		return "Player"
	case dealerTurn:
		return "Dealer"
	}
	return ""
}

type gameState struct {
	Deck   deck.Deck
	Turn   turn
	Player *Player
	Dealer *Player
}

func (gs gameState) String() string {
	dealer := gs.Dealer.DealerHand()
	if gs.Turn == dealerTurn || gs.Turn == gameOver {
		dealer = gs.Dealer.Hand()
	}

	return fmt.Sprintf(`It's %s turn,
Decks remain %d cards,
Player has: [ %s ],
Dealer has: [ %s ]`, gs.Turn, len(gs.Deck), gs.Player.Hand(), dealer)
}

func (gs gameState) PlayerState() string {
	return fmt.Sprintf(`Player's Hand [ %s ],
Score: %d`, gs.Player.Hand(), gs.Player.Score())
}

func (gs gameState) DealerState() string {
	return fmt.Sprintf(`Dealer's Hand [ %s ],
Score: %d`, gs.Dealer.Hand(), gs.Dealer.Score())
}

func (gs gameState) PlayerBusted() bool {
	return gs.Player.Score() > 21
}

func (gs gameState) DealerBusted() bool {
	return gs.Dealer.Score() > 21
}

func NewGame(opts ...deck.DeckOption) gameState {
	if len(opts) == 0 {
		opts = defaultDeckOpts
	}
	deck := deck.New(opts...)
	player := NewPlayer()
	dealer := NewDealer()

	deck.Deal(2, dealer, player)
	return gameState{
		Deck:   deck,
		Turn:   playerTurn,
		Player: player,
		Dealer: dealer,
	}
}

func Deal(gs gameState) gameState {
	switch gs.Turn {
	case playerTurn:
		gs.Player.Hit(gs.Deck.Draw())
	case dealerTurn:
		gs.Dealer.Hit(gs.Deck.Draw())
	}

	return gs
}

func Stand(gs gameState) gameState {
	switch gs.Turn {
	case playerTurn:
		gs.Turn = dealerTurn
	case dealerTurn:
		gs.Turn = gameOver
	}
	return gs
}

func (gs gameState) DealerMove() bool {
	return gs.Dealer.Score() <= 16 || (gs.Dealer.Score() == 17 && gs.Dealer.MinScore() != gs.Dealer.Score())
}

func End(gs gameState) {
	if gs.Turn != gameOver {
		return
	}
	if gs.Player.Score() > gs.Dealer.Score() {
		fmt.Println("YOU WIN")
		return
	}
	if gs.Player.Score() < gs.Dealer.Score() {
		fmt.Println("YOU LOSE")
		return
	}
	fmt.Println("DRAW")
}
