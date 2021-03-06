package deck

import (
	"math/rand"
	"sort"
	"time"
)

type deck []Card

type Player interface {
	Hit(Card)
	Stand()
	Hand() string
}

func (d *deck) Draw() Card {
	card := (*d)[0]
	*d = (*d)[1:]
	return card
}

// Deal n cards to each player
func (d *deck) Deal(n int, players ...Player) {
	for i := 0; i < n; i++ {
		for j := range players {
			players[j].Hit(d.Draw())
		}
	}
}

func New(opts ...DeckOption) deck {
	var cards []Card

	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{
				Suit: suit,
				Rank: rank,
			})
		}
	}

	for _, opt := range opts {
		cards = opt(cards)
	}

	return cards
}

type DeckOption func([]Card) []Card

func Decks(n int) DeckOption {
	return func(deck []Card) []Card {
		result := []Card{}
		for i := 0; i < n; i++ {
			result = append(result, deck...)
		}
		return result
	}
}

func DefaultSort(deck []Card) []Card {
	sort.Slice(deck, Less(deck))
	return deck
}

func Sort(less func(deck []Card) func(i, j int) bool) DeckOption {
	return func(deck []Card) []Card {
		sort.Slice(deck, less(deck))
		return deck
	}
}

func Less(deck []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(deck[i]) < absRank(deck[j])
	}
}

var shuffleRand = rand.New(rand.NewSource(time.Now().Unix()))

func Shuffle(deck []Card) []Card {
	result := make([]Card, len(deck))

	perm := shuffleRand.Perm(len(deck))

	for idx, p := range perm {
		result[idx] = deck[p]
	}

	return result
}

func Jokers(n int) DeckOption {
	return func(deck []Card) []Card {
		for i := 0; i < n; i++ {
			deck = append(deck, Card{
				Suit: joker,
				Rank: rank(i),
			})
		}
		return deck
	}
}

func Filter(f func(Card) bool) DeckOption {
	return func(deck []Card) []Card {
		result := make([]Card, 0)

		for _, c := range deck {
			if !f(c) {
				result = append(result, c)
			}
		}

		return result
	}
}
