package test

import (
	"testing"

	"example.com/xulu"
)

func TestSimpleWordParsing(t *testing.T) {
	x := xulu.WordBag{}
	x.Parse("cccd") // Simple parse test
	if len(x.Letters) != 2 {
		t.Errorf("faulty word parsing, expected %d, got %d", 2, len(x.Letters))
	}
}

func TestVerbParsing(t *testing.T) {
	x := xulu.WordBag{}
	err := x.Parse("dede") // Parse verb should throw error
	if err == nil {
		t.Error("did not throw error on verb parsing")
	}
}

func TestEmptyWordParsing(t *testing.T) {
	x := xulu.WordBag{}
	x.Parse("") // Parse verb should throw error
	if len(x.Letters) != 0 {
		t.Error("empty string parsing doesn't result in 0 words")
	}
}

func TestInvalidSentenceVerb(t *testing.T) {
	x := xulu.Sentence{}
	_, err := x.ParseSentence("abdd abcd") // Parse character should throw error
	if err == nil {
		t.Error("invalid result, expected error, got nil")
	}
}

func TestInvalidSentenceCharacter(t *testing.T) {
	x := xulu.Sentence{}
	_, err := x.ParseSentence("abcd agcr") // Parse character should throw error
	if err == nil {
		t.Error("invalid result, expected error, got nil")
	}
}
