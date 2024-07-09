package main

import (
	"fmt"
	"testing"

	"go-reloaded/transform"
)

func TestConvToDec(t *testing.T) {
	inputText := "Simply add 42 (hex) and 10 (bin) and you will see the result is 68."
	expectedText := "Simply add 66 and 2 and you will see the result is 68."

	str := transform.TransformWords(inputText)

	if str != expectedText {
		t.Fatalf("Test failed. Expected: %s, got: %s", expectedText, str)
	} else {
		fmt.Println("Test passed sucessfully!")
	}
}

func TestIsVowel(t *testing.T) {
	inputText := "There is no greater agony than bearing a untold story inside you."
	expectedText := "There is no greater agony than bearing an untold story inside you."

	str := transform.TransformWords(inputText)

	if str != expectedText {
		t.Fatalf("Test failed. Expected: %s, got: %s", expectedText, str)
	} else {
		fmt.Println("Test passed sucessfully!")
	}
}

func TestNum(t *testing.T) {
	inputText := "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair."
	expectedText := "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair."

	str := transform.TransformWords(inputText)

	if str != expectedText {
		t.Fatalf("Test failed. Expected: %s, got: %s", expectedText, str)
	} else {
		fmt.Println("Test passed sucessfully!")
	}
}

func TestPunctuation(t *testing.T) {
	inputText := "Punctuation tests are ... kinda boring ,don't you think !?"
	expectedText := "Punctuation tests are... kinda boring, don't you think!?"

	str := transform.TransformWords(inputText)

	if str != expectedText {
		t.Fatalf("Test failed. Expected: %s, got: %s", expectedText, str)
	} else {
		fmt.Println("Test passed sucessfully!")
	}
}
