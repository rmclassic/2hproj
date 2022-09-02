package xulu

import (
	"errors"
	"strings"

	"example.com/xulu/interfaces"
)

type Sentence struct {
	Verb  string
	Words []interfaces.IWordSentence
}

func (s Sentence) CalculateNumbers() int {
	if len(s.Words) == 0 {
		return 0
	}

	calculatedNum := s.Words[0].CalculateNumbers() // We store the first result in accumulated result so the multiplication would not be 0*result

	for _, w := range s.Words[1:] {
		calculatedNum = xuluVerbs[s.Verb](calculatedNum, w.CalculateNumbers())
	}

	return calculatedNum
}

func (s *Sentence) ParseSentence(sentence string) (string, error) {

	if !s.isSentence(sentence) {
		return sentence, errors.New("sentence is incorrect")
	}

	verb, words := s.getNextToken(sentence)
	s.Verb = verb

	var err error
	s.Words = make([]interfaces.IWordSentence, 0)

	if s.isSentence(words) { // see if we have another sentence nested in this sentence
		for {
			innerSentence := Sentence{}
			words, err = innerSentence.ParseSentence(words)
			if err != nil {
				break
			}
			s.Words = append(s.Words, innerSentence)
		}
	} else {
		words, err = s.parseWords(words)
	}

	return words, err
}

func (*Sentence) isSentence(s string) bool {
	if len(s) < 4 {
		return false
	}

	_, ok := xuluVerbs[s[:4]]
	return ok
}

func (s *Sentence) parseWords(words string) (string, error) {
	word := ""

	for words != "" && !s.isSentence(words) {
		word, words = s.getNextToken(words)
		wordBag := WordBag{}
		err := wordBag.Parse(word)
		if err != nil {
			return "", err
		}
		s.Words = append(s.Words, &wordBag)
	}

	return words, nil
}

func (*Sentence) getNextToken(words string) (string, string) {
	i := strings.Index(words, " ")
	if i == -1 {
		return words, ""
	}

	word := words[:i]
	nextWords := words[i+1:]
	return word, nextWords
}
