package main

import (
	"fmt"
	"math/rand"
)

const (
	totalFrames int = 10
	totalPins   int = 10
)

var totalRolls int = 2

// roll returns roll score - random integer between 0 and 10.
func roll() int {
	return rand.Intn(totalPins + 1)
}

func main() {
	totalScore := 0
	bonus := false
	for frameIndex := 0; frameIndex < totalFrames; frameIndex++ {
		if frameIndex == totalFrames {
			totalRolls++
		}
		frameScore := 0
		for rollIndex := 0; rollIndex < totalRolls; rollIndex++ {
			rollScore := roll()
			if bonus {
				totalScore = totalScore + rollScore
				bonus = false
			}
			fmt.Printf("rollScore: %d\n", rollScore)
			frameScore = frameScore + rollScore
			if frameScore >= totalPins {
				frameScore = totalPins
				bonus = true
				break
			}
		}
		fmt.Printf("frameScore: %d\n", frameScore)
		totalScore = totalScore + frameScore
		fmt.Printf("totalScore: %d\n", totalScore)
	}
	fmt.Printf("totalScore: %d\n", totalScore)
}
