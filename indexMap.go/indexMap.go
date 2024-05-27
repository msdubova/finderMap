package indexmap

func IndexMap(slice []string) map[string][]int {
	var indexedMap = make(map[string][]int)

	for index, item := range slice {

		indexedMap[item] = append(indexedMap[item], index)

	}

	return indexedMap

}
