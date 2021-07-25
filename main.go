package main

import (
	"fmt"
	"strings"

	"github.com/weeee9/cards/deck"
)

type Hand []deck.Card

func (h Hand) String() string {
	str := make([]string, len(h))
	for i := range h {
		str[i] = h[i].String()
	}
	return strings.Join(str, ", ")
}

func (h Hand) DealerString() string {
	return fmt.Sprintf("%s, ** Hidden **", h[0])
}

func (h Hand) MinScore() int {
	score := 0
	for _, c := range h {
		// jack, queen and king count as 10 point
		score += min(int(c.Rank), 10)
	}
	return score
}

func min(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (h Hand) Score() int {
	minScore := h.MinScore()
	if minScore > 11 {
		return minScore
	}
	for _, c := range h {
		// ace is currently worth 1, and we are change it to be worth 11
		// 11 - 1 = 10
		if c.Rank == 1 {
			return minScore + 10
		}
	}
	return minScore
}

func main() {
	deck := deck.New(deck.Decks(3), deck.Shuffle)

	var player, dealer Hand

	for i := 0; i < 2; i++ {
		for _, hand := range []*Hand{&player, &dealer} {
			*hand = append(*hand, deck.Draw())
		}
	}

	var input string
	for input != "s" {
		fmt.Printf("Player: %s\n", player.String())
		fmt.Printf("Dealer: %s\n", dealer.DealerString())
		fmt.Printf("What would you do?, (h)it or (s)tand\n")

		fmt.Scanf("%s", &input)

		switch input {
		case "h":
			player = append(player, deck.Draw())
		}

		if player.Score() > 21 {
			fmt.Printf("Player: %s\n", player.String())
			fmt.Println("YOU LOSE")
			fmt.Println("GAME OVER")
			return
		}
	}
	fmt.Printf("\n===== Final Hands =====\n")
	fmt.Printf("Player: %s\n", player.String())
	fmt.Printf("Dealer: %s\n", dealer.String())
}
