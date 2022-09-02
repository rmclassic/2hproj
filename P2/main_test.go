package main

import "testing"

func TestStringNormalization(t *testing.T) {
	x := NormalizeString("1 + 2")
	if x != "1 2" {
		t.Error("invalid string normalization: simple")
	}

	x = NormalizeString("1 +,2")
	if x != "1 2" {
		t.Error("invalid string normalization: plus & comma")
	}

	x = NormalizeString("1  +   -2")
	if x != "1 -2" {
		t.Error("invalid string normalization: multiple spaces")
	}

	x = NormalizeString("1 2 ")
	if x != "1 2" {
		t.Error("invalid string normalization: trailing space")
	}

	x = NormalizeString(" 1 2")
	if x != "1 2" {
		t.Error("invalid string normalization: starting space")
	}
}

func TestOperationCalculation(t *testing.T) {
	a, _ := DoOperation("1 ,+ -2")
	if a != -1 {
		t.Error("operation result invalid: negative number")
	}

	a, _ = DoOperation("-1+,+1")
	if a != 0 {
		t.Error("operation result invalid: negative first number")
	}

	_, err := DoOperation("-1+,-+1")
	if err == nil {
		t.Error("err was nil, expected value")
	}
}
