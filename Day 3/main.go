package main

import (
	"fmt"
	"strings"
)

func compartmentalize(rucksacks string) ([]string, []string) {
	var (
		compartment1List []string
		compartment2List []string
	)

	fullSize := len(rucksacks)
	compartmentSize := len(rucksacks) / 2

	compartment1 := rucksacks[:compartmentSize]
	compartment2 := rucksacks[compartmentSize:fullSize]

	compartment1List = append(compartment1List, compartment1)
	compartment2List = append(compartment2List, compartment2)

	return compartment1List, compartment2List
}

func sharedItem(compartment1 []string, comparement2 []string) string {
	var commonItems []string

	compartment1Items := strings.Split(compartment1[0], "")
	compartment2Items := strings.Split(comparement2[0], "")

	// Making the assumption that the rucksacks always has an even amount of items
	for i := 0; i < len(compartment1Items); i++ {
		for j := 0; j < len(compartment2Items); j++ {
			if compartment1Items[i] == compartment2Items[j] {
				commonItems = append(commonItems, compartment1Items[i])
				break
			}
		}
	}
	// Some compartments have 3 or more of the same item so going to lazily deal with that
	priorityItems := commonItems[0]

	return priorityItems
}

// Part 2
func elfGroups(rucksack []string) []string {
	var (
		groupLetter []string
	)
	for i, _ := range rucksack {
		j := i + 1
		if j%3 == 0 {
			groupOfElves := rucksack[j-3 : j]
			thirdElf := strings.Split(groupOfElves[(j-2)%3], "")
			secondElf := strings.Split(groupOfElves[(j-1)%3], "")
			firstElf := strings.Split(groupOfElves[j%3], "")
			badge := isBadge(firstElf, secondElf, thirdElf)
			groupLetter = append(groupLetter, badge)
		}
	}
	return groupLetter
}

func isBadge(elf1 []string, elf2 []string, elf3 []string) string {
	var badge string
	//counter := 0
	for i := 0; i < len(elf1); i++ {
		for j := 0; j < len(elf2); j++ {
			for n := 0; n < len(elf3); n++ {
				switch elf1[i] {
				case elf2[j]:
					switch elf2[j] {
					case elf3[n]:
						badge = elf3[n]
					}
				}
			}
		}
	}
	return badge
}

// Assign priority value to each item
func prioritize(letter string) int {
	var (
		priorityScore int
	)
	asciiValues := []rune(letter)
	for _, asciiValue := range asciiValues {
		// asciiValue - 96 should give us the priority for each lowercase letter
		if int(asciiValue) >= 97 {
			priorityScore = int(asciiValue) - 96
			// Uppercase equivalent
		} else if int(asciiValue) >= 65 && int(asciiValue) < 96 {
			priorityScore = int(asciiValue) - 38
		}
	}
	return priorityScore
}

func totalScore(priorityValues []int) int {
	var finalScore int
	for _, values := range priorityValues {
		finalScore += values
	}
	return finalScore
}

func main() {
	var (
		priorityScores1 []int
		priorityScores2 []int
	)

	inputs, err := ReadInput()
	if err != nil {
		panic(err)
	}

	for _, rucksacks := range inputs {
		if rucksacks == "" {
			break
		}
		// Part 1
		compartment1, compartment2 := compartmentalize(rucksacks)
		commonItems := sharedItem(compartment1, compartment2)
		priorityScores1 = append(priorityScores1, prioritize(commonItems))

	}
	// Part 2
	badges := elfGroups(inputs)
	for _, badge := range badges {
		priorityScores2 = append(priorityScores2, prioritize(badge))
	}
	fmt.Println(totalScore(priorityScores1))
	fmt.Println(totalScore(priorityScores2))
}
