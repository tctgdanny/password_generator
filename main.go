package main

import (
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//go:embed word_lists/*.json
var wordLists embed.FS

func loadWordsList(wordLength string) ([]string, error) {
	// Load the word list file for the amount of letters designated by wordLength
	jsonFile, err := wordLists.Open(fmt.Sprintf("word_lists/%s_letter_words.json", wordLength))
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var wordsList []string
	json.Unmarshal(byteValue, &wordsList)
	return wordsList, nil
}

func generateWordNumber(wordsList []string, count int) {
	// Generate a random word from the loaded list, and a random number with 4 characters
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		word := wordsList[rand.Intn(len(wordsList))]
		word = strings.Title(word)
		number := rand.Intn(10000)
		fmt.Printf("%s+%04d\n", word, number)
	}
}

func main() {
	// Pointer flags for user customizability of readout
	amountPtr := flag.Int("amount", 5, "number of passwords to generate")
	lengthPtr := flag.String("length", "5", "length of the words to use")

	// Alternate shorthand pointer flags
	flag.IntVar(amountPtr, "a", 5, "number of passwords to generate (alternate)")
	flag.StringVar(lengthPtr, "l", "5", "length of the words to use (alternate)")

	// Help output (-h)
	flag.Usage = func() {
		fmt.Println("Usage: pwgen --amount [num] --length [num]")
		fmt.Println("  --amount, -a: number of passwords to generate (default 5)")
		fmt.Println("  --length, -l: length of the words to use (default 5, must be between 4 and 8)")
	}

	flag.Parse()

	// Convert length string to an integer and verify it is a valid number from 4-8
	length, err := strconv.Atoi(*lengthPtr)
	if err != nil || length < 4 || length > 8 {
		fmt.Println("Error: 'length' must be a number from 4 to 8")
		return
	}

	wordsList, err := loadWordsList(*lengthPtr)
	if err != nil {
		fmt.Println("Error: could not get words list")
		return
	}

	generateWordNumber(wordsList, *amountPtr)
}
