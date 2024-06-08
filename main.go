package main

import (
	"mapfinder/scan"
	"mapfinder/search"
)

func Main() {
	lines := scan.Scan()
	search.Search(lines)
}
