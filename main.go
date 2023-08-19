package main

import (
	"encoding/json"
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

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/calculate", calculateHandler)
	http.HandleFunc("/sizes", getPackSizesHandler)
	http.HandleFunc("/sizes/set", setPackSizesHandler)
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	itemsStr := r.URL.Query().Get("items")
	items, err := strconv.Atoi(itemsStr)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	packCount := calculatePacks(items, packSizes)

	res, err := json.MarshalIndent(packCount, "", "  ")
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", res)
}

func getPackSizesHandler(w http.ResponseWriter, r *http.Request) {
	res, err := json.MarshalIndent(packSizes, "", "  ")
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", res)
}

func setPackSizesHandler(w http.ResponseWriter, r *http.Request) {
	packSizesStr := r.URL.Query().Get("sizes")
	packSizesStr = strings.Trim(packSizesStr, "[]")
	packSizesStrs := strings.Split(packSizesStr, ",")
	var newPackSizes []int

	for _, packSizeStr := range packSizesStrs {
		packSize, err := strconv.Atoi(packSizeStr)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		newPackSizes = append(newPackSizes, packSize)
	}

	packSizes = newPackSizes

	res, err := json.MarshalIndent(packSizes, "", "  ")
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", res)
}
