package main

import (
	"fmt"
	"sort"
	"strings"
)

func getTopWords(wordsMap map[string]int, n int) []string {
	numSlice := make([]int, 0, len(wordsMap))
	for _, val := range wordsMap {
		numSlice = append(numSlice, val)
	}
	sort.Ints(numSlice)
	for i, j := 0, len(numSlice)-1; i < j; i, j = i+1, j-1 {
		numSlice[i], numSlice[j] = numSlice[j], numSlice[i]
	}

	topSlice := make([]string, 0, n)
	for i := range n {
		for key := range wordsMap {
			if wordsMap[key] == numSlice[i] {
				topSlice = append(topSlice, key)
				break
			}
		}
	}
	return topSlice
}

func AnalyzeText(text string) {
	wordsMap := make(map[string]int)
	maxVal := -99999
	var maxKey string
	var kUniq = 0
	var textWithOnlySpaces = text

	for _, sep := range []string{".", ",", "!", "?"} {
		textWithOnlySpaces = strings.Replace(textWithOnlySpaces, sep, " ", -1)
	}

	words := strings.Fields(textWithOnlySpaces)
	for _, word := range words {
		word = strings.ToLower(word)
		if _, ok := wordsMap[word]; !ok {
			wordsMap[word] = 1
			kUniq += 1
		} else {
			wordsMap[word]++
		}
		if currVal := wordsMap[word]; currVal > maxVal {
			maxVal = currVal
			maxKey = word
		}
	}

	top5words := getTopWords(wordsMap, 5)

	fmt.Printf("Количество слов: %d\nКоличество уникальных слов: %d\nСамое часто встречающееся слово: \"%s\" (встречается %d раз)\nТоп-5 самых часто встречающихся слов:\n\"%s\": %d раз\n\"%s\": %d раз\n\"%s\": %d раз\n\"%s\": %d раз\n\"%s\": %d раз\n", len(words), kUniq, maxKey, maxVal, top5words[0], wordsMap[top5words[0]], top5words[1], wordsMap[top5words[1]], top5words[2], wordsMap[top5words[2]], top5words[3], wordsMap[top5words[3]], top5words[4], wordsMap[top5words[4]])

}
