package test

import (
	"testing"

	"example.com/xulu"
)

func TestSimpleSentenceParsing(t *testing.T) {
	x := xulu.Sentence{}
	x.ParseSentence("abcd a b")
	if len(x.Words) != 2 {
		t.Errorf("invalid word count, expected %d, got %d", 2, len(x.Words))
	}
}

func TestNestedSentenceParsing(t *testing.T) {
	x := xulu.Sentence{}
	x.ParseSentence("abcd abcd a b")
	if len(x.Words) != 1 {
		t.Errorf("invalid sub-sentence count, expected %d, got %d", 1, len(x.Words))
	}
}

func TestSimpleSentenceCalculation(t *testing.T) {
	x := xulu.Sentence{}
	x.ParseSentence("abcd a b")
	r := x.CalculateNumbers()
	if r != 5 {
		t.Errorf("invalid result, expected %d, got %d", 5, r)
	}
}

func TestComplexSentenceCalculation(t *testing.T) {
	x := xulu.Sentence{}
	x.ParseSentence("abcd dede aab bba abcd cd bcd")
	r := x.CalculateNumbers()
	if r != 190 {
		t.Errorf("invalid result, expected %d, got %d", 190, r)
	}

	x.ParseSentence("abcd abcd aabbc ab a c ccd dede cccd cd")
	r = x.CalculateNumbers()
	if r != 861 {
		t.Errorf("invalid result, expected %d, got %d", 190, r)
	}
}

func TestInvalidSentence(t *testing.T) {
	x := xulu.Sentence{}
	_, err := x.ParseSentence("cbad dwa")
	if err == nil {
		t.Error("expected error, got nil")
	}
}
