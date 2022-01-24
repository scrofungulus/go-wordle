package wordle

import (
	"math/rand"
	"strings"
	"time"
)

func words() []string {
	return strings.Split(wordsStr, ",")
}

func existsInWords(g string) bool {
	for _, w := range words() {
		if w == g {
			return true
		}
	}

	return false
}

func randomWord() string {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	randomIndex := r.Intn(len(words()))
	pick := words()[randomIndex]
	return pick
}
