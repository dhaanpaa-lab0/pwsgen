package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"math/rand"
	"strings"
	"time"
)

var (
	//go:embed words.txt
	wordFile []byte
)

func TitleCaseWord(s string) string {
	caser := cases.Title(language.English)
	return caser.String(s)
}

func GetRandomWord(wordsList []string) string {
	wordFileCount := len(wordsList)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	randomWord := TitleCaseWord(wordsList[r1.Intn(wordFileCount)])
	return randomWord
}

func GetWordsList() []string {
	var words []string
	wordFileLines := bytes.Split(wordFile, []byte{'\n'})
	for _, w := range wordFileLines {
		word := strings.TrimSpace(string(w))
		if word != "" {
			words = append(words, word)
		}

	}
	return words
}

func main() {

	wordList := GetWordsList()
	// Post Process

	fmt.Println(GetRandomWord(wordList))
}
