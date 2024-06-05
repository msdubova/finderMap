package main

import (
	"mapfinder/scan"
	"mapfinder/search"
)

func main() {
	key := scan.Scan()
	search.Search(key)
}
