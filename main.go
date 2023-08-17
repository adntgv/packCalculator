package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var packSizsFlag = flag.String("packs", "250,500,1000,2000,5000", "Comma separated list of pack sizes")
var packSizes []int

func main() {
	flag.Parse()

	packSizesStr := strings.Split(*packSizsFlag, ",")
	for _, packSizeStr := range packSizesStr {
		packSize, err := strconv.Atoi(packSizeStr)
		if err != nil {
			panic(err)
		}
		packSizes = append(packSizes, packSize)
	}

	http.HandleFunc("/calculate", calculateHandler)
	http.ListenAndServe(":8080", nil)
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	itemsStr := r.URL.Query().Get("items")
	items, err := strconv.Atoi(itemsStr)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	packCount := calculatePacks(items, packSizes)

	// Display the result
	for packSize, count := range packCount {
		if count > 0 {
			fmt.Fprintf(w, "%d x %d\n", count, packSize)
		}
	}
}
