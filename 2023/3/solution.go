package solution

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var partNumberRegex = regexp.MustCompile("[0-9]+")

func Solution(body string) int {
	var sum int

	matrix := strings.Split(body, "\n")

	N := len(matrix)
	for i, line := range matrix {
		M := len(line)

		lineInBytes := []byte(line)
		loc := partNumberRegex.FindAllIndex(lineInBytes, -1)
		if loc == nil {
			// no match
			continue
		}

		for _, match := range loc {
			startIdx := match[0]
			endIdx := match[1]

			valid := false

			for j := startIdx; j < endIdx; j++ {
				left := j - 1
				right := j + 1
				bottom := i + 1
				top := i - 1

				if left >= 0 {
					elem := matrix[i][left]
					if isPartNumber(rune(elem)) {
						// log.Println(elem)
						valid = true
						break
					}
				}

				if right < M {
					elem := matrix[i][right]
					if isPartNumber(rune(elem)) {
						// log.Println(elem)
						valid = true
						break
					}
				}

				if bottom < N {
					elem := matrix[bottom][j]
					if isPartNumber(rune(elem)) {
						// log.Println(elem)
						valid = true
						break
					}
				}

				if top >= 0 {
					elem := matrix[top][j]
					if isPartNumber(rune(elem)) {
						// log.Println(elem)
						valid = true
						break
					}
				}

				if top >= 0 && left >= 0 {
					elem := matrix[top][left]
					if isPartNumber(rune(elem)) {
						// log.Println(elem)
						valid = true
						break
					}
				}

				if top >= 0 && right < M {
					elem := matrix[top][right]
					if isPartNumber(rune(elem)) {
						// log.Println(elem)
						valid = true
						break
					}
				}

				if bottom < N && left >= 0 {
					elem := matrix[bottom][left]
					if isPartNumber(rune(elem)) {
						// log.Println(elem)
						valid = true
						break
					}
				}

				if bottom < N && right < M {
					elem := matrix[bottom][right]
					if isPartNumber(rune(elem)) {
						// log.Println(elem)
						valid = true
						break
					}
				}
			}

			if valid {
				num, err := strconv.Atoi(string(lineInBytes[startIdx:endIdx]))
				if err != nil {
					log.Fatal(err)
				}
				// log.Println(num)
				sum += num
			}
		}
	}

	return sum
}

func isPartNumber(r rune) bool {
	return r != '.' && !unicode.IsDigit(r)
}

func Solution2(body string) int {
	var sum int

	matrix := strings.Split(body, "\n")

	type PartNumber struct {
		Value    int
		StartIdx int
		EndIdx   int
		N        int
	}
	partNumbers := []PartNumber{}

	N := len(matrix)
	M := len(matrix[0])
	for i, line := range matrix {

		lineInBytes := []byte(line)
		loc := partNumberRegex.FindAllIndex(lineInBytes, -1)
		if loc == nil {
			// no match
			continue
		}

		for _, match := range loc {
			startIdx := match[0]
			endIdx := match[1]

			valid := false

			for j := startIdx; j < endIdx; j++ {
				left := j - 1
				right := j + 1
				bottom := i + 1
				top := i - 1

				if left >= 0 {
					elem := matrix[i][left]
					if isPartNumber(rune(elem)) {
						valid = true
						break
					}
				}

				if right < M {
					elem := matrix[i][right]
					if isPartNumber(rune(elem)) {
						valid = true
						break
					}
				}

				if bottom < N {
					elem := matrix[bottom][j]
					if isPartNumber(rune(elem)) {
						valid = true
						break
					}
				}

				if top >= 0 {
					elem := matrix[top][j]
					if isPartNumber(rune(elem)) {
						valid = true
						break
					}
				}

				if top >= 0 && left >= 0 {
					elem := matrix[top][left]
					if isPartNumber(rune(elem)) {
						valid = true
						break
					}
				}

				if top >= 0 && right < M {
					elem := matrix[top][right]
					if isPartNumber(rune(elem)) {
						valid = true
						break
					}
				}

				if bottom < N && left >= 0 {
					elem := matrix[bottom][left]
					if isPartNumber(rune(elem)) {
						valid = true
						break
					}
				}

				if bottom < N && right < M {
					elem := matrix[bottom][right]
					if isPartNumber(rune(elem)) {
						// log.Println(elem)
						valid = true
						break
					}
				}
			}

			if valid {
				num, err := strconv.Atoi(string(lineInBytes[startIdx:endIdx]))
				if err != nil {
					log.Fatal(err)
				}
				partNumbers = append(partNumbers, PartNumber{
					Value:    num,
					StartIdx: startIdx,
					EndIdx:   endIdx,
					N:        i,
				})
			}
		}
	}

	type Pos struct {
		N int
		M int
	}

	gearPositions := []Pos{}
	for i, line := range matrix {
		for j, char := range line {
			if char == '*' {
				gearPositions = append(gearPositions, Pos{N: i, M: j})
				log.Println(i, j)
			}
		}
	}

	for _, gearPos := range gearPositions {
		gears := []int{}

		for _, partnum := range partNumbers {
			for j := partnum.StartIdx; j < partnum.EndIdx; j++ {
				left := j - 1
				right := j + 1
				bottom := partnum.N + 1
				top := partnum.N - 1

				if left >= 0 {
					if partnum.N == gearPos.N && left == gearPos.M {
						log.Println(
							partnum.Value,
							"is adjacent to gear in pos",
							gearPos.N,
							gearPos.M,
						)
						gears = append(gears, partnum.Value)
						break
					}
				}

				if right < M {
					if partnum.N == gearPos.N && right == gearPos.M {
						log.Println(
							partnum.Value,
							"is adjacent to gear in pos",
							gearPos.N,
							gearPos.M,
						)
						gears = append(gears, partnum.Value)
						break
					}
				}

				if bottom < N {
					if bottom == gearPos.N && j == gearPos.M {
						log.Println(
							partnum.Value,
							"is adjacent to gear in pos",
							gearPos.N,
							gearPos.M,
						)
						gears = append(gears, partnum.Value)
						break
					}
				}

				if top >= 0 {
					if top == gearPos.N && j == gearPos.M {
						log.Println(
							partnum.Value,
							"is adjacent to gear in pos",
							gearPos.N,
							gearPos.M,
						)
						gears = append(gears, partnum.Value)
						break
					}
				}

				if top >= 0 && left >= 0 {
					if top == gearPos.N && left == gearPos.M {
						log.Println(
							partnum.Value,
							"is adjacent to gear in pos",
							gearPos.N,
							gearPos.M,
						)
						gears = append(gears, partnum.Value)
						break
					}
				}

				if top >= 0 && right < M {
					if top == gearPos.N && right == gearPos.M {
						log.Println(
							partnum.Value,
							"is adjacent to gear in pos",
							gearPos.N,
							gearPos.M,
						)
						gears = append(gears, partnum.Value)
						break
					}
				}

				if bottom < N && left >= 0 {
					if bottom == gearPos.N && left == gearPos.M {
						log.Println(
							partnum.Value,
							"is adjacent to gear in pos",
							gearPos.N,
							gearPos.M,
						)
						gears = append(gears, partnum.Value)
						break
					}
				}

				if bottom < N && right < M {
					if bottom == gearPos.N && right == gearPos.M {
						log.Println(
							partnum.Value,
							"is adjacent to gear in pos",
							gearPos.N,
							gearPos.M,
						)
						gears = append(gears, partnum.Value)
						break
					}
				}
			}
		}

		if len(gears) == 2 {
			sum += gears[0] * gears[1]
		}
	}

	return sum
}
