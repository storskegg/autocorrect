package autocorrect

import (
    "github.com/storskegg/autocorrect/wordcount"
    "sort"
    "strings"
)

const (
    LettersEnglish = "abcdefghijklmnopqrstuvwxyz"
)

func CountWords(body string) wordcount.WordCount {
    wc := wordcount.New()

    for _, word := range strings.Fields(body) {
        wc.Add(word)
    }

    return wc
}

func splitWord(word string) []rune {
    return []rune(strings.ToLower(word))
}

//for(let i = 0; i < word.length; i++){
//    for(let j = 0; j < alphabet.length; j++){
//      let newWord = word.slice();
//      newWord[i] = alphabet[j];
//      results.push(newWord.join(''));
//    }
//  }

func editDistance1(word string, letters string) []string {
    split := splitWord(word)
    holdings := make(map[string]struct{})

    // Additions
    for i := 0; i <= len(split); i++ {
        for j := 0; j < len(letters); j++ {
            newWord := make([]rune, len(split)+1)
            copy(newWord[:i], split[:i])
            newWord[i] = rune(letters[j])
            copy(newWord[i+1:], split[i:])

            holdings[string(newWord)] = struct{}{}
        }
    }

    // Subtractions
    if len(split) > 1 {
        for i := 0; i < len(split); i++ {
            newWord := make([]rune, len(split)-1)
            copy(newWord[:i], split[:i])
            if i+1 < len(split) {
                copy(newWord[i:], split[i+1:])
            }

            holdings[string(newWord)] = struct{}{}
        }
    }

    // Transpositions
    if len(split) > 1 {
        for i := 0; i < len(split)-1; i++ {
            newWord := make([]rune, len(split))
            copy(newWord[:i], split[:i])
            newWord[i], newWord[i+1] = split[i+1], split[i]
            copy(newWord[i+2:], split[i+2:])

            holdings[string(newWord)] = struct{}{}
        }
    }

    // Substitutions
    if len(split) > 1 {
        for i := 0; i < len(split); i++ {
            for j := 0; j < len(letters); j++ {
                if split[i] != rune(letters[j]) {
                    newWord := make([]rune, len(split))
                    copy(newWord, split)
                    newWord[i] = rune(letters[j])

                    holdings[string(newWord)] = struct{}{}
                }
            }
        }
    }

    results := make([]string, len(holdings))
    i := 0
    for result := range holdings {
        results[i] = result
        i++
    }
    sort.Strings(results)

    return results
}

func Correct(word string, dict wordcount.WordCount, letters string) (string, bool) {
    if dict.Has(word) {
        return word, true
    }

    maxCount := 0
    correctWord := word
    distance1Words := editDistance1(word, letters)
    distance2Words := make([]string, 0)

    for _, word := range distance1Words {
        distance2Words = append(distance2Words, editDistance1(word, letters)...)
    }

    for _, word := range distance1Words {
        if dict.Has(word) {
            if dict.Count(word) > maxCount {
                maxCount = dict.Count(word)
                correctWord = word
            }
        }
    }

    maxCount2 := 0
    correctWord2 := correctWord

    for _, word := range distance2Words {
        if dict.Has(word) {
            if dict.Count(word) > maxCount2 {
                maxCount2 = dict.Count(word)
                correctWord2 = word
            }
        }
    }

    if len(word) > dict.MeanWordLength()*2/3 {
        if maxCount2 > 100*maxCount {
            return correctWord2, false
        }
        return correctWord, false
    }

    if maxCount2 > 4*maxCount {
        return correctWord2, false
    }
    return correctWord, false
}
