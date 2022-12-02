package main

import (
	"fmt"
)

const (
	rock     = 1
	paper    = 2
	scissors = 3
	win      = 6
	draw     = 3
)

func roundsPlayed(round []string) ([]int, []int) {
	var totalScorePart1 []int
	var totalScorePart2 []int

	for _, choice := range round {
		roundScorePtrPart1 := &totalScorePart1
		roundScorePtrPart2 := &totalScorePart2

		if choice == "" {
			break
		}
		opponent := string(choice[0])
		player := string(choice[2])

		roundPointPart1 := playScorePart1(opponent, player)
		*roundScorePtrPart1 = append(*roundScorePtrPart1, roundPointPart1)

		roundPointPart2 := playScorePart2(opponent, player)
		*roundScorePtrPart2 = append(*roundScorePtrPart2, roundPointPart2)
	}
	return totalScorePart1, totalScorePart2
}

func playScorePart1(opponentChoice string, playerChoice string) int {
	var (
		score int
	)

	switch opponentChoice {
	case "A":
		switch playerChoice {
		case "X":
			score = rock + draw
		case "Y":
			score = paper + win
		case "Z":
			score += scissors

		}
	case "B":
		switch playerChoice {
		case "X":
			score += rock
		case "Y":
			score = paper + draw
		case "Z":
			score = scissors + win
		}
	case "C":
		switch playerChoice {
		case "X":
			score = rock + win
		case "Y":
			score += paper
		case "Z":
			score = scissors + draw
		}
	}
	return score
}

func playScorePart2(opponentChoice string, matchDecision string) int {
	var (
		score int
	)
	switch matchDecision {
	// Lose
	case "X":
		switch opponentChoice {
		case "A":
			score += scissors
		case "B":
			score += rock
		case "C":
			score += paper
		}
	// Draw
	case "Y":
		switch opponentChoice {
		case "A":
			score = rock + draw
		case "B":
			score = paper + draw
		case "C":
			score = scissors + draw
		}
	// Win
	case "Z":
		switch opponentChoice {
		case "A":
			score = paper + win
		case "B":
			score = scissors + win
		case "C":
			score = rock + win
		}

	}

	return score
}

func sumPoints(pointsList []int) int {
	var totalPoints int
	for _, point := range pointsList {
		totalPoints += point
	}
	return totalPoints
}

func main() {
	inputs, err := ReadInput()
	if err != nil {
		fmt.Println(err)
	}

	roundPointPart1, roundPointPart2 := roundsPlayed(inputs)
	totalPointsPart1 := sumPoints(roundPointPart1)
	fmt.Printf("The total points from part #1: %d\n", totalPointsPart1)
	totalPointsPart2 := sumPoints(roundPointPart2)
	fmt.Printf("The total points from part #2: %d\n", totalPointsPart2)

}
