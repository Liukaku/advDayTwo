package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fileOpen, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("file err", err.Error())
	}

	fileRead, err := ioutil.ReadAll(fileOpen)

	matches := string(fileRead)
	matchSlice := strings.Split(matches, "\n")

	totalScore := 0
	for _, match := range matchSlice {
		hands := strings.Split(match, " ")
		if len(hands) == 2 {
			// fmt.Println(matchCalc(hands))
			totalScore += matchCalc(hands)
		}

	}
	fmt.Println("score : ", totalScore)
}

// A Rock 1
// B Paper 2
// C Scissors 3

// X Lose 0
// Y Draw 3
// Z Win 6

func matchCalc(match []string)int{
	switch match[0] {
	case "A": return calcScore(match[1], [3]int{3, 4, 8})
	case "B": return calcScore(match[1], [3]int{1, 5, 9})
	case "C": return calcScore(match[1], [3]int{2, 6, 7})
	default: return 0
	}
}

func calcScore(hand string, scores [3]int)int{
	switch hand {
		case "X": {
			return scores[0]
		}
		case "Y": {
			return scores[1]
		}
		case "Z": {
			return scores[2]
		}
	}
	return 0
}