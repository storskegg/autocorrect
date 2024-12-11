package autocorrect

import (
    "strings"
)

const (
    letters = "abcdefghijklmnopqrstuvwxyz"
)

func CountWords(body string) WordCount {
    wc := NewWordCount()

    for _, word := range strings.Fields(body) {
        wc.Add(word)
    }

    return wc
}

func splitWord(word string) []rune {
    return []rune(strings.ToLower(word))
}
