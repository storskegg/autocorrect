package autocorrect

import (
    "github.com/stretchr/testify/assert"
    "sort"
    "testing"
)

const testLetters = "zef"

func helperDedupeAndSort(words []string) []string {
    holdings := make(map[string]struct{})
    for _, word := range words {
        holdings[word] = struct{}{}
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

func Test_editDistance1(t *testing.T) {
    type args struct {
        word    string
        letters string
    }
    tests := []struct {
        name string
        args args
        want []string
    }{
        {
            name: "Happy Path: a",
            args: args{
                word:    "a",
                letters: testLetters,
            },
            want: helperDedupeAndSort([]string{
                // Additions
                "za", "ea", "fa",
                "az", "ae", "af",
                // Subtractions
                // -- none: too short
                // Transpositions
                // -- none: too short
                // Substitutions
                // -- none: too short
            }),
        },
        {
            name: "Happy Path: be",
            args: args{
                word:    "be",
                letters: testLetters,
            },
            want: helperDedupeAndSort([]string{
                // Additions
                "zbe", "ebe", "fbe",
                "bze", "bee", "bfe",
                "bez", "bee", "bef",
                // Subtractions
                "e", "b",
                // Transpositions
                "eb",
                // Substitutions
                "ze", "ee", "fe",
                "bz", "bf",
            }),
        },
        {
            name: "Happy Path: cat",
            args: args{
                word:    "cat",
                letters: testLetters,
            },
            want: helperDedupeAndSort([]string{
                // Additions
                "zcat", "ecat", "fcat",
                "czat", "ceat", "cfat",
                "cazt", "caet", "caft",
                "catz", "cate", "catf",
                // Subtractions
                "at", "ct", "ca",
                // Transpositions
                "act", "cta",
                // Substitutions
                "zat", "eat", "fat",
                "czt", "cet", "cft",
                "caz", "cae", "caf",
            }),
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := editDistance1(tt.args.word, tt.args.letters)
            assert.Equal(t, tt.want, got)
        })
    }
}
