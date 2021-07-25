package deck

import (
	"fmt"
)

type suit uint8

const (
	spade suit = iota
	heart
	diamond
	club
	joker
)

var (
	suits = [...]suit{spade, heart, diamond, club}

	suitMap = map[suit]string{
		spade:   "Spade",
		heart:   "Heart",
		diamond: "Diamond",
		club:    "Club",
		joker:   "Joker",
	}
)

func (s suit) IsValid() bool {
	switch s {
	case spade, heart, diamond, club, joker:
		return true
	}
	return false
}

func (s suit) String() string {
	if s.IsValid() {
		return suitMap[s]
	}
	return ""
}

func (s suit) point() int {
	if s == joker {
		return 5
	}
	return 4 - int(s)
}

type rank uint8

const (
	_ rank = iota
	ace
	two
	three
	four
	five
	six
	seven
	eight
	nine
	ten
	jack
	queen
	king

	minRank = ace
	maxRank = king
)

var rankMap = map[rank]string{
	ace:   "Ace",
	two:   "2",
	three: "3",
	four:  "4",
	five:  "5",
	six:   "6",
	seven: "7",
	eight: "8",
	nine:  "9",
	ten:   "9",
	jack:  "Jack",
	queen: "Queen",
	king:  "King",
}

func (r rank) IsValid() bool {
	switch r {
	case ace, two, three, four, five, six,
		seven, eight, nine, ten, jack, queen, king:
		return true
	}
	return false
}

func (r rank) String() string {
	if r.IsValid() {
		return rankMap[r]
	}
	return ""
}

func (r rank) isAce() bool {
	return r == ace
}

func (r rank) point() int {
	return int(r)
}

type Card struct {
	Suit suit
	Rank rank
}

func (c Card) String() string {
	if c.Suit == joker {
		return "Joker"
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func (c Card) IsAce() bool {
	return c.Rank.isAce()
}

func (c Card) Point() int {
	return c.Rank.point()
}

func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}
