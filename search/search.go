package search

import (
	"bufio"
	indexmap "mapfinder/indexMap.go"
	"strings"

	"fmt"
	"os"
)

func Search(lines []string) {
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
	index := indexmap.IndexMap(lines)
	result := index.Search(key)

	if len(result) > 0 {
		fmt.Printf("Слово '%s' що ви шукаєте, знайдене в рядках: %v \n", key, result)
	} else {
		fmt.Printf("Слово '%s' не знайдено\n", key)
	}

}
