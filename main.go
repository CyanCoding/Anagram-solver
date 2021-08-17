package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	gosort "github.com/cyancoding/go-sort"
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

	fmt.Println(hash[gosort.SortString(word)])

	//getAnagrams(word)
}
