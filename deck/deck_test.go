package deck

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDeck(t *testing.T) {
	t.Run("test new deck", func(t *testing.T) {
		assert.Len(t, New(), 13*4)
	})

	t.Run("test new deck with default sort", func(t *testing.T) {
		deck := New(DefaultSort)
		expected := Card{
			Suit: spade,
			Rank: ace,
		}
		assert.Equal(t, expected, deck[0],
			"expected 'Ace of Spades' as first card. Received: '%v'", deck[0])
	})
}

func TestJoker(t *testing.T) {
	t.Run("add 3 jokers", func(t *testing.T) {
		deck := New(Jokers(3))
		cnt := 0
		for _, c := range deck {
			if c.Suit == joker {
				cnt++
			}
		}
		assert.Equal(t, 3, cnt)
	})
}

func TestFilter(t *testing.T) {
	t.Run("filter rank two", func(t *testing.T) {
		filter := func(c Card) bool {
			return c.Rank == two
		}
		deck := New(Filter(filter))

		// full deck (52) without 2s (4)
		require.Len(t, deck, 48)

		for _, c := range deck {
			require.NotEqual(t, c.Rank, two)
		}
	})

	t.Run("filter rank two and three", func(t *testing.T) {
		filter := func(c Card) bool {
			return c.Rank == two || c.Rank == three
		}
		deck := New(Filter(filter))

		// full deck (52) without 2s (4) and 3s (4)
		require.Len(t, deck, 44)

		for _, c := range deck {
			require.NotContains(t, []rank{two, three}, c.Rank,
				"deck should not contain rank '%v'", c.Rank)
		}
	})
}

func TestDecks(t *testing.T) {
	t.Run("test init with 3 decks", func(t *testing.T) {
		deck := New(Decks(3))
		// 4 * 13 * 3
		require.Len(t, deck, 52*3,
			"expected deck contain %d of card, received %d of cards", 52*3, len(deck))
	})
}
