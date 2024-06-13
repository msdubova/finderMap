package indexmap

import "strings"

type Index struct {
	indexedMap map[string][]int
}

func IndexMap(lines []string) Index {
	var indexedMap = make(map[string][]int)

	for lineNumber, lineContent := range lines {
		words := strings.Fields(lineContent)

		for _, word := range words {
			word = strings.ToLower(word)
			indexedMap[word] = append(indexedMap[word], lineNumber+1)
		}

	}

	return Index{indexedMap: indexedMap}

}

func (index Index) Search(key string) []int {
	key = strings.ToLower(key)
	var result []int
	for word, indexes := range index.indexedMap {
		if strings.Contains(word, key) {
			result = append(result, indexes...)
		}
	}
	return result
}
