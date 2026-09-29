package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(original string) []string {
	result := make([]string, 0, 10)

	words := strings.Fields(original)

	mappedWords := make(map[string]int)

	for _, word := range words {
		mappedWords[word]++
	}

	type keyValue struct {
		Key   string
		Value int
	}

	var forSorting = make([]keyValue, 0, len(mappedWords))

	for word, count := range mappedWords {
		forSorting = append(forSorting, keyValue{word, count})
	}

	sort.Slice(forSorting, func(i, j int) bool {
		if forSorting[i].Value == forSorting[j].Value {
			return forSorting[i].Key < forSorting[j].Key
		}
		return forSorting[i].Value > forSorting[j].Value
	})

	for i := 0; i < len(forSorting) && i < 10; i++ {
		result = append(result, forSorting[i].Key)
	}

	// Place your code here.
	return result
}
