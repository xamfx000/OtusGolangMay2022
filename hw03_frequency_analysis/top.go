package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type WordFreqPair struct {
	Word         string
	TimesOccured int
}

func NewWordFreqPair(word string, timesOccured int) WordFreqPair {
	return WordFreqPair{Word: word, TimesOccured: timesOccured}
}

func Top10(s string) []string {
	if s == "" {
		return []string{}
	}
	words := strings.Fields(s)

	wordFreq := map[string]int{}
	for _, word := range words {
		if _, ok := wordFreq[word]; ok {
			wordFreq[word]++
			continue
		}
		wordFreq[word] = 1
	}
	wordFreqPairs := frequencyMapToStructList(wordFreq)

	sort.Slice(wordFreqPairs, func(i int, j int) bool {
		if wordFreqPairs[i].TimesOccured == wordFreqPairs[j].TimesOccured {
			return wordFreqPairs[i].Word < wordFreqPairs[j].Word
		}
		return wordFreqPairs[i].TimesOccured > wordFreqPairs[j].TimesOccured
	})
	return getTop10Words(wordFreqPairs)
}

func frequencyMapToStructList(frequencies map[string]int) []WordFreqPair {
	result := []WordFreqPair{}
	for word, freq := range frequencies {
		result = append(result, NewWordFreqPair(word, freq))
	}
	return result
}

func getTop10Words(words []WordFreqPair) []string {
	result := []string{}
	for _, wordFreqPair := range words {
		result = append(result, wordFreqPair.Word)
	}
	if len(result) < 10 {
		return result
	}
	return result[0:10]
}
