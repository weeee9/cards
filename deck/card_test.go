package deck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCard(t *testing.T) {
	t.Run("Ace of Hearts", func(t *testing.T) {
		c := Card{
			Rank: ace,
			Suit: heart,
		}
		assert.Equal(t, "Ace of Hearts", c.String())
	})

	t.Run("Three of Spades", func(t *testing.T) {
		c := Card{
			Rank: three,
			Suit: spade,
		}
		assert.Equal(t, "Three of Spades", c.String())
	})

	t.Run("Nine of Diamonds", func(t *testing.T) {
		c := Card{
			Rank: nine,
			Suit: diamond,
		}
		assert.Equal(t, "Nine of Diamonds", c.String())
	})

	t.Run("Jack of Clubs", func(t *testing.T) {
		c := Card{
			Rank: jack,
			Suit: club,
		}
		assert.Equal(t, "Jack of Clubs", c.String())
	})

	t.Run("Joker", func(t *testing.T) {
		c := Card{
			Suit: joker,
		}
		assert.Equal(t, "Joker", c.String())
	})
}
