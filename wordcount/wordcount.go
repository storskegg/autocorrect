package wordcount

import (
    "bufio"
    "os"
    "path/filepath"
    "runtime"
    "strings"
    "sync"
)

// WordCount is a thread-safe mechanism that produces a wordcount for a body of
// text.
type WordCount interface {
    Add(string)
    Count(string) int
    All() map[string]int
    Length() int
    MeanWordLength() int
    Reset()
    Has(string) bool
}

type wordCount struct {
    sync.RWMutex

    charsCount int
    words      map[string]int
}

func NewWordCount() WordCount {
    return &wordCount{
        words: make(map[string]int),
    }
}

func NewWordCountFromDictionary(path string) (WordCount, error) {
    wc := NewWordCount()

    pathAbs, err := filepath.Abs(path)
    if err != nil {
        return nil, err
    }

    f, err := os.Open(pathAbs)
    scanner := bufio.NewScanner(f)
    scanner.Split(bufio.ScanWords)

    for scanner.Scan() {
        word := strings.TrimSpace(scanner.Text())
        if word == "" {
            continue
        }
        wc.Add(word)
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return wc, nil
}
func (wc *wordCount) Add(word string) {
    wc.Lock()
    defer wc.Unlock()

    lower := strings.ToLower(word)

    if _, exists := wc.words[lower]; exists {
        wc.words[lower]++
    } else {
        wc.charsCount += len(lower)
        wc.words[lower] = 1
    }
}

func (wc *wordCount) Count(word string) int {
    wc.RLock()
    defer wc.RUnlock()

    return wc.words[strings.ToLower(word)]
}

func (wc *wordCount) All() map[string]int {
    wc.RLock()
    defer wc.RUnlock()

    return wc.words
}

func (wc *wordCount) Length() int {
    wc.RLock()
    defer wc.RUnlock()

    return len(wc.words)
}

func (wc *wordCount) MeanWordLength() int {
    wc.RLock()
    defer wc.RUnlock()

    return wc.charsCount / wc.Length()
}

func (wc *wordCount) Reset() {
    wc.Lock()
    defer wc.Unlock()

    wc.words = make(map[string]int)
    runtime.GC() // Not normally done, but in this case let the runtime know it's time.
}

func (wc *wordCount) Has(word string) bool {
    wc.RLock()
    defer wc.RUnlock()

    return wc.words[strings.ToLower(word)] > 0
}
