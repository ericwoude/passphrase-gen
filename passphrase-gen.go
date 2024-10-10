package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Constants for consonants and vowels
var (
	consonants = []rune{'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n', 'p', 'q', 'r', 's', 't', 'v', 'w', 'x', 'y', 'z'}
	vowels     = []rune{'a', 'e', 'i', 'o', 'u'}
)

// Utility function to capitalize a given string
func capitalize(s string) string {
	return strings.ToUpper(string(s[0])) + s[1:]
}

// Utility function to select a random rune from a slice
func randomChoice(choices []rune) rune {
	return choices[rand.Intn(len(choices))]
}

// Generate a syllable in consonant-vowel-consonant-(consonant)-vowel-consonant fashion
func generateSyllable() string {
	syllable := []rune{
		randomChoice(consonants),
		randomChoice(vowels),
		randomChoice(consonants),
	}

	// 50% chance to add an extra consonant
	if rand.Intn(2) == 0 {
		syllable = append(syllable, randomChoice(consonants))
	}

	// Add vowel - consonant to end the syllable
	syllable = append(syllable, randomChoice(vowels), randomChoice(consonants))

	return string(syllable)
}

// Generates a password with 4 syllables, a capitalized syllable, and a random digit insertion
func generatePassword() string {
	syllables := []string{
		generateSyllable(),
		generateSyllable(),
		generateSyllable(),
		generateSyllable(),
	}

	// Capitalize a random syllable
	randomIndex := rand.Intn(len(syllables))
	syllables[randomIndex] = capitalize(syllables[randomIndex])

	// Add a single digit before or after a random picked syllable
	digit := fmt.Sprintf("%d", rand.Intn(10))
	position := rand.Intn(len(syllables))
	if rand.Intn(2) == 0 {
		syllables[position] = digit + syllables[position]
	} else {
		syllables[position] = syllables[position] + digit
	}

	return strings.Join(syllables, "-")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(generatePassword())
}
