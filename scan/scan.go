package scan

import (
	"bufio"
	"fmt"
	"os"
)

func Scan() []string {
	var file, error = os.Open("File1.txt")
	if error != nil {
		fmt.Println("Помилка відкриття файлу")
	}

	defer file.Close()

	var text = bufio.NewScanner(file)
	var Slice []string
	for text.Scan() {
		var line = text.Text()
		Slice = append(Slice, line)
	}
	// fmt.Println("Довжина зрізу:", len(Slice))

	return Slice
}
