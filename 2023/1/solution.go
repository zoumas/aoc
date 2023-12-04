package solution

import (
	"log"
	"strconv"
	"strings"
)

func Solution1(document string) int {
	digits := map[string]struct{}{
		"1": {},
		"2": {},
		"3": {},
		"4": {},
		"5": {},
		"6": {},
		"7": {},
		"8": {},
		"9": {},
	}

	var sum int
	for _, line := range strings.Split(document, "\n") {
		if line == "" {
			continue
		}

		var first, last string

		for i := 0; i < len(line); i++ {
			c := string(line[i])
			if _, ok := digits[c]; ok {
				first = c
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			c := string(line[i])
			if _, ok := digits[c]; ok {
				last = c
				break
			}
		}

		log.Printf("line: %s\tvalue: %s\n", line, first+last)
		calibrationValue, err := strconv.Atoi(first + last)
		if err != nil {
			log.Fatal(err)
		}
		sum += calibrationValue
	}
	return sum
}

// An idea that didn't work out
// func Solution2(document string) int {
// 	replacer := strings.NewReplacer(
// 		"one", "1",
// 		"two", "2",
// 		"three", "3",
// 		"four", "4",
// 		"five", "5",
// 		"six", "6",
// 		"seven", "7",
// 		"eight", "8",
// 		"nine", "9",
// 	)
// 	return Solution1(replacer.Replace(document))
// }

func Solution2(document string) int {
	match := map[string]string{
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	var sum int
	for _, line := range strings.Split(document, "\n") {
		if line == "" || line == "\n" {
			continue
		}

		var first, second string

		for i := 0; i < len(line); i++ {
			s := line[i:]
			for w := range match {
				if strings.HasPrefix(s, w) {
					first = w
					break
				}
			}
		}

		for i := len(line); i >= 0; i-- {
			s := line[i:]
			for w := range match {
				if strings.HasPrefix(s, w) {
					second = w
					break
				}
			}
		}

		log.Println(match[first] + match[second])
		calibrationValue, err := strconv.Atoi(match[second] + match[first])
		if err != nil {
			log.Println(err)
		}
		sum += calibrationValue
	}
	return sum
}
