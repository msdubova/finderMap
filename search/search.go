package search

import (
	"bufio"
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
	// // var text = scan.Scan()
	// // for _, line := range text {
	// // 	if strings.Contains(strings.ToLower(line), strings.ToLower(key)) {
	// // 		fmt.Printf("✅ Строка %v містить ключ %v\n", line, key)
	// // 	} else {
	// // 		fmt.Printf("❗️ ❗️ ❗️ Строка %v не містить ключ %v\n", line, key)
	// // 	}
	// // }
}
