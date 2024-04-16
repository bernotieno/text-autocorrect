package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestConvertToDecimal(t *testing.T) {
	inputtext := "Simply add 42 (hex) and 10 (bin) and you will see the result is 68."
	expectedoutput := "Simply add 66 and 2 and you will see the result is 68."
	d := strings.Fields(inputtext)
	result := wordsChanger(d)
	finResult := strings.Join(result, " ")

	if finResult == expectedoutput {
		fmt.Println("TEST PASS")
	} else {
		fmt.Println("TEST FAILED")
	}
}

func TestNumberAfterCom(t *testing.T) {
	inputtext := "no no no no no no no no no no no no no no no no no no no no no no (cap, 12)"
	expectedoutput := "no no no no no no no no no no No No No No No No No No No No No No"

	n := strings.Fields(inputtext)
	result := wordsChanger(n)
	finalResult := strings.Join(result, " ")

	if finalResult == expectedoutput {
		fmt.Println("TEST PASS")
	} else {
		fmt.Println("TEST FAILED")
	}
}

func TestVowelCase(t *testing.T) {
	inputtext := "There is no greater agony than bearing a untold story inside you."
	expectedoutput := "There is no greater agony than bearing an untold story inside you."

	v := strings.Fields(inputtext)
	result := wordsChanger(v)
	finalResult := strings.Join(result, " ")

	if finalResult == expectedoutput {
		fmt.Println("TEST PASS")
	} else {
		fmt.Println("TEST FAILED")
	}
}

func TestPunctuations(t *testing.T) {
	inputtext := "Punctuation tests are ... kinda boring ,don't you think !?"
	expectedoutput := "Punctuation tests are... kinda boring, don't you think!?"

	p := strings.Fields(inputtext)
	result := wordsChanger(p)
	finalResult := strings.Join(result, " ")

	if finalResult == expectedoutput {
		fmt.Println("TEST PASS")
	} else {
		fmt.Println("TEST FAILED")
	}
}
