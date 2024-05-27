package indexmap

import "fmt"

func IndexMap(slice []string) map[string][]int {
	var indexedMap = make(map[string][]int)

	for index, item := range slice {

		indexedMap[item] = append(indexedMap[item], index)

	}

	fmt.Println(indexedMap)
	return indexedMap

}
