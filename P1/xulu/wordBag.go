package xulu

import (
	"errors"
	"math"
)

type WordBag struct {
	Letters map[rune]int
}

func (w *WordBag) CalculateNumbers() int {
	calculatedNumber := 0
	for letter, count := range w.Letters {
		letterNumber := xuluLetters[letter]
		letterCount := count

		calculatedNumber += int(math.Pow(float64((letterNumber*letterCount)%5), 2))
	}

	return calculatedNumber
}

func (w *WordBag) Parse(word string) error {
	if _, ok := xuluVerbs[word]; ok {
		return errors.New("value is a verb")
	}

	w.Letters = make(map[rune]int)

	if len(word) < 1 {
		return nil
	}

	for _, letter := range word {
		if letter < 'a' || letter > 'e' {
			return errors.New("invalid character")
		}

		if _, ok := w.Letters[letter]; !ok {
			w.Letters[letter] = 0
		}

		w.Letters[letter]++
	}

	return nil
}
