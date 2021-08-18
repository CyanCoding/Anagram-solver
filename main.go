package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"

	gosort "github.com/cyancoding/go-sort"
	"github.com/slok/gospinner"
)

var dictionary []string
var results []string
var s *gospinner.Spinner

func readWordsFile(wordsFile string) []string {
	byteData, err := ioutil.ReadFile(wordsFile)

	if err != nil {
		fmt.Println(err)
	}

	text := string(byteData)
	words := strings.Fields(text)

	return words
}

func factorial(num int) int {
	factorial := 1

	for i := num; i > 0; i-- {
		factorial *= i
	}

	return factorial
}

func duplicateLetter(word string) map[string]int {
	// Word map of string: int values
	var wordMap = make(map[string]int)

	// For each letter, add to that key
	for i := 0; i < len(word); i++ {
		wordMap[string(word[i])]++
	}

	// In the end you have a map of 1 of each letter.
	// No duplicates are allowed in a map, so
	// we can print the length
	return wordMap
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

func getAnagrams(word string, possibilities int) {
	var alreadyUsed []string

	for len(alreadyUsed) != possibilities {

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
			results = append(results, string(chars))
			s.SetMessage("Finding matches (" + strconv.Itoa(len(results)) +
				"/" + strconv.Itoa(possibilities) + ")")
		}
	}
}

func main() {
	// go func() {
	// 	dictionary = readWordsFile("dictionary.txt")
	// }()

	// Create a hashmap of the dictionary
	dictionary = readWordsFile("dictionary.txt")
	hash := CreateHash(dictionary)

	fmt.Print("Anagram > ")
	var word string = "hello"
	fmt.Scanln(&word)
	lengthFactorial := factorial(len(word))

	divisor := 1

	for _, value := range duplicateLetter(word) {
		divisor *= factorial(value)
	}
	s, _ = gospinner.NewSpinner(gospinner.Dots2)
	s.Start("Finding matches (0/" + strconv.Itoa(lengthFactorial/divisor) + ")")

	getAnagrams(word, lengthFactorial/divisor)

	s.SetMessage("Found all matches!")
	s.Succeed()

	fmt.Println(hash[gosort.SortString(word)])
	for _, a := range results {
		fmt.Println(a)
	}

}
