package main

import (
	indexmap "mapfinder/indexMap.go"
	"mapfinder/scan"
	"mapfinder/search"
)

func main() {
	search.Search()
	indexmap.IndexMap(scan.Scan())
}
