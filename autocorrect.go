package autocorrect

import (
    "github.com/storskegg/autocorrect/wordcount"
    "strings"
)

const (
    letters = "abcdefghijklmnopqrstuvwxyz"
)

func CountWords(body string) wordcount.WordCount {
    wc := wordcount.NewWordCount()

    for _, word := range strings.Fields(body) {
        wc.Add(word)
    }

    return wc
}

func splitWord(word string) []rune {
    return []rune(strings.ToLower(word))
}
