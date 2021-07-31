package main

import (
	"fmt"

	"github.com/weeee9/cards/blackjack"
	"github.com/weeee9/cards/deck"
)

func main() {
	deck := deck.New(deck.Decks(3), deck.Shuffle)

	player, dealer := blackjack.NewPlayer(), blackjack.NewDealer()

	deck.Deal(2, dealer, player)

	var input string
	for input != "s" {

		fmt.Printf("Player: %s\n", player.Hand())
		fmt.Printf("Dealer: %s\n", dealer.DealerHand())
		fmt.Printf("What would you do?, (h)it or (s)tand\n")

		fmt.Scanf("%s", &input)

		switch input {
		case "h":
			player.Hit(deck.Draw())
		}

		if player.Score() > 21 {
			fmt.Printf("Player: %s\n", player.Hand())
			fmt.Println("YOU LOSE, Score:", player.Score())
			fmt.Println("GAME OVER")
			return
		}
	}

	for dealer.Score() <= 16 ||
		// Ace + 6 can count as 7 or 17 (soft 17)
		(dealer.Score() == 17 && dealer.MinScore() != dealer.Score()) {
		dealer.Hit(deck.Draw())
	}

	fmt.Printf("\n===== Final Hands =====\n")
	fmt.Printf("Player: %s\nScore: %d\n", player.Hand(), player.Score())
	fmt.Println("==========================")
	fmt.Printf("Dealer: %s\nScore: %d\n", dealer.Hand(), dealer.Score())

	if player.Score() > dealer.Score() {
		fmt.Println("YOU WIN")
		return
	}

	if player.Score() < dealer.Score() {
		fmt.Println("YOU LOSE")
		return
	}
	fmt.Println("DRAW")

}
