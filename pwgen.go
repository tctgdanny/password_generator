package pwgen

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

//go:embed word_lists/*.json
var wordLists embed.FS

func LoadWords(wordLength string) ([]string, error) {
	// Load the word list file for the amount of letters designated by wordLength
	jsonFile, err := wordLists.Open(fmt.Sprintf("word_lists/%s_letter_words.json", wordLength))
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var wordsList []string
	json.Unmarshal(byteValue, &wordsList)
	return wordsList, nil
}

func GeneratePasswords(wordsList []string, count int) ([]string, error) {

	// Generate a random word from the loaded list, and a random number with 4 characters
	var passwords []string
	for i := 0; i < count; i++ {
		word := wordsList[rand.Intn(len(wordsList))]
		caser := cases.Title(language.Und)
		word = caser.String(word)
		number := rand.Intn(10000)
		pw := fmt.Sprintf("%s+%04d", word, number)
		passwords = append(passwords, pw)
	}

	return passwords, nil

}
