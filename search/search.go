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

	for line, index := range mapData {
		if strings.Contains(strings.ToLower(line), strings.ToLower(key)) {
			fmt.Printf("✅ Слово '%s' знайдено в рядках з індексами: %v\n", key, index)
		} else {
			fmt.Printf("❗️ ❗️ ❗️  Слово '%s' не знайдено в рядках з індексами: %v\n", key, index)
		}
	}

}
