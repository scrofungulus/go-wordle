# go-wordle

A rudimentary implementation of the viral game Wordle. 

**Disclaimer:** I did not research every single rule of the actual game but I believe logic in this implementation is close.

## Data Source

Words are hard coded as a comma delimited list in the binary. In the future I might source words from an API.

5 Letter words were sourced from [charlesreid1](https://github.com/charlesreid1)'s repository [five-letter-words](https://github.com/charlesreid1/five-letter-words).

## How To Play

Guess the word in **6** tries.

Each guess must be exactly five letters.

If the word you guess contains letters of the chosen word the letters will be outputted
in either yellow or green. Green if the letter is in the word and in the right location.
Yellow if the letter is in the word in the wrong location. Otherwise the remaining
letters will show red.

If the word is `point` and you guess `poopy` only one `o` will be highlighted yellow. 
However, if the word was `loops` and you guessed `poopy` both `o`'s would be 
highlighted green.
