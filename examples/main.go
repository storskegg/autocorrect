package main

import (
    "bytes"
    _ "embed"
    "github.com/storskegg/autocorrect"
    "github.com/storskegg/autocorrect/wordcount"
    "log"
)

//go:embed words_alpha.txt
var internalDist []byte

func main() {
    wordGud := "facilitate"
    wordBad := "facillittate"

    buf := bytes.NewBuffer(internalDist)
    wc, err := wordcount.NewFromReader(buf)
    if err != nil {
        panic(err)
    }

    {
        word, ok := autocorrect.Correct(wordGud, wc, autocorrect.LettersEnglish)
        if ok {
            log.Printf("[GOOD] Correctly spelled word: %s", word)
        } else {
            log.Printf("[BAD] Misspelled word '%s' was corrected to '%s'", wordGud, word)
        }
    }

    {
        word, ok := autocorrect.Correct(wordBad, wc, autocorrect.LettersEnglish)
        if ok {
            log.Printf("[GOOD] Correctly spelled word: %s", word)
        } else {
            log.Printf("[BAD] Misspelled word '%s' was corrected to '%s'", wordBad, word)
        }
    }
}
