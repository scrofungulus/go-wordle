package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/scrofungulus/go-wordle/pkg/wordle"
)

func main() {
	game := wordle.New()

	for {
		consoleReader := bufio.NewReader(os.Stdin)
		fmt.Print(">")

		input, _ := consoleReader.ReadString('\n')
		input = strings.ToLower(input)
		stripped := strings.TrimSpace(input)

		lenGuess := len(stripped)
		if lenGuess != 5 {
			fmt.Println(">Your guess must be exactly 5 letters!")
			continue
		}

		correct, err := game.Guess(stripped)
		if err != nil {
			fmt.Println(fmt.Sprintf(">%s", err.Error()))
			continue
		}

		game.PrintGuesses()

		if correct {
			fmt.Println(fmt.Sprintf(">You guessed the correct word: %s!", game.Word()))
			os.Exit(0)
		}

		if game.IsOver() {
			fmt.Println(">Game over!")
			fmt.Println(fmt.Sprintf(">The word was: %s", game.Word()))
			os.Exit(0)
		}

		fmt.Println(fmt.Sprintf(">Number of guesses left: %d", game.GuessesLeft()))
	}
}
