package solution

import (
	"log"
	"strconv"
	"strings"
)

//
// func Solution(body string) int {
// 	lines := strings.Split(body, "\n")
//
// 	var sumPoints int
//
// 	for _, line := range lines {
// 		cardAndNumbers := strings.Split(line, ":")
// 		card := cardAndNumbers[0]
//
// 		numbers := cardAndNumbers[1]
// 		allNumbers := strings.Split(numbers, "|")
// 		winning := strings.Fields(strings.TrimSpace(allNumbers[0]))
// 		gotNums := strings.Fields(strings.TrimSpace(allNumbers[1]))
//
// 		log.Printf("%v %v\n", winning, gotNums)
//
// 		var points int
// 		for _, n := range gotNums {
// 			for _, winNum := range winning {
// 				if n != winNum {
// 					continue
// 				}
//
// 				if points == 0 {
// 					points++
// 				} else {
// 					points *= 2
// 				}
// 			}
// 		}
// 		log.Println(card, points)
// 		sumPoints += points
// 	}
//
// 	return sumPoints
// }

func Solution2(body string) int {
	lines := strings.Split(body, "\n")

	orig := make(map[int]int)
	copies := make(map[int]int)

	var scratchcards int

	for _, line := range lines {
		var matches int
		cardAndNumbers := strings.Split(line, ":")
		card := cardAndNumbers[0]
		cardNum, err := strconv.Atoi(strings.Fields(card)[1])
		if err != nil {
			log.Fatal(err)
		}

		numbers := cardAndNumbers[1]
		allNumbers := strings.Split(numbers, "|")
		winning := strings.Fields(strings.TrimSpace(allNumbers[0]))
		gotNums := strings.Fields(strings.TrimSpace(allNumbers[1]))

		log.Printf("%v %v\n", winning, gotNums)

		for _, n := range gotNums {
			for _, winNum := range winning {
				if winNum == n {
					matches++
				}
			}
		}

		log.Println(card, matches)
		orig[cardNum] = matches
		for i := 1; i <= matches; i++ {
			copies[cardNum+i]++
			copies[cardNum+i] += copies[cardNum]
		}
		log.Println(orig, copies)
	}

	for k := range orig {
		scratchcards++
		scratchcards += copies[k]
	}

	return scratchcards
}
