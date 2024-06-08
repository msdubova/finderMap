package indexmap

import "strings"

func IndexMap(slice []string) map[string][]int {
	var indexedMap = make(map[string][]int)

	for index, line := range slice {
		words := strings.Fields(line)

		for _, line := range words {
			line = strings.ToLower(line)
			indexedMap[line] = append(indexedMap[line], index+1)
		}

	}

	return indexedMap

}
