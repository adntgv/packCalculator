package main

func calculatePacks(items int, packSizes []int) map[int]int {
	packCount := make(map[int]int)

	if items <= 0 {
		return packCount
	}

	packSizes = cleanPackSizes(packSizes)

	if len(packSizes) == 0 {
		return packCount
	}

	for i := len(packSizes) - 1; i > 0; i-- { // Changed the loop condition
		switch {
		case packSizes[i] <= 0:
			continue
		case items > packSizes[i]:
			packCount[packSizes[i]] = items / packSizes[i]
			items = items % packSizes[i]
		case items == packSizes[i]:
			packCount[packSizes[i]] = 1
			items = 0
		}
	}

	if items > packSizes[0] {
		if len(packSizes) > 1 {
			packCount[packSizes[1]] += 1
		} else {
			packCount[packSizes[0]] = items / packSizes[0]
		}
	} else if items > 0 {
		packCount[packSizes[0]] = 1
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
