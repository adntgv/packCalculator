package main

import (
	"golang.org/x/exp/slices"
)

func calculatePacks(items int, packSizes []int) map[int]int {
	packCount := make(map[int]int)

	if items <= 0 {
		return packCount
	}

	packSizes = cleanPackSizes(packSizes)

	if len(packSizes) == 0 {
		return packCount
	}

	slices.Sort(packSizes)

	return calculatePacksRecursive(items, packSizes, packCount)
}

func cleanPackSizes(packSizes []int) []int {
	var cleanedPackSizes []int

	for _, packSize := range packSizes {
		if packSize > 0 {
			cleanedPackSizes = append(cleanedPackSizes, packSize)
		}
	}

	return cleanedPackSizes
}

func calculatePacksRecursive(items int, packSizes []int, packCount map[int]int) map[int]int {
	for i := len(packSizes) - 1; i > 0; i-- { // Changed the loop condition
		packSize := packSizes[i]
		if items >= packSize {
			packCount[packSize] = items / packSize
			items %= packSize
		} else {
			res := calculatePacks(items, packSizes[:i])
			oversell := packSize - items
			resSell := 0

			for k, v := range res {
				resSell += k * v
			}

			resOversell := resSell - items

			if resOversell >= oversell {
				packCount[packSize] = 1
				items = 0
				return packCount
			}
		}
	}

	if items > 0 {
		packSize := packSizes[0]
		count := items / packSize
		mod := items % packSize

		if mod > 0 {
			count++
		}

		if _, ok := packCount[packSize]; !ok {
			packCount[packSize] = count
		} else {
			packCount[packSize] += count
		}
	}

	return packCount
}
