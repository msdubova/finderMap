package indexmap

import "strings"

func IndexMap(slice []string) map[string][]int {
	var indexedMap = make(map[string][]int)

	for index, item := range slice {
		lines := strings.Fields(item)

		for _, line := range lines {
			line = strings.ToLower(line)
			indexedMap[line] = append(indexedMap[line], index+1)
		}

	}

	return indexedMap

}
