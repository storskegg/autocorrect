package wordcount

import (
    "testing"
)

const (
    testDictionary = "../words_alpha.txt"
)

func Test_wordCount_MeanWordLength(t *testing.T) {
    tests := []struct {
        name    string
        path    string
        want    int
        wantErr bool
    }{
        {
            name:    "meanWordLength",
            path:    testDictionary,
            want:    9,
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            wc, err := NewWordCountFromDictionary(tt.path)
            if (err != nil) != tt.wantErr {
                t.Errorf("NewWordCountFromDictionary() error = %v, wantErr %v", err, tt.wantErr)
                return
            }

            if got := wc.MeanWordLength(); got != tt.want {
                t.Errorf("MeanWordLength() = %v, want %v", got, tt.want)
            }
        })
    }
}
