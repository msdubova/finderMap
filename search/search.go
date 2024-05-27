package search

import (
	"bufio"
	indexmap "mapfinder/indexMap.go"
	"mapfinder/scan"
	"strings"

	"fmt"
	"os"
)

func Search() {
	var key string

	fmt.Println("Який текст шукати?")

	var reader = bufio.NewReader(os.Stdin)
	key, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Помилка вводу пошуку:", err)
		return
	}

	key = strings.TrimSpace(key)
	fmt.Println("Ключ пошуку:", key)
	var mapData = indexmap.IndexMap(scan.Scan())
	var result []int
	for line, index := range mapData {
		if strings.Contains(strings.ToLower(line), strings.ToLower(key)) {
			// fmt.Printf("✅ Слово '%s' знайдено в рядках з індексами: %v\n", key, index)
			result = append(index, result...)
		}

		// else {
		// 	fmt.Printf("❗️ ❗️ ❗️  Слово '%s' не знайдено в рядках з індексами: %v\n", key, index)
		// }
	}
	if len(result) > 0 {
		fmt.Printf("Слово '%s' що ви шукаєте, знайдене в рядках: %v \n", key, result)
	} else {
		fmt.Printf("Слово '%s' не знайдено\n", key)
	}
}
