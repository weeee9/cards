package blackjack

import (
	"strings"

	"github.com/weeee9/cards/deck"
)

var _ deck.Player = (*Player)(nil)

type Player struct {
	hand     []deck.Card
	isDealer bool
	stand    bool
}

func NewDealer() *Player {
	p := NewPlayer()
	p.isDealer = true
	return p
}

// NewPlayer
func NewPlayer() *Player {
	return &Player{}
}

func (p *Player) CanMove() bool {
	return !p.stand
}

func (p *Player) Hit(c deck.Card) {
	p.hand = append(p.hand, c)
}

func (p *Player) Stand() {
	p.stand = true
}

func (p *Player) Hand() string {
	str := make([]string, len(p.hand))
	for i, c := range p.hand {
		str[i] = c.String()
	}

	return strings.Join(str, ", ")
}

func (p *Player) DealerHand() string {
	str := make([]string, len(p.hand))
	for i, c := range p.hand {
		if i == 1 {
			str[i] = "**HIDDEN**"
			continue
		}
		str[i] = c.String()
	}
	return strings.Join(str, ", ")
}

func (p *Player) Score() int {
	minScore := p.MinScore()

	// handle Ace

	// if original score is greater than 11
	// then Ace can't count as 11 point
	if minScore > 11 {
		return minScore
	}
	// if original score is less than or equal to 11
	// find if there's any Ace in hands
	// if has, count it as 11 point
	for _, c := range p.hand {
		// ace is currently worth 1, and we are change it to be worth 11
		// 11 - 1 = 10
		if c.IsAce() {
			return minScore + 10
		}
	}

	return minScore
}

func (p *Player) MinScore() int {
	score := 0
	for _, c := range p.hand {
		// jack, queen and king count as 10 point
		score += min(c.Point(), 10)
	}
	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
