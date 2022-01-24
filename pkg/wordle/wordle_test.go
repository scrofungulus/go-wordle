package wordle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordle_Guess(t *testing.T) {
	wordle := &Wordle{word: "stars"}
	wordle.Guess("stars")

	assert.Equal(t, true, wordle.correct())

	wordle.Guess("strap")
	assert.Equal(t, false, wordle.correct())

	wordle.Guess("stars")
	assert.Equal(t, true, wordle.correct())
}

func TestWordle_toLetters(t *testing.T) {
	wordle := &Wordle{word: "stars"}

	letters := []letter{
		{
			val:   "s",
			color: Green,
		},
		{
			val:   "t",
			color: Green,
		},
		{
			val:   "a",
			color: Green,
		},
		{
			val:   "r",
			color: Green,
		},
		{
			val:   "s",
			color: Green,
		},
	}

	assert.Equal(t, letters, wordle.toLetters("stars"))

	letters = []letter{
		{
			val:   "s",
			color: Green,
		},
		{
			val:   "r",
			color: Yellow,
		},
		{
			val:   "t",
			color: Yellow,
		},
		{
			val:   "o",
			color: DarkGray,
		},
		{
			val:   "s",
			color: Green,
		},
	}

	assert.Equal(t, letters, wordle.toLetters("srtos"))

	letters = []letter{
		{
			val:   "w",
			color: DarkGray,
		},
		{
			val:   "o",
			color: DarkGray,
		},
		{
			val:   "o",
			color: DarkGray,
		},
		{
			val:   "d",
			color: DarkGray,
		},
		{
			val:   "y",
			color: DarkGray,
		},
	}

	assert.Equal(t, letters, wordle.toLetters("woody"))
}
