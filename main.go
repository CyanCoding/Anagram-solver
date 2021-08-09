package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

var dictionary []string

func readWordsFile(wordsFile string) []string {
	byteData, err := ioutil.ReadFile(wordsFile)

	if err != nil {
		fmt.Println(err)
	}

	text := string(byteData)
	words := strings.Fields(text)

	return words
}

// Checks to see if a string is in a string array
func containsString(array []string, test string) bool {
	for _, a := range array {
		if a == test {
			return true
		}
	}
	return false
}

func getAnagrams(word string) {
	var alreadyUsed []string

	// Get amount of possibilities
	factorial := 1
	// Get factorial of anagrams
	for i := len(word); i > 0; i-- {
		factorial = i * factorial
	}

	fmt.Println(factorial)

	for len(alreadyUsed) != factorial {

		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Millisecond)

		chars := []rune(word)
		for i := len(word) - 1; i > 0; i-- {
			index := rand.Intn(i + 1)
			letter := chars[index]
			chars[index] = chars[i]
			chars[i] = letter
		}

		if !containsString(alreadyUsed, string(chars)) {
			alreadyUsed = append(alreadyUsed, string(chars))
			fmt.Println(string(chars))
		}
	}
}

func main() {
	go func() {
		dictionary = readWordsFile("dictionary.txt")
	}()

	fmt.Print("Anagram > ")
	var word string = "snacky"
	fmt.Scanln(&word)

	getAnagrams(word)
}
