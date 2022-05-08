package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"
)

var (
	//go:embed words.txt
	wordFile []byte
	words    []string
)

func main() {

	wordFileLines := bytes.Split(wordFile, []byte{'\n'})
	for _, w := range wordFileLines {
		word := strings.TrimSpace(string(w))
		if word != "" {
			words = append(words, word)
		}

	}
	fmt.Println(len(words))
}
