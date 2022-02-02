package wordle

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

type (
	Wordle struct {
		word    string
		guesses [][]letter
	}

	letter struct {
		val   string
		color WordleColor
	}

	WordleColor string
)

const (
	Green    WordleColor = "green"
	Yellow   WordleColor = "yellow"
	DarkGray WordleColor = "dark gray"
)

const TOTAL_GUESSES = 6

func New() *Wordle {
	word := randomWord()
	return &Wordle{
		word: word,
	}
}

func (w *Wordle) Guess(g string) (bool, error) {
	if !existsInWords(g) {
		return false, fmt.Errorf(`"%s" does not exist in list of words!`, g)
	}

	w.addGuess(g)
	return w.correct(), nil
}

func (w *Wordle) Guesses() [][]letter {
	return w.guesses
}

func (w *Wordle) addGuess(g string) {
	w.guesses = append(w.guesses, w.toLetters(g))
}

func (w *Wordle) toLetters(g string) []letter {
	guessChars := charSlice(g)
	letters := make([]letter, len(guessChars))

	wordCharsWithGuesses := charsWithGuesses(w.word)

	for gi, gc := range guessChars {
		if w.wordContains(gc) && w.wordContainsAtPosition(gc, gi) {
			letters[gi] = letter{val: gc, color: Green}
			wordCharsWithGuesses[gi].guessed = true
		}
	}

	for gi, gc := range charSlice(g) {
		if wordCharsWithGuesses[gi].guessed {
			continue
		}

		if !w.wordContains(gc) {
			letters[gi] = letter{val: gc, color: DarkGray}
			continue
		}

		ind := strings.Index(w.word, gc)
		exists := ind > -1
		existsNotGuessed := exists && !wordCharsWithGuesses[ind].guessed && func() bool {
			present := false
			for _, wc := range wordCharsWithGuesses {
				if wc.char == gc && wc.present {
					present = true
				}
			}

			return !present
		}()

		if existsNotGuessed {
			letters[gi] = letter{val: gc, color: Yellow}
			wordCharsWithGuesses[ind].present = true
		} else {
			letters[gi] = letter{val: gc, color: DarkGray}
		}
	}

	return letters
}

type charWithGuess struct {
	char    string
	guessed bool
	present bool
}

func charsWithGuesses(word string) []charWithGuess {
	chars := []charWithGuess{}
	for _, char := range charSlice(word) {
		chars = append(chars, charWithGuess{char: char, guessed: false, present: false})
	}

	return chars
}

func (w *Wordle) Word() string {
	return w.word
}

func (w *Wordle) getWordChars() []string {
	return charSlice(w.word)
}

func (w *Wordle) wordContains(char string) bool {
	for _, c := range w.getWordChars() {
		if c == char {
			return true
		}
	}

	return false
}

func (w *Wordle) wordContainsAtPosition(char string, pos int) bool {
	chars := w.getWordChars()
	return char == chars[pos]
}

func charSlice(g string) []string {
	return strings.Split(g, "")
}

func (w *Wordle) getCurrGuess() []letter {
	return w.guesses[len(w.guesses)-1]
}

func (w *Wordle) getCurrGuessString() string {
	str := ""
	for _, letter := range w.getCurrGuess() {
		str += letter.val
	}

	return str
}

func (w *Wordle) correct() bool {
	return w.getCurrGuessString() == w.word
}

func (w *Wordle) IsOver() bool {
	return len(w.Guesses()) >= TOTAL_GUESSES
}

func (w *Wordle) PrintGuesses() {
	allGuesses := ""
	for _, guess := range w.Guesses() {
		guessStr := "|"
		for i, letter := range guess {
			print := colorFunc(letter.color)
			if i == len(guess)-1 {
				coloredLetter := print("", letter.val)
				guessStr += fmt.Sprintf(" %s |\n", coloredLetter)
			} else {
				coloredLetter := print("", letter.val)
				guessStr += fmt.Sprintf(" %s |", coloredLetter)
			}
		}
		allGuesses += guessStr
	}

	fmt.Println(allGuesses)
}

func colorFunc(c WordleColor) func(a ...interface{}) string {
	switch c {
	case Green:
		return color.New(color.FgGreen).SprintFunc()
	case Yellow:
		return color.New(color.FgYellow).SprintFunc()
	case DarkGray:
		return color.New(color.FgRed).SprintFunc()
	default:
		return nil
	}
}

func (w *Wordle) GuessesLeft() int {
	return TOTAL_GUESSES - len(w.guesses)
}
