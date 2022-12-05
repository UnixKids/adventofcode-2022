package main

import (
	"fmt"
	"strconv"
	"strings"
)

func elfSections(elfRange []string) ([]int, []int) {
	var (
		elfMin    int
		elfMax    int
		elfSlice1 []int
		elfSlice2 []int
	)

	for i, _ := range elfRange {
		elfRangeSplit := strings.Split(elfRange[i], "-")
		elfMin, _ = strconv.Atoi(elfRangeSplit[0])
		elfMax, _ = strconv.Atoi(elfRangeSplit[1])
		elfSlice2 = []int{elfMin, elfMax}
	}
	// Ran into an annoying issue with elfMax and elfMin overwriting each other so added a break and another slice to fix this
	for i, _ := range elfRange {
		elfRangeSplit := strings.Split(elfRange[i], "-")
		elfMin, _ = strconv.Atoi(elfRangeSplit[0])
		elfMax, _ = strconv.Atoi(elfRangeSplit[1])
		elfSlice1 = []int{elfMin, elfMax}
		break
	}
	return elfSlice1, elfSlice2
}

// Part 1
func compareSection(sections [][]int) int {
	var counter int
	for i := 0; i < len(sections); i++ {
		if i+1 < len(sections) {
			currentMin, currentMax := sections[i][0], sections[i][1]
			nextMin, nextMax := sections[i+1][0], sections[i+1][1]
			if currentMin <= nextMin && currentMax >= nextMax {
				counter++
			} else if nextMin <= currentMin && nextMax >= currentMax {
				counter++
			}

		}
		// Skips to the next pair of inputs
		i++
	}
	return counter
}

// Part 2
func overlaps(sections [][]int) int {
	var counter int
	for i := 0; i < len(sections); i++ {
		if i+1 < len(sections) {
			currentMin, currentMax := sections[i][0], sections[i][1]
			nextMin, nextMax := sections[i+1][0], sections[i+1][1]

			if currentMin <= nextMax && currentMax >= nextMax {
				counter++
			} else if nextMax <= currentMin && nextMax >= currentMin {
				counter++
			} else if nextMin <= currentMax && nextMax >= currentMin {
				counter++
			}
		}
		i++
	}
	return counter
}

func main() {
	var (
		elfSectionList [][]int
	)
	inputs, err := ReadInput()
	if err != nil {
		panic(err)
	}
	for _, elfRanges := range inputs {
		if elfRanges == "" {
			break
		}
		elfRange := strings.Split(elfRanges, ",")
		elfSectionCovered2, elfSectionCovered1 := elfSections(elfRange)
		elfSectionList = append(elfSectionList, elfSectionCovered2)
		elfSectionList = append(elfSectionList, elfSectionCovered1)
	}
	fmt.Println(compareSection(elfSectionList))
	fmt.Println(overlaps(elfSectionList))

}
