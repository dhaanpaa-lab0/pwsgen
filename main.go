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
	words    []string
)

func TitleCaseWord(s string) string {
	caser := cases.Title(language.English)
	return caser.String(s)
}

func main() {

	wordFileLines := bytes.Split(wordFile, []byte{'\n'})
	for _, w := range wordFileLines {
		word := strings.TrimSpace(string(w))
		if word != "" {
			words = append(words, word)
		}

	}

	// Post Process
	wordFileCount := len(words)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	randomWord := TitleCaseWord(words[r1.Intn(wordFileCount)])
	fmt.Println(randomWord)
}
