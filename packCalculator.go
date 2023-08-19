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

	for i := len(packSizes) - 1; i > 0; i-- { // Changed the loop condition
		switch {
		case items > packSizes[i]:
			packCount[packSizes[i]] = items / packSizes[i]
			items = items % packSizes[i]
		case items == packSizes[i]:
			packCount[packSizes[i]] = 1
			items = 0
			return packCount
		case items < packSizes[i]:
			res := calculatePacks(items, packSizes[:i])
			oversell := packSizes[i] - items
			resSell := 0

			for k, v := range res {
				resSell += k * v
			}

			resOversell := resSell - items

			if resOversell >= oversell {
				packCount[packSizes[i]] = 1
				items = 0
				return packCount
			} else {
				continue
			}
		}
	}

	if items > 0 {
		count := items / packSizes[0]
		mod := items % packSizes[0]

		if mod > 0 {
			count++
		}

		packCount[packSizes[0]] = count
	}

	return packCount
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
