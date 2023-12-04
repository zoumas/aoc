package solution

import (
	"log"
	"testing"
)

func TestSolution1(t *testing.T) {
	given := `1abc2

pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

	want := 12 + 38 + 15 + 77
	got := Solution1(given)
	if got != want {
		t.Errorf("\ngot: %d\nwant: %d", got, want)
	}
}

func TestSolution2(t *testing.T) {
	given := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

	want := 29 + 83 + 13 + 24 + 42 + 14 + 76
	got := Solution2(given)
	log.Println(got)
	if got != want {
		t.Errorf("\ngot: %d\nwant: %d", got, want)
	}
}
